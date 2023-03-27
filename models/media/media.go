package media

type Media struct {
	ID           uint   `json:"id"`
	TestResultID uint   `json:"test_result_id"`
	URL          string `json:"url"`
	Type         string `json:"type"`
}
