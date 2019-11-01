package libraries

import (
    "bytes"
    "fmt"
    "html/template"
    "net/smtp"
)

type EmailContent struct {
    Title string
    Content string
    To string
    Views string
}

func GoogleEmailSend(setting EmailContent) {
    username := "#########"
    password := "#########"
    gmailAuth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
    temp, _ := template.ParseFiles(setting.Views)

    var body bytes.Buffer
    var header = "MIME-version: 1.0;\nContent-type: text/html;"
    body.Write([]byte(fmt.Sprintf("Subject:%s\n%s\n\n", setting.Title, header)))

    var content = struct{ token string }{
        token: setting.Content,
    }

    _ = temp.Execute(&body, content)
    var err error
    err = smtp.SendMail("smtp.gmail.com:587", gmailAuth, username, []string{setting.To}, body.Bytes())
    fmt.Print(err)
}
