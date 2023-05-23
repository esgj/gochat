package gochat

import (
	"github.com/esgj/gochat/engine"
	"github.com/esgj/gochat/model"
)

func Default(intents []model.Intent) *engine.Engine {
	engine := engine.Engine{
		Intents: intents,
	}

	// Iterates over intent classes and adds them to the bayesian classifier.
	engine.Setup()
	// Takes all intents and runs it through the bayesian algorithm.
	engine.Learn()

	return &engine
}

var TestIntents []model.Intent = []model.Intent{
	{
		Class: "greeting",
		Words: []string{"hi", "hello"},
		Responses: []string{"Hello there!", "Hi, how are you?"},
		Fallback: []string{"I can't answer your question at this time."},
	},
	{
		Class: "weather",
		Words: []string{"weather", "now"},
		Responses: []string{"I'm not sure how the weather is right now", "I'm not sure about the weather, sorry."},
		Fallback: []string{"I cannot answer that question right now"},
	},
	{
		Class: "joke",
		Words: []string{"joke", "joke?"},
		Responses: []string{"I do not have any jokes at this moment, sorry!"},
		Fallback: []string{"I cannot answer that question right now"},
	},
	{
		Class: "joke-tomorrow",
		Words: []string{"joke", "tomorrow", "please"},
		Responses: []string{"I do not have any jokes at this moment, sorry!"},
		Fallback: []string{"I cannot answer that question right now"},
	},
}