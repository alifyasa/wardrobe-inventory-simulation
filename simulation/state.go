package simulation

import (
	"container/heap"
	"fmt"

	"github.com/alifyasa/wardrobe-inventory-simulation/outfits"
)

type SimulationStateType struct {
	Hour          int
	eventQueue    EventQueue
	AllOutfits    *outfits.AllOutfits
	CurrentOutfit *outfits.Outfit
	Metrics       SimulationMetrics
}

var allOutfits = &outfits.AllOutfits{Outfits: make([]*outfits.Outfit, 0)}
var SimulationState = &SimulationStateType{
	0,
	make(EventQueue, 0),
	allOutfits,
	allOutfits.GetCleanOutfit(),
	make(SimulationMetrics, 0),
}

func (s *SimulationStateType) Push(e *Event) {
	heap.Push(&s.eventQueue, e)
}

func (s *SimulationStateType) Pop() *Event {
	return heap.Pop(&s.eventQueue).(*Event)
}

func (s *SimulationStateType) Peek() *Event {
	if len(s.eventQueue) == 0 {
		return nil
	}
	return s.eventQueue[0]
}

func (s *SimulationStateType) ReadableTime() string {
	hour := s.Hour % 24
	day := (s.Hour - hour) / 24
	return fmt.Sprintf("D%04d/H%02d", day, hour)
}
