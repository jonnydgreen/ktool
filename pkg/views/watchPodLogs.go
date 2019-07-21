package views

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/projectjudge/ktool/pkg/config"
	"github.com/projectjudge/ktool/pkg/ktoolgrid"
	"github.com/projectjudge/ktool/pkg/pods"
)

type selectedPod struct {
	Name      string
	Namespace string
}

func podLogs(logs *widgets.List, grid *ui.Grid, kubeconfig string, pod string, namespace string) {
	// Get initial logs
	initialRead := len(pods.GetPodLogs(kubeconfig, pod, namespace))

	// Get rolling logs
	command := "kubectl"
	args := []string{
		"--kubeconfig",
		kubeconfig,
		"-n",
		namespace,
		"logs",
		"-f",
		pod,
	}

	buff := make([]byte, 100)
	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	totalRead := 0
	runBefore := false
	logs.Rows = []string{}
	rollingLogs := ""
	for err == nil {
		n, _ := stdout.Read(buff)
		totalRead += n
		if n > 0 {
			rollingLogs += string(buff[:n])
			logs.Rows = append(strings.Split(rollingLogs, "\n"))
			if initialRead <= totalRead {
				if !runBefore {
					fmt.Println("NOT RUN BEFORE, SCROLLING TO BOTTOM")
					fmt.Printf("                                     INITIAL: %d | TOTAL: %d", initialRead, totalRead)
					logs.ScrollBottom()
					runBefore = true
				}
				ui.Render(grid)
			}
		}
		// fmt.Println("AFTER " + strconv.Itoa(n))
	}
}

// WatchPodLogs view renders the pods logs
func WatchPodLogs() {
	ktoolgrid.InitGrid()
	defer ui.Close()

	kubeconfig, _ := config.KubeConfig()

	allPods := strings.Split(pods.GetPods(kubeconfig), "\n")
	allPods = allPods[1:]

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	selectablePods := widgets.NewList()
	selectablePods.Title = "Context | Select pod logs to follow"
	selectablePods.Rows = allPods
	selectablePods.TextStyle = ui.NewStyle(ui.ColorYellow)
	selectablePods.WrapText = false

	logs := widgets.NewList()
	logs.Title = fmt.Sprintf(" <Namespace> | <Pod> | Logs ")
	logs.Rows = []string{" No logs to show"}
	logs.WrapText = true

	grid.Set(
		ui.NewRow(1.0/6, selectablePods),
		ui.NewRow(5.0/6, logs),
	)

	ui.Render(grid)

	previousKey := ""
	selectedPod := selectedPod{
		Name:      "<Pod>",
		Namespace: "<Namespace>",
	}
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				selectablePods.ScrollDown()
			case "k", "<Up>":
				selectablePods.ScrollUp()
			case "<C-d>":
				logs.ScrollHalfPageDown()
			case "<C-u>":
				logs.ScrollHalfPageUp()
			case "<C-f>":
				logs.ScrollPageDown()
			case "<C-b>":
				logs.ScrollPageUp()
			case "g":
				if previousKey == "g" {
					logs.ScrollTop()
				}
			case "H", "<Home>":
				logs.ScrollTop()
			case "G", "<End>":
				logs.ScrollBottom()
			case "<Enter>":
				space := regexp.MustCompile(`\s+`)
				podString := selectablePods.Rows[selectablePods.SelectedRow]
				podStringAttributes := strings.Split(space.ReplaceAllString(podString, " "), " ")
				if selectedPod.Name != podStringAttributes[1] {
					selectedPod.Name = podStringAttributes[1]
					selectedPod.Namespace = podStringAttributes[0]
					logs.Title = fmt.Sprintf(" %s | %s | Logs ", selectedPod.Namespace, selectedPod.Name)

					// Kick off our go routine
					go podLogs(logs, grid, kubeconfig, selectedPod.Name, selectedPod.Namespace)
				}
			case "<MouseWheelUp>":
				if logs.SelectedRow != 0 {
					logs.ScrollUp()
				}
			case "<MouseWheelDown>":
				if logs.SelectedRow < (len(logs.Rows) - 1) {
					logs.ScrollDown()
				}
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}

			ui.Render(grid)
		case <-ticker:
			// if selectedPod.Name != "<Pod>" {
			// 	logs.Rows = strings.Split(pods.GetPodLogs(kubeconfig, selectedPod.Name, selectedPod.Namespace), "\n")
			// }
			// ui.Render(grid)
			// Get pods
		}
	}
}
