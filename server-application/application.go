package application

// import (
// 	"fmt"
// )


var App *Application


type Rectangle struct {
	Name          string
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}


type Question struct {
	Question		string
	Votes			int
}

type Application struct {
	Questions 		[]Question
}

func (app *Application) AddQuestion(q string) {
	app.Questions = append(app.Questions, Question {q, 0})
}