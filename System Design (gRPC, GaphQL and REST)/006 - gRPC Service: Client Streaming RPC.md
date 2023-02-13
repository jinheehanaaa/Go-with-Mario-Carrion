# Client Streaming RPC
- Client writes sequence of messages & send to the server using Stream

# STEP
## Update WearableService Proto (wearable_service.proto)
- Add new RPC service "ConsumeBeatsPerMinute"
- Add messages for Request & Response
- buf lint, buf generate
- git diff => Get ConsumBeatsPerMinute interface from v1/wearable_service_grpc.pb.go
## [Server](code/006/EX-001/examples/server/wearable_service.go)
- [ConsumeBeatsPerMinute](code/006/EX-001/examples/server/wearable_service.go#L40)
- - stream.Recv
## [Client](code/006/EX-001/examples/client-streaming-client/main.go)
- Connection Implementation
- [ConsumeBeatsPerMinute](code/006/EX-001/examples/client-streaming-client/main.go#L27)
-- stream.Send
-- stream.CloseAndRecv

# Resources
- [YT](https://www.youtube.com/watch?v=wKkaCxdfEh8&list=PL7yAAGMOat_EX1nv8fgltlm0CnJTH8Nwg&index=6)
- [Code Example](https://github.com/MarioCarrion/grpc-microservice-example/tree/ca6e51eabc0857c184467bb23cee8b89d5a2ee7b)