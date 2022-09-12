package util

import (
	"bytes"
	"html/template"
	"net/smtp"
	"okami.auth.backend/constanta"
)

var auth smtp.Auth

// created at 06-21-2022
func SendEmailGeneral(receiver []string, subject string, url string, username string) (result bool) {
	auth = smtp.PlainAuth("", constanta.EmailOkamiProject, constanta.EmailAppPassword, constanta.EmailHostGmail)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: username,
		URL:  url,
	}
	r := newRequest(receiver, subject, "")
	err := r.parseTemplate(constanta.EmailTemplatePath, templateData)
	if err = r.parseTemplate(constanta.EmailTemplatePath, templateData); err == nil {
		ok, _ := r.sendEmail()
		result = ok
	}
	return result
}

// created at 06-21-2022
//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

// created at 06-21-2022
func newRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

// created at 06-21-2022
func (r *Request) sendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := constanta.EmailHostGmailWithPort

	if err := smtp.SendMail(addr, auth, constanta.EmailOkamiProject, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

// created at 06-21-2022
func (r *Request) parseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
