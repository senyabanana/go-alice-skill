package models

const (
	TypeSimpleUtterance = "SimpleUtterance"
)

// Request описывает запрос пользователя.
type Request struct {
	// тут будет, например, строка "Europe/Moscow" для часового пояса Москвы
	Timezone string          `json:"timezone"`
	Request  SimpleUtterance `json:"request"`
	Session  Session         `json:"session"`
	Version  string          `json:"version"`
}

type Session struct {
	New bool `json:"new"`
}

// SimpleUtterance описывает команду, полученную в запросе типа SimpleUtterance.
type SimpleUtterance struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}

// Response описывает ответ сервера.
type Response struct {
	Response ResponsePayload `json:"response"`
	Version  string          `json:"version"`
}

// ResponsePayload описывает ответ, который нужно озвучить.
type ResponsePayload struct {
	Text string `json:"text"`
}
