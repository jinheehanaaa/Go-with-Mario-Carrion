# gRPC
- is a remote procedure call framework

# 4 Kinds of Service Methods
1. Unary RPC
2. Server Streaming RPC
3. Client Streaming RPC
4. Bi-directional Streaming RPC

# Unary RPC
- Client sends single request to the server & gets single response back

# Setup
## Implementation
1. Install protoc gen go
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
2. buf.gen.yaml
- add go-grpc plugin
3. Define "user.proto"
4. Define & work with ("user_service.proto")
- Define UserService
- Define GetUserRequest (Request with user id)
- Define GetUserResponse (Response to user)
5. Call the interface "UserServiceServer" ("main.go")
## Server Run
1. Implement Typical Server Logic ("main.go")
- Use <code>-require_unimplemented_servers=false # XXX</code> for now
- go run main.go

## Client Run
- go run main.go


# CMD
- <code>buf lint</code> for checking the syntax
- <code>buf generate</code> for generating both client and service for gRPC Service

# TODO
- PRACTICE THE gRPC FLOW!

# Resources
- [YT](https://www.youtube.com/watch?v=GgR4JDVgPOc&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=4)
- [Code Example](https://github.com/MarioCarrion/grpc-microservice-example/commit/d0c4f126be2db0d1fa4f458dfc537d39ea1aec5b)