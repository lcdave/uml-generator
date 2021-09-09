FROM golang:1.16-alpine as builder

# create a directory inside the image
WORKDIR /app

# We’ll copy the go.mod and go.sum file into our working directory /app which, 
# owing to our use of WORKDIR, is the current directory (.) inside the image.
COPY go.mod go.sum ./

# execute the command go mod download (install go modules)
RUN go mod download

COPY models/ ./models/
COPY backend/ ./backend/
COPY main.go main.go

# compile our application
RUN go build -o /upsee

# expose ports
EXPOSE 8085 8085

# tell Docker what command to execute when our image is used to start a container
CMD [ "/upsee" ]