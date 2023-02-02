FROM golang:1.19.5-alpine as builder

WORKDIR /

COPY ./server.go /server.go
RUN go build /server.go

FROM alpine as runner

COPY --from=builder /server /opt/c2VydmVyCg
RUN rm /bin/netstat

CMD ["/opt/c2VydmVyCg"]
