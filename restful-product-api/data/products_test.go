package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "buli",
		Price: 1.00,
		SKU:   "abc-def-fhi",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
