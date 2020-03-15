FROM golang:alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apk --no-cache add ca-certificates g++

WORKDIR $GOPATH/src/swapi-go-db/

COPY . .

RUN go mod download

RUN go build -ldflags="-w -s" -o /go/bin/swapi-go-db

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/swapi-go-db /
COPY --from=build /go/src/swapi-go-db/sw_data.db /
COPY --from=build /usr/include /usr/include
COPY --from=build /lib /lib

ENTRYPOINT ["/swapi-go-db"]
