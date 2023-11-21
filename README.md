# https://medium.com/@elmashad/deploying-go-applications-using-docker-on-aws-e6e76e609b49
## aws lightsail push-container-image --region ca-central-1  --service-name aws-container-service-t1 --label node-app-t8 --image node-t8

# try exposing 3001 on https
# create a postgres service in a separate container and connect the echo service via gorm