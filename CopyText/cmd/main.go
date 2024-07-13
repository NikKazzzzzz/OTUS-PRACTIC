package main

import (
	"flag"
	"fmt"
	copytext "github.com/NikKazzzzzz/OTUS-PRACTIC/CopyText/pkg"
	"log"
	"os"
)

func main() {
	from := flag.String("from", "", "Source file")
	to := flag.String("to", "", "Destination file")
	offset := flag.Int("offset", 0, "Offset in source file")
	limit := flag.Int("limit", 0, "Number of bytes to copy")
	flag.Parse()

	if *from == "" || *to == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := copytext.Copy(*from, *to, *limit, *offset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File copied successfully")
}
