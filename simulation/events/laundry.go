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
		// If there are no laundry, schedule it next time
		laundryStartTime := simulation.SimulationState.Hour + helper.IntNormalDistribution(
			config.MEAN_HOUR_UNTIL_NEXT_LAUNDRY,
			config.STD_HOUR_UNTIL_NEXT_LAUNDRY,
		)
		simulation.Schedule(
			LaundryStart,
			laundryStartTime,
		)
	} else {
		laundryEndTime := simulation.SimulationState.Hour + helper.IntNormalDistribution(
			config.MEAN_LAUNDRY_DURATION_HOUR,
			config.STD_LAUNDRY_DURATION_HOUR,
		)
		simulation.Schedule(
			LaundryEnd,
			laundryEndTime,
		)
	}
}

func LaundryEnd() {
	outfits.SetMultipleOutfitState(
		simulation.SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.BEING_CLEANED),
		outfits.CLEAN,
	)
	laundryStartTime := simulation.SimulationState.Hour + helper.IntNormalDistribution(
		config.MEAN_HOUR_UNTIL_NEXT_LAUNDRY,
		config.STD_HOUR_UNTIL_NEXT_LAUNDRY,
	)
	simulation.Schedule(
		LaundryStart,
		laundryStartTime,
	)
}
