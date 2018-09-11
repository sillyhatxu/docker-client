package docker

type Docker struct {

	Endpoint string

}

func NewDockerClient(endpoint string) *Docker {
	return &Docker{Endpoint:endpoint}
}