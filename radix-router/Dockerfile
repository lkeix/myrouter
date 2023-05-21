FROM golang:1.19

WORKDIR /app

COPY . /app

RUN go get github.com/lkeix/myrouter@inspect

CMD go test -bench . -benchmem