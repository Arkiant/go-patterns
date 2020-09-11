package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	tc := []struct {
		name      string
		color     int
		prototype *Shirt
		sku       string
	}{
		{
			name:      "White new item cannot be equal to the white prototype",
			color:     White,
			prototype: whitePrototype,
			sku:       "abbcc",
		},
		{
			name:      "Black new item cannot be equal to the black prototype",
			color:     Black,
			prototype: blackPrototype,
			sku:       "abbcc",
		},
		{
			name:      "Blue new item cannot be equal to the blue prototype",
			color:     Blue,
			prototype: bluePrototype,
			sku:       "abbcc",
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			item, err := shirtCache.GetClone(tt.color)
			if err != nil {
				t.Error(err)
			}

			if item == tt.prototype {
				t.Errorf("Cannot be equal")
			}

			shirt, ok := item.(*Shirt)
			if !ok {
				t.Fatal("Type assertion for shirt couldn't be done successfully")
			}

			shirt.SKU = tt.sku

			item2, err := shirtCache.GetClone(tt.color)
			if err != nil {
				t.Fatal(err)
			}

			shirt2, ok := item2.(*Shirt)
			if !ok {
				t.Fatal("Type assertion for shirt2 couldn't be done successfully")
			}

			if shirt.SKU == shirt2.SKU {
				t.Error("Shirt1 cannot be equal to Shirt2")
			}

			if &shirt == &shirt2 {
				t.Errorf("shirt and shirt2 has same memory positions")
			}
		})
	}
}
