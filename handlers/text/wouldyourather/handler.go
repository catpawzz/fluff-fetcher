package wouldyourather_handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wouldYouRatherQuestions = []string{
	"Would you rather be able to fly or be invisible?",
	"Would you rather live in the mountains or by the beach?",
	"Would you rather have unlimited money or unlimited time?",
	"Would you rather never use social media again or never watch TV again?",
	"Would you rather be famous or be the best friend of someone famous?",
	"Would you rather have the ability to talk to animals or speak all human languages?",
	"Would you rather always be 10 minutes late or always be 20 minutes early?",
	"Would you rather lose all your money or all your photos?",
	"Would you rather have free Wi-Fi wherever you go or free coffee wherever you go?",
	"Would you rather live in the past or the future?",
	"Would you rather be without internet for a week or without your phone?",
	"Would you rather always have to sing instead of speak or always have to dance while walking?",
	"Would you rather never eat your favorite food again or only eat your favorite food?",
	"Would you rather have a pause or a rewind button for your life?",
	"Would you rather be able to read minds or be able to see the future?",
	"Would you rather have unlimited sushi or unlimited tacos?",
	"Would you rather explore space or the deep ocean?",
	"Would you rather be a famous athlete or a famous musician?",
	"Would you rather never have to sleep or never have to eat?",
	"Would you rather always be too hot or always be too cold?",
}

func getRandomWouldYouRather() string {
	return wouldYouRatherQuestions[rand.Intn(len(wouldYouRatherQuestions))]
}

func WouldYouRatherHandler(w http.ResponseWriter, r *http.Request) {
	question := getRandomWouldYouRather()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"handler": "/handlers/text/wouldyourather","question": "%s", "status": "success"}`, question)
}
