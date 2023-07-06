package core

import (
	"QAPI/entities"
	"QAPI/logger"
	"QAPI/models"
	"math"
	"runtime"
	"sync"
	"time"
)

func StartJob() {
	timeStart := time.Now()
	cores := runtime.NumCPU() * 2

	merchant := models.GetMerchantCheckMutation()
	var dividen float64 = 0
	if len(merchant) > 0 {
		var x float64 = float64(len(merchant))
		var y float64 = float64(cores)
		dividen = math.Floor(x / y)
	}

	/**
		RUN THREADS
	**/
	var wg sync.WaitGroup
	wg.Add(cores)

	for c := 0; c < cores; c++ {
		start := 0
		end := int(dividen)

		if c > 0 {
			start = (c * int(dividen))
			end = (c * int(dividen)) + int(dividen)
			if (c + 1) == cores {
				end = len(merchant)
			}
		}

		go func() {
			worker(merchant[start:end])
			wg.Done()
		}()
	}
	wg.Wait()

	/**
		STOP PROCESS
	**/
	timeFinish := time.Since(timeStart)
	logger.Log.Printf("Completed in %s", timeFinish)
}

func worker(merchants []entities.Merchant) {

}
