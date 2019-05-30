package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"strings"
	"time"
)

type S3PoyntStore struct {
	Bucket string
	Pre    string
	Region string
}

func NewS3PoyntStore(b string, p string, r string) *S3PoyntStore {
	return &S3PoyntStore{
		Bucket: b,
		Pre:    p,
		Region: r,
	}
}

func (s *S3PoyntStore) Get(key string) shaper {
	var pc PoyntCollection
	keys := s.keysFromKeyspace(key)
	pc.store = s.KeysToPoynts(key, keys)
	return &pc
}

func (s *S3PoyntStore) Write(key string, p Poynt) bool {
	tojoin := []string{
		s.Keyspace(key),
		TimeToString(p.Logt),
		TimeToString(p.Obst),
		TimeToString(p.Opt),
		"data.json",
	}
	path := strings.Join(tojoin, "/")
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(s.Region)})
	params := &s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	}
	_, err := svc.HeadObject(params)

	switch {
	case err == nil:
		fmt.Println("File exists, not writing point")
		return false
	case err != nil && !(strings.Contains(err.Error(), "404")):
		fmt.Println(err)
		return false
	}
	s.WritePoynt(path, p.Data)
	return true
}

func (s *S3PoyntStore) WritePoynt(path string, payload []byte) error {
	uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String(s.Region)}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(payload),
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	})
	fmt.Println("Uploaded", result.Location)
	return err
}

func (s *S3PoyntStore) KeysToPoynts(key string, keys []string) []Poynt {
	var poynts []Poynt
	for _, v := range keys {
		//removing the keyspace and the data blob, so we just have time dimensions
		mystring := strings.TrimRight(strings.TrimLeft(v, s.Keyspace(key)), "/data.json")
		dims := strings.Split(mystring, "/")
		var times []time.Time
		for _, t := range dims {
			converted, _ := StringToTime(t)
			times = append(times, converted)
		}
		poyntData := s.DataForKey(v)
		p := Poynt{
			Name: key,
			Logt: times[0],
			Obst: times[1],
			Opt:  times[2],
			Data: poyntData}
		poynts = append(poynts, p)
	}
	return poynts
}

func (s *S3PoyntStore) Keyspace(key string) string {
	if s.Pre == "" {
		return key
	} else {
		return s.Pre + "/" + key
	}
}

func (s *S3PoyntStore) List() []string {
	var keys []string

	svc := s3.New(session.New(), &aws.Config{Region: aws.String(s.Region)})
	var next *string
	for {
		params := &s3.ListObjectsInput{Bucket: aws.String(s.Bucket), Marker: next, Prefix: aws.String(s.Pre + "/"), Delimiter: aws.String("/")}
		resp, _ := svc.ListObjects(params)
		for _, obj := range resp.CommonPrefixes {
			keys = append(keys, *obj.Prefix)
		}
		if *resp.IsTruncated {
			next = resp.Contents[len(resp.Contents)-1].Key
		} else {
			break
		}
	}
	return keys
}

func (s *S3PoyntStore) keysFromKeyspace(key string) []string {
	keyspace := s.Keyspace(key)
	fmt.Println("Pulling keys for keyspace:", keyspace)
	var keys []string

	svc := s3.New(session.New(), &aws.Config{Region: aws.String(s.Region)})
	var next *string
	for {
		params := &s3.ListObjectsInput{Bucket: aws.String(s.Bucket), Marker: next, Prefix: &keyspace}
		resp, _ := svc.ListObjects(params)
		for _, obj := range resp.Contents {
			if !(*obj.Key == keyspace) {
				keys = append(keys, *obj.Key)
			}
		}
		if *resp.IsTruncated {
			next = resp.Contents[len(resp.Contents)-1].Key
		} else {
			break
		}
	}
	return keys
}

func (s *S3PoyntStore) DataForKey(key string) json.RawMessage {
	var payload json.RawMessage
	p := aws.NewWriteAtBuffer([]byte{})
	dler := s3manager.NewDownloader(session.New(&aws.Config{Region: aws.String(s.Region)}))
	dler.Download(p,
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(key),
		})
	json.Unmarshal(p.Bytes(), &payload)
	return payload
}
