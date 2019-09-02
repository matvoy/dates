FROM golang:1.12.5 as run
RUN mkdir -p /app
WORKDIR /app
COPY . /app
ENV GO111MODULE on
RUN if [ ! -d "vendor" ]; then go mod vendor; fi
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o /app/bin/dates /app/cmd/...
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
EXPOSE 3333
CMD ["/app/bin/dates"]
