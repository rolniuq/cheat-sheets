package images

const (
	ImagesName = "images"
)

type Images struct{}

func NewImages() *Images {
	return &Images{}
}

func (s *Images) GetCommands() map[string]string {
	return map[string]string{
		"docker build -t <image name>":             "build image from Dockerfile",
		"docker build -t <image name> . -no-cache": "build image from Dockerfile without cache",
		"docker images":                            "list all images",
		"docker rmi <image name>":                  "delete docker image",
		"docker image prune":                       "remove all unused images",
	}
}
