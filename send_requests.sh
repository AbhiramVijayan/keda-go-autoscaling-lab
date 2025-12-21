#!/bin/bash

URL="http://localhost:8081/work"
HOST_HEADER="go-http.local"
REQUESTS=1000

echo "Sending $REQUESTS requests to KEDA HTTP proxy..."

for i in $(seq 1 $REQUESTS); do
  curl -s -H "Host: $HOST_HEADER" "$URL" > /dev/null &
done

wait

echo "Done sending requests."

