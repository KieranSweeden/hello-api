package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kieransweeden/hello/translation"
)

type Response struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")

	translation := translation.Translate(word, language)
	if translation == "" {
		language = ""
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Response{
		Language:    language,
		Translation: translation,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
