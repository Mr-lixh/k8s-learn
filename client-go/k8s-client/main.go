package main

import (
	"flag"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// GetClientFromDir 通过读取本地 kubeconfig 文件生成 k8s 客户端
// 如果默认路径不存在，则使用启动命令传入 kubeconfig 的绝对路径
func GetClientFromDir() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 创建 *rest.Config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	// 创建 *Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset

	// 创建单个资源的 client
	//deploymentClient := clientset.AppsV1().Deployments("")
}

// GetDynamicClient 读取本地 kubeconfig 文件创建一个动态客户端
func GetDynamicClient() dynamic.Interface {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 创建 *rest.Config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// 创建动态客户端
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(client)
	}
	return client
}

// GetFakeClient 返回一个 fake 客户端
// 常用于测试用例
func GetFakeClient() *fake.Clientset {
	return fake.NewSimpleClientset()
}
