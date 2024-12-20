package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func doWork(s string){
	fmt.Println(s)

}

func main() {
	doWork("Kyruus tech interview")
}
