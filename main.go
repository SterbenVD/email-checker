package main

import (
	"bufio"
	"fmt"
	"log"
	_ "net"
	"os"
	_ "strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("domain ...")

	for sc.Scan() {
		checkDomain(sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal("Error in taking input: %v\n", err)
	}
}

func checkDomain(domain string) {

}
