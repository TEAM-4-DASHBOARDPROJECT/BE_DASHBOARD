FROM  golang:1.18

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o BE_DASHBOARD

CMD ["./BE_DASHBOARD"]