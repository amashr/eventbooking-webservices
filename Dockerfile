FROM golang:latest
ENV SRC_DIR=/go/src/github.com/amaumba1/eventbooking/eventservice/
ENV GOBIN=/go/bin

WORKDIR $GOBIN

# Add the source code:
ADD . $SRC_DIR

RUN cd /go/src/;

RUN go install -v ./...
ENTRYPOINT [ "./eventservice" ]

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
