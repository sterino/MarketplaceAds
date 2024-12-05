package email

import (
	"fmt"
	"net/smtp"
)

func SendVerificationCode(toEmail, code string) error {
	from := "egore.chip@example.com" // Ваш email
	password := "khasenov2003"       // Пароль от почты
	smtpHost := "smtp.gmail.com"     // SMTP хост
	smtpPort := "587"                // Порт для SMTP

	// Сообщение для отправки
	subject := "Your Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", code)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	// Авторизация и отправка
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	return err
}
