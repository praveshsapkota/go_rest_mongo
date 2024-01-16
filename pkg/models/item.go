package models

type Create_item struct {
	Name      string `json:"name" bson:"name" validate:"required"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail" validate:"required"`
	Image     string `json:"image" bson:"image"`
	Skus      []Skus `json:"sku" bson:"sku"`
}

type Update_item struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
	Image     string `json:"image,omitempty" bson:"image,omitempty"`
	Skus      []Skus `json:"sku,omitempty" bson:"sku,omitempty"`
}

type Item struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Image     string `json:"image" bson:"image"`
	Skus      []Skus `json:"sku" bson:"sku"`
}

type Skus struct {
	Name             string     `json:"name" bson:"name"`
	Barcode          int64      `json:"barcode" bson:"barcode"`
	Weight           int        `json:"weight" bson:"weight"`
	Unit             string     `json:"unit" bson:"unit"`
	Image            string     `json:"image" bson:"image"`
	Thumbnail        string     `json:"thumbnail" bson:"thumbnail"`
	Quantity         int8       `json:"quantity" bson:"quantity"`
	Price            int8       `json:"price" bson:"price"`
	Discount_percent int8       `json:"discount_percent" bson:"discount_percent"`
	Attributes       any `json:"attributes" bson:"attributes"`
}

// type attributes struct {
// 	Size  any    `json:"size" bson:"size"`
// 	Color string `json:"color" bson:"color"`
// }

type Search struct {
	ID        string `bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Image     string `json:"image" bson:"image"`
	// Skus      []skus `json:"sku" bson:"sku"`
	Sku []struct {
		Name            string `bson:"name"`
		Barcode         int    `bson:"barcode"`
		Image           string `bson:"image"`
		Thumbnail       string `bson:"thumbnail"`
		Quantity        int    `bson:"quantity"`
		Price           int    `bson:"price"`
		DiscountPercent int    `bson:"discount_percent"`
		Attributes      any
	} `bson:"sku"`
}
