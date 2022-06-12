package mail

import (
	"crypto/tls"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/supertokens/supertokens-golang/supertokens"
	"gopkg.in/gomail.v2"
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

type authMailHandler struct{}

func (authMailHandler) Handle(email string, userInputCode, urlWithLinkCode *string, codeLifetime uint64,
	preAuthSessionId string, userContext supertokens.UserContext) error {
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
