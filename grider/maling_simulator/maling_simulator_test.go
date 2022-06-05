package maling_simulator

import (
	"reflect"
	"testing"
)

func Test_malingSimulator(t *testing.T) {
	type args struct {
		items  []Item
		maxCap float64
	}
	tests := []struct {
		name  string
		args  args
		want  []FinalItem
		want1 float64
	}{
		{
			name: "max cap = 20kg",
			args: args{
				items: []Item{
					{
						Name:   "Smart TV",
						Weight: 4800,
						Total:  2,
						Price:  9000000,
					},
					{
						Name:   "Android Tablet",
						Weight: 700,
						Total:  2,
						Price:  3300000,
					},
					{
						Name:   "Nintendo Switch",
						Weight: 400,
						Total:  3,
						Price:  3000000,
					},
					{
						Name:   "Premium Kitchen Knives",
						Weight: 1100,
						Total:  5,
						Price:  4000000,
					},
					{
						Name:   "Premium Golf Stick",
						Weight: 5500,
						Total:  1,
						Price:  9000000,
					},
					{
						Name:   "Laptop",
						Weight: 2000,
						Total:  3,
						Price:  6100000,
					},
					{
						Name:   "Premium Billard Stick",
						Weight: 1000,
						Total:  14,
						Price:  2000000,
					},
					{
						Name:   "PS4",
						Weight: 2800,
						Total:  5,
						Price:  5600000,
					},
					{
						Name:   "Vase",
						Weight: 4400,
						Total:  2,
						Price:  5800000,
					},
					{
						Name:   "Powerbank",
						Weight: 880,
						Total:  7,
						Price:  3800000,
					},
					{
						Name:   "Motorcycle Tire",
						Weight: 2400,
						Total:  8,
						Price:  4100000,
					},
				},
				maxCap: 20000,
			},

			want: []FinalItem{
				{
					Name:     "Nintendo Switch",
					Quantity: 3,
				},
				{
					Name:     "Android Tablet",
					Quantity: 2,
				},
				{
					Name:     "Powerbank",
					Quantity: 7,
				},
				{
					Name:     "Premium Kitchen Knives",
					Quantity: 5,
				},
				{
					Name:     "Laptop",
					Quantity: 2,
				},
				{
					Name:     "Premium Billard Stick",
					Quantity: 1,
				},
			},
			want1: 76400000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := malingSimulator(tt.args.items, tt.args.maxCap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("malingSimulator() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("malingSimulator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
