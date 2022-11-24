package repositories

type Book struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Author    []Author `json:"author"`
	Genre     string   `json:"genre"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
