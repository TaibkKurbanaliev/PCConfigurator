FROM golang:1.23.2-alpine:3.20 AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/TaibkKurbanaliev/PCConfigurator
WORKDIR /github.com/TaibkKurbanaliev/PCConfigurator

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/TaibkKurbanaliev/PCConfigurator/.bin/app .
COPY --from=0 /github.com/TaibkKurbanaliev/PCConfigurator/config/ ./config/

EXPOSE 8040

CMD [ "./app"]

