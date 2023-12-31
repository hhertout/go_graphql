FROM golang:1.20

WORKDIR /app 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go_graphql

CMD [ "/go_graphql" ]