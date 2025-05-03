package jokes_handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

var Categories = []string{"bad", "dad", "programming", "pun", "random"}

func getRandomJoke(category string) string {
	switch strings.ToLower(category) {
	case "bad":
		return BadJokes[rand.Intn(len(BadJokes))]
	case "dad":
		return DadJokes[rand.Intn(len(DadJokes))]
	case "programming":
		return ProgrammingJokes[rand.Intn(len(ProgrammingJokes))]
	case "pun":
		return PunJokes[rand.Intn(len(PunJokes))]
	default:
		allJokes := append(append(append(BadJokes, DadJokes...), ProgrammingJokes...), PunJokes...)
		return allJokes[rand.Intn(len(allJokes))]
	}
}

func JokesHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		category = "random"
	}

	joke := getRandomJoke(category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"handler": "/handlers/text/jokes","category": "%s", "joke": "%s", "status": "success"}`, category, joke)
}

var BadJokes = []string{
	"What do you call a fake noodle? An impasta!",
	"Why don't scientists trust atoms? Because they make up everything!",
	"I told my wife she was drawing her eyebrows too high. She looked surprised.",
	"Why don't eggs tell jokes? They'd crack each other up.",
	"I'm reading a book on anti-gravity. It's impossible to put down!",
	"Did you hear about the guy who invented the knock-knock joke? He won the 'no-bell' prize.",
	"What do you call a bear with no teeth? A gummy bear!",
	"Why did the scarecrow win an award? Because he was outstanding in his field!",
	"What do you call a fish with no eyes? Fsh!",
	"What's orange and sounds like a parrot? A carrot!",
	"Why can't you give Elsa a balloon? Because she'll let it go.",
	"I'm on a whiskey diet. I've lost three days already.",
	"What do you call a parade of rabbits hopping backwards? A receding hare-line.",
	"How do you organize a space party? You planet.",
	"Why don't some couples go to the gym? Because some relationships don't work out.",
	"What do you call a cow with no legs? Ground beef.",
	"How do you make holy water? You boil the hell out of it.",
	"What's the best thing about Switzerland? I don't know, but the flag is a big plus.",
	"I told my wife she should embrace her mistakes. She gave me a hug.",
	"Why couldn't the bicycle stand up by itself? It was two tired.",
	"What's brown and sticky? A stick.",
	"How do you make a tissue dance? You put a little boogie in it.",
	"What do you call cheese that isn't yours? Nacho cheese!",
	"I bought some shoes from a drug dealer. I don't know what he laced them with, but I've been tripping all day.",
	"What do you call a factory that makes good products? A satisfactory.",
	"Why couldn't the leopard play hide and seek? Because he was always spotted.",
	"I used to play piano by ear, but now I use my hands.",
	"What did one wall say to the other? I'll meet you at the corner!",
	"How do you count cows? With a cowculator.",
	"What's Forrest Gump's password? 1forrest1",
	"Why did the tomato turn red? Because it saw the salad dressing!",
	"What did the ocean say to the beach? Thanks for all the sediment.",
	"What did one hat say to another? You stay here, I'll go on ahead!",
	"I tried to catch fog yesterday. Mist.",
	"I used to have a job at a calendar factory, but I got fired because I took a couple of days off.",
	"What did the zero say to the eight? Nice belt!",
	"I finally got rid of that nasty cold. I just unfriended it on Facebook.",
	"Parallel lines have so much in common. It's a shame they'll never meet.",
	"What do you call a pile of cats? A meowntain.",
	"I accidentally swallowed some food coloring. The doctor says I'm OK, but I feel like I've dyed a little inside.",
}

var DadJokes = []string{
	"I'm afraid for the calendar. Its days are numbered.",
	"My wife said I should do lunges to stay in shape. That would be a big step forward.",
	"Why do fathers take an extra pair of socks when they go golfing? In case they get a hole in one!",
	"Singing in the shower is fun until you get soap in your mouth. Then it's a soap opera.",
	"What do you call a factory that makes okay products? A satisfactory.",
	"I've got a joke about construction, but I'm still working on it.",
	"I used to hate facial hair, but then it grew on me.",
	"I decided to sell my vacuum cleaner—it was just gathering dust.",
	"I had a neck brace fitted years ago and I've never looked back since.",
	"Did you hear about the cheese factory explosion? There was nothing left but de-brie.",
	"When two vegans get in an argument, is it still called a beef?",
	"I would tell you a chemistry joke but I know I wouldn't get a reaction.",
	"I only know 25 letters of the alphabet. I don't know y.",
	"I told my son I was named after Thomas Jefferson... He said, 'But dad, your name is Brian.' I said, 'I know, but I was named AFTER Thomas Jefferson.'",
	"A steak pun is a rare medium well done.",
	"What do you call someone with no body and no nose? Nobody knows.",
	"Why did the invisible man turn down the job offer? He couldn't see himself doing it.",
	"I'm so good at sleeping I can do it with my eyes closed.",
	"How do you get a squirrel to like you? Act like a nut.",
	"I bought a ceiling fan the other day. Complete waste of money. It just stands there applauding.",
	"I tell dad jokes, but I don't have any kids. I'm a faux pa.",
	"Someone asked me to name two structures that hold water. I said 'Well, dam'.",
	"Two goldfish are in a tank. One says to the other, 'Do you know how to drive this thing?'",
	"What did the baby corn say to the mama corn? 'Where's pop corn?'",
	"I ordered a chicken and an egg from Amazon. I'll let you know which comes first.",
	"My wife told me to stop impersonating a flamingo. I had to put my foot down.",
	"How do you make a Kleenex dance? Put a little boogie in it.",
	"I asked my wife if I was the only one she's been with. She said, 'Yes, the others were nines and tens.'",
	"I just watched a documentary about beavers. It was the best dam show I ever saw!",
	"If a child refuses to nap, are they guilty of resisting a rest?",
	"What do you call a belt made out of dollar bills? Cash pants.",
	"My wife said I need to stop acting like a flamingo, so I had to put my foot down.",
	"Why don't crabs donate? Because they're shellfish.",
	"What's E.T. short for? Because he's only got little legs.",
	"What did the janitor say when he jumped out of the closet? 'It's my time to shine!'",
	"What do you call a hippie's wife? Mississippi.",
	"Why did the scarecrow win the Nobel Prize? He was outstanding in his field.",
	"Where do boats go when they're sick? To the dock.",
	"When does a joke become a 'dad joke'? When it becomes apparent.",
	"I burned 2000 calories today. I left my pizza in the oven too long.",
}

var ProgrammingJokes = []string{
	"Why do programmers always mix up Halloween and Christmas? Because Oct 31 == Dec 25.",
	"There are only 10 types of people in the world: those who understand binary, and those who don't.",
	"A SQL query walks into a bar, walks up to two tables and asks, 'Can I join you?'",
	"Why did the programmer quit his job? Because he didn't get arrays.",
	"Why do Java developers wear glasses? Because they don't C#.",
	"How many programmers does it take to change a light bulb? None, that's a hardware problem.",
	"What's the object-oriented way to become wealthy? Inheritance.",
	"Why was the function sad after a party? It didn't get called.",
	"Knock knock. Who's there? [very long pause] Java.",
	"Why do programmers prefer dark mode? Because light attracts bugs.",
	"Why was the JavaScript developer sad? Because he didn't Node how to Express himself.",
	"A programmer puts two glasses on his bedside table before going to sleep. A full one, in case he gets thirsty, and an empty one, in case he doesn't.",
	"Why did the developer go broke? Because he used up all his cache.",
	"!false - It's funny because it's true.",
	"What is a programmer's favorite eyewear? CSS glasses, because they help you see styles.",
	"Why are assembly programmers always soaking wet? They work below C level.",
	"When a programmer is born, the first thing the doctor says is 'Hello, World!'",
	"Why do programmers hate nature? It has too many bugs and no debugger.",
	"Why was the constant always stressed? It was assigned too many values at a young age and couldn't change.",
	"Dev1: We should name this variable 'data'. Dev2: No, let's call it 'information'. Manager: I've solved your naming problem; call it 'datainformation'.",
	"What's a programmer's favorite place to hang out? Foo Bar.",
	"Two bytes meet. The first byte asks, 'Are you ill?' The second byte replies, 'No, just feeling a bit off.'",
	"I'd tell you a joke about UDP, but you might not get it.",
	"A byte walks into a bar and orders a pint. Bartender asks, 'What's wrong?' The byte says, 'Parity error.' Bartender nods and says, 'Yeah, I thought you looked a bit off.'",
	"Why did the programmer go broke? Because he lost his domain in a bet.",
	"What do you call 8 hobbits? A hobbyte.",
	"Why do C# and Java developers keep breaking their keyboards? Because they use a strongly typed language.",
	"What's a pirate's favorite programming language? R!",
	"What happens when a programmer dies? Their life pointer is null.",
	"How do you tell HTML from HTML5? Try it on IE. If it works, it's not HTML5.",
	"The glass is neither half-full nor half-empty. The glass is twice as large as it needs to be.",
	"Have you heard about the new Cray super computer? It's so fast, it executes an infinite loop in 6 seconds.",
	"A programmer had a problem. He decided to use Java. Now he has a ProblemFactory.",
	"What's the second movie about a database engineer called? The SQL.",
	"Why doesn't a developer ever make plans? Because they're always undefined.",
	"If you put a million monkeys at a million keyboards, one of them will eventually write a Java program. The rest of them will write Perl programs.",
	"Debugging is like being the detective in a crime movie where you're also the murderer.",
	"The six stages of debugging: 1. That can't happen. 2. That doesn't happen on my machine. 3. That shouldn't happen. 4. Why does that happen? 5. Oh, I see. 6. How did that ever work?",
	"Real programmers count from 0.",
	"My code doesn't work, I have no idea why. My code works, I have no idea why.",
}

var PunJokes = []string{
	"I'm on a seafood diet. Every time I see food, I eat it.",
	"What did the grape say when it got stepped on? Nothing, it just let out a little wine.",
	"What's the best time to go to the dentist? Tooth-hurty!",
	"I was going to make a belt made out of watches, but then I realized it would be a waist of time.",
	"I couldn't quite remember how to throw a boomerang, but eventually, it came back to me.",
	"Did you hear about the guy who got hit in the head with a can of soda? He was lucky it was a soft drink.",
	"I used to be afraid of hurdles, but I got over it.",
	"I was wondering why the ball was getting bigger. Then it hit me.",
	"I'm friends with all electricians. We have great current connections.",
	"The past, present, and future walked into a bar. It was tense.",
	"I don't trust stairs because they're always up to something.",
	"When life gives you melons, you're probably dyslexic.",
	"What do you call a fish wearing a crown? King of the sea!",
	"I would avoid the sushi if I were you. It's a little fishy.",
	"I can't believe I got fired from the calendar factory. All I did was take a day off.",
	"I went to buy some camouflage trousers yesterday but couldn't find any.",
	"What did the shark say when he ate the clownfish? This tastes a little funny.",
	"What do you call a deer with no eyes? No idea (no-eye deer).",
	"People who use selfie sticks really need to stick to themselves.",
	"I told my computer I needed a break, and now it won't stop sending me vacation ads.",
	"What do you call a pony with a sore throat? A little horse.",
	"What did the sea say to the shore? Nothing, it just waved.",
	"Have you heard about the corduroy pillow? It's making headlines everywhere!",
	"What do you call a belt made of watches? A waist of time.",
	"Time flies like an arrow. Fruit flies like a banana.",
	"When the power went out at the school, the principal was de-lighted.",
	"The shovel was a ground-breaking invention.",
	"I used to be a banker, but I lost interest.",
	"I'm reading a book on the history of glue. I just can't seem to put it down.",
	"The rotation of earth really makes my day.",
	"I stayed up all night wondering where the sun went. Then it dawned on me.",
	"I used to be addicted to the hokey pokey, but I turned myself around.",
	"I tried to catch some fog, but I mist.",
	"A bicycle can't stand on its own because it's two-tired.",
	"I used to be afraid of hurdles, but I got over it.",
	"A chicken crossing the road is poultry in motion.",
	"When a clock is hungry, it goes back four seconds.",
	"What's the definition of a good farmer? A man outstanding in his field.",
	"I did a theatrical performance about puns. It was just a play on words.",
	"England doesn't have a kidney bank, but it does have a Liverpool.",
	"Did you hear about the fire at the circus? It was in tents!",
	"I've been told I'm condescending. That means I talk down to people.",
	"Atheism is a non-prophet organization.",
	"RIP boiled water. You will be mist.",
	"Need an ark? I Noah guy.",
}
