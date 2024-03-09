FROM golang:1.22.0

RUN mkdir /app
RUN mkdir /app/models
RUN mkdir /app/worker



ADD ./worker /app/worker
ADD ./models /app/models


WORKDIR /app
RUN go mod init l0
RUN go get ./models/nats
RUN go get ./models/database
RUN go get ./models/cache
RUN go get ./models/order

WORKDIR /app/worker
RUN go build -o worker ./worker.go

CMD ["/app/worker/worker"]

