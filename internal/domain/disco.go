package domain

type Disco struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Artist    string  `json:"artist"`
	Year      int     `json:"year"`
	Genre     string  `json:"genre"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Available bool    `json:"available"`
}
