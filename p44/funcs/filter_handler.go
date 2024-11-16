package funcs

import (
	"html/template"
	"net/http"
	"strconv"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	//r.ParseForm()
	////
	
	/////
	Location := r.FormValue("location")
	
	Members := r.Form["member"]
	///////
	CreationStart := r.FormValue("CreationDate-start")
	CreationStartInt, _ := strconv.Atoi(CreationStart)

	/////
	CreationEnd := r.FormValue("CreationDate-end")
	CreationEndInt, _ := strconv.Atoi(CreationEnd)

	////
	FirstAlbumStart := r.FormValue("FirstAlbum-start")
	FirstAlbumStartInt, _ := strconv.Atoi(FirstAlbumStart)

	//////
	FirstAlbumEnd := r.FormValue("FirstAlbum-end")
	FirstAlbumEndInt, _ := strconv.Atoi(FirstAlbumEnd)

	Research(Members, Location, CreationStartInt, CreationEndInt, FirstAlbumStartInt, FirstAlbumEndInt)

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	errr := tmpl.Execute(w, Artis)
	if errr != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

// func Research(Members []string, Location string, CreationStart int, CreationEnd int, FirstAlbumStart int, FirstAlbumEnd int) {
// 	var ArtistList []ArtistData

// 	for index, artist := range Artis.Artists {

// 		foundCreation := CreationStart <= artist.CreationDate && CreationEnd >= artist.CreationDate

// 		foundMember := Members == nil
// 		for _, Member := range Members {

// 			MemberInt, err := strconv.Atoi(Member)
// 			if err != nil {
// 				return
// 			}
// 			if len(artist.Members) == MemberInt {
// 				foundMember = true
// 				break
// 			}
// 		}

// 		foundLocation := Location == ""
// 		locations := Artis.AllLocation.Index[index].Locations
// 		for _, location := range locations {
// 			if location == Location {
// 				foundLocation = true
// 				break
// 			}
// 		}

// 		albumYear := IsValid(artist.FirstAlbum)
// 		foundFirstAlbum := (albumYear >= FirstAlbumStart && albumYear <= FirstAlbumEnd)

// 		if foundCreation && foundMember && foundLocation && foundFirstAlbum {
// 			ArtistList = append(ArtistList, Artis.Artists[index])
// 		}

// 	}
// 	if ArtistList == nil {
// 		Artis.SearchArt = Artis.Artists
// 	}
// 	Artis.SearchArt = ArtistList
// }

func Research(Members []string, Location string, CreationStart int, CreationEnd int, FirstAlbumStart int, FirstAlbumEnd int) {
	var ArtistList []ArtistData

	for index, artist := range Artis.Artists {

		foundCreation := CreationStart == 1987 && CreationEnd == 1987
		for i := CreationStart; i <= CreationEnd; i++ {
			if i == artist.CreationDate {
				foundCreation = true
				break
			}
		}

		foundMember := Members == nil
		for _, Member := range Members {

			MemberInt, err := strconv.Atoi(Member)
			if err != nil {
				return
			}
			if len(artist.Members) == MemberInt {
				foundMember = true
				break
			}
		}

		foundLocation := Location == ""
		locations := Artis.AllLocation.Index[index].Locations
		for _, location := range locations {
			if location == Location {
				foundLocation = true
				break
			}
		}

		foundFirstAlbum := FirstAlbumStart == 1987 && FirstAlbumEnd == 1987
		// fmt.Println(FirstAlbumStart)
		// fmt.Println(FirstAlbumEnd)
		for i := FirstAlbumStart; i <= FirstAlbumEnd; i++ {

			if i == IsValid(artist.FirstAlbum){
				foundFirstAlbum = true
				break
			}
		}

		if foundCreation && foundMember && foundLocation && foundFirstAlbum {
			ArtistList = append(ArtistList, Artis.Artists[index])
		}

	}

	Artis.SearchArt = ArtistList
}

func IsValid(albumYear string) int {
	TheYear := albumYear[len(albumYear)-4:]

	// fmt.Println(TheYear)

	AlbumYear, _ := strconv.Atoi(TheYear)
	return AlbumYear
}
