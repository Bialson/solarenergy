FROM golang:latest

RUN apt-get update && apt-get install -y

WORKDIR /usr/src/solarenergy

COPY go.mod go.sum ./

RUN go mod download && go mod verify

ENV GOPATH /go

ENV PATH $PATH:/go/bin:$GOPATH/bin

COPY . .

RUN cd solarenergy-cli go build -v -o /usr/local/bin/solarenergy/solarenergy-cli

RUN go install ./solarenergy-cli

RUN cd solarenergy-server && go build -v -o /usr/local/bin/solarenergy/solarenergy-server

RUN go install ./solarenergy-server

COPY run.sh /

WORKDIR /

ENTRYPOINT [ "/run.sh" ]