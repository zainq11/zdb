# Key value store in Go
A simple implementation inspired by Redis for learning purposes.

Tasks
- [ ] Build CLI with the following subcommand
      - [ ] `server` to start a grpc server that servers string commands.
      - [ ] read config from a config file 
      - [ ] grpc client to tunnel string commands to the server. Basically, default client mode is to connect to localhost.
      - [ ] implement [SET](https://redis.io/commands/set/)
      - [ ] implement [GET](https://redis.io/commands/get/)





To learn
- [ ] How to do grpc client?
- [ ] `log` vs `fmt` packages
- [ ] Read code of log and net packages to understand a bit more.
- [ ] Mocking libraries in golang.


Protobuf command
```bash
 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative api/v1/database.proto
```