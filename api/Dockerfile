FROM golang:1.17-stretch AS builder

RUN apt update && apt install git

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o /app/bin/api cmd/api/*.go

FROM alpine

COPY --from=builder /app/bin/api /usr/bin/
RUN chmod +x /usr/bin/api

ENTRYPOINT ["/usr/bin/api"]
