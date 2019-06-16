FROM golang:1.12

WORKDIR /go/src/github.com/amaumba1/eventbooking
COPY . .

# Download all the dependencies 
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181

# Run the executable
CMD [/"bookingservice"]