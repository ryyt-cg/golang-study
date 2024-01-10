# Greeting gRPC
* gRPC uses Protocol Buffers for communications.
* Protobuf has much less payload size comparing to JSON that leading to use much less Network Bandwidth.
* Parsing JSON requires more CPU usage. 
* Parsing Protobuf (binary format) requires less CPU
* Mobile devices friendly because payload size and CPU usage
* Multiplexing communication
* Support many languages

## Why Protobuf?
* Easy to define message
* API definition is independent from the implementation
* Stub can be generated to any languages
* Binary format, small payload
* Evolving without breaking existing clients, no versioning requires

## HTTP/2
* gRPC uses HTTP/2
* https://imagekit.io/demo/http2-vs-http1
* HTTP/2 performs much quicker on slow network bandwidth


## 4 Types of API
1. Unary like REST API
2. Client Streaming
3. Server Streaming
4. Bidirectional Streaming

