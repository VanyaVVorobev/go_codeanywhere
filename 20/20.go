package task20

import (
	"fmt"
	"go_code_anywhere/utils"
	"io"
	"os"
	"time"
)

func Task20(fileName string) {
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

		dur := time.Duration(utils.ParseTimeToInt(line))
		done := make(chan bool, 1)
		tasksProgress = append(tasksProgress, done)
		go utils.MakeTask2(dur, i, done)
	}

	for _, done := range tasksProgress {
		<-done
	}
}
