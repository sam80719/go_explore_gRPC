# go_explore_gRPC

## gRPC environment install [gRPC GO](https://grpc.io/docs/languages/go/quickstart/)
1.  install protobuf on your system
```shell
# Linux 
$ apt install -y protobuf-compiler

# Mac
$ brew install protobuf

$ protoc --version  # Ensure compiler version is 3+
```

2. Go plugins for the protocol compiler
```shell
# Install the protocol compiler plugins
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Install the protocol compiler plugins for Go
$ export PATH="$PATH:$(go env GOPATH)/bin"
```


##  create proto by protocol compiler
```shell
## how to use protoc to gen proto
protoc -Igreeting/proto --go_opt=module=${YOUR_MODULE_NAME} --go_out=${WHERE_TO_GENERATE} --go-grpc_opt=module=${YOUR_MODULE_NAME} --go-grpc_out=${WHERE_TO_GENERATE} greeting/proto/dummy.proto
## example: you use this cmd, then /greeting/proto will add two document "greeting.pb.go" and "greeting_grpc.pb.go"
protoc -Igreeting/proto --go_out=. --go_opt=module=github.com/sam80719/go_explore_gRPC --go-grpc_out=. --go-grpc_opt=module=github.com/sam80719/go_explore_gRPC greeting/proto/greeting.proto
```


## TEST
simple to test blog function 
```shell
$ go test ./blog/server

# individual test
$ go test ./blog/server/server_test.go ./blog/server/server.go
```

## Note
- folder greeting is learning Grpc
- folder calculator is used to implement the grpc just learned
- folder blog is used to realize the interaction between gRPC and the database