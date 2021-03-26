package plugins

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html"
	"reflect"
	"strconv"
	"strings"
	"tenet/lib/db"
	"tenet/lib/model"
	"tenet/lib/utlis"
	"time"
)

type AddDatabaseData struct {
	Link    string
	Content string
}

func (d *AddDatabaseData) Do(c *gin.Context, content string) (string, bool) {
	// 从数据库中拿出来数据，并替换内容
	var bd bytes.Buffer
	doc, err := html.Parse(bytes.NewReader([]byte(content)))
	if err != nil {
		return content, true
	}

	err = html.Render(&bd, doc)
	if err != nil {
		return content, true
	}
	back := bd.String()
	recursionGetDataUseMarkInfo(doc, d.Link, &back)
	return back, true
}

func (d *AddDatabaseData) Init(c *gin.Context) {
	d.Link = fmt.Sprintf("http://%s%s", c.Request.Host, c.Request.RequestURI)
}

func recursionGetDataUseMarkInfo(node *html.Node, link string, back *string) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if attr.Key == "data-re" {
				var b bytes.Buffer
				err := html.Render(&b, node)
				if err != nil {
					continue
				}

				dataRe := attr.Val
				mark, err := utlis.GetMarksFrom(dataRe, link)
				if err != nil {
					continue
				}

				var replace []string
				contents := getContentFrom(mark)
				for _, content := range contents {
					item := b.String()
					item = strings.Replace(item, fmt.Sprintf(`data-re="%s"`, html.EscapeString(dataRe)), "", 1)
					item = strings.Replace(item, fmt.Sprintf(`data-re='%s'`, html.EscapeString(dataRe)), "", 1)

					data := dbContentToString(content)
					markType := reflect.TypeOf(mark)
					markVal := reflect.ValueOf(mark)

					for i := 0; i < markVal.NumField(); i++ {
						tagVal := markType.Field(i).Tag.Get("json")
						if markVal.Field(i).String() != "" && tagVal != "" {
							if val, ok := data[model.MCMap[tagVal]]; ok {
								item = strings.ReplaceAll(item, markVal.Field(i).String(), val)
							}
						}
					}

					replace = append(replace, item)
				}

				if len(replace) > 0 {
					*back = strings.Replace(*back, b.String(), strings.Join(replace, "\n"), -1)
				}
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		recursionGetDataUseMarkInfo(c, link, back)
	}
}

func getContentFrom(mark model.Mark) []model.Content {
	var contents []model.Content

	defer func() {
		utlis.ReverseAny(contents)
	}()

	page := mark.Page
	pageSize := mark.PageSize

	switch mark.RealID {
	case 0:
		err := db.DB.Limit(pageSize).Offset(utlis.Paginate(page, pageSize)).Where("key = ?", mark.Key).
			Order("created_at DESC").Find(&contents).Error
		if err == gorm.ErrRecordNotFound {
			return contents
		}
	case -1:
		err := db.DB.Limit(pageSize).Offset(utlis.Paginate(page, pageSize)).Where("key = ?", mark.Key).
			Order("created_at DESC").Find(&contents).Error
		if err == gorm.ErrRecordNotFound {
			return contents
		}
	default:
		err := db.DB.Limit(pageSize).Offset(utlis.Paginate(page, pageSize)).Where("id = ?", mark.RealID).
			Order("created_at DESC").Find(&contents).Error
		if err == gorm.ErrRecordNotFound {
			return contents
		}
	}

	return contents
}

func dbContentToString(content model.Content) map[string]string {
	data := map[string]string{}

	contentType := reflect.TypeOf(content)
	contentVal := reflect.ValueOf(content)

	for i := 0; i < contentVal.NumField(); i++ {
		baseTagVal := contentType.Field(i).Tag.Get("json")
		if contentVal.Field(i).String() != "" && baseTagVal != "" {
			switch contentVal.Field(i).Type().String() {
			case "time.Time":
				data[baseTagVal] = contentVal.Field(i).Interface().(time.Time).Format("2006-01-02")
			case "int64":
				data[baseTagVal] = strconv.Itoa(int(contentVal.Field(i).Interface().(int64)))
			case "*time.Time":
				if !utlis.IsNil(contentVal.Field(i).Interface()) {
					data[baseTagVal] = contentVal.Field(i).Interface().(*time.Time).Format("2006-01-02")
				} else {
					data[baseTagVal] = ""
				}
			default:
				data[baseTagVal] = contentVal.Field(i).String()
			}
		}
	}
	return data
}
