FROM golang:1.19.0

WORKDIR /usr/src/app

COPY . .

# install air
RUN go install github.com/cosmtrek/air@latest

# check if go.mod is up to date
RUN go mod tidy

# installing vim in the container
# RUN ["apt-get", "update"]
# RUN ["apt-get", "install", "-y", "vim"]
