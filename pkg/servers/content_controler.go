package servers

import (
	"encoding/json"
	"fmt"
	"github.com/Garbrandt/tenet/pkg/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (s *Server) DeleteContent(c *gin.Context) {
	id := c.Param("id")
	s.DB.Where("id = ?", id).Delete(&model.Content{})
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", id)})
}

func getFormFrom(c *gin.Context) ([]model.Form, error) {
	var forms []model.Form

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err.Error(),
		})
		return forms, err
	}

	err = json.Unmarshal(body, &forms)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err.Error(),
		})
		return forms, err
	}

	return forms, nil
}

func (s *Server) CreateContent(c *gin.Context) {
	forms, err := getFormFrom(c)
	if err != nil {
		return
	}

	for _, form := range forms {
		form.Submit.Content.Key = form.Mark.Key

		if form.Submit.Content.ID == 0 {
			if form.Mark.Multiple {
				s.DB.Model(&model.Content{}).Create(&form.Submit.Content)
			} else {
				db := s.DB.Model(&model.Content{}).Where("key = ?", form.Mark.Key).FirstOrCreate(&form.Submit.Content)
				if err := db.Error; err != nil {
					continue
				}
			}
		} else {
			if s.DB.Model(&model.Content{}).Where("id = ?", form.Submit.Content.ID).Updates(&form.Submit.Content).RowsAffected == 0 {
				s.DB.Create(&form.Submit.Content)
			}
		}

		for key, connections := range form.Submit.Connections {
			s.DB.Debug().Where("content_id = ? AND connection_content_type = ?", form.Submit.Content.ID, key).Delete(&[]model.Connection{})
			for _, connection := range connections {
				connection := model.Connection{
					Content:               model.Content{Type: connection.Type},
					ContentId:             form.Submit.Content.ID,
					ContentType:           form.Submit.Content.Key,
					ConnectionContentId:   connection.ID,
					ConnectionContentType: connection.Key,
				}

				db := s.DB.Model(&model.Connection{}).
					Where("content_id = ? AND connection_content_id = ?",
						connection.ContentId,
						connection.ConnectionContentId).FirstOrCreate(&connection)
				if err := db.Error; err != nil {
					continue
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
