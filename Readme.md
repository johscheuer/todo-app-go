# Build

```Bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo-app-go .
sudo docker build -t johscheuer/todo-app-go .
```

# Todo App

A simple app which creates todo entries inside a redis database which is in master-slave mode and runs on [Mesos](https://mesos.apache.org) over [Marathon](https://mesosphere.github.io)
