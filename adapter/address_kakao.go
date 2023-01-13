package adapter

import (
	"context"
	"html/template"
	"net/http"
)

type addressSvcKakao struct {
	htmlFile     string
	htmlTemplate *template.Template
}

func NewAddressSvcKakao(htmlFile string) *addressSvcKakao {
	return &addressSvcKakao{
		htmlFile:     htmlFile,
		htmlTemplate: template.Must(template.ParseFiles(htmlFile)),
	}
}

func (a *addressSvcKakao) FindAddress(ctx context.Context, w http.ResponseWriter) error {
	return a.htmlTemplate.Execute(w, nil)
}
