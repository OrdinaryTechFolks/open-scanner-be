FROM ghcr.io/hybridgroup/opencv:4.9.0

RUN mkdir /app
WORKDIR /app

COPY . .

RUN make setup && make build