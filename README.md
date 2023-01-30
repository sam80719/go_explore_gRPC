# go_explore_gRPC

## 1. create proto
```shell
protoc -Igreeting/proto --go_opt=module=${YOUR_MODULE_NAME} --go_out=${WHERE_TO_GENERATE} --go-grpc_opt=module=${YOUR_MODULE_NAME} --go-grpc_out=${WHERE_TO_GENERATE} greeting/proto/dummy.proto
protoc -Igreeting/proto --go_out=. --go_opt=module=github.com/sam80719/go_explore_gRPC --go-grpc_out=. --go-grpc_opt=module=github.com/sam80719/go_explore_gRPC greeting/proto/dummy.proto
```

## Note
- folder greeting is learning Grpc
- folder calculator is used to implement the grpc just learned
- folder blog is used to realize the interaction between gRPC and the database