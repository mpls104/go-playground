FROM golang:latest

WORKDIR /go

CMD ["go", "run", "main.go"]
