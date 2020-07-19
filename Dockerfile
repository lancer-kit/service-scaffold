# Compile stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git bash

#ENV GOPROXY=direct
ENV GO111MODULE=on
ENV GOPRIVATE=gitlab.com

ARG CONFIG="master"

WORKDIR /service
ADD . .
COPY ./env/${CONFIG}.config.yaml /config.yaml
RUN go get -v && ./build.sh /app

# Final stage
FROM alpine:3.7

# Port 8080 belongs to our application
EXPOSE 8080

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache ca-certificates bash

WORKDIR /

COPY --from=build-env /app /
COPY --from=build-env /config.yaml /

# Run app
CMD ["/app", "serve"]
