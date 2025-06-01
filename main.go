package main 

import (
  "fmt"
  "log"
 "encoding/json"
 "math/rand"
 "net/http"
 "strconv"
 "github.com/gorilla/mux"
)

type Movie struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}

type Director struct {
     FirstName string `json:"firstname"`
  LastName string `json:"lastname"`

}


var Movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  if err := json.NewEncoder(w).Encode(Movies); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}



func DeleteMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  id := params["id"]
  for index, movie := range Movies{
    if movie.ID == id{
      Movies = append(Movies[:index], Movies[index+1:]...)
      break
    }
  }
  if err := json.NewEncoder(w).Encode(Movies); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}


func GetMovie (w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  params := mux.Vars(r)
  id := params["id"]

  for _, movie := range Movies{
    if(movie.ID == id){
      if err := json.NewEncoder(w).Encode(movie); err != nil{
        http.Error(w, err.Error(), http.StatusInternalServerError)
      }
      return
    }
  }

  http.Error(w, "Movie not found", http.StatusNotFound)
}



func CreateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    defer r.Body.Close()

    var newMovie Movie
    if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Validate required fields
  if newMovie.Title == "" || newMovie.Isbn == "" || newMovie.Director.FirstName == "" || newMovie.Director.LastName == "" {
        http.Error(w, "Missing required movie fields", http.StatusBadRequest)
        return
    }

    // Assign a random ID
    newMovie.ID = strconv.Itoa(rand.Intn(1000))
    Movies = append(Movies, newMovie)

    if err := json.NewEncoder(w).Encode(newMovie); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}



func UpdateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]

    for index, existingMovie := range Movies {
        if existingMovie.ID == id {
            // Remove old movie
            Movies = append(Movies[:index], Movies[index+1:]...)

            var updatedMovie Movie
            if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "Invalid request body",
                })
                return
            }

            // Validate required fields
            if updatedMovie.Isbn == "" || updatedMovie.Title == "" ||
                updatedMovie.Director == nil || updatedMovie.Director.FirstName == "" || updatedMovie.Director.LastName == "" {
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "ID, ISBN, Title, and Director (firstname, lastname) are required fields",
                })
                return
            }

            updatedMovie.ID = id
            Movies = append(Movies, updatedMovie)

            json.NewEncoder(w).Encode(updatedMovie)
            return
        }
    }

    // Movie not found: send JSON error instead of plain text
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(map[string]string{
        "error": "Movie not found",
    })
}



func main() {
  r := mux.NewRouter()

  Movies = append(Movies, Movie{ID:"1", Isbn: "438743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
  Movies = append(Movies, Movie{ID:"2", Isbn: "438744", Title: "Movie Two", Director: &Director{FirstName: "Jane", LastName: "Doe"}})
  Movies = append(Movies, Movie{ID:"3", Isbn: "438745", Title: "Movie Three", Director: &Director{FirstName: "Jim", LastName: "Beam"}})
  Movies = append(Movies, Movie{ID:"4", Isbn: "43546", Title: "Movie Four", Director: &Director{FirstName: "Jack", LastName: "Daniels"}})


  r.HandleFunc("/movies", GetMovies).Methods("GET")
  r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
  r.HandleFunc("/movies", CreateMovie).Methods("POST")
  r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
  r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

  fmt.Println("Starting server on port 8000...")
  log.Fatal(http.ListenAndServe(":8000", r))
}
