package contactform

import (
	_ "embed"
	"html/template"
	"net/mail"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

type NotifyData struct {
	AppURL  string
	Sender  *mail.Address
	Message string
}

//go:embed notify.html.gotmpl
var notifyTmpl string

func Notify(app *pocketbase.PocketBase) func(e *core.ModelEvent) error {
	return func(e *core.ModelEvent) error {
		users, err := app.FindAllRecords("users")
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

		t, err := template.New("").Parse(notifyTmpl)
		if err != nil {
			return err
		}

		record := e.Model.(*core.Record)

		var buf strings.Builder
		data := NotifyData{
			AppURL: app.Settings().Meta.AppURL,
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

		if err := app.NewMailClient().Send(message); err != nil {
			return err
		}

		return e.Next()
	}
}
