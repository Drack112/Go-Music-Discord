FROM golang:1.19-bullseye as builder

WORKDIR /go/src/app

RUN apt-get update

# build-essential and gcc to compile libopus-dev, both to work with gopus
RUN apt-get install bash gcc ffmpeg python3 python3-pip libopus-dev build-essential -y

# instead of use youtube-dl, yt-dlp is a better option since he's a fork with better perfomance and less errors
RUN pip3 install yt-dlp

COPY go.mod /go/src/app/
COPY go.sum /go/src/app/

RUN go mod download
RUN go mod tidy

COPY . /go/src/app/

RUN go build -o /go/src/app/main

CMD ["/go/src/app/main"]