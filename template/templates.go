package template

type ListItem struct {
	Image       TemplateImage `json:"image"`
	Token       string        `json:token`
	TextContent TemplateText  `json:textContent`
}

type TemplateText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ImageSource struct {
	Url          string `json:"url"`
	WidthPixels  int    `json:"widthPixels"`
	HeightPixels int    `json:"heightPixels"`
	Size         string `json:"size"`
}

type TemplateImage struct {
	ContentDescription string        `json:"contentDescription"`
	Sources            []ImageSource `json:"sources"`
}
