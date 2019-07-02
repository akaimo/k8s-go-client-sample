package main

import (
	"fmt"
	"log"
	"os"

	app_v1 "k8s.io/api/apps/v1"
	api_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

	pods, err := client.CoreV1().Pods("").List(meta_v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	fmt.Println()

	pods, err = client.CoreV1().Pods("").List(meta_v1.ListOptions{LabelSelector: "job-name=pi"})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	fmt.Println()

	job, err := client.BatchV1().Jobs(api_v1.NamespaceDefault).Get("pi", meta_v1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(job.Name)
	fmt.Println()

	f, err := os.Open("sample.yaml")
	if err != nil {
		log.Fatal(err)
	}
	deployment := &app_v1.Deployment{}
	err = yaml.NewYAMLOrJSONDecoder(f, 4096).Decode(deployment)
	if err != nil {
		log.Fatal(err)
	}
	deploymentsClient := client.AppsV1().Deployments(api_v1.NamespaceDefault)
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Name)
	fmt.Println(result.Spec.Template.Spec.Containers[0].Name)
}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}
