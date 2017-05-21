package house

var testVersion = 1

var objects []string = []string{
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}
var actions []string = []string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

func constructVerse(num int, current int) string {
	if current < 0 {
		return ""
	}
	if current == num {
		return "This is " + objects[current] + constructVerse(num, current-1)
	}

	return "\nthat " + actions[current] + " " + objects[current] + constructVerse(num, current-1)
}

func Verse(num int) string {
	return constructVerse(num-1, num-1)
}

func Song() string {
	total := Verse(1)
	for i := 2; i <= 12; i++ {
		total = total + "\n\n" + Verse(i)
	}

	return total
}
