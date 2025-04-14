#build stage
FROM golang:1.24-alpine AS builder
#upx cto comprise image size
RUN apk add --no-cache git upx
#set WD
WORKDIR /app
COPY go.mod go.sum ./
#download dependencies
RUN go mod download
#copy source code
COPY . ./

#build the go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/rest-books-server/main.go

# Compress binary to reduce size
RUN upx --best --lzma main

#final stage
FROM alpine:3.21
RUN apk update --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/main .

#copy config file
COPY configs/config.yaml ./configs/config.yaml
# Copiar el directorio de migraciones
COPY scripts/migrations ./scripts/migrations
# command to start the api with configFile flag
ENTRYPOINT ["./main", "--configFile=./configs/config.yaml"]