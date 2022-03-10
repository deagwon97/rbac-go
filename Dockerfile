FROM golang:latest AS builder

RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    curl -sL https://deb.nodesource.com/setup_16.x | bash - &&\
    apt-get install nodejs -y

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /root/workdir/src

COPY ./src /root/workdir/src

RUN go mod download

RUN go build -o main ./

WORKDIR /dist

RUN cp /root/workdir/src/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]