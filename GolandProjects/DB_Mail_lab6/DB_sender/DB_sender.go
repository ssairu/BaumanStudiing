package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
)

const smtpHost = "mail.nic.ru"
const smtpPort = 465
const username = "dts21@dactyl.su"
const password = "12345678990DactylSUDTS"

var tempMessage = "<!doctype html><html xmlns=\"http://www.w3.org/1999/xhtml\" xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:o=\"urn:schemas-microsoft-com:office:office\"><head><title></title><!--[if !mso]><!--><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><!--<![endif]--><meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\"><meta name=\"viewport\" content=\"width=device-width,initial-scale=1\"><style type=\"text/css\">#outlook a { padding:0; }\n          body { margin:0;padding:0;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%; }\n          table, td { border-collapse:collapse;mso-table-lspace:0pt;mso-table-rspace:0pt; }\n          img { border:0;height:auto;line-height:100%; outline:none;text-decoration:none;-ms-interpolation-mode:bicubic; }\n          p { display:block;margin:13px 0; }</style><!--[if mso]>\n        <noscript>\n        <xml>\n        <o:OfficeDocumentSettings>\n          <o:AllowPNG/>\n          <o:PixelsPerInch>96</o:PixelsPerInch>\n        </o:OfficeDocumentSettings>\n        </xml>\n        </noscript>\n        <![endif]--><!--[if lte mso 11]>\n        <style type=\"text/css\">\n          .mj-outlook-group-fix { width:100% !important; }\n        </style>\n        <![endif]--><!--[if !mso]><!--><link href=\"https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700\" rel=\"stylesheet\" type=\"text/css\"><style type=\"text/css\">@import url(https://fonts.googleapis.com/css?family=Ubuntu:300,400,500,700);</style><!--<![endif]--><style type=\"text/css\">@media only screen and (min-width:480px) {\n        .mj-column-per-100 { width:100% !important; max-width: 100%; }\n      }</style><style media=\"screen and (min-width:480px)\">.moz-text-html .mj-column-per-100 { width:100% !important; max-width: 100%; }</style><style type=\"text/css\"></style></head><body style=\"word-spacing:normal;background-color:#00FFFF;\"><div style=\"background-color:#00FFFF;\"><!--[if mso | IE]><table align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" class=\"\" style=\"width:600px;\" width=\"600\" ><tr><td style=\"line-height:0px;font-size:0px;mso-line-height-rule:exactly;\"><![endif]--><div style=\"margin:0px auto;max-width:600px;\"><table align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"width:100%;\"><tbody><tr><td style=\"direction:ltr;font-size:0px;padding:20px 0;text-align:center;\"><!--[if mso | IE]><table role=\"presentation\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\"><tr><td class=\"\" style=\"vertical-align:top;width:600px;\" ><![endif]--><div class=\"mj-column-per-100 mj-outlook-group-fix\" style=\"font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;\"><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"vertical-align:top;\" width=\"100%\"><tbody><tr><td align=\"left\" style=\"font-size:0px;padding:10px 25px;word-break:break-word;\"><div style=\"font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;font-weight:bold;line-height:1;text-align:left;color:#000000;\">HELLO</div></td></tr><tr><td align=\"left\" style=\"font-size:0px;padding:10px 25px;word-break:break-word;\"><div style=\"font-family:Ubuntu, Helvetica, Arial, sans-serif;font-size:13px;font-style:italic;line-height:1;text-align:left;color:#000000;\">MESSAGE</div></td></tr></tbody></table></div><!--[if mso | IE]></td></tr></table><![endif]--></td></tr></tbody></table></div><!--[if mso | IE]></td></tr></table><![endif]--></div></body></html>\n"

type Item struct {
	Name    string
	Email   string
	Message string
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func sendMes(person Item) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", smtpHost, smtpPort), tlsConfig)
	if err != nil {
		fmt.Println("Error connecting to SMTP server:", err)
		os.Exit(1)
	}

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

	if err := client.Mail(username); err != nil {
		fmt.Println("Error setting sender:", err)
		os.Exit(1)
	}

	if err := client.Rcpt(person.Email); err != nil {
		fmt.Println("Error setting recipient:", err)
		os.Exit(1)
	}

	w, err := client.Data()
	if err != nil {
		fmt.Println("Error opening data connection:", err)
		os.Exit(1)
	}
	defer w.Close()

	messageBody := strings.Replace(tempMessage, "HELLO", "Hello, "+person.Name, -1)
	messageBody = strings.Replace(messageBody, "MESSAGE", person.Message, -1)
	subject := "Test"
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html;\r\n%s", person.Email, subject, messageBody)
	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing message:", err)
		os.Exit(1)
	}
	fmt.Println("Email sent successfully!")
	client.Quit()
}

func main() {
	fmt.Println("start connect to bd")
	db, err := sql.Open("mysql", "iu9networkslabs:Je2dTYr6@tcp(students.yss.su)/iu9networkslabs")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("end connect to bd")
	fmt.Println("start get from bd")
	rows, err := db.Query("select * from iu9networkslabs.Penkin_mail_sender")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("end get from bd")

	Items := []Item{}

	for rows.Next() {
		p := Item{}
		var id int32
		err := rows.Scan(&id, &p.Name, &p.Email, &p.Message)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Items = append(Items, p)
	}

	for {
		for _, person := range Items {
			sendMes(person)
		}
		time.Sleep(time.Duration(3*randInt(4, 9)) * time.Second)
	}
}
