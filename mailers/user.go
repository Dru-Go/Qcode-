package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

// SendUser sends message to a user
func SendUser(from, message string, to []string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = message
	m.From = from
	m.To = to
	err := m.AddBody(r.HTML("user.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
