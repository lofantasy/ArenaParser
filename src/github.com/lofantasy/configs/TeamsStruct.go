package configs

// MatchStats for full on stats of the game
type MatchStats struct {
	Game               int64
	Filename           string
	MatchingA          int64
	MatchingB          int64
	NewRating          int64
	RatingChange       int64
	EnemyRatingChangeA int64
	EnemyRatingChangeB int64
	EnemyClassA        string
	EnemySpecA         string
	EnemyClassB        string
	EnemySpecB         string
	Heal               HealerStats
	Dps                DpsStats
}

// HealerStats - for Healer Comps.
type HealerStats struct {
	HealComp bool
	Win      bool
	Loss     bool
}

// DpsStats - for Healer Comps.
type DpsStats struct {
	DpsComp bool
	Win     bool
	Loss    bool
}

// TeamComp for Enemey team stats.
type TeamComp struct {
	PlayerAClass        string
	PlayerASpec         string
	PlayerBClass        string
	PlayerBSpec         string
	PlayerARatingChange int64
	PlayerBRatingChange int64
}

// PlayerComp for Players team stats.
type PlayerComp struct {
}

// GampeComp do things.
type GameComp struct {
	MatchingA    int64
	MatchingB    int64
	NewRating    int64
	RatingChange int64
}
