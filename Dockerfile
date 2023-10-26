FROM golang:1.20-alpine AS builder

RUN mkdir -p /app

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o coreService cmd/api/.

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

EXPOSE 3200

# ENTRYPOINT ["/usr/bin"]

FROM alpine:latest

WORKDIR /usr/bin

COPY --from=builder /app/coreService .

CMD [ "coreService" ]