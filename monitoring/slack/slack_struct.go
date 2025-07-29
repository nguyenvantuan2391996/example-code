package slack

import "fmt"

const (
	FormatTextLinkSlack = "<%s|%s>"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	ImageUrl string  `json:"image_url,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type SlackMessage struct {
	Text        string        `json:"text"`
	IconEmoji   string        `json:"icon_emoji"`
	Attachments []*Attachment `json:"attachments"`
}

type TextMessageNotifyRun struct {
	CurrentTime string `json:"currentTime"`
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	IsRunning   string `json:"isRunning"`
	Note        string `json:"note"`
}

type TextMessageNotifySummary struct {
	CurrentTime      string `json:"currentTime"`
	SportType        string `json:"sport_type"`
	Name             string `json:"name"`
	Distance         string `json:"distance"`
	MovingTime       string `json:"moving_time"`
	AverageSpeed     string `json:"average_speed"`
	MaxSpeed         string `json:"max_speed"`
	AverageHeartrate string `json:"average_heartrate"`
	MaxHeartrate     string `json:"max_heartrate"`
	Note             string `json:"note"`
}

type TextMessageNotifyStatistical struct {
	ImageUrl string `json:"image_url"`
}

type TextMessageDailyCodingChallenge struct {
	Date       string `json:"date"`
	Difficulty string `json:"difficulty"`
	TopicTags  string `json:"topic_tags"`
	Link       string `json:"link"`
	Title      string `json:"title"`
}

func (s *TextMessageNotifyRun) ToAttachment() []*Attachment {
	fields := make([]Field, 0)
	fields = append(fields, Field{
		Short: true,
		Title: "Current Time",
		Value: s.CurrentTime,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Location",
		Value: s.Location,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Temperature",
		Value: s.Temperature,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Weather",
		Value: s.Weather,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Running",
		Value: s.IsRunning,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Note",
		Value: s.Note,
	})

	return []*Attachment{{Fields: fields}}
}

func (s *TextMessageNotifySummary) ToAttachment() []*Attachment {
	fields := make([]Field, 0)
	fields = append(fields, Field{
		Short: true,
		Title: "Current Time",
		Value: s.CurrentTime,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Sport Type",
		Value: s.SportType,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Name",
		Value: s.Name,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Kilometers",
		Value: s.Distance,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Average Speed",
		Value: s.AverageSpeed,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Max Speed",
		Value: s.MaxSpeed,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Average Heartrate",
		Value: s.AverageHeartrate,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Max Heartrate",
		Value: s.MaxHeartrate,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Moving Time",
		Value: s.MovingTime,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Note",
		Value: s.Note,
	})

	return []*Attachment{{Fields: fields}}
}

func (s *TextMessageNotifyStatistical) ToAttachment() []*Attachment {
	return []*Attachment{{ImageUrl: s.ImageUrl}}
}

func (s *TextMessageDailyCodingChallenge) ToAttachment() []*Attachment {
	fields := make([]Field, 0)
	fields = append(fields, Field{
		Short: true,
		Title: "Date",
		Value: s.Date,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Difficulty",
		Value: s.Difficulty,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Topic Tags",
		Value: s.TopicTags,
	})
	fields = append(fields, Field{
		Short: true,
		Title: "Link",
		Value: fmt.Sprintf(FormatTextLinkSlack, s.Link, s.Title),
	})

	return []*Attachment{{Fields: fields}}
}