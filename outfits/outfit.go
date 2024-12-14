package outfits

type OutfitState int

const (
	CLEAN OutfitState = iota
	DIRTY
	BEING_CLEANED
)

type Outfit struct {
	state OutfitState
}

func NewOutfit() Outfit {
	return Outfit{CLEAN}
}
