package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	isExportPrice bool
}

func (p *Product) MarshalJSON() ([]byte, error) {
	if p.isExportPrice {
		return json.Marshal(*p)
	}

	m := map[string]interface{}{
		"name":  p.Name,
	}
	return json.Marshal(m)
}

func main() {
	p1 := &Product{Name: "Paper", Price: 50, isExportPrice: false}
	p2 := &Product{Name: "Gold", Price: 100, isExportPrice: true}

	o1, _ := json.Marshal(p1)
	fmt.Println(string(o1))

	o2, _ := json.Marshal(p2)
	fmt.Println(string(o2))
}
