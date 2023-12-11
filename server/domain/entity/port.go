package entity

type Port struct {
	ID          string    `json:"ID"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
