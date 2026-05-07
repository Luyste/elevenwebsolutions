package handlers

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

type contactFormData struct {
	Name    string
	Email   string
	Message string
	Error   string
}

func Contact(c echo.Context) error {
	name := strings.TrimSpace(c.FormValue("name"))
	email := strings.TrimSpace(c.FormValue("email"))
	message := strings.TrimSpace(c.FormValue("message"))

	if name == "" || email == "" || message == "" {
		return c.Render(200, "contact_error", contactFormData{
			Name:    name,
			Email:   email,
			Message: message,
			Error:   "All fields are required.",
		})
	}

	log.Printf("Contact submission — name=%q email=%q message=%q", name, email, message)

	return c.Render(200, "contact_success", nil)
}
