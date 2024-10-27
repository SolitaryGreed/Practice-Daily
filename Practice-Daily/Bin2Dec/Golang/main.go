package main

import (
	"fmt"
	"strconv"
)

func main() {

	var bin string

	for {
		_, err := fmt.Scan(&bin)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		resBin, err := strconv.ParseInt(bin, 2, 32)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Parsed bin: %d\n", resBin)

	}
}
