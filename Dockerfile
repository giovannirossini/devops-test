FROM golang:1.19.5-alpine

COPY ./server.go /opt/server.go

RUN cd /opt; go build /opt/server.go
RUN rm /opt/server.go

CMD ["/opt/server"]