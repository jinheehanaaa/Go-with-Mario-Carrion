# Bi-directional Streaming RPC
- Both streams are independent to each other
- Server & Client can read & write messages to each other in any order

# STEP
## Define New Method in Service (wearable_service.proto)
- CalculateBeatsPerMinute rpc in service
- Message for Request & Response
- buf lint & buf generate
- Copy Method "CalculateBeatsPerMinute" from wearable_service_grpc.pb.go
## [Server]
- Implementation for CalculateBeatsPerMinute (wearable_service.go)
- - stream.Recv
- - EOF in case the client is not sending message anymore
- - Implement Calculation for sending 5 messages
## [Client]
- Connecting to the Server
- client.CalculateBeatsPerMinutes()
- - stream.Send()
- - stream.CloseSend()
- - stream.Recv()

## RUN
go run examples/server/main.go
go run examples/bidirectional-client/main.go

# Resources
- [YT](https://www.youtube.com/watch?v=Z9SHd1K1k5Y&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=7)
- [Code Example](https://github.com/MarioCarrion/grpc-microservice-example/tree/c38dbd7cc0afa727a0b6b87ab47283601d60b852)