package services

import (
	"github.com/minhaj/library-system/models"
)

var (
	firstNames = []string{
		"John", "Sarah", "Michael", "Emily", "David",
		"Jessica", "Robert", "Jennifer", "William", "Lisa",
	}

	lastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Miller", "Davis", "Garcia", "Rodriguez", "Wilson",
	}

	domains = []string{
		"gmail.com", "yahoo.com", "outlook.com", "protonmail.com", "icloud.com",
	}
)

func randomFirstName() string {
	return firstNames[Random.Intn(len(firstNames))]
}

func randomLastName() string {
	return lastNames[Random.Intn(len(lastNames))]
}

func randomEmail(first, last string) string {
	return first + "." + last + "@" + domains[Random.Intn(len(domains))]
}

func ListOfUsers() []*models.User {
	users := make([]*models.User, 0, 20)

	for i := 0; i < 20; i++ {
		firstName := randomFirstName()
		lastName := randomLastName()
		email := randomEmail(firstName, lastName)

		users = append(users, models.NewUser(
			i+1,
			firstName+" "+lastName,
			email,
		))
	}

	return users
}
