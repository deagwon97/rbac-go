FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /src

COPY ./src /src

RUN go mod download

RUN go build -o main ./

WORKDIR /dist

RUN cp /src/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]