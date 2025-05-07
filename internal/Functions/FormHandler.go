package BA

import (
	"html/template"
	"net/http"
)

type PageData struct {
	AsciiArt  string
	UserInput string
	Banner    string
	FontSize  string
	FontColor string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("userinput")
		inputStyle := r.FormValue("banner")
		fontSize := r.FormValue("fontsize")
		fontColor := r.FormValue("color")

		fontMap := LoadFont(inputStyle)
		asciiResult, err := PrintAscii(input, fontMap)
		if err != nil {
			http.Error(w, "HTTP Status 500 - Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("../internal/frontend/MainPage.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		data := PageData{
			AsciiArt:  asciiResult,
			UserInput: input,
			Banner:    inputStyle,
			FontSize:  fontSize,
			FontColor: fontColor,
		}
		tmpl.Execute(w, data)
		return
	}

	// Handle initial GET request with default values
	data := PageData{
		FontSize:  "16px",
		FontColor: "#00ffcc",
		Banner:    "standard",
	}
	tmpl, err := template.ParseFiles("../internal/frontend/MainPage.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
