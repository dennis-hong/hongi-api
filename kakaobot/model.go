package kakaobot

type KakaoReq struct {
	UserRequest UserRequest `json:"userRequest"`
	UserAction Action `json:"action"`
}

type Action struct {
	Params map[string]string `json:"params"`
}

type KakaoRes struct {
	Contents []KakaoText `json:"contents"`
}

type KakaoText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type UserRequest struct {
	User      User   `json:"user"`
	ChatId    string `json:"chatId"`
	Utterance string `json:"utterance"`
}

type User struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Properties string `json:"properties"`
}