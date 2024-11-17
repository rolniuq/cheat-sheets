package containers

const (
	ContainersName = "containers"
)

type Containers struct{}

func NewContainers() *Containers {
	return &Containers{}
}

func (s *Containers) GetCommands() map[string]string {
	return map[string]string{
		"docker run --name <container name> <image name>":         "create and run container from image",
		"docker run -p <host port>:<container port> <image name>": "run container with publish container's port to the host",
		"docker run -d <image name>":                              "run container in background",
		"docker start|stop <container name> (or <container id>)":  "start or stop existing container",
		"docker rm <container name>":                              "remove a stopped container",
		"docker exec -it <container name> sh":                     "open shell inside running container",
		"docker logs -f <container name>":                         "fetch and follow logs from container",
		"docker inspect <container name> (or <container id>)":     "inspect running container",
		"docker ps":              "list currently running containers",
		"docker ps --al":         "list all docker containers",
		"docker container stats": "view resource usage stats",
	}
}
