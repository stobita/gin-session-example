FROM golang:1.14-alpine3.12
WORKDIR /api
COPY . /api
RUN apk add --update --no-cache \
      git && \
      go get github.com/cortesi/modd/cmd/modd
CMD ["modd", "-f", "./modd.conf"]
