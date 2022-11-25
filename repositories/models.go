package repositories

type Book struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Author []Author `json:"author"`
	Genre  string   `json:"genre"`
}

type Author struct {
	ID    string `json:"id" copier:"AuthorID"`
	Name  string `json:"name"`
	Email string `json:"email" copier:"Mail"`
}

type Genre struct {
	ID   string `json:"id" copier:"GenreID"`
	Name string `json:"name"`
}
