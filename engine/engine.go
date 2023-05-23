package engine

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/esgj/gochat/model"
	"github.com/esgj/gochat/utils"
	"github.com/jbrukh/bayesian"
)

type Engine struct {
	Intents       []model.Intent
	currentIntent model.Intent
	classifier *bayesian.Classifier
}

func (e *Engine) GetResponse(message string) string {
	e.calcNewIntent(message);
	rand.Seed(time.Now().Unix())

	if score := e.getScoreByCurrentIntent(message); score > 0.5 {
		randIndex := rand.Intn(len(e.currentIntent.Responses))
		return e.currentIntent.Responses[randIndex]
	}

	return e.currentIntent.Fallback[rand.Intn(len(e.currentIntent.Fallback))]
}

func (e *Engine) calcNewIntent(message string) {
	// using the bayesian classifier to get likely intent
	_, likelyIntentIndex, _ := e.classifier.LogScores(strings.Split(getParsedMessage(message), " "))

	if (likelyIntentIndex >= 0 && likelyIntentIndex < len(e.Intents)) {
		e.currentIntent = e.Intents[likelyIntentIndex]
	}
}

func (e *Engine) getScoreByCurrentIntent(message string) float32 {
	var sum float32

	words := strings.Split(getParsedMessage(message), " ")

	for _, word := range e.currentIntent.Words {
		for _, givenWord := range words {
			if score := utils.CompareTwoStrings(word, givenWord); score > 0.5 {
				sum += score
			}
		}
	}

	fmt.Println("Intent: ", e.currentIntent.Class)
	fmt.Println("Score: ", sum)

	return sum
}

func (e *Engine) Setup() {
	var classes []bayesian.Class

	for _, intent := range e.Intents {
		classes = append(classes, intent.Class)
	}

	e.classifier = bayesian.NewClassifier(classes...)
}

func (e *Engine) Learn() {
	for _, intent := range e.Intents {
		e.classifier.Learn(intent.Words, intent.Class)
	}
}

func getParsedMessage(message string) string {
	p := strings.ReplaceAll(message, ",", "")
	p = strings.ReplaceAll(p, "?", "")
	p = strings.ToLower(p)

	return p
}