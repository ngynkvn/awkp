package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/kballard/go-shellquote"

	"github.com/rivo/tview"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	path    = kingpin.Arg("path", "Path to file.").Required().String()
)

const DEFAULT_TEXT = "Enter an awk expression to get a preview of the results."

var previewer = tview.NewTextView().SetText(DEFAULT_TEXT)
var app = tview.NewApplication()
var setPreview = func(text string) {
	app.QueueUpdateDraw(func() {
		previewer.SetText(text)
	})
}

func execAwk(text string) {
	if len(text) == 0 {
		setPreview(DEFAULT_TEXT)
	}
	args, err := shellquote.Split(text)
	if err != nil {
		setPreview(fmt.Sprintf("Shellquote Error: %v", err.Error()))
		return
	}
	log.Printf("Trying awk %v %v", text, *path)
	args = append(args, *path)
	output, err := exec.Command("awk", args...).Output()
	if err != nil {
		exitErr, isExitErr := err.(*exec.ExitError)
		if isExitErr {
			setPreview(string(exitErr.Stderr))
		} else {
			setPreview(err.Error())
		}
	} else {
		setPreview(string(output))
	}
}

var dots = []string{".", "..", "..."}

func start_dots(ticker *time.Ticker) {
	i := 0
	for range ticker.C {
		app.QueueUpdateDraw(func() {
			previewer.SetText(dots[i])
			i++
			i %= 3
		})
	}
}

func debounced(callback func(text string)) func(text string) {
	var handle *time.Timer
	ticker := time.NewTicker(time.Millisecond * 200)
	ticker.Stop()
	go start_dots(ticker)
	return func(text string) {
		if handle != nil {
			handle.Stop()
			ticker.Reset(time.Millisecond * 200)
		}
		handle = time.AfterFunc(time.Millisecond*500, func() {
			callback(text)
			ticker.Stop()
			handle = nil
		})
	}
}

func main() {
	kingpin.Parse()

	logOutput := tview.NewTextView()
	logOutput.Box.SetBorder(true).SetTitle("DEBUG LOG")
	log.SetOutput(logOutput)

	readline := tview.NewInputField()
	readline.SetLabel("awk: ")

	runAwk := debounced(execAwk)
	readline.SetChangedFunc(runAwk)

	appMain := tview.NewFlex().SetDirection(tview.FlexRow)
	appMain.Box.SetBorder(true).SetTitle("Awk Preview")
	appMain.
		AddItem(readline, 3, 1, true).
		AddItem(previewer, 0, 1, false).
		AddItem(logOutput, 10, 1, false)

	if err := app.SetRoot(appMain, true).SetFocus(appMain).Run(); err != nil {
		panic(err)
	}
}
