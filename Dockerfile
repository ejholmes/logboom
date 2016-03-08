FROM golang

COPY . /go/src/github.com/ejholmes/logboom
RUN go install github.com/ejholmes/logboom

ENTRYPOINT ["/go/bin/logboom"]
