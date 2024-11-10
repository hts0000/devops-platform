package mysqltesting

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	image = ""
	port  = "23306/tcp"
)

func RunMySQLInDocker(m *testing.M) int {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx := context.Background()
	cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		// 容器
		ExposedPorts: nat.PortSet{
			port: {},
		},
	}, &container.HostConfig{}, &network.NetworkingConfig{}, &v1.Platform{}, "devops-platform-mysql-test")

	return m.Run()
}
