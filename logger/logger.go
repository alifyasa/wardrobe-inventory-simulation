package logger

import (
	"fmt"
	"log"

	"github.com/alifyasa/wardrobe-inventory-simulation/simulation"
)

var logger = log.Default()

func init() {
	logger.SetFlags(0)
}

func Log(format string, v ...any) {
	printStr := fmt.Sprintf(format, v...)
	logger.Printf("[%s] %s", simulation.SimulationState.ReadableTime(), printStr)
}
