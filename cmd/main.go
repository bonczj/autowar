package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bonczj/autowar/internal/war"
	"github.com/jszwec/csvutil"
)

func main() {
	var wars = flag.Int("c", 1000, "number of wars to run")
	var routines = flag.Int("r", 10, "number of routines to run")
	flag.Parse()
	log.Printf("Running %d wars across %d go routines", *wars, *routines)
	ch := make(chan bool, *wars)

	for i := 0; i < *wars; i++ {
		ch <- true
	}

	close(ch)
	var wg sync.WaitGroup
	results := make(chan war.Result, *wars)

	for i := 0; i < *routines; i++ {
		wg.Add(1)
		go run(&wg, ch, results)
	}

	wg.Wait()
	close(results)

	// get the output file ready
	outputFile, err := os.OpenFile("results.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	output := make([]war.Result, 0)

	for result := range results {
		output = append(output, result)
	}

	b, err := csvutil.Marshal(output)
	if err != nil {
		fmt.Println("error:", err)
	}
	outputFile.Write(b)
}

func run(wg *sync.WaitGroup, input chan bool, output chan war.Result) {
	defer wg.Done()

	for range input {
		w, err := war.NewWar()
		if err != nil {
			log.Fatalln(err)
		}

		switch w.Play() {
		case war.WinnerPlayerOne:
			output <- war.Result{
				Win:    war.WinnerPlayerOne,
				Rounds: w.Rounds,
			}
		case war.WinnerPlayerTwo:
			output <- war.Result{
				Win:    war.WinnerPlayerTwo,
				Rounds: w.Rounds,
			}
		case war.WinnerTie:
			output <- war.Result{
				Win:    war.WinnerTie,
				Rounds: w.Rounds,
			}
		case war.Continue:
			log.Fatalln("game returned an unexpected 'continue' as the winner")
		case war.Error:
			log.Fatalln("game returned an error as the winner")
		case war.Loop:
			output <- war.Result{
				Win:    war.Loop,
				Rounds: w.Rounds,
			}
		}
	}
}
