package worker

import (
	"context"
	"fmt"
	"golangbase/tasks/multithreading/fileprovider"
	"log"
	"strconv"
	"time"
)

type TaskHandler struct {
	worker       Worker
	globalCancel context.CancelFunc
	fProvider    fileprovider.FileProvider
}

func NewTaskHandler(worker Worker, globalCancel context.CancelFunc, fProvider fileprovider.FileProvider) *TaskHandler {
	return &TaskHandler{
		worker:       worker,
		globalCancel: globalCancel,
		fProvider:    fProvider,
	}
}

func (h *TaskHandler) VisitForExit() {
	fmt.Printf("%d: Exit\n", h.worker.GetNum())
	h.globalCancel()
}

func (h *TaskHandler) VisitForCurrentTime() {
	fmt.Printf("%d: CurrentTime %s\n", h.worker.GetNum(), time.Now().Local().UTC())
}

func (h *TaskHandler) VisitForCurrentTimeWait() {
	fmt.Printf("%d: CurrentTime %s\n", h.worker.GetNum(), time.Now().Local().UTC())
	fmt.Printf("%d: Wait\n", h.worker.GetNum())
	time.Sleep(time.Second)
	fmt.Printf("%d: CurrentTimeWait %s\n", h.worker.GetNum(), time.Now().Local().UTC())
}

func (h *TaskHandler) VisitForReadWriteToFile() {
	wNum := h.worker.GetNum()
	fmt.Printf("%d: ReadWriteToFile\n", wNum)
	err := h.fProvider.Open()
	if err != nil {
		fmt.Printf("%d: ReadWriteToFile Error '%s'\n", h.worker.GetNum(), err.Error())
		log.Fatal(err)
	}

	defer func() {
		err := h.fProvider.Close()
		if err != nil {
			fmt.Printf("%d: File closing error '%s'\n", wNum, err.Error())
			return
		}
		fmt.Printf("%d: File closed\n", wNum)
	}()

	txt := h.fProvider.Read()
	fmt.Printf("%d: ReadWriteToFile Content '%s'\n", wNum, txt)

	err = h.fProvider.WriteString(strconv.Itoa(wNum))
	if err != nil {
		fmt.Printf("%d: ReadWriteToFile Error '%s'\n", h.worker.GetNum(), err.Error())
		return
	}
	fmt.Printf("%d: ReadWriteToFile Wrote '%d'\n", wNum, wNum)
}

func (h *TaskHandler) VisitForWorkerShutDown() {
	fmt.Printf("%d: ShutDown\n", h.worker.GetNum())
	h.worker.Cancel()
}
