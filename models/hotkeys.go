package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/MakeNowJust/hotkey"
)

//Hotkeys to be assigned to counter items
type Hotkeys struct {
	Increase string
	Decrease string
	Reset    string
}

//ParseHotkey returns the modifier and key from a string
func ParseHotkey(str string) (hotkey.Modifier, uint32, error) {
	_, modifiers := getModifiers(str)
	character, err := getKey(str)
	if err != nil {
		return hotkey.Alt + hotkey.Ctrl, 'F', errors.New("Failed To Interpret")
	}
	fmt.Printf("%s; Mod: %d; Key: %d\n", str, modifiers, character)
	return modifiers, character, nil
}

func getModifiers(str string) (bool, hotkey.Modifier) {
	modifier := hotkey.None

	modifierRegex := regexp.MustCompile(`(Ctrl\+|Alt\+|Win\+|Shift\+)`)
	found := modifierRegex.FindAll([]byte(str), -1)
	for _, mod := range found {
		parsedMod, err := parseModifier(mod)
		if err == nil {
			modifier += parsedMod
		}
	}

	return len(found) > 0, modifier
}

func getKey(str string) (uint32, error) {
	key := uint32(0x00)
	tokens := strings.Split(str, "+")
	lastidx := len(tokens) - 1
	if lastidx < 0 {
		return key, errors.New("No Tokens Found")
	}
	if len(tokens[lastidx]) == 1 {
		char := tokens[lastidx][0]
		return uint32(char), nil
	}
	switch tokens[lastidx] {
	case "Back":
		return hotkey.BACK, nil
	case "Tab":
		return hotkey.TAB, nil
	case "Clear":
		return hotkey.CLEAR, nil
	case "Return":
		return hotkey.RETURN, nil
	case "Pause":
		return hotkey.PAUSE, nil
	case "Caps":
		return hotkey.CAPITAL, nil
	case "Esc":
		return hotkey.ESCAPE, nil
	case "Space":
		return hotkey.SPACE, nil
	case "Prior":
		return hotkey.PRIOR, nil
	case "Next":
		return hotkey.NEXT, nil
	case "End":
		return hotkey.END, nil
	case "Home":
		return hotkey.HOME, nil
	case "Left":
		return hotkey.LEFT, nil
	case "Up":
		return hotkey.UP, nil
	case "Right":
		return hotkey.RIGHT, nil
	case "Down":
		return hotkey.DOWN, nil
	case "Select":
		return hotkey.SELECT, nil
	case "Print":
		return hotkey.PRINT, nil
	case "Snapshot":
		return hotkey.SNAPSHOT, nil
	case "Insert":
		return hotkey.INSERT, nil
	case "Del":
		return hotkey.DELETE, nil
	case "Numpad0":
		return hotkey.NUMPAD0, nil
	case "Numpad1":
		return hotkey.NUMPAD1, nil
	case "Numpad2":
		return hotkey.NUMPAD2, nil
	case "Numpad3":
		return hotkey.NUMPAD3, nil
	case "Numpad4":
		return hotkey.NUMPAD4, nil
	case "Numpad5":
		return hotkey.NUMPAD5, nil
	case "Numpad6":
		return hotkey.NUMPAD6, nil
	case "Numpad7":
		return hotkey.NUMPAD7, nil
	case "Numpad8":
		return hotkey.NUMPAD8, nil
	case "Numpad9":
		return hotkey.NUMPAD9, nil
	case "Num*":
		return hotkey.MULTIPLY, nil
	case "Num+":
		return hotkey.ADD, nil
	case "Num-":
		return hotkey.SUBTRACT, nil
	case "Num.":
		return hotkey.DECIMAL, nil
	case "Num/":
		return hotkey.DIVIDE, nil
	case "F1":
		return hotkey.F1, nil
	case "F2":
		return hotkey.F2, nil
	case "F3":
		return hotkey.F3, nil
	case "F4":
		return hotkey.F4, nil
	case "F5":
		return hotkey.F5, nil
	case "F6":
		return hotkey.F6, nil
	case "F7":
		return hotkey.F7, nil
	case "F8":
		return hotkey.F8, nil
	case "F9":
		return hotkey.F9, nil
	case "F10":
		return hotkey.F10, nil
	case "F11":
		return hotkey.F11, nil
	case "F12":
		return hotkey.F12, nil
	case "F13":
		return hotkey.F13, nil
	case "F14":
		return hotkey.F14, nil
	case "F15":
		return hotkey.F15, nil
	case "F16":
		return hotkey.F16, nil
	case "F17":
		return hotkey.F17, nil
	case "F18":
		return hotkey.F18, nil
	case "F19":
		return hotkey.F19, nil
	case "F20":
		return hotkey.F20, nil
	case "F21":
		return hotkey.F21, nil
	case "F22":
		return hotkey.F22, nil
	case "F23":
		return hotkey.F23, nil
	case "F24":
		return hotkey.F24, nil
	case "Oem+":
		return hotkey.OEM_PLUS, nil
	case "Oem,":
		return hotkey.OEM_COMMA, nil
	case "Oem-":
		return hotkey.OEM_MINUS, nil
	case "Oem.":
		return hotkey.OEM_PERIOD, nil
	}
	return key, nil
}

func parseModifier(str []byte) (hotkey.Modifier, error) {
	var modifier hotkey.Modifier
	switch string(str) {
	case "Ctrl+":
		return hotkey.Ctrl, nil
	case "Alt+":
		return hotkey.Alt, nil
	case "Win+":
		return hotkey.Win, nil
	case "Shift+":
		return hotkey.Shift, nil
	}
	return modifier, nil
}
