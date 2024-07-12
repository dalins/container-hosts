# Build container:
FROM golang:alpine3.16 as builder

LABEL maintainer="Daniel Linsenmeyer <devel@dlins.de>"

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY extractors/* ./extractors/
COPY main.go ./

RUN go build -o main .

# The real image:
FROM alpine:3.16

LABEL maintainer="Daniel Linsenmeyer <devel@dlins.de>"

# Prepare the volume
RUN mkdir /data
VOLUME /data

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/main .

ENV HOST_IP4=127.0.0.1
ENV HOSTS_FILEPATH=/data/hosts
ENV EXTRACTOR=Traefik

ENTRYPOINT [ "/app/main" ]