package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	err := unix.Mkfifo("dingle", 0600)
	if err != nil && err.Error() != "file exists" {
		log.Fatal(err)
	}

	for {
		fi, err := os.Open("dingle")
		if err != nil {
			log.Fatal(err)
		}
		in := bufio.NewReader(fi)

	readloop:
		for {
			line, err := in.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}

				log.Println("reopening")
				break readloop
			}
			log.Printf("you said: %s", line)
		}
	}
}
