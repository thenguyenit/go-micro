Generate go-micro
```bash
protoc -I/usr/local/include -I. -I$GOPATH/src --micro_out=. greeter.proto
```

Generate gRPC gateway
```bash
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --micro_out=. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. greeter.proto
```

Testing

Start Service
```bash
go run srv/main.go --server_address=localhost:9090
```

Start Gateway
```bash
go run gateway/main.go
```

CURL to Testing
```bash
curl -X POST 'http://localhost:8080/greeter/hello?name=NTN'
```