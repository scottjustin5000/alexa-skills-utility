package template

type TemplateBuilder interface {
	SetType(string) TemplateBuilder
	SetTextContent(string, string, string) TemplateBuilder
	SetImage(TemplateImage) TemplateBuilder
	AddListItem(TemplateImage, string, string, string, string) TemplateBuilder
	SetTitle(string) TemplateBuilder
	SetToken(string) TemplateBuilder
	SetBackgroundImage(TemplateImage) TemplateBuilder
	SetBackButtonBehavior(string) TemplateBuilder
	Build() Template
}

type tempateBuilder struct {
	Type            string
	Title           string
	TextContent     TemplateText
	Image           TemplateImage
	Items           []ListItem
	ListItems       []ListItem
	BackButton      string
	BackgroundImage TemplateImage
}
