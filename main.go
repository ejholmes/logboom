package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var d = flag.Duration("d", time.Second, "How long to log for.")
var r = flag.Int("r", 1000, "The number of lines per second to log.")

func main() {
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()+"\n")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	timeout := time.After(*d)
	ticker := time.Tick(time.Duration(int(time.Second) / *r))

	for {
		select {
		case <-timeout:
			os.Exit(0)
		case <-ticker:
			m := lines[rand.Intn(len(lines))]
			io.WriteString(os.Stdout, m)
		}
	}
}
