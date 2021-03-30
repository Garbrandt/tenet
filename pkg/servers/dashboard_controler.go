package servers

import (
	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/Garbrandt/tenet/pkg/model"
	"github.com/Garbrandt/tenet/pkg/utlis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func getMarksFrom(path string) []model.Mark {
	content := utlis.Convert(path)
	templates := utlis.GetMarksForm(content)
	return templates
}

func generateNavsFrom(marks []model.Mark) []map[string]string {
	var navs []map[string]string
	var navKeys []string
	for _, mark := range marks {
		if mark.Section == "" {
			continue
		}

		if !utlis.Contains(navKeys, mark.Section) {
			navs = append(navs,
				map[string]string{
					"Section":      mark.Section,
					"SectionLabel": mark.SectionLabel,
				},
			)
			navKeys = append(navKeys, mark.Section)
		}
	}
	return navs
}

func generateFormsFrom(marks []model.Mark) map[string][]model.Form {
	forms := map[string][]model.Form{}
	var keys []string
	for _, mark := range marks {
		if mark.Section == "" {
			continue
		}

		if forms[mark.Section] == nil {
			forms[mark.Section] = []model.Form{}
		}

		if !utlis.Contains(keys, mark.Key) {
			keys = append(keys, mark.Key)

			var content = make(map[string][]model.Content, 0)
			for _, connection := range mark.Relations {
				content[connection.Key] = make([]model.Content, 0)
			}

			forms[mark.Section] = append(forms[mark.Section],
				model.Form{
					Mark: mark,
					Submit: model.Submit{
						Content:     model.Content{},
						Connections: content,
					},
				},
			)

			continue
		}
	}

	return forms
}

func getMarks() []model.Mark {
	files := utlis.Work(config.SiteWebPath)
	var marks []model.Mark
	for _, file := range files {
		marks = append(marks, getMarksFrom(file)...)
	}

	return marks
}

func (s *Server) getContentBy(key string) ([]model.Content, error) {
	var contents []model.Content

	limit := 10
	err := s.DB.Where("key = ?", key).Order("created_at DESC").Limit(limit).Find(&contents).Error
	if err == gorm.ErrRecordNotFound || err != nil {
		return contents, err
	}
	return contents, nil
}

func (s *Server) GetDashboard(c *gin.Context) {
	marks := getMarks()

	var dashboard = model.Dashboard{
		Navigations: generateNavsFrom(marks),
		Forms:       generateFormsFrom(marks), // 某个分类下的所有表单
		Rows:        map[string][]model.Row{},
	}

	for _, form := range dashboard.Forms {
		for index, mark := range form {
			contents, err := s.getContentBy(mark.Mark.Key)
			if err != nil {
				continue
			}

			// 我们通过获取到的content来填充内容
			for _, content := range contents {
				if !mark.Mark.Multiple {
					dashboard.Forms[mark.Mark.Section][index] = model.Form{
						Mark:   dashboard.Forms[mark.Mark.Section][index].Mark,
						Submit: model.Submit{},
					}
				}

				// 添加关联数据
				var connections = make(map[string][]model.Content, 0)
				for _, conn := range mark.Mark.Relations {
					var c []model.Content
					err := s.DB.Debug().Joins("JOIN connections ON connections.connection_content_id = contents.id").
						Where("connections.content_id = ? AND connections.connection_content_type = ? AND connections.deleted_at IS NULL",
							content.ID, conn.Key).
						Order("created_at DESC").Where("").Find(&c).Error
					if err == gorm.ErrRecordNotFound || err != nil {
						return
					}
					connections[conn.Key] = c
				}

				row := model.Row{
					Content:     content,
					Connections: connections,
				}

				dashboard.Rows[mark.Mark.Key] = append(dashboard.Rows[mark.Mark.Key], row)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   dashboard,
	})
}
