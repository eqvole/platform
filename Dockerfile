FROM golang
WORKDIR /go/src/github.com/eqvole/platform
COPY . ./
RUN go mod tidy
EXPOSE 8080
CMD ["go","run","cmd/http/main.go"]
