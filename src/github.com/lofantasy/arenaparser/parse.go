package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lofantasy/configs"
)

type games struct {
	Games []configs.MatchStats
}

func main() {
	fmt.Println("Starting")

	//day1Games := parseFile("e:/programming/go/wowparser/day1.txt")
	//day2Games := parseFile("e:/programming/go/wowparser/day2.txt")
	//day2Games := parseFile("e:/programming/go/wowparser/day2.txt")
	//day2Games := parseFile("e:/programming/go/wowparser/day2.txt")
	//day4Games := parseFile("e:/programming/go/wowparser/day4.txt")
	day5Games := parseFile("e:/programming/go/wowparser/day5.txt")

	results5 := parseGameResults(day5Games, 1368, 5)

	// rst := echoResults(day1Games, 0)
	// echoResults(day2Games, rst)
	//rst := echoResults(day4Games, 1349)
	fmt.Println(echoResults(day5Games, 1368, results5)) // 1368
	fmt.Println(results5)
}

func echoResults(gamesPlayed games, ratingChange int64, dayStats configs.DayStats) int64 {
	// var ratingChange int64
	var healerComps int64
	var dpsComps int64
	var totalTeams int64
	var wins int64
	var winsVsHeals, lossesVsHeals, winsVsDps, lossesVsDps int64
	var losses int64

	fmt.Println("Day Starting rating: ", ratingChange)
	for _, elem := range gamesPlayed.Games {
		// fmt.Println(elem.RatingChange)
		fmt.Println(elem)

		ratingChange += elem.RatingChange
		isHeal, isDps, healsWin, healsLoss, dpsWin, dpsLoss := teamTypeCheck(elem.EnemySpecA, elem.EnemySpecB, elem.RatingChange)

		healerComps += isHeal
		dpsComps += isDps
		totalTeams++

		winsVsHeals += healsWin
		lossesVsHeals += healsLoss
		winsVsDps += dpsWin
		lossesVsDps += dpsLoss

		if elem.RatingChange > 0 {
			wins++
		} else {
			losses++
		}
	}
	// var deltaW, deltaL float64
	// w1 := float64(wins)
	// l1 := float64(losses)
	// total := float64(totalTeams)

	// deltaW = (w1 / total)
	// deltaL = (l1 / total)

	deltaW, deltaL := computeWinLoss(totalTeams, wins, losses)

	fmt.Println("  Day ending rating: ", ratingChange)
	fmt.Println("              Teams: ", totalTeams)
	fmt.Println("       HEaler Comps: ", healerComps)
	fmt.Println("               Wins: ", winsVsHeals)
	fmt.Println("             Losses: ", lossesVsHeals)
	fmt.Println("          DPS Comps: ", dpsComps)
	fmt.Println("               Wins: ", winsVsDps)
	fmt.Println("             Losses: ", lossesVsDps)
	fmt.Println("            Overall: ")
	fmt.Println("               Wins: ", wins)
	fmt.Printf("                   : %0.2f\n", (deltaW * 100))
	fmt.Println("             Losses: ", losses)
	fmt.Printf("                   : %0.2f\n", (deltaL * 100))

	fmt.Println("---------------")
	fmt.Println("Day #", 1)
	fmt.Println("Starting Rating: ", dayStats.StartingRating)
	fmt.Println("  Ending Rating: ", dayStats.EndingRating)
	fmt.Println("    Rating Gain: ", (dayStats.EndingRating - dayStats.StartingRating))
	fmt.Println("Overall")
	fmt.Println("  Games: ", dayStats.TotalGames)
	fmt.Println("   Wins: ", dayStats.TotalWins)
	fmt.Println(" Losses: ", dayStats.TotalLosses)
	fmt.Println("   Win%: ", dayStats.TotalWinPercent)
	fmt.Printf("   Loss: %0.2f\n", (dayStats.TotalLossPercent * 100))
	fmt.Println("")
	fmt.Println("Healers")
	fmt.Println("  Games: ", dayStats.HealerGames)
	fmt.Println("   Wins: ", dayStats.HealerWins)
	fmt.Println(" Losses: ", dayStats.HealerLosses)
	fmt.Printf("    Win: %0.2f\n", (dayStats.HealerWinPercent * 100))
	fmt.Printf("   Loss: %0.2f\n", (dayStats.HealerLossPercent * 100))
	fmt.Println("")
	fmt.Println("DPS")
	fmt.Println("  Games: ", dayStats.DPSGames)
	fmt.Println("   Wins: ", dayStats.DPSWins)
	fmt.Println(" Losses: ", dayStats.DPSLosses)
	fmt.Printf("    Win: %0.2f\n", (dayStats.DPSWinPercent * 100))
	fmt.Printf("   Loss: %0.2f\n", (dayStats.DPSLossPercent * 100))

	fmt.Println("")

	return ratingChange
}

func parseGameResults(gamesPlayed games, startingRating int64, day int64) configs.DayStats {
	results := configs.DayStats{}

	results.Day = day
	var endingRating int64
	endingRating += startingRating
	results.StartingRating = startingRating

	for _, elem := range gamesPlayed.Games {
		fmt.Println("parsing: ", elem)
		endingRating += elem.RatingChange
		results.TotalGames++

		if elem.RatingChange > 0 {
			results.TotalWins++
		} else {
			results.TotalLosses++
		}

		if elem.Heal.HealComp {
			results.HealerGames++
			if elem.Heal.Win {
				results.HealerWins++
			} else {
				results.HealerLosses++
			}
		}

		if elem.Dps.DpsComp {
			results.DPSGames++
			if elem.Dps.Win {
				results.DPSWins++
			} else {
				results.DPSLosses++
			}
		}
	}

	deltaW, deltaL := computeWinLoss(results.TotalGames, results.TotalWins, results.TotalLosses)

	results.EndingRating = endingRating
	results.TotalWinPercent = deltaW
	results.TotalLossPercent = deltaL

	deltaHW, deltaHL := computeWinLoss(results.HealerGames, results.HealerWins, results.HealerLosses)
	results.HealerWinPercent = deltaHW
	results.HealerLossPercent = deltaHL

	deltaDW, deltaDL := computeWinLoss(results.DPSGames, results.DPSWins, results.DPSLosses)
	results.DPSWinPercent = deltaDW
	results.DPSLossPercent = deltaDL

	fmt.Println("Ending REsults: ", results)
	return results
}

func computeWinLoss(totalGames int64, wins int64, losses int64) (float64, float64) {
	var deltaW, deltaL float64
	w1 := float64(wins)
	l1 := float64(losses)
	total := float64(totalGames)

	deltaW = (w1 / total)
	deltaL = (l1 / total)

	return deltaW, deltaL
}

func teamTypeCheck(spec1 string, spec2 string, ratingChange int64) (int64, int64, int64, int64, int64, int64) {

	var heals, dps, healsWin, healsLoss, dpsWin, dpsLoss int64

	if spec1 == "Resto" || spec1 == "Holy" || spec1 == "MW" || spec1 == "Disc" ||
		spec2 == "Resto" || spec2 == "Holy" || spec2 == "MW" || spec2 == "Disc" {
		heals = 1

		if ratingChange > 0 {
			healsWin++
		} else {
			healsLoss++
		}
	} else {
		dps = 1

		if ratingChange > 0 {
			dpsWin++
		} else {
			dpsLoss++
		}
	}

	return heals, dps, healsWin, healsLoss, dpsWin, dpsLoss
}

func parseFile(infile string) games {
	games := games{}

	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var i int64 = 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		fileName := line[0]
		matchingA, _ := strconv.ParseInt(line[1], 10, 64)
		matchingB, _ := strconv.ParseInt(line[2], 10, 64)
		newRating, _ := strconv.ParseInt(line[3], 10, 64)
		ratingChange, _ := strconv.ParseInt(line[4], 10, 64)
		enemyRatingChangeA, _ := strconv.ParseInt(line[5], 10, 64)
		enemyRatingChangeB, _ := strconv.ParseInt(line[6], 10, 64)
		enemyClassA := line[7]
		enemySpecA := line[8]
		enemyClassB := line[9]
		enemySpecB := line[10]
		healerComp := configs.HealerStats{
			false, false, false,
		}
		dpsComp := configs.DpsStats{
			false, false, false,
		}

		isHeal, isDps, healsWin, healsLoss, dpsWin, dpsLoss := teamTypeCheck(enemySpecA, enemySpecB, ratingChange)

		if isHeal == 1 {
			healerComp.HealComp = true
			healerComp.Win = healsWin == 1
			healerComp.Loss = healsLoss == 1
		}
		if isDps == 1 {
			dpsComp.DpsComp = true
			dpsComp.Win = dpsWin == 1
			dpsComp.Loss = dpsLoss == 1
		}

		games.Games = append(games.Games, configs.MatchStats{
			i,
			fileName,
			matchingA,
			matchingB,
			newRating,
			ratingChange,
			enemyRatingChangeA,
			enemyRatingChangeB,
			enemyClassA,
			enemySpecA,
			enemyClassB,
			enemySpecB,
			healerComp,
			dpsComp,
		})
		i++
	}

	return games
}
