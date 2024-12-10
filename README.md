## Project Specs:

1. Gin Framework
2. gRPC

## Endpoint List & Description

1.  Project One

    a. Check Health

        URL: localhost:8080/health
        Method: GET

    b. Send Message

        URL: localhost:8080/message
        Method: POST
        Body:
          {
            "message": "Hello World"
          }

2.  Project Two

    a. Check Health

        URL: localhost:8081/health
        Method: GET

    b. Receiver RPC - Send Message

        URL: localhost:50051
        Proto: message.proto
        Message:
          {
            "message": "Hello World"
          }

## Regenerate gRPC

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative {proto_file_name}.proto
```
