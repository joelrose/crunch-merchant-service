FROM golang:1.19 as builder

WORKDIR /go/src/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/app /app
COPY --from=builder /go/src/certificates /certificates

CMD ["/app"]