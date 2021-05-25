FROM golang:1 as builder

WORKDIR /go/src/geizhals-api-go

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o geizhals-api-go .

FROM alpine:latest
COPY --from=builder /go/src/geizhals-api-go .

EXPOSE 8080
ENV GIN_MODE=release

CMD ["./geizhals-api-go"]