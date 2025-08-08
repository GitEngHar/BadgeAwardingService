aws dynamodb scan \
  --table-name badge-service \
  --endpoint-url http://localhost:8000 \
  --output json | jq .
