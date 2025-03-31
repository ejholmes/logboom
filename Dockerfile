FROM golang

COPY . /go/src/github.com/ejholmes/logboom
RUN cd /go/src/github.com/ejholmes/logboom && go install .

ENTRYPOINT ["/go/bin/logboom"]
