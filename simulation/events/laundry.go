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
		// If there are no laundry, schedule it next time
		fmt.Printf("[%s] No Laundry, Scheduling Next Time\n", simulation.SimulationState.ReadableTime())
		laundryStartTime := simulation.SimulationState.Hour + helper.NonNegativeIntNormalDistribution(
			config.MEAN_HOUR_UNTIL_NEXT_LAUNDRY,
			config.STD_HOUR_UNTIL_NEXT_LAUNDRY,
		)
		simulation.Schedule(
			LaundryStart,
			laundryStartTime,
		)
	} else {
		fmt.Printf("[%s] Doing Laundry\n", simulation.SimulationState.ReadableTime())
		laundryEndTime := simulation.SimulationState.Hour + helper.NonNegativeIntNormalDistribution(
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
	fmt.Printf("[%s] Laundry Done\n", simulation.SimulationState.ReadableTime())
	outfits.SetMultipleOutfitState(
		simulation.SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.BEING_CLEANED),
		outfits.CLEAN,
	)
	laundryStartTime := simulation.SimulationState.Hour + helper.NonNegativeIntNormalDistribution(
		config.MEAN_HOUR_UNTIL_NEXT_LAUNDRY,
		config.STD_HOUR_UNTIL_NEXT_LAUNDRY,
	)
	simulation.Schedule(
		LaundryStart,
		laundryStartTime,
	)
}
