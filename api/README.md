# Hestia API

## Compile proto files

```bash
make proto
```

or

```bash
protoc --go_out=./pb --go-grpc_out=./pb proto/*.proto
```
