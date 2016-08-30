package omdbapi

const (
	omdbURL = "http://omdbapi.com"
)

type Movie struct {
	Title     string
	Year      string
	Rated     string
	Poster    string
	Metascore string
}
