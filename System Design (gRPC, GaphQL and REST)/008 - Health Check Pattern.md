# Health Check Pattern
- Consist of implementing a way to expose the health of our service to indicate if everything is performing correctly
- EX: Connection to the services, CPU, Memory

# STEP
## SERVER
- Use gRPC Health Protocol
- - Import grpc heath package
- - Instantiate health.NewServer
- - Register healthServer
- Implement Heath Check to our Server
- - healthServer.SetServingStatus
- Use goroutine to constantly check the health
- go run examples/server/main.go
- go install github.com/grpc-ecosystem/grpc-health-probe@v0.4.14
- grpc-health-probe -addr="0.0.0.0:9879" -service=""

# Resources
- [YT](https://www.youtube.com/watch?v=3iPVSGvpC-k&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=8)
- [Code Example](https://github.com/MarioCarrion/grpc-microservice-example/tree/81042d9af4f3a65d2671e6d9f827d93d119a3274)