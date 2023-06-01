FROM golang:1.19

WORKDIR /app

COPY . /app

RUN go get github.com/lkeix/myrouter@v0.0.0

CMD go test -bench . -benchmem
