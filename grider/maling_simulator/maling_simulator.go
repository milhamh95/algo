package maling_simulator

import (
	"sort"
)

type Item struct {
	Name           string
	Weight         float64
	Total          int
	Price          float64
	PricePerWeight float64
}

type FinalItem struct {
	Name     string
	Quantity int
}

func malingSimulator(items []Item, maxWeight float64) ([]FinalItem, float64) {
	// price per weight
	result := []FinalItem{}

	for i, v := range items {
		pricePerWeight := v.Price / v.Weight
		items[i].PricePerWeight = pricePerWeight
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].PricePerWeight > items[j].PricePerWeight
	})

	totalWeight := float64(0)
	totalPrice := float64(0)
	for _, v := range items {
		qty := 0
		for qty < v.Total && totalWeight < maxWeight {
			tmpWeight := totalWeight + v.Weight
			if tmpWeight > maxWeight {
				break
			}

			totalPrice += v.Price
			totalWeight = tmpWeight
			qty++
		}

		if qty > 0 {
			result = append(result, FinalItem{
				Name:     v.Name,
				Quantity: qty,
			})
		}
	}

	return result, totalPrice
}
