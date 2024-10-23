package types

type Component struct {
	Type          string `json:"type"`
	Status        string `json:"status"`
	MainText      string `json:"mainText"`
	SecondaryText string `json:"secondaryText"`
}
