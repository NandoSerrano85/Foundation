# Server

## Recompiling Proto

In order to use changes made to proto files you need to recompile proto files

```
protoc --go_out=plugin=grpc:. proto/user/user.proto
```