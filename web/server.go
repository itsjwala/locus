package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/rakyll/statik/fs"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	_ "github.com/itsjwala/locus/web/statik"
)

type input struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/execute", execCode).Methods("POST")

	statikFS, err := fs.New()
	if err != nil {
	log.Fatal(err)
	}
	router.PathPrefix("").Handler(http.StripPrefix("/", http.FileServer(statikFS)))

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("Error while listening to the specified port. Reason:", err)
	}
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

	go containerCleanupDaemon(cli,ctx)

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

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, buf.String())
}

func containerCleanupDaemon(cli *client.Client,ctx context.Context){
	
	filter := filters.NewArgs()
	filter.Add("type", "container")
	filter.Add("event", "die")

	msg,errChan := cli.Events(ctx,types.EventsOptions{
		Filters: filter,
	})
	
	for {
		select {
		case err := <-errChan:
			panic(err)
		case message := <-msg:
			cli.ContainerRemove(ctx,message.ID, types.ContainerRemoveOptions{Force : true,})
		}
	}

}
