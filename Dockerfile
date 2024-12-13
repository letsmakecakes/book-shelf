FROM golang:alpine

WORKDIR /book-shelf
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/book-shelf/bin/api"]
EXPOSE 8080