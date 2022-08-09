package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/forbearing/k8s"
	"github.com/forbearing/k8s/types"
)

var (
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute*10)
	namespace   = "test"
	kubeconfig  = filepath.Join(os.Getenv("HOME"), ".kube/config")
	deployFile  = "../../testdata/examples/deployment.yaml"
	deployName  = "mydep"
	podFile     = "../../testdata/examples/pod.yaml"
	podName     = "mypod"
)

func main() {
	deployHandler, err := k8s.NewDeployment(ctx, kubeconfig, namespace)
	if err != nil {
		panic(err)
	}
	podHandler, err := k8s.NewPod(ctx, kubeconfig, namespace)
	if err != nil {
		panic(err)
	}
	cleanup(deployHandler)
	cleanup(podHandler)

	deploy, err := deployHandler.Create(deployFile)
	checkErr("create deployment from file", deploy.Name, err)
	pod, err := podHandler.Create(podFile)
	checkErr("create pod from file", pod.Name, err)
}

func checkErr(name string, val interface{}, err error) {
	if err != nil {
		log.Printf("%s failed: %v\n", name, err)
	} else {
		log.Printf("%s success: %v\n", name, val)
	}
}

// cleanup will delete or prune created deployments.
func cleanup(handler types.Deleter) {
	handler.Delete(deployName)
	handler.Delete(podName)
}