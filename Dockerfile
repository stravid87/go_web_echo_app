FROM golang:1.21.4-alpine3.17

RUN mkdir /build

COPY . /build

WORKDIR /build

RUN go build

EXPOSE 3001

CMD ["./go-web-echo-app"]

# https://medium.com/@elmashad/deploying-go-applications-using-docker-on-aws-e6e76e609b49
## aws lightsail push-container-image --region ca-central-1  --service-name aws-container-service-t1 --label node-app-t8 --image node-t8

# try exposing 3001 on https
# create a postgres service in a separate container and connect the echo service via gorm