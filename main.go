package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/libcompose/config"
	"github.com/docker/libcompose/lookup"
	"github.com/docker/libcompose/project"
)

func main() {
	var composeFile string

	if len(os.Args) == 2 {
		composeFile = os.Args[1]
	} else {
		composeFile = "docker-compose.yml"
	}

	ctx := project.Context{
		ComposeFiles: []string{composeFile},
		ProjectName:  "libcompose-sample",
	}

	cwd, _ := os.Getwd()

	ctx.ResourceLookup = &lookup.FileConfigLookup{}
	ctx.EnvironmentLookup = &lookup.ComposableEnvLookup{
		Lookups: []config.EnvironmentLookup{
			&lookup.EnvfileLookup{
				Path: filepath.Join(cwd, ".env"),
			},
			&lookup.OsEnvLookup{},
		},
	}

	prj := project.NewProject(&ctx, nil, nil)

	if err := prj.Parse(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, key := range prj.ServiceConfigs.Keys() {
		if svc, ok := prj.ServiceConfigs.Get(key); ok {
			fmt.Println("=== " + key)

			for _, env := range svc.Environment {
				fmt.Println(env)
			}
		}
	}
}
