# Message Queue Example
This is just example for learning purpose so, just follow the instruction ^^.
The use case I took, is a simple API to hit and consume Notification Info.

## Running the App
We have two service that work as API and Consumer, The Notification API and Notification Consumer/Worker.

### Build the API Service
At the first time you just need to build the application for the Notification API.

```bash
cd service-api/

# build the application using Docker
docker build . -t notification-api:demo
```

### Build the Consumer Service
And the scond, you need to build the application for the Notification Consumer.

```bash
cd service-api/

# build the application using Docker
docker build . -t notification-consumer:demo
```

### Run the Apps
This application, need AMQP Protocol and we're using RabbitMQ as the server. So, you need to run the RabbitMQ Server first on the hostname `rabbitmq` and the network is `rabbitnet`.

```bash
# check the network first
docker network ls

# create the network if do not exist
docker network create rabbitnet

# run the rabbitmq server
docker run -it --rm --name rabbitmq --hostname rabbitmq -p 5672:5672 -p 15672:15672 --network rabbitnet rabbitmq:3.9-management

# run the notification api
docker run --network rabbitnet -p 8080:8080 notification-api:demo

# run the notification consumer
docker run --network rabbitnet -p 5672:5672 notification-consumer:demo
```
### Sample Request
You may do request using this CURL:
```bash
curl --location --request POST 'http://:8080/api/v1/send/notification' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Title Push Notification Payment Product 4",
    "body": "This is just for learning purpose.",
    "data": {
        "sound":"notif.wav"
    },
    "receivers":["x0kU2vSEmI", "TOiUwz8GVr"]
}'
```
## Reference
RabbitMQ Documentation: https://www.rabbitmq.com/documentation.html  
Docker Documentation: https://docs.docker.com/  
Echo Golang Framework: https://echo.labstack.com/guide/
