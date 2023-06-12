package repository

import "fmt"

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8").AutoMigrate(
		&User{},
	)
	if err != nil {
		fmt.Println("migration err", err)
	}
}
