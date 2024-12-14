package events

import (
	"fmt"

	"github.com/alifyasa/wardrobe-inventory-simulation/config"
	"github.com/alifyasa/wardrobe-inventory-simulation/helper"
	"github.com/alifyasa/wardrobe-inventory-simulation/outfits"
	"github.com/alifyasa/wardrobe-inventory-simulation/simulation"
)

func LaundryStart() {
	numChanged := outfits.SetMultipleOutfitState(
		simulation.SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.DIRTY),
		outfits.BEING_CLEANED,
	)
	if numChanged <= 0 {
		// If there are no laundry, schedule 5am tomorrow
		fmt.Printf("[%s] No Laundry, Scheduling at 5 AM Tomorrow\n", simulation.SimulationState.ReadableTime())
		laundryStartTime := getNextTimeAtHour(5)
		simulation.Schedule(
			LaundryStart,
			laundryStartTime,
		)
	} else {
		fmt.Printf("[%s] Doing Laundry for %d Outfits\n", simulation.SimulationState.ReadableTime(), numChanged)
		laundryEndTime := simulation.SimulationState.Hour + helper.IntExponentialDistribution(
			config.MEAN_LAUNDRY_DURATION_HOUR,
			config.MIN_LAUNDRY_DURATION_HOUR,
		)
		simulation.Schedule(
			LaundryEnd,
			laundryEndTime,
		)
	}
}

func LaundryEnd() {
	numChanged := outfits.SetMultipleOutfitState(
		simulation.SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.BEING_CLEANED),
		outfits.CLEAN,
	)
	fmt.Printf("[%s] Laundry Done for %d Outfits\n", simulation.SimulationState.ReadableTime(), numChanged)
	laundryStartTime := getNextTimeAtHour(5)
	simulation.Schedule(
		LaundryStart,
		laundryStartTime,
	)
}
