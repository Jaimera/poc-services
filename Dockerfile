FROM golang:1.19-alpine

WORKDIR /app

ADD . .

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o /poc-services

EXPOSE 8080

CMD [ "/poc-services" ]