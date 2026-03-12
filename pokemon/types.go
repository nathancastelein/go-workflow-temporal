package pokemon

// Pokemon represents a Pokemon with its current and maximum health points.
type Pokemon struct {
	Name  string
	Type  string
	HP    int
	MaxHP int
}

// CaptureResult represents the outcome of a capture attempt.
type CaptureResult struct {
	Success bool
	Pokemon Pokemon
}

// EvolutionResult represents the outcome of the evolution chamber.
type EvolutionResult struct {
	Pokemon Pokemon
	Evolved bool
	Trigger string // "timer", "feed", "cancelled"
}

// BattleResult represents the outcome of a single battle.
type BattleResult struct {
	Winner string
	Loser  string
}

// TournamentResult represents the full tournament results.
type TournamentResult struct {
	Bracket  [][]string // rounds of winners
	Champion string
}
