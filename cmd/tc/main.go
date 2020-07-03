package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const usage = `Usage:
  tc <unix timestamp>
  Ex) tc 1593766111
      2020-07-03T17:38:23+09:00

  Ex) TZ=America/New_York tc 1593766111
      2020-07-03T09:38:23+01:00`


func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println(usage)
	case 2:
		tm := getTime(os.Args[1])
		print(tm, defaultLayout)
	default:
		fmt.Println(usage)
	}
}

func getTime(s string) time.Time {
	parsed, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return time.Unix(parsed, 0)
}

var defaultLayout = time.RFC3339

func print(tm time.Time, layout string) {
	fmt.Println(tm.Format(layout))
}