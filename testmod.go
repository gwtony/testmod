package testmod

import (
	"errors"
	"fmt"
)

// Hi returns a friendly greeting in language lang
func Hi(name, lang string) (string, error) {
	switch lang {
	case "cn":
		return fmt.Sprintf("Hello, ba%s!", name), nil
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
