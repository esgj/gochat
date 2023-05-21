package gochat

import (
	"github.com/esgj/gochat/engine"
	"github.com/esgj/gochat/model"
)

func Default(intents []model.Intent, classes []model.IntentClass) *engine.Engine {
	engine := engine.Engine{
		Classes: classes,
		Intents: intents,
	}

	return &engine
}

var TestIntents []model.Intent = []model.Intent{
	{
		Name: "greeting",
		Match: []string{"hi", "hello"},
		Respones: []string{"Hello there!", "Hi, how are you?"},
		Fallback: []string{"I can't answer your question at this time."},
	},
	{
		Name: "weather",
		Match: []string{"weather", "how is the weather", "what is the weather like", "weather tomorrow", "weather today"},
		Respones: []string{"I cannot check that right now!", "I'm not sure, sorry."},
		Fallback: []string{"I cannot answer that question right now"},
	},
}

var TestClasses []model.IntentClass = []model.IntentClass{
	{
		Intent: "greeting",
		Words: []string{"hi", "hello"},
	},
	{
		Intent: "weather",
		Words: []string{"weather", "how is the weather", "what is the weather like", "weather tomorrow", "weather today"},
	},
}