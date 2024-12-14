package outfits

type AllOutfits struct {
	Outfits []*Outfit
}

func (w *AllOutfits) GetCleanOutfit() *Outfit {
	// Search for clean outfit
	for _, outfit := range w.Outfits {
		if outfit.state == CLEAN {
			return outfit
		}
	}
	// If not found, get a new one
	newOutfit := NewOutfit()
	w.Outfits = append(w.Outfits, &newOutfit)
	return &newOutfit
}

func (w *AllOutfits) GetAllOutfitsWithState(state OutfitState) []*Outfit {
	filteredOutfits := make([]*Outfit, 0)
	for _, outfit := range w.Outfits {
		if outfit.state == state {
			filteredOutfits = append(filteredOutfits, outfit)
		}
	}
	return filteredOutfits
}

func SetMultipleOutfitState(outfits []*Outfit, state OutfitState) int {
	for _, outfit := range outfits {
		outfit.state = state
	}
	return len(outfits)
}
