package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/gomail.v2"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var Validate = validator.New()


func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func SendEmail(user_email string) {

	email := "matcherx1337@gmail.com";
	email_pass := os.Getenv("EMAIL_PASS");
	
	mail := gomail.NewMessage();
	mail.SetHeader("From", email);
	mail.SetHeader("To", user_email);
	mail.SetHeader("Subject", "MatcherX account verification");

	body := fmt.Sprintf(`<div><a href="%s"><b>Clicki 3la had lb3ar!</b></a> <br> <img src="%s" alt="img" /></div>`, "https://abder.vercel.app", "https://media.makeameme.org/created/fact-no-verification.jpg");
	mail.SetBody("text/html", body);

	d := gomail.NewDialer("smtp.gmail.com", 587, email, email_pass);
	if err := d.DialAndSend(mail); err != nil {
		log.Print(err);
	}
	
}