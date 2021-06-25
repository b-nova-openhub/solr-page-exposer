FROM golang:1.16.5-alpine

WORKDIR $GOPATH/src/github.com/b-nova-openhub/sopagex

COPY . .

RUN go get -d -v ./... \
    && go build -o bin/sopagex cmd/sopagex/main.go \
    && go install -v ./... \
    && chmod +x sopagex.sh

EXPOSE 8080

CMD ["sh", "sopagex.sh"]