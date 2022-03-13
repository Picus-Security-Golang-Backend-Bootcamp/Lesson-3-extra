package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://google.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*45))
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
