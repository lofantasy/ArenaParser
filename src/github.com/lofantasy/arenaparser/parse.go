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
	fmt.Println("Starting Day 1")

	//day1Games := parseFile("e:/programming/go/wowparser/day1.txt")
	//day2Games := parseFile("e:/programming/go/wowparser/day2.txt")
	//day2Games := parseFile("e:/programming/go/wowparser/day2.txt")
	//day2Games := parseFile("e:/programming/go/wowparser/day2.txt")
	day4Games := parseFile("e:/programming/go/wowparser/day4.txt")
	day5Games := parseFile("e:/programming/go/wowparser/day5.txt")

	// rst := echoResults(day1Games, 0)
	// echoResults(day2Games, rst)
	echoResults(day4Games, 1349)
	echoResults(day5Games, 1368)
}

func echoResults(gamesPlayed games, ratingChange int64) int64 {
	// var ratingChange int64
	var healerComps int64
	var dpsComps int64
	var totalTeams int64
	var wins int64
	var losses int64

	for _, elem := range gamesPlayed.Games {
		// fmt.Println(elem.RatingChange)
		ratingChange += elem.RatingChange
		isHeal, isDps := teamTypeCheck(elem.EnemySpecA, elem.EnemySpecB)

		healerComps += isHeal
		dpsComps += isDps
		totalTeams++

		if elem.RatingChange > 0 {
			wins++
		} else {
			losses++
		}
	}
	var deltaW, deltaL float64
	w1 := float64(wins)
	l1 := float64(losses)
	total := float64(totalTeams)

	deltaW = (w1 / total)
	deltaL = (l1 / total)

	fmt.Println("Day ending rating: ", ratingChange)
	fmt.Println("            Teams: ", totalTeams)
	fmt.Println("     HEaler Comps: ", healerComps)
	fmt.Println("        DPS Comps: ", dpsComps)
	fmt.Println("             Wins: ", wins)
	fmt.Printf("                  : %0.2f\n", (deltaW * 100))
	fmt.Println("           Losses: ", losses)
	fmt.Printf("                  : %0.2f\n", (deltaL * 100))
	return ratingChange
}

func teamTypeCheck(spec1 string, spec2 string) (int64, int64) {

	var heals int64
	var dps int64

	if spec1 == "Resto" || spec1 == "Holy" || spec1 == "MW" || spec1 == "Disc" ||
		spec2 == "Resto" || spec2 == "Holy" || spec2 == "MW" || spec2 == "Disc" {
		heals = 1
	} else {
		dps = 1
	}

	return heals, dps
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
		})
		i++
	}

	return games
}
