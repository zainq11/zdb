# zdb
A key-value store that implements Redis commands. A learning exercise to build a database from the ground up.

## Protobuf command
```bash
 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative api/v1/database.proto
```


## Compile
```bash
go build .
```

## Run Server
```shell
./zdb server
```

## Run Client
```shell
./zdb client
```
