FROM golang:1.19

WORKDIR /app
COPY . /app

RUN go get github.com/lkeix/myrouter@section2-trie-router

CMD go test -bench . -benchmem