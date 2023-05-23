package dto

type MailData struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Name    string   `json:"name"`
	Action  string   `json:"action"`
	Url     string   `json:"url"`
}

type InviteMailData struct {
	To          []string `json:"to"`
	Subject     string   `json:"subject"`
	Name        string   `json:"name"`
	InvitedId   string   `json:"invitedId"`
	Url         string   `json:"url"`
	ProjectName string   `json:"projectName"`
	ProjectId   string   `json:"projectId"`
}
