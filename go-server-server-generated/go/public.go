package sw

import (
	application "server-application"
	"fmt"
	"net/http"
	"encoding/json"
)

type Public struct {

}


func Index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World 2!!!")
}

func AddQuestion(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// TODO: Create a convention to plug non-generated code into the generated one with minimal effort
		// https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler
		// The best is to have a MyAddQuestion(w http.ResponseWriter, r *http.Request)
		// tpo match every generated method. 
		// Access to global variables (and imports) can then be dealt with in non-generated files...
		// An alternative to this would be to ignore this file completely and replace it
		// (using same names to get help from compiler...)
		application.App.AddQuestion("whatever")
		w.WriteHeader(http.StatusOK)
}

func ListQuestions(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")

		// w.WriteHeader(http.StatusOK)

		questionsJson, _ := json.Marshal(application	.App.Questions)
		// if err != nil {
		// 	fmt.Fprintf(os.Stdout, "Cannot encode to JSON %s", err)
		// }
		// fmt.Fprintf(os.Stdout, "%s", questionsJson)
		fmt.Fprintf(w, string(questionsJson))

		
}

func Upvote(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

