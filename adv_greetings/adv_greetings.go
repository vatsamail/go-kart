package adv_greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hi(name string) (string, error) {
	if name == "" {
		return name, errors.New("Error: Hello function needs a name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hi(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		">>>>> Hi, %v. Welcome",
		">>>>> Greeting to you, %v",
		">>>>> Hail %v, well met",
	}
	return formats[rand.Intn(len(formats))]
}
