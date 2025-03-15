/*
Interface - a type that defines a set of method signatures without providing implementations.
*/
package intermediate

import (
	"errors"
	"fmt"
	"time"
)

type Storage interface {
	Save(key string, value string) error
	Load(key string) (string, error)
}

type MemoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]string)}
}

func (m *MemoryStorage) Save(key string, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Load(key string) (string, error) {
	if value, ok := m.data[key]; ok {
		return value, nil
	}
	return "", errors.New("Key not found")
}

func ExampleInterfaces() {
	var storage Storage = NewMemoryStorage()

	storage.Save("firstName", "Daniil")
	storage.Save("lastName", "Kalts")

	fmt.Println("Trying to figure out my first name..")
	time.Sleep(time.Second)
	if value, err := storage.Load("firstName"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("My first name is:", value)
	}

	fmt.Println()

	fmt.Println("Trying to figure out my age..")
	time.Sleep(time.Second * 2)
	if value, err := storage.Load("age"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("My age is:", value)
	}
}
