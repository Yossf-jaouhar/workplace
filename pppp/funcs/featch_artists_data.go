package funcs

import "net/http"


	 
func FeatchArtists(){
	Respons , err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if  err != nil {
		return
	}
	defer Respons.Body.Close()
	
} 