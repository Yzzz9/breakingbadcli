package data

var Episode int

// Response struct
type Response struct {
    Eid int `json:"episode_id"`
    Title string `json:"title"`
    Season string `json:"season"`
    Epnum string `json:"episode"`
    Series string `json:"series"`
    Airdate string `json:"air_date"`
    Characters []string `json:"characters"`
}
