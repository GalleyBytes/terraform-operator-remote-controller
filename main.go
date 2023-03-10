package main

import (
	"fmt"
	"log"
	"os"

	"github.com/galleybytes/monitor/projects/terraform-operator-remote-controller/pkg/tfhandler"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func kubernetesConfig(kubeconfigPath string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal("Failed to get config for clientset")
	}
	return config
}

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	proto := os.Getenv("TFO_API_PROTOCOL")
	host := os.Getenv("TFO_API_HOST")
	port := os.Getenv("TFO_API_PORT")
	user := os.Getenv("TFO_API_LOGIN_USER")
	password := os.Getenv("TFO_API_LOGIN_PASSWORD")
	url := fmt.Sprintf("%s://%s:%s", proto, host, port)
	config := kubernetesConfig(kubeconfig)
	// client := kubernetes.NewForConfigOrDie(config)
	dynamicClient := dynamic.NewForConfigOrDie(config)
	tfinformer := tfhandler.NewInformer(dynamicClient, url, user, password)

	tfinformer.Run()

	os.Exit(1) // should this be 0 instead?
}
