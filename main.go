package main

import (
	"fmt"
	"os"
	"time"
)

// main function of all
func main() {
	os.Mkdir("tests", os.ModePerm)
	target := os.Args[1]
	checkpoints := getRoutePoints(target)
	folder := fmt.Sprintf("tests/%v_%s", time.Now().Unix(), target)
	os.Mkdir(folder, os.ModePerm)
	for {
		for _, point := range checkpoints {
			c := make(chan Pinged)
			go pingAdress(c, point, 1)
			res := <-c
			file := fmt.Sprintf("%s/%s.log", folder, res.Adress)
			fmt.Print(res.String())
			writeFile(file, res.String())
		}
		fmt.Println("<------------------------------------->")
		time.Sleep(10 + time.Second)
	}
}
