package main

import (
	"fmt"
	"golangbase/tasks/statistics"
)

func main() {
	st := statistics.NewFileStatistician("input.txt")
	stat, err := st.GetStatistics()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println(stat)
}
