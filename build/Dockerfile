FROM golang:latest

ENV PATH=$PATH:$GOPATH/bin:/.local/bin
ENV PB_REL=https://github.com/protocolbuffers/protobuf/releases
ENV PB_VERSION=3.17.3

RUN apt update && apt install -y unzip

RUN curl -LO $PB_REL/download/v$PB_VERSION/protoc-$PB_VERSION-linux-x86_64.zip
RUN unzip protoc-$PB_VERSION-linux-x86_64.zip -d /.local
RUN rm protoc-$PB_VERSION-linux-x86_64.zip

WORKDIR /database-system

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest

RUN go generate ./...
RUN go build -o database-system ./cmd/grpc

CMD ["./database-system"]

