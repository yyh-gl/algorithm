FROM golang:latest AS builder

WORKDIR /go/src/github.com/yyh-gl/Algorithm/LeetCode

ENV TZ="Asia/Tokyo"

COPY . .

RUN go install github.com/budougumi0617/leetgode/cmd/leetgode@latest
