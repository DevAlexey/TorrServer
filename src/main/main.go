package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"server"
	"server/log"
	"server/settings"
	"server/version"
)

type args struct {
	Port     string `arg:"-p" help:"web server port"`
	Path     string `arg:"-d" help:"database path"`
	Add      string `arg:"-a" help:"add torrent link and exit"`
	PlaylistDir string   `arg:"-l" help:"playlists directory path"`
	RDB      bool   `arg:"-r" help:"start in read-only DB mode"`
	DontKill bool   `arg:"-k" help:"dont kill program on signal"`
}

func (args) Version() string {
	return "TorrServer " + version.Version
}

var params args

func main() {
	arg.MustParse(&params)

	if params.Path == "" {
		params.Path, _ = os.Getwd()
	}

	if params.Port == "" {
		params.Port = "8090"
	}

	if params.Add != "" {
		add()
	}
	settings.Path = params.Path
	settings.PlaylistDir = params.PlaylistDir
	Preconfig(params.DontKill)

	server.Start(params.Port, params.RDB)
	log.TLogln(server.WaitServer())
	time.Sleep(time.Second * 3)
	os.Exit(0)
}

func add() {
	err := addRemote()
	if err != nil {
		fmt.Println("Error add torrent:", err)
		os.Exit(-1)
	}

	fmt.Println("Added ok")
	os.Exit(0)
}

func addRemote() error {
	url := "http://localhost:" + params.Port + "/torrent/add"
	fmt.Println("Add torrent link:", params.Add, "\n", url)

	json := `{"Link":"` + params.Add + `"}`
	resp, err := http.Post(url, "text/html; charset=utf-8", bytes.NewBufferString(json))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}
