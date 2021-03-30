package model

type FilePond struct {
	ID      int64  `json:"id"`
	Source  string `json:"source"`
	Options struct {
		Type string `json:"type"`
	} `json:"options"`
}
