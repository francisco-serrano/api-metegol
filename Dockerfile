FROM golang:latest

ENV API_DIR $GOPATH/src/github.com/api-metegol
ENV API_PORT 8080

WORKDIR $API_DIR

COPY ./controllers $API_DIR/controllers
COPY ./routers $API_DIR/routers
COPY ./views $API_DIR/views
COPY ./main.go $API_DIR/

COPY ./vendor/vendor.json $API_DIR/vendor/

RUN go get -u github.com/kardianos/govendor

RUN govendor init
RUN govendor sync
RUN go build main.go

EXPOSE 8080

CMD ["./main"]