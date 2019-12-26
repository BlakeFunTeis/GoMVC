package libraries

import (
    "bytes"
    "html/template"
    "log"
    "net/smtp"
)

type EmailRequest struct {
    Title string
    content string
    To string
    Subject string
    View string
    Data interface{}
}
//signature
var account = "xxxxxx"

var password = "xxxxxx"

var auth smtp.Auth

func (request *EmailRequest) GoogleEmailSend() {

    auth = smtp.PlainAuth("", account, password, "smtp.gmail.com")
    if err := request.parseTemplate(request.View, request.Data); err == nil {
        result, _ := request.sendEmail()
        log.Println(result)
    } else {
        log.Println(err)
    }
}

func (request *EmailRequest) parseTemplate(templateFile string, data interface{}) error {
    var err error
    view, err := template.ParseFiles(templateFile)
    if err != nil {
        return err
    }
    buf := new(bytes.Buffer)
    if err = view.Execute(buf, data); err != nil {
        return err
    }
    request.content = buf.String()
    return nil
}

func (request *EmailRequest) sendEmail() (bool, error) {
    log.Println(request.Subject)
    mime := "Message-ID: "+MakeMd5(MakeSnowId().String())+"\nReply-To: "+account+"\nFrom: Journeyfoxx <"+account+">"+"\nMIME-version: 1.0;\nTo: "+request.To+"\nContent-Type: text/html; charset=\"UTF-8\";\n"
    subject := "Subject: " + request.Subject + "\n"
    msg := []byte(subject + mime + "\n" + request.content)
    addr := "smtp.gmail.com:587"

    if err := smtp.SendMail(addr, auth, account, []string{request.To}, msg); err != nil {
        log.Println(err.Error())
        return false, err
    }

    return true, nil
}

func EmailNewRequest(to, title, view string, data interface{}) *EmailRequest {
    return &EmailRequest{
        To: to,
        Subject: title,
        View: view,
        Data: data,
    }
}