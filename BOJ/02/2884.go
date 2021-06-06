package main

import "fmt"

func main() {
	var hour, minute int
	fmt.Scanf("%d %d", &hour, &minute)

	var alarmTime = 60*hour + minute - 45
	if alarmTime < 0 {
		alarmTime += 60 * 24
	}

	fmt.Printf("%d %d", alarmTime/60, alarmTime%60)
}
