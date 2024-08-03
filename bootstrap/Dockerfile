FROM --platform=linux/amd64 golang:1.22-alpine AS builder

WORKDIR /app
RUN apk --no-cache update && \
    apk --no-cache add git gcc libc-dev

# Kafka Go client is based on the C library librdkafka
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor

RUN export GO111MODULE=on

COPY . .
RUN go mod vendor
RUN go build -tags musl --ldflags "-extldflags -static" -o main .
# -tags musl: Tells Go to use the musl C library instead of the default glibc. This is common for Alpine Linux, which uses musl.
# --ldflags "-extldflags -static": Statically links the Go binary. This is useful for creating a self-contained binary that doesn't rely on dynamic libraries.

CMD [ "./main" ]