FROM golang:1.14.9-alpine
ADD app/ /app/
WORKDIR /app
RUN go build
EXPOSE 80
CMD ["./app"]
