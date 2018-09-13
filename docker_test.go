package docker

import (
	"testing"
	log "github.com/xushikuan/microlog"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestDockerImages(t *testing.T)  {
	client := NewDockerClient("unix:///var/run/docker.sock")
	images,err := client.Images()
	assert.Nil(t, err)
	for _,image := range images{
		log.Info("ID:",image.ID,"  RepoTags:",image.RepoTags,"  Created:",image.Created,"  Size:",image.Size,"  VirtualSize:",image.VirtualSize,"  ParentID:",image.ParentID,"  RepoDigests:",image.RepoDigests,"  Labels:",image.Labels)
	}
}

func TestDockerImageInspect(t *testing.T)  {
	client := NewDockerClient("unix:///var/run/docker.sock")
	image,err := client.ImageInspect("473ce030434e")
	assert.Nil(t, err)
	log.Info(image.ID)
	log.Info(image.RepoTags)
	log.Info(image.Parent)
	log.Info(image.Comment)
	log.Info(image.Created)
	log.Info(image.Container)
	log.Info(image.ContainerConfig)
	log.Info(image.DockerVersion)
	log.Info(image.Author)
	log.Info(image.Config)
	log.Info(image.Architecture)
	log.Info(image.Size)
	log.Info(image.VirtualSize)
	log.Info(image.RepoDigests)
	log.Info(image.RootFS)
	log.Info(image.OS)
}


func TestDockerSearchImages(t *testing.T)  {
	client := NewDockerClient("unix:///var/run/docker.sock")
	searchImageArray,err := client.SearchImages("consul")
	assert.Nil(t, err)
	for _,searchImage := range searchImageArray{
		log.Info("Name:",searchImage.Name,"  StarCount:",searchImage.StarCount,"  IsOfficial:",searchImage.IsOfficial,"  IsAutomated:",searchImage.IsAutomated,"  Description:",searchImage.Description)
	}
}


func TestDockerServiceLSFilter(t *testing.T)  {
	client := NewDockerClient("unix:///var/run/docker.sock")
	//serviceArray,err := client.ServiceLSFilter("test_consul")
	serviceArray,err := client.ServiceLSFilter("test_api-b")
	assert.Nil(t, err)
	for _,service := range serviceArray{
		log.Info("ID:",service.ID,"  Meta:",service.Meta,"  Spec:",service.Spec,"  PreviousSpec:",service.PreviousSpec,"  Endpoint:",service.Endpoint,"  UpdateStatus:",service.UpdateStatus)
	}
}

func TestDockerServiceLSFilterFuzzyMatching(t *testing.T)  {
	client := NewDockerClient("unix:///var/run/docker.sock")
	//serviceArray,err := client.ServiceLSFilter("test_consul")
	serviceArray,err := client.ServiceLSFilterFuzzyMatching("test_api")
	assert.Nil(t, err)
	for _,service := range serviceArray{
		log.Info("ID:",service.ID,"  Meta:",service.Meta,"  Spec:",service.Spec,"  PreviousSpec:",service.PreviousSpec,"  Endpoint:",service.Endpoint,"  UpdateStatus:",service.UpdateStatus)
	}
}

func TestDockerNetworkLSFilter(t *testing.T)  {
	for i:=1;i<=10;i++{
		log.Info("------------------  ",i,"  ------------------")
		client := NewDockerClient("unix:///var/run/docker.sock")
		networkArray,err := client.NetworkLSFilter("test_default")
		assert.Nil(t, err)
		for _,network := range networkArray{
			log.Info("ID:",network.ID,"  Name:",network.Name,"  Scope:",network.Scope,"  Driver:",network.Driver,"  IPAM:",network.IPAM,"  Containers:",network.Containers,"  Options:",network.Options,"  Internal:",network.Internal,"  Labels:",network.Labels,"  EnableIPv6:",network.EnableIPv6)
		}
		time.Sleep(1*time.Second)
	}
}


func TestDockerNetworkLSFilterFuzzyMatching(t *testing.T)  {
	client := NewDockerClient("unix:///var/run/docker.sock")
	networkArray,err := client.NetworkLSFilterFuzzyMatching("test")
	assert.Nil(t, err)
	for _,network := range networkArray{
		log.Info("ID:",network.ID,"  Name:",network.Name,"  Scope:",network.Scope,"  Driver:",network.Driver,"  IPAM:",network.IPAM,"  Containers:",network.Containers,"  Options:",network.Options,"  Internal:",network.Internal,"  Labels:",network.Labels,"  EnableIPv6:",network.EnableIPv6)
	}
}

