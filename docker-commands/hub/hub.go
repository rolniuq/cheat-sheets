package hub

const (
	HubName = "hub"
)

type Hub struct{}

func NewHub() *Hub {
	return &Hub{}
}

func (s *Hub) GetCommands() map[string]string {
	return map[string]string{
		"docker login -u  <username>":         "login to docker",
		"docker push <username>/<image name>": "publish image to docker hub",
		"docker search <image name>":          "search hub for an image",
		"docker pull <image name>":            "pull image from docker hub",
	}
}
