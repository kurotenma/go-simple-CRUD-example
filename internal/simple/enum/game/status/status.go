package gameStatus

type Enum struct {
	Type  string
	Value string
}

var (
	NotRegistered = Enum{"NOT_REGISTERED", "Not Registered"}
	Registered    = Enum{"VERIFIED", "Verified"}
	Unknown       = Enum{"UNKNOWN", "Unknown"}
)

func GetEnum(g string) Enum {
	switch g {
	case NotRegistered.Type:
		return NotRegistered
	case Registered.Type:
		return Registered
	}
	return Unknown
}
