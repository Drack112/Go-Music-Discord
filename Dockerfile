FROM golang:1.19-bullseye as builder

WORKDIR /go/src/app

RUN apt-get update
RUN apt-get install git bash curl gcc ffmpeg youtube-dl libopus-dev build-essential -y

COPY go.mod /go/src/app/
COPY go.sum /go/src/app/

RUN go mod download
RUN go mod tidy

COPY . /go/src/app/

RUN go build -o /go/src/app/main

CMD ["/go/src/app/main"]