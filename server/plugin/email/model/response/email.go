package response

type Email struct {
	To      string `json:"to"`      // Who to send the email to
	Subject string `json:"subject"` // Email title
	Body    string `json:"body"`    // Email content
}
