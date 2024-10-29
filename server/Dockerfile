FROM golang:alpine

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

EXPOSE 4221

CMD [ "/server" ]