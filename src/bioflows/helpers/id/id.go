package id

import "github.com/aidarkhanov/nanoid"

const (
	FLOWS_ALPHABET = nanoid.DefaultAlphabet
	SIZE = 15
)

func NewID() (string,error){
	return nanoid.GenerateString(FLOWS_ALPHABET,SIZE)
}