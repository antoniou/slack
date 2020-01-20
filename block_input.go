package slack

// https://api.slack.com/reference/block-kit/blocks#input

const (
	IBEPlainText   InputBlockElementType = "plain_text_input"
	IBESelect      InputBlockElementType = "static_select"
	IBEMultiSelect InputBlockElementType = "multi_static_select"
	IBEDatePicker  InputBlockElementType = "datepicker"
)

type InputBlock struct {
	Type     MessageBlockType   `json:"type"`
	Label    *TextBlockObject   `json:"label,omitempty"`
	Element  *InputBlockElement `json:"element,omitempty"`
	BlockID  string             `json:"block_id,omitempty"`
	Hint     *TextBlockObject   `json:"hint,omitempty"`
	Optional bool               `json:"multiline,omitempty"`
}

// BlockType returns the type of the block
func (i InputBlock) BlockType() MessageBlockType {
	return i.Type
}

func NewInputBlock(blockID string, label *TextBlockObject, element *InputBlockElement, hint *TextBlockObject, optional bool) *InputBlock {
	return &InputBlock{
		BlockID:  blockID,
		Type:     MBTInput,
		Label:    label,
		Element:  element,
		Hint:     hint,
		Optional: optional,
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
