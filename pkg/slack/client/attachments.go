package client

// Attachment ...
type Attachment struct {
	Fallback     string   `json:"fallback"`
	Color        string   `json:"color"`
	AuthorName   string   `json:"author_name"`
	Title        string   `json:"title"`
	Text         string   `json:"text"`
	ThumbnailURI string   `json:"thumb_url,omitempty"`
	MarkDownIn   []string `json:"mrkdwn_in,omitempty"`
	Footer       string   `json:"footer,omitempty"`
	FooterIcon   string   `json:"footer_icon,omitempty"`
	Timestamp    int64    `json:"ts"`
}

// Attachments , list of attachment
type Attachments []*Attachment

// Add ...
func (atts *Attachments) Add(att *Attachment) {
	*atts = append(*atts, att)
}

// NewAttachment ...
func NewAttachment() *Attachment {
	return &Attachment{
		// AuthorName: "author",
		Color: "##dddddd",
		// Title:      title,
		// Text:       fmt.Sprintf("```%s```", string(jsonBytes)),
		// Footer:     footer,
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		MarkDownIn: []string{"text"},
	}
}
