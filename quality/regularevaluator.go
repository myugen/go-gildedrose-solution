package quality

import (
	"github.com/myugen/go-gildedrose-solution/gildedrose"
)

type RegularEvaluator struct {
	minimumQuality int
}

func (e RegularEvaluator) Evaluate(item *gildedrose.Item) {
	amountToDecrease := 1
	if item.SellIn <= 0 {
		amountToDecrease = 2
	}
	if item.Quality == e.minimumQuality {
		amountToDecrease = 0
	}

	item.Quality -= amountToDecrease
}

func NewRegularEvaluator() RegularEvaluator {
	return RegularEvaluator{minimumQuality: RegularMinimumQuality}
}