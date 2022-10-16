FROM golang:1.19-alpine as build

RUN apk add --update make git build-base
RUN apk --no-cache add ca-certificates

COPY . /bot/
WORKDIR /bot/

RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /bot/
COPY --from=build /bot/bot .
COPY --from=build /bot/bot.yaml .

CMD ["./bot"]
