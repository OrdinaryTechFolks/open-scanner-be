FROM ghcr.io/hybridgroup/opencv:4.9.0

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum Makefile ./
RUN go mod download && make setup

COPY . .

RUN make build