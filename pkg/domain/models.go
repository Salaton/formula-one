package domain

type DataResponse struct {
	MRData MRData `json:"MRData"`
}
type Location struct {
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Locality string `json:"locality"`
	Country  string `json:"country"`
}
type Circuit struct {
	CircuitID   string   `json:"circuitId"`
	URL         string   `json:"url"`
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"Location"`
}
type FirstPractice struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
type SecondPractice struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
type ThirdPractice struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
type Qualifying struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
type Races struct {
	Season         string         `json:"season"`
	Round          string         `json:"round"`
	URL            string         `json:"url"`
	RaceName       string         `json:"raceName"`
	Circuit        Circuit        `json:"Circuit"`
	Date           string         `json:"date"`
	Time           string         `json:"time"`
	FirstPractice  FirstPractice  `json:"FirstPractice"`
	SecondPractice SecondPractice `json:"SecondPractice"`
	ThirdPractice  ThirdPractice  `json:"ThirdPractice"`
	Qualifying     Qualifying     `json:"Qualifying"`
}
type RaceTable struct {
	Season string  `json:"season"`
	Races  []Races `json:"Races"`
}
type MRData struct {
	Xmlns     string    `json:"xmlns"`
	Series    string    `json:"series"`
	Limit     string    `json:"limit"`
	Offset    string    `json:"offset"`
	Total     string    `json:"total"`
	RaceTable RaceTable `json:"RaceTable"`
}
