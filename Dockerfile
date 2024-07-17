FROM golang:1-22-alpine

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

EXPOSE 8000

CMD ["air"]