FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/orijtech/asana/v1

RUN go get github.com/odeke-em/asana/v1

WORKDIR /go/src/github.com/heaptracetechnology/microservice-asana

ADD . /go/src/github.com/heaptracetechnology/microservice-asana

RUN go install github.com/heaptracetechnology/microservice-asana

ENTRYPOINT microservice-asana

EXPOSE 3000