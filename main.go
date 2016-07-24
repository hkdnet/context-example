package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("START request to backs...")
	hosts := []string{"back1", "back2", "back3"}
	for _, host := range hosts {
		req(fmt.Sprintf("http://%s/", host))
	}
	fmt.Println("END request to backs...")
}

func req(url string) error {
	fmt.Printf("start req to %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
	fmt.Printf("\nend req to %s\n", url)

	return nil
}
