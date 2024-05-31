package shell

import (
	"os"
	"strings"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm"
)

// DefaultInteractiveTextInput is the default InteractiveTextInput printer.
var DefaultInteractivePromptInput = InteractiveTextPromptPrinter{
	DefaultText: "shell",
	Delimiter:   ": ",
	TextStyle:   &pterm.ThemeDefault.PrimaryStyle,
	Mask:        "",
}

// InteractiveTextInputPrinter is a printer for interactive select menus.
type InteractiveTextPromptPrinter struct {
	TextStyle       *pterm.Style
	DefaultText     string
	DefaultValue    string
	Delimiter       string
	MultiLine       bool
	Mask            string
	OnInterruptFunc func()

	input         []string
	cursorXPos    int
	cursorYPos    int
	text          string
	startedTyping bool
	valueStyle    *pterm.Style
}

// WithDefaultText sets the default text.
func (p InteractiveTextPromptPrinter) WithDefaultText(text string) *InteractiveTextPromptPrinter {
	p.DefaultText = text
	return &p
}

// WithDefaultValue sets the default value.
func (p InteractiveTextPromptPrinter) WithDefaultValue(value string) *InteractiveTextPromptPrinter {
	p.DefaultValue = value
	return &p
}

// WithTextStyle sets the text style.
func (p InteractiveTextPromptPrinter) WithTextStyle(style *pterm.Style) *InteractiveTextPromptPrinter {
	p.TextStyle = style
	return &p
}

// WithMask sets the mask.
func (p InteractiveTextPromptPrinter) WithMask(mask string) *InteractiveTextPromptPrinter {
	p.Mask = mask
	return &p
}

// WithOnInterruptFunc sets the function to execute on exit of the input reader
func (p InteractiveTextPromptPrinter) WithOnInterruptFunc(exitFunc func()) *InteractiveTextPromptPrinter {
	p.OnInterruptFunc = exitFunc
	return &p
}

// WithDelimiter sets the delimiter between the message and the input.
func (p InteractiveTextPromptPrinter) WithDelimiter(delimiter string) *InteractiveTextPromptPrinter {
	p.Delimiter = delimiter
	return &p
}

// Show shows the interactive select menu and returns the selected entry.
func (p InteractiveTextPromptPrinter) Show(text ...string) (string, error) {
	// should be the first defer statement to make sure it is executed last
	// and all the needed cleanup can be done before
	cancel, exit := NewCancelationSignal(p.OnInterruptFunc)
	defer exit()

	var areaText string

	if len(text) == 0 || text[0] == "" {
		text = []string{p.DefaultText}
	}

	if p.MultiLine {
		areaText = p.TextStyle.Sprintfln("%s %s %s", text[0], pterm.ThemeDefault.SecondaryStyle.Sprint("[Press tab to submit]"), p.Delimiter)
	} else {
		areaText = p.TextStyle.Sprintf("%s%s", text[0], p.Delimiter)
	}

	p.text = areaText
	area := cursor.NewArea()
	area.Update(areaText)
	area.StartOfLine()

	if !p.MultiLine {
		cursor.Right(runewidth.StringWidth(pterm.RemoveColorFromString(areaText)))
	}

	if p.DefaultValue != "" {
		p.input = append(p.input, pterm.Gray(p.DefaultValue))
		p.updateArea(&area)
	}

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if !p.MultiLine {
			p.cursorYPos = 0
		}
		if len(p.input) == 0 {
			p.input = append(p.input, "")
		}

		switch key.Code {
		case keys.Tab:
			if p.MultiLine {
				area.Bottom()
				return true, nil
			}
		case keys.Enter:
			if p.DefaultValue != "" && !p.startedTyping {
				for i := range p.input {
					p.input[i] = pterm.RemoveColorFromString(p.input[i])
				}

				if p.MultiLine {
					area.Bottom()
				}
				return true, nil
			}

			if p.MultiLine {
				if key.AltPressed {
					p.cursorXPos = 0
				}
				appendAfterY := append([]string{}, p.input[p.cursorYPos+1:]...)
				appendAfterX := string(append([]rune{}, []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...))
				p.input[p.cursorYPos] = string(append([]rune{}, []rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos]...))
				p.input = append(p.input[:p.cursorYPos+1], appendAfterX)
				p.input = append(p.input, appendAfterY...)
				p.cursorYPos++
				p.cursorXPos = -GetStringMaxWidth(p.input[p.cursorYPos])
				cursor.StartOfLine()
			} else {
				return true, nil
			}
		case keys.RuneKey:
			if !p.startedTyping {
				p.input = []string{""}
				p.startedTyping = true
			}
			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(key.String()), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))
		case keys.Space:
			if !p.startedTyping {
				p.input = []string{" "}
				p.startedTyping = true
			}
			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(" "), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))
		case keys.Backspace:
			if !p.startedTyping {
				p.input = []string{""}
				p.startedTyping = true
			}
			if len([]rune(p.input[p.cursorYPos]))+p.cursorXPos > 0 {
				p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))-1+p.cursorXPos], []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...))
			} else if p.cursorYPos > 0 {
				p.input[p.cursorYPos-1] += p.input[p.cursorYPos]
				appendAfterY := append([]string{}, p.input[p.cursorYPos+1:]...)
				p.input = append(p.input[:p.cursorYPos], appendAfterY...)
				p.cursorXPos = 0
				p.cursorYPos--
			}
		case keys.Delete:
			if !p.startedTyping {
				p.input = []string{""}
				p.startedTyping = true
				return false, nil
			}
			if len([]rune(p.input[p.cursorYPos]))+p.cursorXPos < len([]rune(p.input[p.cursorYPos])) {
				p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos+1:]...))
				p.cursorXPos++
			} else if p.cursorYPos < len(p.input)-1 {
				p.input[p.cursorYPos] += p.input[p.cursorYPos+1]
				appendAfterY := append([]string{}, p.input[p.cursorYPos+2:]...)
				p.input = append(p.input[:p.cursorYPos+1], appendAfterY...)
				p.cursorXPos = 0
			}
		case keys.CtrlC:
			cancel()
			return true, nil
		case keys.Down:
			if !p.startedTyping {
				p.input = []string{""}
				p.startedTyping = true
			}
			if p.cursorYPos+1 < len(p.input) {
				p.cursorXPos = (GetStringMaxWidth(p.input[p.cursorYPos]) + p.cursorXPos) - GetStringMaxWidth(p.input[p.cursorYPos+1])
				if p.cursorXPos > 0 {
					p.cursorXPos = 0
				}
				p.cursorYPos++
			}
		case keys.Up:
			if !p.startedTyping {
				p.input = []string{""}
				p.startedTyping = true
			}
			if p.cursorYPos > 0 {
				p.cursorXPos = (GetStringMaxWidth(p.input[p.cursorYPos]) + p.cursorXPos) - GetStringMaxWidth(p.input[p.cursorYPos-1])
				if p.cursorXPos > 0 {
					p.cursorXPos = 0
				}
				p.cursorYPos--
			}
		}

		if GetStringMaxWidth(p.input[p.cursorYPos]) > 0 {
			switch key.Code {
			case keys.Right:
				if p.cursorXPos < 0 {
					p.cursorXPos++
				} else if p.cursorYPos < len(p.input)-1 {
					p.cursorYPos++
					p.cursorXPos = -GetStringMaxWidth(p.input[p.cursorYPos])
				}
			case keys.Left:
				if p.cursorXPos+GetStringMaxWidth(p.input[p.cursorYPos]) > 0 {
					p.cursorXPos--
				} else if p.cursorYPos > 0 {
					p.cursorYPos--
					p.cursorXPos = 0
				}
			}
		}

		p.updateArea(&area)

		return false, nil
	})
	if err != nil {
		return "", err
	}

	// Add new line
	pterm.Println()

	for i, s := range p.input {
		if i < len(p.input)-1 {
			areaText += s + "\n"
		} else {
			areaText += s
		}
	}

	if !p.startedTyping {
		return p.DefaultValue, nil
	}

	return strings.ReplaceAll(areaText, p.text, ""), nil
}

func (p InteractiveTextPromptPrinter) updateArea(area *cursor.Area) string {
	if !p.MultiLine {
		p.cursorYPos = 0
	}
	areaText := p.text

	for i, s := range p.input {
		if i < len(p.input)-1 {
			areaText += s + "\n"
		} else {
			areaText += s
		}
	}

	if p.Mask != "" {
		areaText = p.text + strings.Repeat(p.Mask, GetStringMaxWidth(areaText)-GetStringMaxWidth(p.text))
	}

	if p.cursorXPos+GetStringMaxWidth(p.input[p.cursorYPos]) < 1 {
		p.cursorXPos = -GetStringMaxWidth(p.input[p.cursorYPos])
	}

	area.Update(areaText)
	area.Top()
	area.Down(p.cursorYPos + 1)
	area.StartOfLine()
	if p.MultiLine {
		cursor.Right(GetStringMaxWidth(p.input[p.cursorYPos]) + p.cursorXPos)
	} else {
		cursor.Right(GetStringMaxWidth(areaText) + p.cursorXPos)
	}
	return areaText
}

func GetStringMaxWidth(s string) int {
	var max int
	ss := strings.Split(s, "\n")
	for _, s2 := range ss {
		s2WithoutColor := color.ClearCode(s2)
		if runewidth.StringWidth(s2WithoutColor) > max {
			max = runewidth.StringWidth(s2WithoutColor)
		}
	}
	return max
}

func NewCancelationSignal(interruptFunc func()) (func(), func()) {
	canceled := false

	cancel := func() {
		canceled = true
	}

	exit := func() {
		if canceled {
			if interruptFunc != nil {
				interruptFunc()
			} else {
				os.Exit(1)
			}
		}
	}

	return cancel, exit
}
