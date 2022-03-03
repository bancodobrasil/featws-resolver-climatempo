FROM golang:1.17-alpine AS BUILD

WORKDIR /app

COPY go.mod /app

COPY go.sum /app

RUN go mod download

COPY . /app

RUN go build -o resolver

FROM alpine:3.15

COPY --from=BUILD /app/resolver /bin/

CMD [ "resolver" ] 



