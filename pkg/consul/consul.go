package consul

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"github.com/NpoolPlatform/go-service-framework/pkg/envconf"

	"github.com/hashicorp/consul/api"
)

type Client struct {
	*api.Client
	envConf *envconf.EnvConf
}

func NewConsulClient() (*Client, error) {
	envConf, err := envconf.NewEnvConf()
	if err != nil {
		return nil, xerrors.Errorf("fail to create environment configuration: %v", err)
	}

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", envConf.ConsulHost, envConf.ConsulPort)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, xerrors.Errorf("fail to create consul client: %v", err)
	}

	return &Client{
		Client:  client,
		envConf: envConf,
	}, nil
}

// IP is parsed from package envconf
type RegisterInput struct {
	ID          uuid.UUID
	Name        string
	Tags        []string
	Port        int
	HealthzPort int
}

func (c *Client) RegisterService(input RegisterInput) error {
	reg := api.AgentServiceRegistration{
		ID:   fmt.Sprintf("%v", input.ID),
		Name: input.Name,
		Tags: input.Tags,
		Port: input.Port,
	}

	chk := api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%v:%v/healthz", c.envConf.IP, input.Port),
		Timeout:                        "20s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "60s",
	}

	if c.envConf.ContainerID != envconf.NotRunInContainer {
		chk.DockerContainerID = c.envConf.ContainerID
	}

	reg.Check = &chk

	return c.Agent().ServiceRegister(&reg)
}

func (c *Client) DeregisterService(id uuid.UUID) error {
	return c.Agent().ServiceDeregister(fmt.Sprintf("%v", id))
}
