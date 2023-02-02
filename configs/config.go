package configs

type Config struct {
	StutterChance int      `yaml:"stutter_chance,omitempty"`
	EmojiChance   int      `yaml:"emoji_chance,omitempty"`
	Emojis        []string `yaml:"emojis,omitempty"`
	Replacer      []string `yaml:"replacer,omitempty"`
}

func (config Config) RunicEmojis() [][]rune {
	var runes [][]rune
	for _, emoji := range config.Emojis {
		runes = append(runes, []rune(emoji))
	}
	return runes
}
