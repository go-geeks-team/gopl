package movie

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Poul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Stive McQeen", "Jacqueline Bisset"}},
}
