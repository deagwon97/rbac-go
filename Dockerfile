FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /root/src

COPY ./src /root/src

RUN go mod download

RUN go build -o main ./

WORKDIR /dist

RUN cp /root/src/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]