package model

type Item struct {
	ID       string   `json:"id"`
	UserID   string   `json:"user_id"`
	Title    string   `json:"title"`
	Image    string   `json:"image"`
	Status   string   `json:"status"`
	AuthorID string   `json:"author"`
	TagIDs   []string `json:"tags,omitempty"`
	Date     string   `json:"date"`
}
