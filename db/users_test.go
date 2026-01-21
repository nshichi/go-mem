package main

import (
	"fmt"
	"testing"
)

func BenchmarkAddUsersXxx(b *testing.B) {
	ConnectDB()

	var users []User
	for i := range 1_0000_0000 {
		users = append(users, User{
			Name: fmt.Sprintf("user%d", i),
		})
	}

	b.ResetTimer()

	result := db.CreateInBatches(&users, 1000)
	if result.Error != nil {
		panic(result.Error)
	}
}

func BenchmarkTrancate(b *testing.B) {
	ConnectDB()

	b.ResetTimer()

	if err := db.Exec("TRUNCATE TABLE users").Error; err != nil {
		panic(err)
	}
}
