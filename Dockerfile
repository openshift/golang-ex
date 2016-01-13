FROM golang:1.5

USER nobody

ENV GO15VENDOREXPERIMENT=1

RUN mkdir -p /go/src/github.com/openshift/golang-ex
WORKDIR /go/src/github.com/openshift/golang-ex

COPY . /go/src/github.com/openshift/golang-ex
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]
