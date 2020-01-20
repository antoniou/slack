package slack

// https://api.slack.com/reference/block-kit/blocks#input

const (
	IBEPlainText   InputBlockElementType = "plain_text_input"
	IBESelect      InputBlockElementType = "static_select"
	IBEMultiSelect InputBlockElementType = "multi_static_select"
	IBEDatePicker  InputBlockElementType = "datepicker"
)

type InputBlock struct {
	Type      MessageBlockType   `json:"type"`
	Label     *TextBlockObject   `json:"label,omitempty"`
	Element   *InputBlockElement `json:"element,omitempty"`
	Text      string             `json:text`
	Multiline bool               `json:multiline`
}

// BlockType returns the type of the block
func (i InputBlock) BlockType() MessageBlockType {
	return i.Type
}

func NewInputBlock(text string, multiline bool) *InputBlock {
	return &InputBlock{
		Type:      MBTInput,
		Text:      text,
		Multiline: multiline,
	}
}

type InputBlockElementType string

type InputBlockElement struct {
	PlainTextInputElement   *PlainTextInputBlockElement
	SelectInputElement      *SelectBlockElement
	MultiSelectInputElement *MultiSelectBlockElement
	DatePickerInputElement  *DatePickerBlockElement
}

type PlainTextInputBlockElement struct {
	Type         string `json:"type,omitempty"`
	ActionID     string `json:"action_id,omitempty"`
	Placeholder  string `json:"placeholder"`
	InitialValue string `json:"initial_value,omitempty"`
	Multiline    bool   `json:"multiline,omitempty"`
	MinLength    int    `json:"min_length,omitempty"`
	MaxLength    int    `json:"max_length,omitempty"`
}

type MultiSelectBlockElement struct{}
