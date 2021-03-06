# Building the binary of the App
FROM golang:1.17.2-alpine AS build

ENV MYSQL_HOST=${MYSQL_HOST}

ENV MYSQL_HOST=${MYSQL_HOST} 
ENV MYSQL_PORT=${MYSQL_PORT} 
ENV MYSQL_USER=${MYSQL_USER} 
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD} 
ENV MYSQL_DBNAME=${MYSQL_DBNAME}
							
# `codefood` should be replaced with your project name
WORKDIR /go/src/codefood

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o codefood -installsuffix cgo .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /

# `codefood` should be replaced here as well
COPY --from=build /go/src/codefood .
EXPOSE 3030
CMD ["./codefood"]