FROM golang:1.16-alpine
WORKDIR /app
COPY . .
ARG DEFAULT_PORT 
ENV port $DEFAULT_PORT
EXPOSE $port
 # VOLUME [ "/data" ] # anonymous volumes
RUN go mod vendor
RUN go build -o /notifier
CMD ["/notifier"]