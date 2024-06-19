package entities

type Broth struct {
	ID               string  `json:"id"`
	ImageInactiveURL string  `json:"imageInactive"`
	ImageActiveURL   string  `json:"imageActive"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Price            float32 `json:"price"`
}
