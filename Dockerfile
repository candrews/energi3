# Build energi3 in a stock Go builder container
FROM golang:1.15.8-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /energi3
RUN cd /energi3 && make energi

# Pull energi3 into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /energi3/build/bin/energi /usr/local/bin/

EXPOSE 39796 39795 39797 39797/udp
ENTRYPOINT ["energi"]
