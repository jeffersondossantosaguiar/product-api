FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && \
  go build -o main cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD [ "./main" ]