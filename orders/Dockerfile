FROM golang:1.21.0 as builder
WORKDIR /usr/src/app/orders

COPY . .
RUN go mod download

COPY . .

#EXPOSE 8080

CMD ["go", "run", "/usr/src/app/orders", "."]