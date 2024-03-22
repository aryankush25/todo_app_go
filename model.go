package main

// Todo represents a single Todo item.
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}
