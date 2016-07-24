package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/context"
)

const (
	key = "ctxUrlKey"
)

func main() {
	fmt.Println("START request to backs...")
	hosts := []string{"back1", "back2", "back3"}
	ctx := context.Background()
	sendRequests(ctx, hosts)
	fmt.Println("END request to backs...")
}
func sendRequests(ctx context.Context, hosts []string) error {
	child, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	errCh := make(chan error, 1)
	doneCh := make(chan string, 1)
	go func() {
		for {
			select {
			case err := <-errCh:
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					cancel()
				}
			case msg := <-doneCh:
				fmt.Println(msg)
			}
			wg.Done()
		}
	}()
	for _, host := range hosts {
		wg.Add(1)
		go func(host string) {
			msg, err := req(context.WithValue(child, key, fmt.Sprintf("http://%s/", host)))
			if err != nil {
				errCh <- err
				return
			}
			doneCh <- msg
		}(host)
	}
	wg.Wait()
	return nil
}

func req(ctx context.Context) (string, error) {
	url := ctx.Value(key).(string)
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	errCh := make(chan error, 1)
	doneCh := make(chan string, 1)
	go func() {
		fmt.Printf("start req to %s\n", url)
		res, err := client.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errCh <- err
			return
		}
		fmt.Printf("end req to %s\n", url)
		doneCh <- string(b)
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		return "", ctx.Err()
	case err := <-errCh:
		return "", err
	case msg := <-doneCh:
		return msg, nil
	}
}
