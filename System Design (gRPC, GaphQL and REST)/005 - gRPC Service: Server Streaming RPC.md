# Server Streaming RPC
- Client requests to the server & gets streaming to read a sequence of messages back

# Hypothetical Wearable Device
- Measures heat beat per minute

# STEP
## Implement Protocol Buffer
1. Use "stream" keword as return in service (wearable_service.proto)
2. buf lint, buf generate
3. Look for interface type "WearableServiceServer" & Copy (wearable_service_grpc.pb.go)
## [Server](code/005/EX-001/examples/server/main.go)
1. Paste interface in server (wearable_service.go)
- Implement BetsPerminute for sending the data
2. Continue send data constantly (wearable_service.go)
- for loop with select/case
- - Stream End
- - Default for sending data
## Instantiate the Server
1.  Instantiation & Registration (main.go)
- Instantiate wearableServer
- Register wearableServer to gRPC Server
## [Client](code/005/EX-001/examples/stream-client/main.go)
- Connecting to WearableServiceClient (main.go)
- Invoke the connection
- Use goroutine
- - stream.Recv (Blocking method)
- - Update
- - Render
- Support the stream.Context coming from the client
- go run examples/server/main.go
- go run examples/stream-client/main.go

# CMD
- buf lint
- buf generate

# Resources
- [YT](https://www.youtube.com/watch?v=l_74x_qQZB8&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=5)
- [Code Example](https://github.com/MarioCarrion/grpc-microservice-example/tree/edee38ce37f13033af764a3120ad9079bcfb802c)