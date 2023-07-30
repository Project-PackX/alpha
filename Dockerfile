FROM golang as builder

WORKDIR /go/src/github.com/packx/backend-alpha

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

LABEL author="Dominik Szilágyi <dominik.szilagyi@gmail.com>"
LABEL maintainer="Károly Szakály <karoly.szakaly2000@gmail.com>"

WORKDIR /root/
COPY --from=builder /go/src/github.com/packx/backend-alpha/app .

CMD [ "./app" ]

EXPOSE $PORT