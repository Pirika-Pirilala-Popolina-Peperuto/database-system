# database-system
Combine ent and gRPC, using Postgres as DBMS

## How to develop
### Install tools
```bash
brew install protobuf

go mod tidy

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
  go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest && \
  go install github.com/silenceper/gowatch
 
git config core.hooksPath githooks
```
