package uwuify

import (
	"git-uwu/configs"
	"math/rand"
	"strings"
	"unicode"
)

var Emojis = configs.RunningConfiguration.RunicEmojis()

// StutterChance determines the chances of a stutter, for example, the default value is 10 which means
// there is a 1 out of 10 chance that a stutter will be encountered.
var StutterChance = configs.RunningConfiguration.StutterChance
var EmojiChance = configs.RunningConfiguration.EmojiChance

// Replacer handles replacing different keywords into cuter keywords.
var Replacer = strings.NewReplacer(configs.RunningConfiguration.Replacer...)

// From creates an uwuified variant of the given text based on a simplified implementation of https://github.com/Daniel-Liu-c0deb0t/uwu
// which is less performant and efficient but does the work nicely.
func From(text string) string {
	var result []rune

	var currentCharacter rune
	var previousCharacter rune

	// IMPORTANT: Hold a reference copy of the emojis. Although this will waste memory in this application's implementation,
	// but for external applications, this will be needed especially when there is a chance of the emojis being modified.
	emojis, emojisLen := Emojis, len(Emojis)

	// IMPORTANT: Inefficiently replaces all the keywords into a cuter variant before we start processing.
	text = Replacer.Replace(text)
	runes := []rune(text)

	ignoring := false
	for index, char := range runes {
		previousCharacter = currentCharacter
		currentCharacter = unicode.ToLower(char)

		// ADDITIONAL RULE: To ignore a specific region  of text, you have to add <: at the start and end with :>
		if previousCharacter == '<' && currentCharacter == ':' {
			result = result[:(len(result) - 1)]

			ignoring = true
			continue
		}

		if ignoring {
			if previousCharacter == ':' && currentCharacter == '>' {
				result = result[:(len(result) - 1)]
				ignoring = false
			} else {
				result = append(result, currentCharacter)
			}
			continue
		}

		// ADDITIONAL RULE: Ignore all symbols since we cannot make modifications on them.
		if !unicode.IsLetter(currentCharacter) {
			result = append(result, currentCharacter)
			continue
		}

		// ADDITIONAL RULE: We shouldn't replace words such as 'ur' and related since that sounds weird.
		if currentCharacter == 'l' || currentCharacter == 'r' && (previousCharacter != 'u' && runes[index+1] != ' ') {
			result = append(result, 'w')
			continue
		}

		if currentCharacter == 'o' && previousCharacter == 'n' {
			result = append(result, 'y', 'o')
			continue
		}

		if currentCharacter == ' ' &&
			(previousCharacter == '!' || previousCharacter == '.' || previousCharacter == '?') &&
			rand.Intn(EmojiChance) == 1 {
			result = append(result, ' ')
			result = append(result, emojis[rand.Intn(emojisLen)]...)

			result = append(result, currentCharacter)
			continue
		}

		if previousCharacter == ' ' && unicode.IsLetter(currentCharacter) && rand.Intn(StutterChance) == 1 {
			result = append(result, currentCharacter, '-', currentCharacter)
			continue
		}

		result = append(result, currentCharacter)
	}
	return string(result)
}
