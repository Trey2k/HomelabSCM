package wasm_ui

import (
	_ "embed"
	"fmt"
	"syscall/js"

	"homelabscm.com/scm/internal/pkg/jquery_wasm"
)

//go:embed settings_ui.html
var settingsTemplate string

type InputType int
const (
	INPUT_TYPE_TEXT InputType = iota
	INPUT_TYPE_NUMBER
	INPUT_TYPE_PASSWORD
	INPUT_TYPE_STR_ARRAY
)

type Input interface {
	GetValue() js.Value
	SetValue(value js.Value)
	GetType() InputType
	GetID() string
}

type SettingsUI struct {
	id string
	inputs []Input
}

func NewSettingsUI(id string, parent js.Value) (*SettingsUI, error) {
	template_data := map[string]string{
		"ID": id,
	}

	html_src, err := ParseTemplate(settingsTemplate, "settings_ui", template_data)
	if err != nil {
		return nil, err
	}

	jquery_wasm.Append(parent, html_src)

	settings := &SettingsUI{
		id: id,
		inputs: make([]Input, 0),
	}

	return settings, nil
}

func (s *SettingsUI) AddButton(label string, id string, handler func(this js.Value, args []js.Value) any) error {
	template_data := map[string]string{
		"Label": label,
		"ID": id,
	}

	html_src, err := ParseTemplate(settingsTemplate, "button", template_data)
	if err != nil {
		return err
	}

	jquery_wasm.Append(js.ValueOf("#"+s.id), html_src)

	jquery_wasm.OnClick(js.ValueOf("#"+id), handler)

	return nil
}


func (s *SettingsUI) GetInputs() []Input {
	return s.inputs
}

func (s *SettingsUI) addInput(input Input) {
	s.inputs = append(s.inputs, input)
}

type textInput struct {
	id string
}

func (t *textInput) GetValue() js.Value {
	return jquery_wasm.GetValue(js.ValueOf("#" + t.id))
}

func (t *textInput) SetValue(value js.Value) {
	jquery_wasm.SetValue(js.ValueOf("#" + t.id), value)
}

func (t *textInput) GetType() InputType {
	return INPUT_TYPE_TEXT
}

func (t *textInput) GetID() string {
	return t.id
}

func (s *SettingsUI) AddTextInput(label string, defaulValue js.Value, id string) error {
	template_data := map[string]string{
		"Label": label,
		"DefaultValue": defaulValue.String(),
		"ID": id,
	}

	html_src, err := ParseTemplate(settingsTemplate, "text_input", template_data)
	if err != nil {
		return err
	}

	jquery_wasm.Append(js.ValueOf("#"+s.id), html_src)
	s.addInput(&textInput{id: id})

	return nil
}

type numberInput struct {
	id string
}

func (n *numberInput) GetValue() js.Value {
	return jquery_wasm.GetValue(js.ValueOf("#" + n.id))
}

func (n *numberInput) SetValue(value js.Value) {
	jquery_wasm.SetValue(js.ValueOf("#" + n.id), value)
}

func (n *numberInput) GetType() InputType {
	return INPUT_TYPE_NUMBER
}

func (n *numberInput) GetID() string {
	return n.id
}

func (s *SettingsUI) AddNumberInput(label string, defaulValue js.Value, id string) error {
	template_data := map[string]interface{}{
		"Label": label,
		"DefaultValue": defaulValue,
		"ID": id,
	}

	html_src, err := ParseTemplate(settingsTemplate, "number_input", template_data)
	if err != nil {
		return err
	}

	jquery_wasm.Append(js.ValueOf("#"+s.id), html_src)
	s.addInput(&numberInput{id: id})

	return nil
}

type passwordInput struct {
	id string
}

func (p *passwordInput) GetValue() js.Value {
	return jquery_wasm.GetValue(js.ValueOf("#" + p.id))
}

func (p *passwordInput) SetValue(value js.Value) {
	jquery_wasm.SetValue(js.ValueOf("#" + p.id), value)
}

func (p *passwordInput) GetType() InputType {
	return INPUT_TYPE_PASSWORD
}

func (p *passwordInput) GetID() string {
	return p.id
}

func (s *SettingsUI) AddPasswordInput(label string, id string) error {
	template_data := map[string]string{
		"Label": label,
		"ID": id,
	}

	html_src, err := ParseTemplate(settingsTemplate, "password_input", template_data)
	if err != nil {
		return err
	}

	jquery_wasm.Append(js.ValueOf("#"+s.id), html_src)
	s.addInput(&passwordInput{id: id})

	return nil
}

type strArrayInput struct {
	id string
	input_count int
}

func (s *strArrayInput) GetValue() js.Value {
	childern := jquery_wasm.GetChildren(js.ValueOf("#" + s.id), "input")
	values := make([]string, len(childern))
	for i, child := range childern {
		values[i] = jquery_wasm.GetValue(child).String()
	}

	return js.ValueOf(values)
}

func (s *strArrayInput) SetValue(value js.Value) {
	jquery_wasm.ClearChildren(js.ValueOf("#" + s.id), ".setting-list-item")
	
	s.input_count = value.Length()
	for i := 0; i < value.Length(); i++ {
		template_data := map[string]interface{}{
			"ID": s.id,
			"Num": i,
			"Value": value.Index(i).String(),
		}

		html_src, err := ParseTemplate(settingsTemplate, "setting_list_item", template_data)
		if err != nil {
			panic(err)
		}

		jquery_wasm.Append(js.ValueOf("#" + s.id), html_src)

		jquery_wasm.OnClick(js.ValueOf(fmt.Sprintf("#remove%s_%d", s.id, i)), remove_handler)
	}
}

func (s *strArrayInput) GetType() InputType {
	return INPUT_TYPE_STR_ARRAY
}

func (s *strArrayInput) GetID() string {
	return s.id
}

func (s *SettingsUI) AddStrArrayInput(label string, defaulValue []any, id string) error {
	template_data := map[string]interface{}{
		"Label": label,
		"ID": id,
	}

	html_src, err := ParseTemplate(settingsTemplate, "str_array_input", template_data)
	if err != nil {
		return err
	}

	jquery_wasm.Append(js.ValueOf("#"+s.id), html_src)
	input := &strArrayInput{
		id: id,
		input_count: len(defaulValue),
	}
	s.addInput(input)

	input.SetValue(js.ValueOf(defaulValue))

	jquery_wasm.OnClick(js.ValueOf("#add"+id), func(this js.Value, args []js.Value) any {
		num := input.input_count
		input.input_count++

		fmt.Printf("Adding %s_%d\n", id, num)

		template_data := map[string]any{
			"ID": id,
			"Num": num,
			"Value": "",
		}

		html_src, err := ParseTemplate(settingsTemplate, "setting_list_item", template_data)
		if err != nil {
			panic(err)
		}

		jquery_wasm.Append(js.ValueOf("#" + id), html_src)

		jquery_wasm.OnClick(js.ValueOf(fmt.Sprintf("#remove%s_%d", id, num)), remove_handler)

		return nil
	})

	return nil
}

func remove_handler(this js.Value, args []js.Value) any {
	jquery_wasm.Remove(this.Get("parentNode"))

	return nil
}