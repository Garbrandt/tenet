package model

var ConnectionTableName = "connections"

type Connection struct {
	Content

	ContentId             int64  `json:"content_id"`
	ContentType           string `json:"content_type"`
	ConnectionContentId   int64  `json:"connection_content_id"`
	ConnectionContentType string `json:"connection_content_type"`
}

func (content *Connection) TableName() string {
	return ConnectionTableName
}
