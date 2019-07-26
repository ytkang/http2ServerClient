# Golang http/2 (H2C) server client example
H2C is http/2 golang library without tls.
In this example, the server can handle http/2 with http/1.1 requests

## Server
```go
h2s := &http2.Server{}
mux := NewMultiplexer()
mux.HandleFunc("/test", TestRequestHandler)

server := &http.Server{
    Addr:    "0.0.0.0:9998",
    Handler: h2c.NewHandler(mux.Handler, h2s),
}

fmt.Printf("Listening [0.0.0.0:9998]...\n")
checkErr(server.ListenAndServe(), "while listening")
```
### Build
go build -o server ./http2Server

## Client
```go
host := "127.0.0.1:9998"
// http2 request
client := NewClient()
sendHttp2Requeest(client, host, "/test", "YT")
sendHttp2Requeest(client, host, "/test", "SJ")
sendHttp2Requeest(client, host, "/test", "JH")

// http1.1 request
sendHttp1Request(host, "/test", "GOGO")

```
### Build
go build -o client ./http2Client

## Protocol
easyjson --all ./http2Protocol/protocol.go

