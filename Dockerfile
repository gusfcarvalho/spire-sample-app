FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o spire-app main.go


FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/spire-app .
ENV SPIFFE_ENDPOINT_SOCKET=/run/spire/sockets/public/api.sock
ENTRYPOINT ["./spire-app"]
