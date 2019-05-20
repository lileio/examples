FROM golang:alpine AS build-env
WORKDIR /email
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /email/go.mod
COPY go.sum /email/go.sum
RUN go mod download
COPY . /email
RUN CGO_ENABLED=0 GOOS=linux go build -o build/email ./email


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /email/build/email /
ENTRYPOINT ["/email"]
CMD ["up", "--grpc-port=80"]
