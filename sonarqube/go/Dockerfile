FROM golang:1.19

WORKDIR /app

RUN go mod init teste

COPY ./sonarqube/go .

RUN go build -o sum

CMD ["./sum"]

