FROM golang:1.19.5-alpine as builder

WORKDIR /

COPY ./server.go /server.go

RUN go build /server.go


FROM alpine as runner

WORKDIR /opt

COPY --from=builder /server /opt/server
RUN rm /bin/netstat

CMD ["/opt/server"]