package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ParseTimeToInt(time string) int {
	var res = 0
	var buff = ""
	for ind := len(time) - 1; ind >= 0; ind = ind - 1 {
		if '0' <= time[ind] && time[ind] <= '9' {
			buff = string(time[ind]) + buff
		} else if buff != "" {
			i, err := strconv.Atoi(buff)
			if err != nil {
				return 0
			}
			if time[ind] == 'm' {
				res = res + i
				buff = ""
			}
			if time[ind] == 'h' {
				res = res + i*60
				buff = ""
			}
		}
	}
	i, err := strconv.Atoi(buff)
	if err != nil {
		return res
	}
	if buff != "" {
		res = res + i*60*60
	}
	return res
}

func MakeTask1(duration time.Duration, taskNum int) {
	num := strconv.Itoa(taskNum)
	fmt.Println("task " + num + " is started")
	time.Sleep(duration * time.Second)
	fmt.Println("task " + num + " is ended")
}

func MakeTask2(duration time.Duration, taskNum int, done chan bool) {
	MakeTask1(duration, taskNum)
	done <- true
}
