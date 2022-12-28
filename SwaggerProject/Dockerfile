# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
WORKDIR /server
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o /SwaggerProject
EXPOSE 8080
CMD [ "/SwaggerProject" ]
