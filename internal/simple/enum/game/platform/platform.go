package gamePlatform

type Enum struct {
	Type  string
	Value string
}

var (
	PC      = Enum{"PC", "PC"}
	Mobile  = Enum{"MOBILE", "Mobile"}
	Console = Enum{"CONSOLE", "Console"}
	Special = Enum{"SPECIAL", "Special"}
	Unknown = Enum{"UNKNOWN", "Unknown"}
)

func GetEnum(g string) Enum {
	switch g {
	case PC.Type:
		return PC
	case Mobile.Type:
		return Mobile
	case Console.Type:
		return Console
	case Special.Type:
		return Special
	}
	return Unknown
}
