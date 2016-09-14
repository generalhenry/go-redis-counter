FROM golang:1.7-alpine

RUN apk --no-cache add curl git
RUN curl https://glide.sh/get | sh
WORKDIR /go/src/github.com/generalhenry/go-redis-counter
COPY glide.lock /go/src/github.com/generalhenry/go-redis-counter/
COPY glide.yaml /go/src/github.com/generalhenry/go-redis-counter/
RUN glide install

COPY . /go/src/github.com/generalhenry/go-redis-counter
RUN go install
CMD ["/go/bin/go-redis-counter"]