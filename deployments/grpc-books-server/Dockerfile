#build stage
FROM golang:1.24-alpine AS builder
#upx cto comprise image size
RUN apk add --no-cache git upx
#set WD
WORKDIR /service
COPY go.mod go.sum ./
#download dependencies
RUN go mod download
#copy source code
COPY . ./

#build the go service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./build/out/grpc-books-server cmd/grpc-books-server/main.go

# Compress binary to reduce size
RUN upx --best --lzma ./build/out/grpc-books-server

#final stage
FROM alpine:3.21
RUN apk update --no-cache add ca-certificates
WORKDIR /service

COPY --from=builder /service/build/out/grpc-books-server .

#copy config file
COPY ./configs/grpc-books-server.yaml configs/grpc-books-server.yaml
# Copiar el directorio de migraciones
COPY scripts/migrations ./scripts/migrations
# command to start the api with configFile flag
ENTRYPOINT ["./grpc-books-server", "--configFile=./configs/grpc-books-server.yaml"]