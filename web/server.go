package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

type input struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/execute", execCode).Methods("POST")
	fmt.Println("done")
	http.ListenAndServe(":8090", router)
}

func execCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var in input
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &in)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: fmt.Sprintf("itsjwala/locus_runner-%s", in.Language),
		Cmd:   []string{string(b)},
		Tty:   true,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true}
	out, err := cli.ContainerLogs(ctx, resp.ID, options)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)

	io.Copy(os.Stdout, out)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, buf.String())
}
