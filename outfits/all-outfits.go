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

func (w *AllOutfits) GetAllDirtyOutfits() []*Outfit {
	dirtyOutfits := make([]*Outfit, 0)
	for _, outfit := range w.Outfits {
		if outfit.state == DIRTY {
			dirtyOutfits = append(dirtyOutfits, outfit)
		}
	}
	return dirtyOutfits
}

func (w *AllOutfits) SetMultipleOutfitState(outfits []*Outfit, state OutfitState) {
	for _, outfit := range outfits {
		outfit.state = state
	}
}
