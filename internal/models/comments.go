package models

type Comment struct {
	CommentID      int    `json:"comment_id"`
	CommentContent string `json:"comment_content"`
	PostID         int    `json:"post_id"`
	PostDate    string `json:"comment_date"`
}
