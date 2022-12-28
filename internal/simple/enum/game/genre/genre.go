package gameGenre

type Enum struct {
	Type  string
	Value string
}

var (
	Action  = Enum{"ACTION", "Action"}
	RPG     = Enum{"RPG", "RPG"}
	JRPG    = Enum{"JRPG", "JRPG"}
	MMORPG  = Enum{"MMORPG", "MMORPG"}
	Gacha   = Enum{"Gacha", "GACHA"}
	Idle    = Enum{"Idle", "IDLE"}
	Ecchi   = Enum{"Ecchi", "Ecchi"}
	Puzzle  = Enum{"PUZZLE", "Puzzle"}
	Unknown = Enum{"UNKNOWN", "Unknown"}
)

func GetEnum(g string) Enum {
	switch g {
	case Action.Type:
		return Action
	case RPG.Type:
		return RPG
	case JRPG.Type:
		return JRPG
	case MMORPG.Type:
		return MMORPG
	case Gacha.Type:
		return Gacha
	case Idle.Type:
		return Idle
	case Ecchi.Type:
		return Ecchi
	case Puzzle.Type:
		return Puzzle
	}
	return Unknown
}
