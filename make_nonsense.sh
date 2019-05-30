#!/bin/bash

while read line; do curl -XPOST -d "${line}" localhost:8080/v1/poynt/brian; done < many_poynts.json

# Array of executables to check for
REQBIN=( jq )

function check_for {
	echo -n "Checking for $1..."
	which $1 > /dev/null
	result=$?
	if [ "$result" -eq "1" ]
	then
		echo -n "Error"
		echo
		exit 2
	fi
	echo "ok"
}

echo
echo "brew install jq or it will fail from here on out"

for K in "${!REQBIN[@]}"; do
	check_for ${REQBIN[$K]}
done

# Show all keys in a namespace
curl -XGET localhost:8080/v1/poynt/brian | jq .

# sort is always Logical, Observational, Operational
curl -XGET localhost:8080/v1/poynt/brian?sort_by=a,a,a | jq
curl -XGET localhost:8080/v1/poynt/brian?sort_by=a,a,a | jq
curl -XGET localhost:8080/v1/poynt/brian?sort_by=d,a,a | jq
curl -XGET localhost:8080/v1/poynt/brian?sort_by=a,a,d | jq

# Every keyspace has a log / obs / ops path
curl -XGET 'localhost:8080/v1/poynt/brian/20180402000100.000' | jq .
curl -XGET 'localhost:8080/v1/poynt/brian/20180402000100.000/20180511134900.000' | jq .

# Logical Ascending, Observational Ascending, Operational Decending, return only 1 result
curl -XGET 'localhost:8080/v1/poynt/brian?sort_by=a,a,d&limit=1'  | jq .

# Logical Ascending, Observational Ascending, Operational Ascending, return only 1 result
curl -XGET 'localhost:8080/v1/poynt/brian?sort_by=a,a,d&limit=1'  | jq .

# Logical Ascending, Observational Ascending, Operational Ascending, return only 1 result
# Also filter 'data' so it only returns "other" field
curl -XGET 'localhost:8080/v1/poynt/brian?sort_by=a,a,d&limit=1&col=other'  | jq .

# Logical time and Obs time identical, will return all the poynts with different op times.
curl -XGET 'localhost:8080/v1/poynt/brian/20180402000100.000/20180411131000.000' | jq .

# use COL to get the column you're interested in if the nested data is JSON.
# if it's not JSON this is will like blow up in your face horribly.
curl -XGET 'localhost:8080/v1/poynt/brian?col=other&sort_by=d,d,d' | jq .
