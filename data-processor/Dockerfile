FROM golang:1.17-stretch AS builder

RUN apt update && apt install git

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o /app/bin/processor cmd/processor/*.go

FROM alpine

COPY --from=builder /app/bin/processor /usr/bin/
RUN chmod +x /usr/bin/processor

ENTRYPOINT ["/usr/bin/processor"]
