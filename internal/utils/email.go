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

func SendEmail(
	toEmail string,
	toName string,
	subject string,
	content string,
) error {

	payload := EmailPayload{}

	payload.Sender.Name =
		os.Getenv(
			"BREVO_SENDER_NAME",
		)

	payload.Sender.Email =
		os.Getenv(
			"BREVO_SENDER_EMAIL",
		)

	payload.To = append(
		payload.To,
		struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			Email: toEmail,
			Name:  toName,
		},
	)

	payload.Subject =
		subject

	payload.Html =
		content

	jsonData, err :=
		json.Marshal(
			payload,
		)

	if err != nil {
		return err
	}

	req, err :=
		http.NewRequest(
			"POST",
			"https://api.brevo.com/v3/smtp/email",
			bytes.NewBuffer(
				jsonData,
			),
		)

	if err != nil {
		return err
	}

	req.Header.Set(
		"accept",
		"application/json",
	)

	req.Header.Set(
		"api-key",
		os.Getenv(
			"BREVO_API_KEY",
		),
	)

	req.Header.Set(
		"content-type",
		"application/json",
	)

	// Force IPv4
	transport :=
		&http.Transport{
			DialContext: (&net.Dialer{
				FallbackDelay: -1,
			}).DialContext,
		}

	client :=
		&http.Client{
			Transport: transport,
		}

	resp, err :=
		client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ :=
		io.ReadAll(
			resp.Body,
		)

	fmt.Println(
		"BREVO STATUS:",
		resp.Status,
	)

	fmt.Println(
		"BREVO RESPONSE:",
		string(body),
	)

	if resp.StatusCode >=
		300 {

		return errors.New(
			"failed send email",
		)
	}

	return nil
}
