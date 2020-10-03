package main

import (
	"fmt"
	"os"
	"time"

	. "./src/localfile"
	. "./src/routeping"
	. "./src/structs"
)

// main function of all
func main() {
	os.Mkdir("tests", os.ModePerm)
	target := os.Args[1]
	for {
		checkpoints := GetRoutePoints(target)
		folder := fmt.Sprintf("tests/%v_%s", time.Now().Unix(), target)
		os.Mkdir(folder, os.ModePerm)
		for i := 0; i < 1000; i++ {
			for _, point := range checkpoints {
				c := make(chan Pinged)
				go PingAddress(c, point, 1)
				res := <-c
				file := fmt.Sprintf("%s/%s.log", folder, res.Address)
				fmt.Print(res.String())
				WriteFile(file, res.String())
			}
			fmt.Println("<------------------------------------->")
			time.Sleep(10 + time.Second)
		}
	}
}
