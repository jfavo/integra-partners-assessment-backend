package constants

import "github.com/jfavo/integra-partners-assessment-backend/internal/models"

var (
	TestUsers = []models.User{
		{
			UserId:     1,
			Username:   "testUser",
			Firstname:  "test",
			Lastname:   "user",
			Email:      "test@user.com",
			UserStatus: "A",
			Department: "sales",
		},
		{
			UserId:     2,
			Username:   "testUser2",
			Firstname:  "test2",
			Lastname:   "user",
			Email:      "test2@user.com",
			UserStatus: "T",
			Department: "management",
		},
	}
)