package main

import (
	"fmt"
	"net/smtp"
)

// smtp server data to smtp server
type smtpServer struct {
	host string
	port string
}
// serverName URI to smtp server
func (s *smtpServer) serverName() string{
	return s.host + ":" + s.port
}

func main() {
	// sender data
	from := "kharansyah1@gmail.com"
	passApp := "migezuxnqwyblzsl"
	// receiver email address
	to := []string {
		"kharansyah.thawaddu@student.umn.ac.id",
		"rizkyzaldi@gmail.com",
		"wannabstart2020@gmail.com",
		"gerysantos03@gmail.com",
		"dekaprimatio@gmail.com",
	}
	// smtp server configuration
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// message
	message := []byte("Coba kirim email lewat smtp :D !")
	// Authentication
	auth := smtp.PlainAuth("", from, passApp, smtpServer.host)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
