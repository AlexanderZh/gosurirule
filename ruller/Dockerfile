FROM golang:1.18 AS build
RUN apt install -y ca-certificates
WORKDIR /go/src

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux 
ENV GOARCH=amd64

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./go ./go
COPY main.go .

RUN go build -a -installsuffix cgo -o gosurirule main.go

FROM scratch AS runtime
COPY --from=build /go/src/gosurirule ./
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 9090/tcp
ENTRYPOINT ["./gosurirule"]
