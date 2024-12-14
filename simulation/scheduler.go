package simulation

func Schedule(event func(), hour int) {
	SimulationState.Push(&Event{
		Hour:      hour,
		EventFunc: event,
	})
}

func Execute() {
	currentEvent := SimulationState.Pop()
	SimulationState.Hour = currentEvent.Hour
	currentEvent.EventFunc()
}
