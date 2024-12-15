package events

import (
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
		simulation.Log("No laundry, scheduling again tomorrow")
		laundryStartTime := getNextTimeAtHour(5)
		simulation.Schedule(
			LaundryStart,
			laundryStartTime,
		)
	} else {
		simulation.Log("Doing laundry for %d outfits", numChanged)
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
	simulation.Log("Doing laundry for %d outfits", numChanged)
	laundryStartTime := getNextTimeAtHour(5)
	simulation.Schedule(
		LaundryStart,
		laundryStartTime,
	)
}
