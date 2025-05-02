package truthordare_handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func TruthOrDareHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/text/truthordare/")
	path = strings.ToLower(path)

	var question string
	var questionType string

	switch path {
	case "truth":
		question = getRandomTruth()
		questionType = "truth"
	case "dare":
		question = getRandomDare()
		questionType = "dare"
	case "":
		if rand.Intn(2) == 0 {
			question = getRandomTruth()
			questionType = "truth"
		} else {
			question = getRandomDare()
			questionType = "dare"
		}
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"handler": "/handlers/text/truthordare","title": "Error 404","message": "Unknown request type. Use /truth, /dare, or no suffix for random.", "status": "failure"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"handler": "/handlers/text/truthordare","type": "%s","question": "%s", "status": "success"}`, questionType, question)
}


var truthQuestions = []string{
	"What's the most embarrassing thing you've ever done?",
	"What's your biggest fear?",
	"What's your biggest secret?",
	"What's the last lie you told?",
	"What's the most childish thing you still do?",
	"What's something you hope your family never finds out about?",
	"What's your biggest regret?",
	"What's the worst thing you've ever done?",
	"What's something you're still trying to get over?",
	"What's the most hurtful thing someone has ever said to you?",
	"What's the most trouble you've been in?",
	"What's a secret you've never told anyone?",
	"What's the most embarrassing thing in your search history?",
	"What's the worst date you've ever been on?",
	"What's your biggest insecurity?",
}

var dareQuestions = []string{
	"Do your best impression of a celebrity.",
	"Call a friend and tell them you miss them.",
	"Speak in an accent for the next 10 minutes.",
	"Do 20 jumping jacks.",
	"Text your crush and tell them you like them.",
	"Let someone else post on your social media.",
	"Show everyone your most recent photo.",
	"Dance without music for 30 seconds.",
	"Let someone else do your hair.",
	"Eat a spoonful of hot sauce.",
	"Call a random number and sing happy birthday.",
	"Try to lick your elbow.",
	"Do your best animal impression.",
	"Wear your clothes backwards for the rest of the day.",
	"Let someone else text anyone from your phone.",
}

func getRandomTruth() string {
	return truthQuestions[rand.Intn(len(truthQuestions))]
}

func getRandomDare() string {
	return dareQuestions[rand.Intn(len(dareQuestions))]
}