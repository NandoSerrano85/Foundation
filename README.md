# Server

## Recompiling Proto

In order to use changes made to proto files you need to recompile proto files

```
protoc --proto_path=. --go_out=api/user --go_opt=paths=import ./proto/user/user.proto 
```