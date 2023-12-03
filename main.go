package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: loggap <threshold>")
		os.Exit(1)
	}

	threshold, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("Invalid threshold: " + os.Args[1])
		os.Exit(1)
	}

	processLogs(os.Stdin, os.Stdout, threshold)
}

func processLogs(input io.Reader, output io.Writer, threshold time.Duration) {
	scanner := bufio.NewScanner(input)

	var prevTime time.Time

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, " ", 3)
		if len(parts) >= 3 {
			ts, err := time.Parse("2006-01-02 15:04:05.000000-07:00", parts[0]+" "+parts[1])
			if err == nil && !prevTime.IsZero() {
				diff := ts.Sub(prevTime)
				if diff > threshold {
					fmt.Fprintln(output, ".......... "+diff.String()+" later")
				}
			}
			prevTime = ts
		}

		fmt.Fprintln(output, line)
	}
}
