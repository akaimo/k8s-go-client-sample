package main

import (
	"fmt"
	"log"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	job, err := client.BatchV1().Jobs("default").Get("pi", meta_v1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(job.Name)
}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}
