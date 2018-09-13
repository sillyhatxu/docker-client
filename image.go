package docker


import (
	dockerClient "github.com/fsouza/go-dockerclient"
	"github.com/docker/docker/api/types/swarm"
)

func (docker Docker) Images() ([]dockerClient.APIImages, error) {
	client, err := dockerClient.NewClient(docker.Endpoint)
	if err != nil {
		return nil,err
	}
	imageArray, err := client.ListImages(dockerClient.ListImagesOptions{All: false})
	if err != nil {
		return nil,err
	}
	return imageArray,nil
}

func (docker Docker) ImageInspect(imageId string) (*dockerClient.Image, error) {
	client, err := dockerClient.NewClient(docker.Endpoint)
	if err != nil {
		return nil,err
	}
	image, err := client.InspectImage(imageId)
	if err != nil {
		return nil,err
	}
	return image,nil
}

//Delete
func (docker Docker) SearchImages(serviceName string) ([]dockerClient.APIImageSearch, error) {
	client, err := dockerClient.NewClient(docker.Endpoint)
	if err != nil {
		return nil,err
	}
	imageSearchArray, err := client.SearchImages(serviceName)
	if err != nil {
		return nil,err
	}
	return imageSearchArray,nil
}


func (docker Docker) ServiceLSFilter(serviceName string) ([]swarm.Service, error) {
	serviceArray, err := docker.ServiceLSFilterFuzzyMatching(serviceName)
	if err != nil {
		return nil,err
	}
	var resultArray []swarm.Service
	for _,service := range serviceArray{
		if service.Spec.Name == serviceName{
			resultArray = append(resultArray,service)
		}
	}
	return resultArray,nil
}

//docker service ls --filter "name=test_consul"
func (docker Docker) ServiceLSFilterFuzzyMatching(serviceName string) ([]swarm.Service, error) {
	client, err := dockerClient.NewClient(docker.Endpoint)
	if err != nil {
		return nil,err
	}
	//check name
	serviceArray, err := client.ListServices(dockerClient.ListServicesOptions{Filters:map[string][]string{"name":{serviceName}}})
	if err != nil {
		return nil,err
	}
	return serviceArray,nil
}