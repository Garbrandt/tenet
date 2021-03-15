package model

type MarkContentMapping map[string]string

var MCMap MarkContentMapping

func init() {
	MCMap = map[string]string{
		"mark_content_id":     "id",
		"mark_created_at":     "created_at",
		"mark_updated_at":     "updated_at",
		"mark_key":            "key",
		"mark_title":          "title",
		"mark_type":           "type",
		"mark_abstract":       "abstract",
		"mark_content_url":    "content_url",
		"mark_body":           "body",
		"mark_source_url":     "source_url",
		"mark_size":           "size",
		"mark_copyright_stat": "copyright_stat",
	}
}
