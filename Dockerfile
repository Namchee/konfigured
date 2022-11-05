FROM golang:1.18

WORKDIR /ci

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN  go build -a -o setel .

RUN chmod +x /ci/setel

CMD ["/ci/setel"]
