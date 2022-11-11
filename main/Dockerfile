FROM golang:alpine as builder
LABEL maintainer="Elena Khodosova"

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir /root/reports

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8000

CMD ["./main"]