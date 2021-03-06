#!/bin/bash
PAYLOAD="{\"url\": \"$1\"}"
URL="localhost:8080/urls/$RANDOM"
echo "Submitting $1 to $URL"
curl -H "Content-type: application/json" -XPUT $URL -d "$PAYLOAD"
