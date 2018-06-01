Bidirectional streaming 双向数据流模式的 gRPC
客户端将连续的数据流发送到服务端，服务端返回交互的数据流。
client 依次请求 1、2、3 的用户数据流，服务端依次返回 1、2、3 的用户数据流：