package model

type Relation struct {
	ContentID     string `json:"mark_content_id"`
	Title         string `json:"mark_title"`
	Type          string `json:"mark_type"`
	Abstract      string `json:"mark_abstract"`
	ContentUrl    string `json:"mark_content_url"`
	Body          string `json:"mark_body"`
	SourceUrl     string `json:"mark_source_url"`
	Cover         string `json:"mark_cover"`
	Size          int64  `json:"mark_size"`
	CopyrightStat string `json:"mark_copyright_stat"`
	CreatedAt     string `json:"mark_created_at"`
	UpdatedAt     string `json:"mark_updated_at"`

	Section      string `json:"mark_section"`
	SectionLabel string `json:"mark_section_label"`
	Env          string `json:"mark_env"`
	EnvLabel     string `json:"mark_env_label"`

	Multiple string `json:"multiple"`

	Page     string `json:"mark_page"`
	PageSize string `json:"mark_page_size"`

	DateRe string `json:"date_re"`
	Key    string `json:"key"`
}

// 我们需要达成一个共识
// 就是如果我mark中如果设置了content_id=-1,就代表着你可以用全局的来替换我。
// 如果我设置了content_id=0，就代表着你不可以替换我。
type Mark struct {
	Section      string `json:"backend_section"`
	SectionLabel string `json:"backend_section_label"`
	Env          string `json:"backend_env"`
	EnvLabel     string `json:"backend_env_label"`

	ShowOnDashboard bool `json:"mark_show_on_dashboard"`

	RealID        int    `json:"content_id"`
	ContentID     string `json:"mark_content_id"`
	Title         string `json:"mark_title"`
	Type          string `json:"mark_type"`
	Abstract      string `json:"mark_abstract"`
	ContentUrl    string `json:"mark_content_url"`
	Body          string `json:"mark_body"`
	SourceUrl     string `json:"mark_source_url"`
	Size          int64  `json:"mark_size"`
	CopyrightStat string `json:"mark_copyright_stat"`
	CreatedAt     string `json:"mark_created_at"`
	UpdatedAt     string `json:"mark_updated_at"`

	Multiple bool `json:"multiple"`

	Page     int `json:"mark_page"`
	PageSize int `json:"mark_page_size"`

	DateRe string `json:"date_re"`
	Key    string `json:"key"`

	Relations []Relation `json:"relations"`
}
