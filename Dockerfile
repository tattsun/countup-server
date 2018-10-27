FROM golang:1.11-alpine as builder
COPY ./main.go ./
RUN go build -o app ./main.go

FROM alpine
EXPOSE 8080
COPY --from=builder /go/app .
ENTRYPOINT ["/app"]