FROM golang:alpine AS builder

WORKDIR /go/src/github.com/amaumba1/eventbooking
COPY . .

WORKDIR /go/src/github.com/amaumba1/eventbooking

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=0 /go/src/github.com/amaumba1/eventbooking/src/bookingservice/bookingservice /bookingservice

WORKDIR /src/bookingservice/eventservice

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181

CMD [ "./main" ]


