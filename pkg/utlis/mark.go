package utlis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Garbrandt/tenet/pkg/model"
	"github.com/joncalhoun/qson"
	"golang.org/x/net/html"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func GetMarksForm(content []byte) (marks []model.Mark) {
	marks = make([]model.Mark, 0)

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return marks
	}

	recursion(doc, &marks)
	return marks
}

func recursion(node *html.Node, marks *[]model.Mark) {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if attr.Key == "data-re" {
				var b bytes.Buffer
				err := html.Render(&b, node)
				if err != nil {
					continue
				}

				dataRe := attr.Val
				mark, err := GetMarksFrom(dataRe, "")
				if err != nil {
					return
				}
				if mark.ShowOnDashboard {
					*marks = append(*marks, mark)
				}
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		recursion(c, marks)
	}
}

func stringToMark(dataRe string) model.Mark {
	mark := model.Mark{}

	if strings.Contains(dataRe, "http") {
		var data []string
		r, _ := url.Parse(dataRe)
		query := r.Query()
		for key, values := range query {
			for _, value := range values {
				data = append(data, fmt.Sprintf("%s=%s", key, value))
			}
		}
		if len(data) > 0 {
			dataRe = strings.Join(data, "&")
		}
	}

	b, err := qson.ToJSON(dataRe)
	if err != nil {
		log.Println(err, dataRe)
	}
	err = json.Unmarshal(b, &mark)
	if err != nil {
		log.Println(err, dataRe)
	}

	return mark
}

func GetMarksFrom(dataRe string, link string) (model.Mark, error) {
	mark := stringToMark(dataRe)
	linkQuery, err := url.ParseQuery(link)
	if err != nil {
		return mark, err
	}

	if len(linkQuery["content_id"]) > 0 && mark.RealID != -1 {
		mark.RealID, _ = strconv.Atoi(linkQuery["content_id"][0])
	}

	markQuery, err := url.ParseQuery(dataRe)
	if err != nil {
		return mark, err
	}

	for _, value := range markQuery["connections"] {
		keys := strings.Split(value, "#")
		mark.Relations = append(mark.Relations, model.Mark{
			Type:    keys[0],
			Section: keys[1],
			Env:     keys[2],
			Key:     fmt.Sprintf("%s#%s", keys[1], keys[2]),
		})
	}

	mark = setMark(mark, linkQuery)
	return mark, nil
}

func setMark(mark model.Mark, values url.Values) model.Mark {
	if strings.TrimSpace(mark.Env) == "" && len(values["backend_env"]) > 0 {
		mark.Env = values["backend_env"][0]
	}

	if strings.TrimSpace(mark.Section) == "" && len(values["backend_section"]) > 0 {
		mark.Section = values["backend_section"][0]
	}

	if strings.TrimSpace(mark.Key) == "" {
		mark.Key = mark.Section + "#" + mark.Env
	}
	return mark
}
