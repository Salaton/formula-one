package model

type DataResponse struct {
	MRData MRData `json:"mrData"`
}

type Location struct {
	Lat      string `json:"lat,omitempty"`
	Long     string `json:"long,omitempty"`
	Locality string `json:"locality,omitempty"`
	Country  string `json:"country,omitempty"`
}

type Circuit struct {
	CircuitID   string   `json:"circuitId,omitempty"`
	URL         string   `json:"url,omitempty"`
	CircuitName string   `json:"circuitName,omitempty"`
	Location    Location `json:"location,omitempty"`
}

type FirstPractice struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}

type SecondPractice struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}

type ThirdPractice struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}

type Qualifying struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}

type Races struct {
	Season         string         `json:"season,omitempty"`
	Round          string         `json:"round,omitempty"`
	URL            string         `json:"url,omitempty"`
	RaceName       string         `json:"raceName,omitempty"`
	Circuit        Circuit        `json:"circuit,omitempty"`
	Date           string         `json:"date,omitempty"`
	Time           string         `json:"time,omitempty"`
	FirstPractice  FirstPractice  `json:"firstPractice,omitempty"`
	SecondPractice SecondPractice `json:"secondPractice,omitempty"`
	ThirdPractice  ThirdPractice  `json:"thirdPractice,omitempty"`
	Qualifying     Qualifying     `json:"qualifying,omitempty"`
}

type RaceTable struct {
	Season string   `json:"season,omitempty"`
	Races  []*Races `json:"races,omitempty"`
}

type MRData struct {
	Xmlns     string     `json:"xmlns,omitempty"`
	Series    string     `json:"series,omitempty"`
	Limit     string     `json:"limit,omitempty"`
	Offset    string     `json:"offset,omitempty"`
	Total     string     `json:"total,omitempty"`
	RaceTable *RaceTable `json:"raceTable,omitempty"`
}
