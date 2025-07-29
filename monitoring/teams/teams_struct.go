package teams

const (
	MaxRetryTimes           = 5
	TypeMessageCard         = "MessageCard"
	TypeAdaptiveCard        = "AdaptiveCard"
	TypeMessage             = "message"
	TypeImage               = "Image"
	TypeTextBlock           = "TextBlock"
	ContentTypeAdaptiveCard = "application/vnd.microsoft.card.adaptive"
	SizeMedium              = "Medium"
	ActivityImage           = "https://adaptivecards.io/content/cats/3.png"
)

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Section struct {
	ActivityTitle    string  `json:"activityTitle"`
	ActivitySubtitle string  `json:"activitySubtitle"`
	ActivityImage    string  `json:"activityImage"`
	Facts            []*Fact `json:"facts"`
	Markdown         bool    `json:"markdown"`
}

type TemplateTeamsMessage struct {
	Type     string     `json:"@type"`
	Summary  string     `json:"summary"`
	Sections []*Section `json:"sections"`
}

type Body struct {
	Type string `json:"type"`
	URL  string `json:"url,omitempty"`
	Size string `json:"size,omitempty"`
	Text string `json:"text"`
	Wrap bool   `json:"wrap"`
}

type Content struct {
	Type string  `json:"type"`
	Body []*Body `json:"body"`
}

type Attachment struct {
	Content     *Content `json:"content"`
	ContentType string   `json:"contentType"`
}

type TemplateTeamsMessageV2 struct {
	Type        string        `json:"type"`
	Summary     string        `json:"summary"`
	Attachments []*Attachment `json:"attachments"`
}
