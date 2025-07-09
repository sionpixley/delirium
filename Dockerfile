FROM golang:1.24.5-alpine AS build

WORKDIR /delirium

COPY . .

RUN go build -o delirium ./cmd/delirium

FROM alpine:latest

COPY --from=build /delirium/delirium /usr/local/bin/delirium

CMD [ "sleep", "infinity" ]
