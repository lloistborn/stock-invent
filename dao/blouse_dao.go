package dao

import (
	"context"
)

// BlouseDao object.
type BlouseDao struct{}

// NewBlouseDao used to access member of dao object.
func NewBlouseDao(ctx context.Context) *BlouseDao {
	return &BlouseDao{}
}

// FindAll return all array of blouse[].
func (blouseDao BlouseDao) FindAll(ctx context.Context) ([]Blouse, error) {
	db := initDB()
	defer db.Close()

	var blouse []Blouse
	err := db.Find(&blouse).Error

	return blouse, err
}

// Insert will create a record from Blouse object.
func (blouseDao BlouseDao) Insert(ctx context.Context, blouse Blouse) error {
	db := initDB()
	defer db.Close()

	err := db.Create(&blouse).Error

	return err
}
