package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	language := "python"
	input := `{"language":"python", "code":"print(\"helloworld\")"}`
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: fmt.Sprintf("itsjwala/locus_runner-%s", language),
		Cmd:   []string{input},
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, resp.ID, options)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
