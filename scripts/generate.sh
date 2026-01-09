#!/bin/bash

cd "$(dirname "$0")/.." || exit

# Event proto
protoc -I ./event-service/api \
  -I ./event-service/api/google/api \
  --go_out=./event-service/internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./event-service/internal/pb --go-grpc_opt=paths=source_relative \
  ./event-service/api/event_api/event.proto ./event-service/api/models/event_model.proto

protoc -I ./event-service/api \
  -I ./event-service/api/google/api \
  --grpc-gateway_out=./event-service/internal/pb \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt logtostderr=true \
  ./event-service/api/event_api/event.proto

protoc -I ./event-service/api \
  -I ./event-service/api/google/api \
  --openapiv2_out=./event-service/internal/pb/swagger \
  --openapiv2_opt logtostderr=true \
  ./event-service/api/event_api/event.proto

# Notify proto (similar)
protoc -I ./notify-service/api \
  -I ./notify-service/api/google/api \
  --go_out=./notify-service/internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./notify-service/internal/pb --go-grpc_opt=paths=source_relative \
  ./notify-service/api/notify_api/notify.proto ./notify-service/api/models/subscription_model.proto

protoc -I ./notify-service/api \
  -I ./notify-service/api/google/api \
  --grpc-gateway_out=./notify-service/internal/pb \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt logtostderr=true \
  ./notify-service/api/notify_api/notify.proto

protoc -I ./notify-service/api \
  -I ./notify-service/api/google/api \
  --openapiv2_out=./notify-service/internal/pb/swagger \
  --openapiv2_opt logtostderr=true \
  ./notify-service/api/notify_api/notify.proto