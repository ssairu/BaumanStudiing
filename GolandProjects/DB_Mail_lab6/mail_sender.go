package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <to_email> <subject> <message_body>")
		os.Exit(1)
	}

	// Параметры для подключения к SMTP серверу
	smtpHost := "mail.nic.ru"
	smtpPort := 465
	username := "dts21@dactyl.su"
	password := "12345678990DactylSUDTS"
	toEmail := os.Args[1]
	subject := os.Args[2]
	messageBody := os.Args[3]

	// Формирование сообщения
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", toEmail, subject, messageBody)

	// Настройка подключения с поддержкой SSL
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Настройте это правильно для безопасного использования в боевом режиме
		ServerName:         smtpHost,
	}

	// Подключение к SMTP серверу
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpHost, smtpPort), tlsConfig)
	if err != nil {
		fmt.Println("Error connecting to SMTP server:", err)
		os.Exit(1)
	}

	// Авторизация на SMTP сервере
	auth := smtp.PlainAuth("", username, password, smtpHost)
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		os.Exit(1)
	}

	if err := client.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		os.Exit(1)
	}

	// Отправка сообщения
	if err := client.Mail(username); err != nil {
		fmt.Println("Error setting sender:", err)
		os.Exit(1)
	}

	if err := client.Rcpt(toEmail); err != nil {
		fmt.Println("Error setting recipient:", err)
		os.Exit(1)
	}

	w, err := client.Data()
	if err != nil {
		fmt.Println("Error opening data connection:", err)
		os.Exit(1)
	}
	defer w.Close()

	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing message:", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully!")

	// Завершение соединения
	client.Quit()
}
