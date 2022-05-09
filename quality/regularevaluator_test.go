package quality_test

import (
	"testing"

	"github.com/myugen/go-gildedrose-solution/gildedrose"
	"github.com/myugen/go-gildedrose-solution/quality"
	"github.com/stretchr/testify/assert"
)

func TestRegularEvaluator_Evaluate(t *testing.T) {
	evaluator := quality.NewRegularEvaluator()

	type args struct {
		item gildedrose.Item
	}
	type want struct {
		quality int
	}
	testcases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "should decrease by 1 the quality value when an item is provided with sell-in value greater than 0",
			args: args{item: gildedrose.Item{Name: "Regular-evaluated item", SellIn: 1, Quality: 9}},
			want: want{quality: 8},
		},
		{
			name: "should decrease by 2 the quality value when an item is provided with sell-in value equal than 0",
			args: args{item: gildedrose.Item{Name: "Regular-evaluated item", SellIn: 0, Quality: 8}},
			want: want{quality: 6},
		},
		{
			name: "should decrease by 2 the quality value when an item is provided with sell-in value lower than 0",
			args: args{item: gildedrose.Item{Name: "Regular-evaluated item", SellIn: -1, Quality: 6}},
			want: want{quality: 4},
		},
		{
			name: "should keep the quality value when an item is provided with quality value equal than minimum quality (0)",
			args: args{item: gildedrose.Item{Name: "Regular-evaluated item", SellIn: -1, Quality: quality.RegularMinimumQuality}},
			want: want{quality: quality.RegularMinimumQuality},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			evaluator.Evaluate(&testcase.args.item)
			assert.Equal(t, testcase.want.quality, testcase.args.item.Quality)
		})
	}
}