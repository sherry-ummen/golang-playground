FROM golang:latest

RUN mkdir -p /go/src/api_with_kubernetes

WORKDIR /go/src/api_with_kubernetes

COPY . /go/src/api_with_kubernetes

RUN go install

CMD /go/bin/api_with_kubernetes

EXPOSE 8080

