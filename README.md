# Server

## Recompiling Proto

In order to use changes made to proto files you need to recompile proto files

```
protoc --go_out=/Users/fserrano/Documents/Projects/Exprience_App/backend/api/user  \
       --go-grpc_out=/Users/fserrano/Documents/Projects/Exprience_App/backend/api/user \
       proto/user/user.proto
```