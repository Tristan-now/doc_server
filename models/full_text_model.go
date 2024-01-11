package models

type ESIndexInterFace interface {
	Index() string
	Mapping() string
}

// FullTextModel es全文搜索的model
type FullTextModel struct {
	DocID uint   `json:"docID"`
	ID    string `json:"id"`    // es的id
	Title string `json:"title"` // 标题
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 跳转地址，可以由docID和title拼接而来
}

func (FullTextModel) Index() string {
	return "gvd_server_full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "body": { 
        "type": "text" 
      },
      "title": {
        "type": "text",
	    "fields": {
          "keyword": {
              "type": "keyword",
			  "ignore_above": 256
          }
        }
	  },
	  "slug": { 
        "type": "keyword" 
      },
      "docID": {
        "type": "integer"
      }
    }
  }
}
`
}
