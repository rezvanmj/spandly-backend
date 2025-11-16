package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"uniqueIndex"` //unique email
	Password     string `json:"password"`
	Income       int    `json:"income"`
	SavingGoal   int    `json:"savingGoal"`
	SavedMoney   int    `json:"savedMoney"`
	CurrentMoney int    `json:"currentMoney"`
}
