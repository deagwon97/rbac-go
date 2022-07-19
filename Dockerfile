FROM golang:latest AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    curl -sL https://deb.nodesource.com/setup_16.x | bash - &&\
    apt-get install nodejs -y

RUN useradd -ms /bin/bash rbac

COPY ./src /home/rbac/workdir/src

USER root
WORKDIR /home/rbac/workdir/src/admin
RUN npm install
RUN npm run build

WORKDIR /home/rbac/workdir/src
RUN go mod tidy
RUN go build -o main ./

WORKDIR /dist
RUN cp /home/rbac/workdir/src/main /dist/main
RUN mkdir -p /dist/admin/build
RUN cp -r /home/rbac/workdir/src/admin/build /dist/admin

FROM scratch AS production
COPY --from=builder /dist .
ENTRYPOINT ["/main"]