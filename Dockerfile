FROM scratch
MAINTAINER Johannes Scheuermann<johannes.scheuermann@inovex.de>
ADD bin/todo-app-go todo-app-go
ENTRYPOINT ["/todo-app-go"]
