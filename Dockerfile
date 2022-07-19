# build chat-room image and expose port

FROM golang:alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

ENV PORT=3000

EXPOSE 3000

RUN go build

CMD ["./Chat-Room"]