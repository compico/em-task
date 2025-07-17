FROM golang:1.24.5-alpine

WORKDIR /app

RUN apk add make bash

SHELL ["/bin/bash", "-c"]

COPY go.mod go.sum Makefile ./

RUN make download_toolchain
RUN go mod download -x

RUN make wire

EXPOSE 8000

CMD ["air"]
