package models

import "gorm.io/gorm"

type Product struct {
	ProductID   int     `json:"productId" gorm:"primaryKey"`
	ProductName string  `json:"productName"`
	SupplierID  *int    `json:"supplierId"`
	CategoryID  *int    `json:"categoryId"`
	QuantityPerUnit string `json:"quantityPerUnit"`
	UnitPrice   float64 `json:"unitPrice"`
	UnitsInStock int    `json:"unitsInStock"`
	UnitsOnOrder int    `json:"unitsOnOrder"`
	ReorderLevel int    `json:"reorderLevel"`
	Discontinued bool   `json:"discontinued"`
}

