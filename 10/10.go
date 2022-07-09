package task10

import (
	"fmt"
	"go_code_anywhere/utils"
	"io"
	"os"
	"time"
)

func Task10(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var line string
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
		utils.MakeTask1(dur, i)
	}
}
