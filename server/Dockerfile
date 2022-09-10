FROM golang:alpine as builder
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crasher cmd/crasher/main.go

FROM alpine
COPY --from=builder app/crasher /usr/bin/crasher
WORKDIR /usr/bin
CMD ["./crasher"]
