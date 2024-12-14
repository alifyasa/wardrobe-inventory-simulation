package events

import (
	"fmt"
	"math"

	"github.com/alifyasa/wardrobe-inventory-simulation/config"
	"github.com/alifyasa/wardrobe-inventory-simulation/constant"
	"github.com/alifyasa/wardrobe-inventory-simulation/outfits"
	"github.com/alifyasa/wardrobe-inventory-simulation/simulation"
)

func Shower() {
	fmt.Printf("[%s] Showered\n", simulation.SimulationState.ReadableTime())
	outfits.SetMultipleOutfitState(
		[]*outfits.Outfit{
			simulation.SimulationState.CurrentOutfit,
		},
		outfits.DIRTY,
	)
	simulation.SimulationState.CurrentOutfit = simulation.SimulationState.AllOutfits.GetCleanOutfit()

	simulation.Schedule(Shower, GetNextShowerTime())
}

func GetNextShowerTime() int {
	minNextShowerTime := math.MaxInt
	for _, showerHour := range config.SHOWER_TIME {
		nextShowerTime := getNextShowerTimeAtHour(showerHour)
		if nextShowerTime < minNextShowerTime {
			minNextShowerTime = nextShowerTime
		}
	}
	return minNextShowerTime
}

func getNextShowerTimeAtHour(hourOfDay int) int {
	currentHourOfDay := simulation.SimulationState.Hour % constant.HOUR_IN_DAY
	if currentHourOfDay == hourOfDay {
		return simulation.SimulationState.Hour + constant.HOUR_IN_DAY
	}

	nextShowerTime := simulation.SimulationState.Hour + (hourOfDay-currentHourOfDay+constant.HOUR_IN_DAY)%constant.HOUR_IN_DAY
	return nextShowerTime
}
