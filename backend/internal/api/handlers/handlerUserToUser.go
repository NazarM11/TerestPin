package handlers

import "github.com/NazarM11/TerestPin/internal/database"

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		Email:     dbUser.Email,
		HashedPassword:  dbUser.HashedPassword,
	}
}
