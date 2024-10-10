package model

import (
	"fmt"

	"github.com/google/uuid"
)

type ModelCacher interface {
	AddToCache(*Todo)
	GetCache() []*Todo
	FindUser(*Todo) *User
	FindUserById(string) (*User, error)
	CreateSubscriber() chan []*Todo
	AssignTodo(string, string) (*Todo, error)
	CreateUser(string) *User
	GetAllUsers() []*User
}

func NewModelCache() ModelCacher {
	return &ModelCache{
		cache: []*Todo{},
	}
}

type ModelCache struct {
	users       []*User
	cache       []*Todo
	subscribers []chan []*Todo
}

func (m *ModelCache) notifySubscribers(t *Todo) {
	for _, subscriber := range m.subscribers {
		subscriber <- []*Todo{t}
	}
}

func (m *ModelCache) AddToCache(t *Todo) {
	m.cache = append(m.cache, t)
	m.notifySubscribers(t)
}

func (m *ModelCache) GetCache() []*Todo {
	cacheCopy := make([]*Todo, len(m.cache))
	copy(cacheCopy, m.cache)
	return cacheCopy
}

func (m *ModelCache) FindUser(t *Todo) *User {
	for _, cachedTodo := range m.cache {
		if cachedTodo.ID == t.ID {
			return cachedTodo.User
		}
	}
	return nil
}

func (m *ModelCache) CreateSubscriber() chan []*Todo {
	newSubscriberChannel := make(chan []*Todo)
	m.subscribers = append(m.subscribers, newSubscriberChannel)
	return newSubscriberChannel
}

func (m *ModelCache) AssignTodo(todoId, userId string) (*Todo, error) {
	foundUser, findUserErr := m.findUser(userId)
	if findUserErr != nil {
		return nil, findUserErr
	}
	for _, todo := range m.cache {
		if todo.ID == todoId {
			todo.User = foundUser
			m.notifySubscribers(todo)
			return todo, nil
		}
	}
	return nil, fmt.Errorf("cannot find todo with id %v", todoId)
}

func (m *ModelCache) findUser(userId string) (*User, error) {
	for _, user := range m.users {
		if user.ID == userId {
			return user, nil
		}
	}
	return nil, fmt.Errorf("Cannot find user with id %v", userId)
}

func (m *ModelCache) CreateUser(userName string) *User {
	newUser := &User{
		ID:   uuid.NewString(),
		Name: userName,
	}
	m.users = append(m.users, newUser)
	return newUser
}

func (m *ModelCache) FindUserById(userId string) (*User, error) {
	for _, user := range m.users {
		if user.ID == userId {
			return user, nil
		}
	}
	return nil, fmt.Errorf("cannot find user with id %v", userId)
}

func (m *ModelCache) GetAllUsers() []*User {
	usersCopy := make([]*User, len(m.users))
	copy(usersCopy, m.users)
	return usersCopy
}
