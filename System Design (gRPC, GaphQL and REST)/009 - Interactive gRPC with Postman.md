# CMD
- snap install postman
- go run examples/server/*

# Test 1
- Create gRPC API from URL https://raw.githubusercontent.com/MarioCarrion/grpc-microservice-example/063818577723ce82e2dc57a0660a947ab84a911d/wearable/v1/wearable_service.proto
- go run examples/server/main.go
- localhost:9869 | BeatsPerMinute
- localhost:9869 | ConsumeBeatsPerMinute

# Test 2
- Create additional gRPC API from file
- - Import user/v1/user_service.proto
- - Add Import Path: import "user/v1/user.proto" (That's how we defined)
- - localhost:9869 | UserService/GetUser

# Test 3 Reflection
- Reflection setup
- -Import grpc reflection for exposing all grpc services
- -Register reflection with grpcServedr
- -go run examples/server/*
- Create New gRPC API from POSTMAN
- Use server reflection (We don't have to import anything)

# Resources
- [YT](https://www.youtube.com/watch?v=yluYiCj71ss&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=9)
- [Code Example](https://github.com/MarioCarrion/grpc-microservice-example/tree/063818577723ce82e2dc57a0660a947ab84a911d)