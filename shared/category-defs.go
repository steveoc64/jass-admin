package shared

import "fmt"

type Category struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Descr string `db:"descr"`
}

func (t Category) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t Category) RootURL() string {
	return "/api/category"
}

type Product struct {
	ID               int    `db:"id"`
	Name             string `db:"name"`
	Descr            string `db:"descr"`
	Image            string `db:"image"`
	ShippingVolumeML int    `db:"shipping_volume_ml"`
	ShippingWeightG  int    `db:"shipping_weight_g"`
	ShippingCode     int    `db:"shipping_code"`
}

func (t Product) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t Product) RootURL() string {
	return "/api/product"
}
