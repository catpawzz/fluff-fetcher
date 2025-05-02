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
	"What's the weirdest thing you've ever eaten?",
	"What's your worst online shopping mistake?",
	"What's the longest you've gone without showering?",
	"What's the most embarrassing message you've sent to the wrong person?",
	"What's the most ridiculous thing you've done while home alone?",
	"What's the last thing you searched for on the internet?",
	"What's your most toxic trait?",
	"What's a weird habit you have that nobody knows about?",
	"What was your most awkward moment in a voice chat?",
	"What's the most embarrassing nickname you've ever had?",
	"What's your go-to excuse when you don't want to do something?",
	"What's a lie you've told that got way out of hand?",
	"What's the weirdest dream you've ever had?",
	"What's your most unpopular opinion?",
	"What's the most ridiculous thing you believed as a child?",
	"What's the worst game you've ever spent money on?",
	"What's the silliest argument you've ever had?",
	"What's something you pretend to understand but actually don't?",
	"What's the laziest thing you've ever done?",
	"What's the strangest thing you've done to impress someone?",
	"What's the most desperate thing you've done because you were bored?",
	"What's the most embarrassing thing your parents have caught you doing?",
	"What's your most irrational fear?",
	"What's your worst fashion choice that you thought was cool at the time?",
	"What's a skill you claim to have but don't actually possess?",
	"What's the worst advice you've ever given someone?",
	"What's the worst advice you've ever followed?",
	"What's your comfort video or stream you watch when you're sad?",
	"What's the pettiest thing you've ever done?",
	"What's something you would never admit to your friends?",
	"What's the weirdest thing you've done in a voice call?",
	"What's a secret talent that you have?",
	"What's the most cringe-worthy thing in your Discord history?",
	"What's the most embarrassing thing you've done in a game?",
	"What's a weird food combination you enjoy?",
	"What's the worst username you've ever had online?",
	"What's your go-to Discord status when you're avoiding someone?",
	"What's a message you drafted but were too afraid to send?",
	"What's the longest time you've spent stalking someone's profile?",
	"What's the most embarrassing emoji you use regularly?",
	"What's the weirdest Discord server you've joined?",
	"What's something you've done in a game that you're ashamed of?",
	"What's the most awkward voice chat you've experienced?",
	"What's the real reason you left that server?",
	"What's the longest you've pretended to be AFK?",
	"What's the most embarrassing thing you've accidentally shared on stream?",
	"What's a Discord notification you received that made your heart race?",
	"What's your most controversial gaming opinion?",
	"What's the worst rage quit you've ever had?",
	"What's the most you've spent on in-game purchases?",
	"What's a server drama you secretly enjoyed watching?",
	"What's the weirdest DM you've ever received?",
	"What's the real reason you have notifications turned off?",
	"What's something you posted that you immediately regretted?",
	"What's your most used Discord emoji and why?",
	"What's a server rule you've broken without getting caught?",
	"What's a channel you lurk in but never participate in?",
	"What's the longest Discord call you've ever been in?",
	"What's the most annoying habit of someone in your server?",
	"What's a Discord trend you secretly can't stand?",
	"What's the strangest thing you've done to get someone's attention online?",
	"What's the most desperate thing you've done to revive a dead chat?",
	"What's a game everyone loves that you think is overrated?",
	"What's the most embarrassing typo you've ever made?",
	"What's your real reaction when someone pings everyone unnecessarily?",
	"What's a Discord feature you use that you think nobody else does?",
	"What's the real reason you mute your mic sometimes?",
	"What's a voice you use in voice chat that isn't your normal speaking voice?",
	"What's the most embarrassing song you've played during Discord activities?",
	"What's the weirdest thing you've done while you thought your mic was muted?",
	"What's your instant judgment when you see someone's profile picture?",
	"What's your most used Discord shortcut that you think others don't know about?",
	"What's the real reason you decline game invites sometimes?",
	"What's a friendship that started on Discord but didn't end well?",
	"What's something you've screenshotted from Discord that you'd be embarrassed if others knew?",
	"What's a Discord memory that makes you cringe thinking about it?",
	"What's the longest you've spent customizing your Discord profile?",
	"What's something you've said in Discord that you would never say in real life?",
	"What's the weirdest reason you've blocked someone?",
	"What's a message you immediately regretted after sending?",
	"What's the real reason you changed your username?",
	"What's the most ridiculous thing you've done to get a specific role in a server?",
	"What's a Discord bot command you've abused?",
	"What's something about your gaming setup you're embarrassed about?",
	"What's a game you pretend to be good at but actually aren't?",
	"What's your honest opinion about someone in this Discord?",
	"What's something you've lied about in your Discord bio?",
	"What's the real reason you use that profile picture?",
	"What's the most awkward interaction you've had with a mod or admin?",
	"What's a Discord habit you know is weird but can't stop doing?",
	"What's the most personal information you've accidentally revealed in chat?",
	"What's something you do on Discord that you think is unique to you?",
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
	"Change your Discord profile picture to whatever the other person chooses for 24 hours.",
	"Send the last photo in your camera roll to the Discord chat.",
	"Type with your elbow for the next 5 minutes in chat.",
	"Send a voice message singing your favorite song.",
	"Write a short poem about the other person and post it in a mutual server.",
	"DM three people a really bad joke.",
	"Send a screenshot of your most recent emojis used.",
	"Change your Discord status to something embarrassing for an hour.",
	"Record yourself saying the alphabet backwards.",
	"Draw something with your eyes closed and share it.",
	"Use only emoji to communicate for the next 10 minutes.",
	"Create a meme about the other person and share it.",
	"Send your Spotify playlist or recently played songs.",
	"Tell a fictional story about how you and the other person met.",
	"Speak in rhymes for the next 5 messages.",
	"Record yourself doing your best dance move.",
	"Share the oldest selfie on your phone.",
	"Do a dramatic reading of the other person's Discord bio.",
	"Write a testimonial about the other person as if you're reviewing a product.",
	"Share your screen while the other person controls what you search for.",
	"Make the other person's profile picture your phone wallpaper for a day.",
	"Send a message to a mutual friend asking a random weird question.",
	"Change your Discord name to something the other person chooses for 24 hours.",
	"Do a dramatic reading of the lyrics to your favorite song.",
	"Call the other person and talk in a whisper the entire time.",
	"Send the funniest meme you have saved.",
	"Post a public compliment about the other person in a mutual server.",
	"Send a voice message of you doing your best impression of the other person.",
	"Share your most embarrassing gaming moment.",
	"Type everything backwards for the next 10 minutes.",
	"Draw yourself from memory and share the result.",
	"Send a voice message to the other person in slow motion.",
	"Create and share an acrostic poem using the other person's name.",
	"Call and sing the happy birthday song to them even if it's not their birthday.",
	"Change all your server nicknames to something ridiculous for a day.",
	"Start a voice chat and try to beatbox for 30 seconds.",
	"Send a DM to someone you haven't talked to in months saying 'I've been thinking about you'.",
	"Change your Discord theme to light mode for a day.",
	"Create a new server emoji based on your facial expression right now.",
	"Join a random public Discord server and introduce yourself.",
	"Post your Discord status using only keyboard symbols.",
	"Send three voice messages in three different accents.",
	"Write a haiku about your favorite Discord emote.",
	"Record yourself doing a dramatic reading of the server rules.",
	"Stream a game but play with your non-dominant hand.",
	"React with the same emoji to the last 10 messages in a shared server.",
	"Send a friend request to someone the other person suggests.",
	"Create a poll in a server about something absurdly specific.",
	"Leave a voice message where you try to imitate five different Discord notification sounds.",
	"Send a screenshot of your Discord DM list (hiding names for privacy).",
	"Share your Discord nitro perks with someone for a month if you have it.",
	"Record yourself explaining something complex while eating crackers.",
	"Go to your oldest Discord message you can find and react to it.",
	"Make a fake advertisement voice message for the Discord server you're in.",
	"Send the most obscure GIF you can find in one minute of searching.",
	"Type in only caps lock for the next hour in all Discord chats.",
	"Call and try to explain your favorite game without naming it.",
	"Send a voice message reciting the alphabet with a different emotion for each letter.",
	"Share your Discord account creation date and tell a story about what you were doing then.",
	"Create a custom status that's a movie quote the other person has to guess.",
	"Take a screenshot of your Discord home page and circle all the unread notifications.",
	"Set your Discord banner to whatever image the other person chooses for a week.",
	"Join a voice channel and speak in third person for five minutes.",
	"Send a message using the predictive text feature on your phone repeatedly.",
	"Record yourself trying to name as many Discord features as possible in 30 seconds.",
	"Share your screen and let the other person navigate through your favorite website.",
	"Start a group call and introduce the other person as a celebrity.",
	"Send a message in a server pretending to be a bot.",
	"Add a random person to a group chat and immediately leave.",
	"Send the same message to five different servers simultaneously.",
	"Create a new server dedicated to something absurdly specific.",
	"Write a short story using only Discord emojis.",
	"Host an impromptu voice channel karaoke session for at least one song.",
	"Send a voice message trying to make the other person laugh without speaking.",
	"Do a stage channel presentation about a completely made-up topic.",
	"Play a game where you can only use your nose to control inputs while streaming.",
	"Join a voice chat and only communicate through soundboard effects.",
	"Change your server avatar to match someone else's for a day.",
	"Create a role in a server with a ridiculous name and permissions.",
	"Make a tier list ranking your Discord friends and share it.",
	"Send a friend request to someone with a similar username to yours.",
	"Put your phone on speaker and call a friend to tell them you're playing Truth or Dare.",
}

func getRandomTruth() string {
	return truthQuestions[rand.Intn(len(truthQuestions))]
}

func getRandomDare() string {
	return dareQuestions[rand.Intn(len(dareQuestions))]
}