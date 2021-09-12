package models 

// `Deck` object / model

type Deck struct {
	Deck_id		string 	`json:"deck_id"`
	Shuffled	bool 		`json:"shuffled"`
	Remaining int 		`json:"remaining"`
	Cards 		[]Card 	`json:"cards,omitempty"`
}
