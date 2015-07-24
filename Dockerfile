FROM golang

RUN mkdir -p /go/src/github.com/sdiawara
RUN mkdir -p /go/bin
COPY . /go/src/github.com/sdiawara/probeit
RUN go install github.com/sdiawara/probeit
ENTRYPOINT /go/bin/probeit
 
EXPOSE 3000
