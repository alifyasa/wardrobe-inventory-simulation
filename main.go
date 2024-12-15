package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/alifyasa/wardrobe-inventory-simulation/config"
	"github.com/alifyasa/wardrobe-inventory-simulation/simulation"
	"github.com/alifyasa/wardrobe-inventory-simulation/simulation/events"
)

func main() {
	simulation.Schedule(events.Shower, events.GetNextShowerTime())
	simulation.Schedule(events.LaundryStart, 0)

	for simulation.SimulationState.Hour <= config.SIMULATION_DURATION_HOUR {
		simulation.Execute()
		simulation.Track()
	}

	PrintSummary()
}

func PrintSummary() {
	simulationTime := time.Now().Format(time.RFC3339)
	fmt.Printf("WARDROBE INVENTORY SIMULATION (%s)\n\n", simulationTime)
	fmt.Printf("Simulated Time (Hours)\t: %d\n", config.SIMULATION_DURATION_HOUR)
	fmt.Printf("Total Outfit\t\t: %d\n", len(simulation.SimulationState.AllOutfits.Outfits))

	logsPath, _ := WriteMetricsToFile(simulationTime)
	fmt.Printf("Logs Path\t\t: %s\n", logsPath)
}

func WriteMetricsToFile(simulationTime string) (string, error) {
	jsonData, err := json.MarshalIndent(simulation.SimulationState.Metrics, "", "  ")
	if err != nil {
		fmt.Println("Error Marshalling Simulation Metrics")
	}

	logsDir := "logs"
	logsFilename := fmt.Sprintf("%s.json", simulationTime)
	logsPath := filepath.Join(logsDir, logsFilename)
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		fmt.Println("Error creating logs directory:", err)
		return "", fmt.Errorf("error creating directory %s. %s", logsDir, err.Error())
	}
	file, err := os.Create(logsPath)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return "", fmt.Errorf("error creating log file %s. %s", logsPath, err.Error())
	}
	defer file.Close()
	if _, err := file.Write(jsonData); err != nil {
		fmt.Println("Error writing to log file:", err)
		return "", fmt.Errorf("error writing to log file %s. %s", logsPath, err.Error())
	}
	return logsPath, nil
}
