FROM golang:1.24.5-alpine

WORKDIR /app

RUN apk add make bash

SHELL ["/bin/bash", "-c"]

RUN go install github.com/air-verse/air@latest && \
    go install github.com/google/wire/cmd/wire@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN make wire

EXPOSE 8000

CMD ["air"]
