package engine

import (
	"errors"
	"math/rand"
	"time"

	"github.com/esgj/gochat/model"
	"github.com/esgj/gochat/utils"
)

type Engine struct {
	Classes       []model.IntentClass
	Intents       []model.Intent
	currentIntent model.Intent
	currentClass  model.IntentClass
}

func (e *Engine) GetResponse(message string) string {
	if (e.currentIntent.Name == model.Intent{}.Name) {
		e.currentIntent = getIntent(e.Classes[0], e.Intents)
	}

	var result string

	for i := 0; i < 2; i++ {
		if res, err := e.calcResult(message); err != nil {

			if intentClass := matchNewIntentClass(message, e.Classes); (intentClass.Intent != model.IntentClass{}.Intent && e.currentClass.Intent != intentClass.Intent) {
				e.currentClass = intentClass
				e.currentIntent = getIntent(e.currentClass, e.Intents)
				continue
			} else {
				rand.Seed(time.Now().Unix())
				randIndex := rand.Intn(len(e.currentIntent.Fallback))
				result = e.currentIntent.Fallback[randIndex]
			}
		} else {
			result = res
		}
	}

	return result
}

func getIntent(class model.IntentClass, intents []model.Intent) model.Intent {
	for _, intent := range intents {
		if intent.Name == class.Intent {
			return intent
		}
	}

	return intents[0]
}

func matchNewIntentClass(word string, classes []model.IntentClass) model.IntentClass {
	for index, class := range classes {
		for _, classWord := range classes[index].Words {
			if utils.CompareTwoStrings(classWord, word) > 0.5 {
				return class
			}
		}
	}

	return model.IntentClass{}
}

func (e *Engine) calcResult(message string) (string, error) {
	var result string
	for _, keyword := range e.currentIntent.Match {
		if utils.CompareTwoStrings(keyword, message) >= 0.5 {
			rand.Seed(time.Now().Unix())
			randIndex := rand.Intn(len(e.currentIntent.Responses))
			result = e.currentIntent.Responses[randIndex]
			break
		}
	}

	if result == "" {
		return "", errors.New("No match")
	}

	return result, nil
}
