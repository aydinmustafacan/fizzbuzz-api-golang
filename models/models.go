package models



import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

type Morse struct {

	MorseCode string `json:"morsecode,omitempty"`
}

type FizzBuzz struct {
	FizzBuzz string `json:"fizzBuzz,omitempty"`

}

