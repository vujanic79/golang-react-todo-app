package db

import "testing"

func TestMapDbUserToUser(t *testing.T) {
	dbU := generateDbUser("John", "Doe", "john.doe@gmail.com")
	want := generateUser(dbU)

	u := mapDbUserToUser(dbU)
	areEqual := checkUserEquality(want, u)

	if !areEqual {
		t.Errorf("MapDbUserToUser(dbU) = %v, want %v", u, want)
	}
}
