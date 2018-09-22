FROM golang:latest

WORKDIR /go/src/app
COPY . . 

RUN go get github.com/urfave/cli
RUN go get -d -v ./ ...
RUN go install -v ./ ...

ENTRYPOINT ["app"]
