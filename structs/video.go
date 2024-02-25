package structs

type VideoPayload struct {
	Email   string `json:"email"`
	Video   string `json:"video"`
	Audio   string `json:"audio"`
	Text    string `json:"text"`
	Summary string `json:"summary"`
}
