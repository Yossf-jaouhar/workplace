package funcs

import (
	"encoding/json"
	"errors"
	"net/http"
)

func GGet(path string) (*http.Response, error) {
	resp, err := http.Get(path)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, errors.New("response not ok")
	}
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// This function fetches a list of artists from the API and stores them in the Artists slice.
func fetchArtists() {
	resp, err := GGet("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&Artis.Artists)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
}

// A method on the ArtistData struct that fetches the relations of an artist from their Relations link.
func (a *ArtistData) fetchRelations() {
	resp, err := GGet(a.Relations)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&a.RelationData)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
}

// A method that fetches the artist's locations data.
func (a *ArtistData) fetchLocation() {
	resp, err := GGet(a.Locations)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&a.LocationsDATA)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
}

// A method that fetches the concert dates of the artist.
func (a *ArtistData) fetchDates() {
	resp, err := GGet(a.ConcertDates)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&a.ConcertDatesDATA)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
}

// This is a wrapper method that calls the other fetch methods (fetchRelations, fetchDates, and fetchLocation) for an artist.
func (a *ArtistData) FeatchAll() {
	a.fetchRelations()
	a.fetchDates()
	a.fetchLocation()
}

func SearchForAllLocations() {
	resp, err := GGet("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&Artis.AllLocation)
	if err != nil {
		Artis.FEATCHINGerror = err
		return
	}
}
