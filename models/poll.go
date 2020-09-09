package models

type PollOption struct {
	Option string `json:"option"`
	Users  []User `json:"users"`
}

type Poll struct {
	Id      int64        `json:"id"`
	Title   string       `json:"title"`
	Options []PollOption `json:"options"`
}
