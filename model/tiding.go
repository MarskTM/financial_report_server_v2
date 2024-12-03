package model

type Tiding struct {
	ID int32 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
    Timestamp int64 `json:"timestamp"`
}
