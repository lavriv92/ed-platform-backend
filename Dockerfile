FROM golang:1.7

# create container directory
RUN mkdir /go/src/app

# make container directory workdir
WORKDIR /go/src/app

# copy source code to container
COPY . /go/src/app

# install packages
RUN go get

# build application
RUN go build ./main.go

CMD ["./main"]
