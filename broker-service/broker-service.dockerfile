# FROM golang:1.24.5-alpine as builder

# RUN mkdir app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=O go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp

# # build tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]
