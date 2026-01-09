module github.com/Timofey0880/GoProject2

go 1.21

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.3
	github.com/jackc/pgx/v5 v5.7.6
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.6.1
	github.com/samber/lo v1.52.0
	github.com/segmentio/kafka-go v0.4.49
	github.com/stretchr/testify v1.11.1
	github.com/swaggo/http-swagger v1.3.4
	go.yaml.in/yaml/v4 v4.0.0-rc.2
	google.golang.org/genproto/googleapis/api v0.0.0-20250929231259-57b25ae835d4
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.10
)

require ( // same as example
	// ...
)

tool (
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	google.golang.org/protobuf/cmd/protoc-gen-go
)