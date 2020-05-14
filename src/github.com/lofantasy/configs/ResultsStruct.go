package configs

// DayStats - for the day
type DayStats struct {
	Day               int64
	StartingRating    int64
	EndingRating      int64
	TotalGames        int64
	TotalWins         int64
	TotalLosses       int64
	TotalWinPercent   float64
	TotalLossPercent  float64
	HealerGames       int64
	HealerWins        int64
	HealerLosses      int64
	HealerWinPercent  float64
	HealerLossPercent float64
	DPSGames          int64
	DPSWins           int64
	DPSLosses         int64
	DPSWinPercent     float64
	DPSLossPercent    float64
}
