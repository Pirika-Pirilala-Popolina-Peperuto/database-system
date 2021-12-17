# database-system
Combine ent and gRPC, using Postgres as DBMS

## How to develop
### Install tools
```bash
brew install protobuf

go mod tidy

go install \
  google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
  entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest \
  github.com/silenceper/gowatch
 
git config core.hooksPath githooks
```
