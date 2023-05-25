package models

type Contract struct {
    ID          int      `json:id`
    Title       string   `json:"title"`
	Type		string	 `json:type`
    Content     string   `json:"content"`
    Author      string   `json:"author"`
    Description string   `json:"description"`
}
