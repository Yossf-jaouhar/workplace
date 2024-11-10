package funcs

type AllArtistData struct {
	Id          int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`

	ConcertDates string `json:"concertDates"`

	Relations string `json:"relations"`

	FEATCHINGerror error
}

var Artist []AllArtistData
