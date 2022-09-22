FROM  golang:alpine

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o project-2/skeleton

CMD ["./project-2/skeleton"]