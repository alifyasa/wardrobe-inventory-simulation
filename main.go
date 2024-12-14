package main

import (
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

}
