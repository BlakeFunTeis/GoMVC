package libraries

import (
	"gopkg.in/jdkato/prose.v2"
	"log"
	"strings"
)

func GetKeyword(text string) []string {
	doc, err := prose.NewDocument(text)
	temp1 := make([]string, 0)
	if err != nil {
		log.Println(err.Error())
		return temp1
	}

	array := make([]string, 0, len(doc.Entities()))
	temp := make(map[string]bool)
	count := make(map[string]int)
	for _, tok := range doc.Tokens() {
		if tok.Tag == "NN" {
			count[tok.Text]++
		}

		if count[tok.Text] >= 1 {
			if _, ok := temp[tok.Text]; !ok {
				temp[tok.Text] = true
				if strings.Contains(tok.Text, "Does") || strings.Contains(tok.Text, "Do") || strings.Contains(tok.Text, "Are") {
					continue
				}
				array = append(array, tok.Text)
			}
		}
	}

	for _, ent := range doc.Entities() {
		if _, ok := temp[ent.Text]; !ok {
			temp[ent.Text] = true
			if strings.Contains(ent.Text, "Does") || strings.Contains(ent.Text, "Do") || strings.Contains(ent.Text, "Are") {
				continue
			}
			array = append(array, ent.Text)
		}
	}

	return array
}