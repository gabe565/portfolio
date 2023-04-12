package contact_form

import (
	_ "embed"
	"html/template"
	"net/mail"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

type NotifyData struct {
	AppUrl  string
	Sender  *mail.Address
	Message string
}

//go:embed notify.html.gotmpl
var notifyTmpl string

func Notify(app *pocketbase.PocketBase) func(e *core.ModelEvent) error {
	return func(e *core.ModelEvent) error {
		users, err := app.Dao().FindRecordsByExpr("users")
		if err != nil {
			return err
		}

		emails := make([]mail.Address, 0, len(users))
		for _, user := range users {
			emails = append(emails, mail.Address{
				Name:    user.GetString("name"),
				Address: user.Email(),
			})
		}

		t, err := template.New("contact_form_notify").Parse(notifyTmpl)
		if err != nil {
			return err
		}

		record := e.Model.(*models.Record)

		var buf strings.Builder
		data := NotifyData{
			AppUrl: app.Settings().Meta.AppUrl,
			Sender: &mail.Address{
				Name:    record.GetString("name"),
				Address: record.GetString("email"),
			},
			Message: record.GetString("message"),
		}
		if err := t.Execute(&buf, data); err != nil {
			return err
		}

		message := &mailer.Message{
			From: mail.Address{
				Address: app.Settings().Meta.SenderAddress,
				Name:    app.Settings().Meta.SenderName,
			},
			Bcc: emails,
			Headers: map[string]string{
				"Reply-To": data.Sender.String(),
			},
			Subject: "Contact Form Submission From " + data.Sender.Name,
			HTML:    buf.String(),
		}

		return app.NewMailClient().Send(message)
	}
}
