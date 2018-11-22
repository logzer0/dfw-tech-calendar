FROM golang:1.10-alpine3.7 as build

WORKDIR /go/src/app

COPY . /go/src/app

RUN ls /go/src/app

RUN go build -o app

FROM alpine:3.7

RUN apk --update add ca-certificates

COPY --from=build /go/src/app/app /usr/local/bin/app
COPY --from=build /go/src/app/keys.json /usr/local/bin/keys.json
COPY --from=build /go/src/app/client_secret.json /usr/local/bin/client_secret.json

RUN ls /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/app"]
