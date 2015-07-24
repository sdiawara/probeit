FROM golang

RUN mkdir -p /go/src/github.com/sdiawara
RUN mkdir -p /go/bin
RUN git clone https://github.com/sdiawara/probeit /go/src/github.com/sdiawara/probeit
RUN go install github.com/sdiawara/probeit
ENTRYPOINT /go/bin/probeit
 
EXPOSE 3000
