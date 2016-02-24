package reveldocker

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/revel/revel"
)

var client *docker.Client

// Init initializes the docker client.
func Init() {
	endpoint, ok := revel.Config.String("docker.endpoint")
	if !ok {
		revel.ERROR.Fatal("No docker endpoint specified.")
	}

	c, err := docker.NewClient(endpoint)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	client = c
}

// DockerController allows embedding a docker client into a revel controller.
type DockerController struct {
    *revel.Controller
	DockerClient *docker.Client
}

// Before adds the docker client to the controller.
func (c *DockerController) Before() revel.Result {
	c.DockerClient = client
	return nil
}
