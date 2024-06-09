# Common

This is the common code used by all Mosha services.

## Protos/gRPC

Prerequisites:

* Have protoc installed: https://grpc.io/docs/protoc-installation/

The communication between services is done using gRPC. 
All service protos definitions are located in `proto` directory.

To regenerate the gRPC code, run in the `proto` folder:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  ./**/*.proto
```
