package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

type EmailPayload struct {
	Sender struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"sender"`

	To []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`

	Subject string `json:"subject"`
	Html    string `json:"htmlContent"`
}

func SendEmail(toEmail string, toName string, subject string, content string) error {
	payload := EmailPayload{}
	payload.Sender.Name = os.Getenv("BREVO_SENDER_NAME")
	payload.Sender.Email = os.Getenv("BREVO_SENDER_EMAIL")
	payload.To = append(payload.To, struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}{
		Email: toEmail,
		Name:  toName,
	})
	payload.Subject = subject
	payload.Html = content

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := "https://api.brevo.com/v3/smtp/email"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", os.Getenv("BREVO_API_KEY"))
	req.Header.Set("content-type", "application/json")

	// Force IPv4
	dialer := &net.Dialer{FallbackDelay: -1}
	transport := &http.Transport{DialContext: dialer.DialContext}
	client := &http.Client{Transport: transport}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("BREVO STATUS:", resp.Status)
	fmt.Println("BREVO RESPONSE:", string(body))

	if resp.StatusCode >= 300 {
		return errors.New("failed send email")
	}
	return nil
}
