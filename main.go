package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/metrics", http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			for _, image := range os.Args[1:] {
				response, err := http.Get(
					"https://hub.docker.com/v2/repositories/" + image,
				)
				if err != nil {
					log.Println(err)
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer response.Body.Close()
				var result struct {
					PullCount int `json:"pull_count"`
				}
				err = json.NewDecoder(response.Body).Decode(&result)
				if err != nil {
					log.Println(err)
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(
					writer,
					"dockerhub_image_pull_count{image=\"%s\"} %v\n",
					image,
					result.PullCount,
				)
			}
		},
	))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN"), nil))
}
