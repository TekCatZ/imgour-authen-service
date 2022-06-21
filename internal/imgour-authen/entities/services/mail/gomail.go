package mail

import (
	"bytes"
	"crypto/tls"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/supertokens/supertokens-golang/supertokens"
	"gopkg.in/gomail.v2"
	"html/template"
	"strconv"
)

var (
	dialer     *gomail.Dialer
	sendCloser *gomail.SendCloser
	from       string
)

type EmailHandler interface {
	Handle(email string, userInputCode, urlWithLinkCode *string, codeLifetime uint64,
		preAuthSessionId string, userContext supertokens.UserContext) error
}

type AuthMailParam struct {
	OTP string
	URL string
	TTL string
}

type authMailHandler struct{}

func (authMailHandler) Handle(email string, userInputCode, urlWithLinkCode *string, codeLifetime uint64,
	preAuthSessionId string, userContext supertokens.UserContext) error {
	// convert codeLifetime (milliseconds) to duration string with format "x minutes"
	duration := strconv.Itoa(int(codeLifetime/1000/60)) + " minutes"

	params := &AuthMailParam{
		OTP: *userInputCode,
		URL: *urlWithLinkCode,
		TTL: duration,
	}

	// read template from file in web/template
	t, err := template.ParseFiles("web/template/auth_mail.tmpl.html")
	if err != nil {
		return err
	}

	// execute template and write to buffer
	buf := new(bytes.Buffer)
	err = t.Execute(buf, params)
	if err != nil {
		return err
	}

	err = SendEmail(email, "Imgour Authen - Verify your account", buf.String())
	if err != nil {
		return err
	}

	return nil
}

func Setup(config configs.MailConfig) error {
	// parse config.Port to int and return error if not success
	port, err := strconv.Atoi(config.Port)
	if err != nil {
		return err
	}

	dialer = gomail.NewDialer(config.Server, port, config.Username, config.Password)
	s, err := dialer.Dial()
	if err != nil {
		return err
	}
	sendCloser = &s

	from = config.Username

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // disable this on prod

	return nil
}

func SendEmail(to, subject, content string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", content)
	return gomail.Send(*sendCloser, message)
}

func GetAuthEmailHandler() EmailHandler {
	return authMailHandler{}
}
