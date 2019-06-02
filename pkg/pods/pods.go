package pods

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/manifoldco/promptui"
	"github.com/projectjudge/ktool/pkg/utils"
)

// WatchPods watch pods
func WatchPods() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	var out bytes.Buffer
	command := "kubectl"
	args := []string{
		"--kubeconfig",
		*kubeconfig,
		"get",
		"pods",
		"--all-namespaces",
	}
	for true {
		out = bytes.Buffer{}
		cmd := exec.Command(command, args...)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		utils.CallClear()
		fmt.Printf("Every 1.0s: %s %s\n\n", command, strings.Join(args, " "))
		fmt.Printf("%s\n", out.String())
		time.Sleep(time.Second)
	}
}

// WatchPodLogs allows a user to select a pod an get its logs
func WatchPodLogs() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	podNames := []string{}
	for _, pod := range pods.Items {
		podNames = append(podNames, fmt.Sprintf("%s/%s", pod.GetNamespace(), pod.GetName()))
	}

	prompt := promptui.Select{
		Label: "Select pods to get the logs of",
		Items: podNames,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	namespace := strings.Split(result, "/")[0]
	pod := strings.Split(result, "/")[1]
	_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		// var out bytes.Buffer
		command := "kubectl"
		args := []string{
			"--kubeconfig",
			*kubeconfig,
			"-n",
			namespace,
			"logs",
			"-f",
			pod,
		}

		fmt.Printf("Reading logs for pod: %s/%s\n----------------------------------------------------------\n", namespace, pod)
		buff := make([]byte, 100)
		cmd := exec.Command(command, args...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		for err == nil {
			n, _ := stdout.Read(buff)
			if n > 0 {
				fmt.Printf("%s", string(buff[:n]))
			}
		}
	}
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func int32Ptr(i int32) *int32 { return &i }
