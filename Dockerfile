FROM golang:1.24.3

WORKDIR /tgApp

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tgbot ./cmd

EXPOSE 2266

CMD ["./tgbot"]