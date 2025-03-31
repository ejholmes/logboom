FROM golang

COPY . /go/src/github.com/ejholmes/logboom
RUN cd /go/src/github.com/ejholmes/logboom && go install .

RUN useradd -u 999 john
USER logboom

ENTRYPOINT ["/go/bin/logboom"]
