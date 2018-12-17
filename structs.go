package main

import (
	"fmt"
	"time"
)

//Pinged strcuture of ping response
type Pinged struct {
	Adress  string
	Average string
	Max     string
	Min     string
}

//Pinged to string
func (p Pinged) String() string {
	return fmt.Sprintf(
		"%v | %s --> Max:%s , Average:%s , Min:%s \n",
		time.Now().Unix(),
		p.Adress,
		p.Max,
		p.Average,
		p.Min,
	)
}
