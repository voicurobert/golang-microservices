package main

import (
	"bufio"
	"fmt"
	"github.com/voicurobert/golang-microservices/src/api/domain/repositories"
	"github.com/voicurobert/golang-microservices/src/api/services"
	"github.com/voicurobert/golang-microservices/src/api/utils/errors"
	"os"
	"sync"
)

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/Users/robert/requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{Name: line}
		result = append(result, request)
	}
	return result
}

var (
	success map[string]string
	failed  map[string]errors.ApiError
)

func main() {
	requests := getRequests()
	fmt.Printf("about to process %d requests", len(requests))

	var wg sync.WaitGroup
	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)

	go handleResults(input, &wg)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(request, input, buffer)
	}

	wg.Wait()

	close(input)
}

func handleResults(input chan createRepoResult, wg *sync.WaitGroup) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
			continue
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(request repositories.CreateRepoRequest, output chan createRepoResult, buffer chan bool) {
	result, err := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buffer // read from the buffer chan so that we can empty it
}

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

//func main() {
//	c := make(chan string, 3)
//	go func(input chan string) {
//		fmt.Println("sending to the channel")
//		input <- "hello1"
//		input <- "hello2"
//		input <- "hello3"
//		input <- "hello4"
//	}(c)
//
//	fmt.Println("receiving from the channel")
//	for greeting := range c {
//		fmt.Println(greeting)
//	}
//	//greeting := <-c
//
//	//go helloWorld()
//	//time.Sleep(1 * time.Millisecond)
//}
//
//func helloWorld() {
//	fmt.Println("Hello World!")
//}
