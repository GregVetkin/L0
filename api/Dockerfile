FROM golang:1.22.0

RUN mkdir /app
RUN mkdir /app/api
RUN mkdir /app/models


ADD ./api /app/api
ADD ./models /app/models


WORKDIR /app
RUN go mod init l0
RUN go get ./models/nats
RUN go get ./models/database
RUN go get ./models/cache
RUN go get ./models/order



WORKDIR /app/api
RUN go build -o api-server ./api-server.go

CMD ["/app/api/api-server"]

