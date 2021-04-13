FROM golang:alpine AS build

RUN apk --no-cache add gcc g++ make
RUN apk add git

WORKDIR /go/src/app

COPY ./api .

RUN go mod download

CMD go test ./...