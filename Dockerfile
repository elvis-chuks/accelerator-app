# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

#ARG robin_env=prod

ENV DB_URL=""
ENV SIGNING_KEY=""

RUN go build -o ./accelerator-app ./main.go

## Deploy
FROM alpine:3.11.3

WORKDIR /

COPY --from=builder /app/accelerator-app .

EXPOSE 5001

#USER nonroot:nonroot

CMD [ "/accelerator-app" ]