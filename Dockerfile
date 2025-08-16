FROM golang:1.25.0-alpine AS build

WORKDIR /delirium

COPY . .

RUN go build -o delirium ./cmd/delirium

FROM alpine:latest

COPY --from=build /delirium/delirium /usr/local/bin/delirium

CMD [ "sleep", "infinity" ]
