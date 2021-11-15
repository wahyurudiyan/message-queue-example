# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go get -v && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bin/app .

# generate clean, final image for end users
FROM alpine:latest
COPY --from=builder /bin/app .

# executable
ENTRYPOINT [ "./app" ]

# expose amqp port
EXPOSE 5672