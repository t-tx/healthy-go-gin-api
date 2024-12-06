package user

type UserDiary struct {
	Username  string `json:"username"`
	CreatedAt string `json:"time"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
