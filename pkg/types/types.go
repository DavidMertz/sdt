package types

// Structs to hold options and configurations
type (
	Options struct {
		Status      bool
		Semantic    bool
		Parsetree   bool
		Glob        string
		Dumbterm    bool
		Verbose     bool
		Source      string
		Destination string
	}

	Config struct {
		Description string             `toml:"description"`
		Commands    map[string]Command `toml:"commands"`
		Glob        string             `toml:"glob"`
	}

	Command struct {
		Executable string   `toml:"executable"`
		Switches   []string `toml:"switches"`
		Options    string   `toml:"options"`
	}

	Highlights struct {
		Add     string
		Del     string
		Header  string
		Info    string
		Clear   string
		Neutral string
	}
)

// In places, github.com/fatih/color is used, but raw ANSI is easier
// for writing custom reports based on sergi/go-diff/diffmatchpatch
var Colors Highlights = Highlights{
	Add:     "\x1b[32m", // green
	Del:     "\x1b[31m", // red
	Header:  "\x1b[33m", // yellow
	Info:    "\x1b[36m", // cyan
	Clear:   "\x1b[0m",
	Neutral: "",
}

var Dumbterm Highlights = Highlights{
	Add:     "{{+",
	Del:     "{{-",
	Header:  "",
	Info:    "",
	Clear:   "}}",
	Neutral: "{{_",
}

var PlainASCII Highlights = Highlights{
	Add:     "",
	Del:     "",
	Header:  "",
	Info:    "",
	Clear:   "",
	Neutral: "",
}

// Keep "enum" of types of parse trees we can handle here (<=256 for now)
type ParseType uint8

const (
	Ruby ParseType = iota
	Python
	JavaScript
	Golang
	SomeOtherLanguage
)
