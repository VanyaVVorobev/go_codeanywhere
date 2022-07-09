package task30

import (
	"fmt"
	"go_code_anywhere/utils"
	"io"
	"os"
	"time"
)

func Task30(fileName string) {
	fmt.Print("Enter n of Threads: ")
	var n int
	fmt.Scanf("%d", &n)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var line string
	var tasksProgress []chan bool
	for i := 0; ; i = i + 1 {
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		if len(tasksProgress) >= n {
			for index, done := range tasksProgress {
				if <-done {
					tasksProgress = append(tasksProgress[:index], tasksProgress[index+1:]...)
					dur := time.Duration(utils.ParseTimeToInt(line))
					newDone := make(chan bool, 1)
					tasksProgress = append(tasksProgress, newDone)
					go utils.MakeTask2(dur, i, newDone)
				}
			}
		} else {
			dur := time.Duration(utils.ParseTimeToInt(line))
			newDone := make(chan bool, 1)
			tasksProgress = append(tasksProgress, newDone)
			go utils.MakeTask2(dur, i, newDone)
		}
	}

	for _, done := range tasksProgress {
		<-done
	}
}
