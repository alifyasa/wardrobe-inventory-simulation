package main

import (
	"fmt"
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
	}

	PrintSummary()
}

func PrintSummary() {
	fmt.Printf("\nWARDROBE INVENTORY SIMULATION (%s)\n\n", time.Now().Format(time.RFC3339))
	fmt.Printf("Simulated Time (Hours)\t: %d\n", config.SIMULATION_DURATION_HOUR)
	fmt.Printf("Total Outfit\t\t: %d\n", len(simulation.SimulationState.AllOutfits.Outfits))
	fmt.Printf("\n")
}
