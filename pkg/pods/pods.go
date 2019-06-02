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

	"github.com/projectjudge/ktool/pkg/utils"

	"k8s.io/client-go/util/homedir"
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
