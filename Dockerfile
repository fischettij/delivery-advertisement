FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go build -o advertisement ./cmd/advertisement

EXPOSE 8080

CMD ["./advertisement"]
