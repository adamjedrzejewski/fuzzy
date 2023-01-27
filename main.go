package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"fuzzy/lib"
)

func main() {
	var hour, minute int

	args := os.Args[1:]
	if len(args) == 0 {
		currentTime := time.Now()
		hour = currentTime.Hour()
		minute = currentTime.Minute()
	} else if len(args) == 2 {
		var err error
		hour, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid hour: %s.\n", args[0])
			os.Exit(1)
		}
		minute, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid minute: %s.\n", args[1])
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Zero or two arguments expected.")
		os.Exit(1)
	}

	time, err := fuzzy.SayTime(hour, minute)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(time)
}

