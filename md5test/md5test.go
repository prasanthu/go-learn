package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func checkMD5(filename string, checksum string, out chan bool) {
	fmt.Printf("Checking %s for %s\n", filename, checksum)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		out <- false
		return
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
		out <- false
		return
	}

	calculated := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Printf("%s %s calculated %s\n", filename, checksum, calculated)
	if calculated != checksum {
		fmt.Printf("Checksum didn't match for %s\n", filename)
		out <- false
	}
	out <- true
}

func checksumMap(filename string) (map[string]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	checksums := map[string]string{}
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		if len(tokens) < 2 {
			err := errors.New("Invalid input file format")
			return nil, err
		}
		checksums[tokens[1]] = tokens[0]
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return checksums, nil
}

func main() {
	workingFolder := "/Users/prasanth/tmp/nasa-logs/"
	file := "md5sum.txt"
	c, err := checksumMap(workingFolder + file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ch := make(chan bool)
	for key, value := range c {
		go checkMD5(workingFolder+key, value, ch)
	}
	for range c {
		out := <-ch
		fmt.Println(out)
	}

}
