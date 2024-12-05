package email

import (
	"fmt"
	"net/smtp"
)

func SendVerificationCode(toEmail, code string) error {
	from := "your_email@example.com"  // Ваш email
	password := "your_email_password" // Пароль от почты
	smtpHost := "smtp.example.com"    // SMTP хост
	smtpPort := "587"                 // Порт для SMTP

	// Сообщение для отправки
	subject := "Your Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", code)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	// Авторизация и отправка
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	return err
}
