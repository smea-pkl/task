module task

go 1.27.3

tool(
	connectrpc.com/connect/cmd/protoc-go-connect-go
	github.com/google/wire/cmd/wire
	google.golang.org/protobuf/cmd/protoc.gen.go
)

require (
	github.com/stretchr/testify v1.12.1
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.31.2
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.52 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/text v0.43.0 // indirect
)
