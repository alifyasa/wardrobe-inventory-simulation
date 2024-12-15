package simulation

import (
	"reflect"
	"runtime"

	"github.com/alifyasa/wardrobe-inventory-simulation/outfits"
)

type SimulationMetric struct {
	Hour                int
	ReadableTime        string
	NumOfOutfitsByState map[outfits.OutfitState]int
	NextEvent           string
}

type SimulationMetrics []SimulationMetric

func Track() {
	SimulationState.Metrics = append(SimulationState.Metrics, SimulationMetric{
		Hour:         SimulationState.Hour,
		ReadableTime: SimulationState.ReadableTime(),
		NextEvent:    runtime.FuncForPC(reflect.ValueOf(SimulationState.Peek().EventFunc).Pointer()).Name(),
		NumOfOutfitsByState: map[outfits.OutfitState]int{
			outfits.CLEAN:         len(SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.CLEAN)),
			outfits.DIRTY:         len(SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.DIRTY)),
			outfits.BEING_CLEANED: len(SimulationState.AllOutfits.GetAllOutfitsWithState(outfits.BEING_CLEANED)),
		},
	})
}
