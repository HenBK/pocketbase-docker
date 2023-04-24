FROM golang:1.20.3-alpine3.17 
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build main.go
RUN chmod +x /app/main 
EXPOSE 8090
CMD [ "/app/main", "serve", "--http", "0.0.0.0:8090" ]