FROM golang:1.16.4-buster AS builder

ARG VERSION=0.1.0

WORKDIR /go/src/app
ENV GOPROXY https://goproxy.io
COPY . .
RUN go mod download
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go 

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]
