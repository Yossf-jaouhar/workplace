package funcs

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Search handles search requests for artists.
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}


	

	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	Searchh := r.FormValue("SeArch")
	Searchh = strings.ToLower(Searchh)
	var ArtistList []ArtistData

	for index, artist := range Artis.Artists {
		found := false

		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(Searchh)) ||
			strings.HasPrefix(strings.ToLower(Searchh), strings.ToLower(artist.Name)) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), Searchh) ||
			artist.FirstAlbum == Searchh {

				ArtistList = append(ArtistList, Artis.Artists[index])
			continue
		}
		locations := Artis.AllLocation.Index[index].Locations
		for _, location := range locations {
			if strings.Contains(location, Searchh) {
				ArtistList = append(ArtistList, Artis.Artists[index])
				found = true
				break
			}
		}

		if found {
			continue
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(Searchh)) {
				ArtistList = append(ArtistList, Artis.Artists[index])
				break
			}
		}
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	Artis.SearchArt = ArtistList
	errr := tmpl.Execute(w, Artis)
	if errr != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
