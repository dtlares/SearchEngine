FROM golang:latest

# Module files
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download

RUN go build /app/main.go
EXPOSE 9091
ENTRYPOINT [ "/app/main" ]
