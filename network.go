package docker


import (
	dockerClient "github.com/fsouza/go-dockerclient"
)

//docker network ls --filter "name=test_default"
func (docker Docker) NetworkLSFilterFuzzyMatching(networkName string) ([]dockerClient.Network, error) {
	client, err := dockerClient.NewClient(docker.Endpoint)
	if err != nil {
		return nil,err
	}
	opts := dockerClient.NetworkFilterOpts{
		"name": map[string]bool{networkName: true},
	}
	networkArray, err := client.FilteredListNetworks(opts)
	if err != nil {
		return nil,err
	}
	return networkArray,nil
}

func (docker Docker) NetworkLSFilter(networkName string) ([]dockerClient.Network, error) {
	networkArray, err := docker.NetworkLSFilterFuzzyMatching(networkName)
	if err != nil {
		return nil,err
	}
	var resultArray []dockerClient.Network
	for _,network := range networkArray{
		if network.Name == networkName{
			resultArray = append(resultArray,network)
		}
	}
	return resultArray,nil
}

func (docker Docker) NetworkName(networkName string) ([]dockerClient.Network, error) {
	client, err := dockerClient.NewClient(docker.Endpoint)
	if err != nil {
		return nil,err
	}
	opts := dockerClient.NetworkFilterOpts{
		"name": map[string]bool{networkName: true},
	}
	networkArray, err := client.FilteredListNetworks(opts)
	if err != nil {
		return nil,err
	}
	return networkArray,nil
}