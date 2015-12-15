package goask

type AlexaOutputSpeechType string

const (
	OutputSpeechTypeText AlexaOutputSpeechType = "PlainText"
	OutputSpeechTypeSSML AlexaOutputSpeechType = "SSML"
)

type AlexaCardType string

const (
	CardTypeSimple AlexaCardType = "Simple"
	CardTypeLink   AlexaCardType = "LinkAccount"
)

type AlexaOutputSpeech struct {
	Type AlexaOutputSpeechType `json:"type,omitempty"`
	Text string                `json:"text,omitempty"`
	SSML string                `json:"ssml,omitempty"`
}

type AlexaCard struct {
	Type    AlexaCardType `json:"type,omitempty"`
	Title   string        `json:"title,omitempty"`
	Content string        `json:"content,omitempty"`
}

type AlexaResponse struct {
	ShouldEndSession bool               `json:"shouldEndSession"`
	OutputSpeech     *AlexaOutputSpeech `json:"outputSpeech,omitempty"`
	Card             *AlexaCard         `json:"card,omitempty"`
	Reprompt         *struct {
		OutputSpeech *AlexaOutputSpeech `json:"outputSpeech,omitempty"`
	} `json:"reprompt,omitempty"`
}

type AlexaResponseContainer struct {
	Version           string            `json:"version"`
	SessionAttributes SessionAttributes `json:"sessionAttributes,omitempty"`
	Response          AlexaResponse     `json:"response"`
}

func NewEmptyResponse() *AlexaResponseContainer {
	return &AlexaResponseContainer{
		Version: "1.0",
		Response: AlexaResponse{
			ShouldEndSession: true,
		},
	}
}
