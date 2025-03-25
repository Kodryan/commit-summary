package client

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Store    bool      `json:"store"`
	Messages []Message `json:"messages"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
}

type Message struct {
	Role        string   `json:"role"`
	Content     string   `json:"content"`
	Refusal     *string  `json:"refusal"`
	Annotations []string `json:"annotations"`
}