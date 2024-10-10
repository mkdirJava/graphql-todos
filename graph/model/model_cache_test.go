package model

import "testing"

func TestMain(t *testing.M) {
	unit = NewModelCache()
	t.Run()
}

var unit ModelCacher

func Test_ModelCacher(t *testing.T) {
	t.Run("should be able to create and retrieve Users", func(t *testing.T) {
		initialAllUsersResult := unit.GetAllUsers()
		if len(initialAllUsersResult) != 0 {
			t.Fatal("There should not be any users")
		}
		userName := "bob"
		unit.CreateUser(userName)
		allUsersResult := unit.GetAllUsers()
		if len(allUsersResult) != 1 {
			t.Fatal("There should only be one element in retrieve all users")
		}
		if allUsersResult[0].Name != userName {
			t.Fatalf("user with %v should be present", userName)
		}
		if len(allUsersResult[0].ID) == 0 {
			t.Fatal("the user created should have a created Id")
		}
	})
}
