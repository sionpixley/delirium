FROM golang:1.23.6-alpine AS build

WORKDIR /delirium

COPY . .

RUN go build -o delirium ./cmd/delirium

FROM alpine:latest

COPY --from=build /delirium/delirium /usr/local/bin/delirium

CMD [ "sleep", "infinity" ]
