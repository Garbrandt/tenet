// @Title  后台接口结构体
// @Description  管理者登陆后台后，后台所有需要的数据，都通过本结构体包含返回
// @Author  魏帅
// @Update  魏帅
package model

type Connect struct {
	Content Content `json:"content"`
}

type Row struct {
	Content     Content              `json:"content"`
	Connections map[string][]Content `json:"connections"`
}

type Relationship struct {
	Relation    Mark      `json:"relation"`
	Connections []Content `json:"connections"`
}

type Submit struct {
	Content     Content              `json:"content"`
	Connections map[string][]Content `json:"connections"`
}

type Form struct {
	Mark   Mark   `json:"mark"`
	Submit Submit `json:"submit"`
}

type Dashboard struct {
	Navigations []map[string]string `json:"navigations"`
	Forms       map[string][]Form   `json:"forms"`
	Rows        map[string][]Row    `json:"rows"`
}
