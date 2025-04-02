package db

import "testing"

func TestMapDbUserToUser(t *testing.T) {
	dbUser := generateDbUser("John", "Doe", "john.doe@gmail.com")
	want := generateUser(dbUser)

	user := mapDbUserToUser(dbUser)
	areEqual := checkUserEquality(want, user)

	if !areEqual {
		t.Errorf("MapDbUserToUser(dbUser) = %v, want %v", user, want)
	}
}
