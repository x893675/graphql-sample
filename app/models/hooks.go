package models

import (
	"context"
	"gorm.io/gorm"
	"log"
)

func (mod *Question) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Logger.Info(context.TODO(), "before question create")
	log.Println("before question create")
	return nil
}

func (mod *QuestionOption) BeforeCreate(tx *gorm.DB) error {
	tx.Logger.Info(context.TODO(), "before question option create")
	log.Println("before question option create")
	return nil
}

func (mod *Answer) BeforeCreate(tx *gorm.DB) error {
	tx.Logger.Info(context.TODO(), "before answer create")
	log.Println("before answer create")
	return nil
}
