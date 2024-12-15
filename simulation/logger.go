package simulation

import (
	"fmt"
	"log"
)

var logger = log.Default()

func init() {
	logger.SetFlags(0)
}

func Log(format string, v ...any) {
	printStr := fmt.Sprintf(format, v...)
	logger.Printf("[%s] %s", SimulationState.ReadableTime(), printStr)
}
