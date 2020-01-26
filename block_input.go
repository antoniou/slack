package slack

import "encoding/json"

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
	PlainTextInputElement *PlainTextInputBlockElement
	SelectElement         *SelectBlockElement
	MultiSelectElement    *MultiSelectBlockElement
	DatePickerElement     *DatePickerBlockElement
}

// MarshalJSON implements the Marshaller interface for InputBlockElement so that any JSON
// marshalling is delegated and proper type determination can be made before marshal
func (i *InputBlockElement) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(inputToBlockElement(i))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func inputToBlockElement(element *InputBlockElement) BlockElement {
	if element.PlainTextInputElement != nil {
		return element.PlainTextInputElement
	}
	if element.SelectElement != nil {
		return element.SelectElement
	}
	// if element.MultiSelectElement != nil {
	// 	return element.MultiSelectElement
	// }
	if element.DatePickerElement != nil {
		return element.DatePickerElement
	}

	return nil
}

// NewAccessory returns a new Accessory for a given block element
func NewInputBlockElement(element BlockElement) *InputBlockElement {
	switch element.(type) {
	case *PlainTextInputBlockElement:
		return &InputBlockElement{PlainTextInputElement: element.(*PlainTextInputBlockElement)}
	case *SelectBlockElement:
		return &InputBlockElement{SelectElement: element.(*SelectBlockElement)}
	case *DatePickerBlockElement:
		return &InputBlockElement{DatePickerElement: element.(*DatePickerBlockElement)}
	}

	return nil
}

type PlainTextInputBlockElement struct {
	Type         MessageElementType `json:"type,omitempty"`
	ActionID     string             `json:"action_id,omitempty"`
	Placeholder  *TextBlockObject   `json:"placeholder,omitempty"`
	InitialValue string             `json:"initial_value,omitempty"`
	Multiline    bool               `json:"multiline,omitempty"`
	MinLength    uint               `json:"min_length,omitempty"`
	MaxLength    uint               `json:"max_length,omitempty"`
}

func (s PlainTextInputBlockElement) ElementType() MessageElementType {
	return s.Type
}

func NewPlainTextInputBlockElement(actionID string, placeholder *TextBlockObject, initialValue string, multiline bool, minLength, maxLength uint) *PlainTextInputBlockElement {
	return &PlainTextInputBlockElement{
		Type:         METPlainTextInput,
		ActionID:     actionID,
		Placeholder:  placeholder,
		InitialValue: initialValue,
		Multiline:    multiline,
		MinLength:    minLength,
		MaxLength:    maxLength,
	}
}

type MultiSelectBlockElement struct{}
