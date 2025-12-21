#!/bin/bash

echo "Sending 1000 requests to KEDA HTTP proxy..."

for i in {1..1000}; do
  curl -s -H "Host: go-http.local" http://localhost:8081/work > /dev/null &
done

wait

echo "Done sending requests."

