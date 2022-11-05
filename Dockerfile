FROM golang:1.19

WORKDIR /ci

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN  go build -a -o konfigured .

RUN chmod +x /ci/konfigured

CMD ["/ci/konfigured"]
