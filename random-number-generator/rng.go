package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func main() {
	var min, max int

	fmt.Fprintln(os.Stdout, "Please enter a minimum number:")
	_, err := fmt.Fscanf(os.Stdin, "%d\n", &min)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "Please enter a maximum number:")
	_, err = fmt.Fscanf(os.Stdin, "%d\n", &max)
	if err != nil {
		log.Fatal(err)
	}

	if min >= max {
		log.Fatal("Invalid input - shutting down...")
	}

	source := rand.NewSource(int64(min))
	r := rand.New(source)
	n := r.Intn(max)
	fmt.Println(n)
}
