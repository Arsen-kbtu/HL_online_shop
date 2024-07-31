FROM golang:1.21.0 as builder
WORKDIR /usr/src/app/api-gateway
COPY . .
RUN go mod download

#EXPOSE 8080

CMD ["go", "run", "/usr/src/app/api-gateway", "."]