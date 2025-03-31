package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var d = flag.Duration("d", time.Duration(0), "How long to log for.")
var r = flag.Int("r", 1000, "The number of lines per second to log.")
var m = flag.String("m", "", "The message to log.")

func main() {
	flag.Parse()

	var rd io.Reader
	if *m != "" {
		rd = strings.NewReader(*m)
	} else {
		rd = os.Stdin
	}

	var lines []string
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()+"\n")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var timeout <-chan time.Time
	if *d != time.Duration(0) {
		timeout = time.After(*d)
	} else {
		// This will block forever.
		timeout = make(<-chan time.Time)
	}

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
