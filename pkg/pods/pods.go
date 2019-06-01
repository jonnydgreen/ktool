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
	"runtime"
	"strings"
	"time"

	"k8s.io/client-go/util/homedir"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CallClear clears the termninal screen
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// WatchPods watch pods
func WatchPods() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// os
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
		CallClear()
		fmt.Printf("Every 1.0s: %s %s\n\n", command, strings.Join(args, " "))
		fmt.Printf("%s\n", out.String())
		time.Sleep(time.Second)
	}

	// tests := map[string]struct {
	// 	obj    runtime.Object
	// 	expect string
	// }{
	// 	"singleObject": {
	// 		&apiv1.Pod{
	// 			TypeMeta: metav1.TypeMeta{
	// 				Kind: "Pod",
	// 			},
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Name: "foo",
	// 			},
	// 		},
	// 		"pod/foo\n"},
	// }

	// docker        compose-7cf768cb84-zr5dn                 1/1     Running   0          53d
	// fmt.Printf("NAMESPACE     NAME                                     READY   STATUS    RESTARTS   AGE\n")
	// // printer := printers.GetNewTabWriter(os.Stdout)
	// for _, p := range list.Items {
	// 	podName := p.GetName()
	// 	// get.TablePrinter()
	// 	// pod, err := podsClient.Get(podName, metav1.GetOptions{})
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }
	// 	print := printers.ResourcePrinterFunc(pod, os.Stdout)
	// 	err = print.PrintObj(pod, os.Stdout)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// if err != nil {
	// 	panic(err)
	// }
	// tw := utilprinters.GetNewTabWriter(os.Stdout)
	// podString := fmt.Sprintf("%s\t%s\t%s\n", p.GetNamespace(), p.GetName(), p.GetCreationTimestamp().String())
	// tw.Write([]byte(podString))

	// pod := p.DeepCopyObject()
	// print := printers.NewHumanReadablePrinter(nil, printers.PrintOptions{})
	// err := print.PrintObj(pod, os.Stdout)
	// if err != nil {
	// 	panic(err)
	// }
	// podString := fmt.Sprintf("%s\t%s\t%s\n", p.GetNamespace(), p.GetName(), p.GetCreationTimestamp().String())
	// err := ps.ResourcePrinter.PrintObj(p.DeepCopyObject(), os.Stdout)
	// if err != nil {
	// 	panic(err)
	// }
	// printer.Write([]byte(podString))
	// printers.ResourcePrinter.PrintObj(p.DeepCopyObject(), os.Stdout)
	// fmt.Printf("%s\t%s\t%s\n", p.GetNamespace(), p.GetName(), p.GetCreationTimestamp().String())
	// }

	// // Delete Deployment
	// prompt()
	// fmt.Println("Deleting deployment...")
	// deletePolicy := metav1.DeletePropagationForeground
	// if err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{
	// 	PropagationPolicy: &deletePolicy,
	// }); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Deleted deployment.")
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
