package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	jobName := flag.String("jobname", "test-job", "The name of job")
	containerImage := flag.String("image", "ubuntu:latest", "The name of the container image")
	entryCommand := flag.String("command", "ls", "The command to run inside the container")

	flag.Parse()

	fmt.Printf("Args: %s %s %s\n", *jobName, *containerImage, *entryCommand)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)

	// clientset := connectToK8s()
	// launchK8sJob(clientset, jobName, containerImage, entryCommand)
}
