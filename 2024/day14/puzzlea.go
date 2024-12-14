package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/nitehawk/advent-of-code/aoclib"
)

type robo struct {
	sx, sy int
	vx, vy int
	ex, ey int
}

func parseBots(botdesc []string) []robo {
	bots := make([]robo, 0)
	// p=0,4 v=3,-3
	re := regexp.MustCompile(`p=([0-9]+),([0-9]+) v=([\-0-9]+),([\-0-9]+)`)
	for _, line := range botdesc {
		matches := re.FindAllStringSubmatch(line, -1)
		bot := robo{}
		bot.sx, _ = strconv.Atoi(matches[0][1])
		bot.sy, _ = strconv.Atoi(matches[0][2])
		bot.vx, _ = strconv.Atoi(matches[0][3])
		bot.vy, _ = strconv.Atoi(matches[0][4])
		bot.ex = bot.sx
		bot.ey = bot.sy

		bots = append(bots, bot)
	}
	return bots
}

// Step the bot forward, return the new bot status
func botFlyer(bot robo) robo {
	for t := 0; t < bottimer; t++ {
		bot.ex += bot.vx
		bot.ey += bot.vy

		// Wrap around
		if bot.ex < 0 {
			bot.ex += spacew
		} else if bot.ex >= spacew {
			bot.ex -= spacew
		}
		if bot.ey < 0 {
			bot.ey += spaceh
		} else if bot.ey >= spaceh {
			bot.ey -= spaceh
		}
	}
	return bot
}

func workerFlyBots(wg *sync.WaitGroup, botinfo chan robo, flythem []robo) {
	defer wg.Done()
	for _, bot := range flythem {
		botinfo <- botFlyer(bot)
	}
}

func workerQCounter(wg *sync.WaitGroup, botinfo chan robo, factor chan int) {
	defer wg.Done()

	qbots := [4]int{0, 0, 0, 0}
	safety := 1
	botscounted := 0

	midx := (spacew - 1) / 2
	midy := (spaceh - 1) / 2

	fmt.Printf("Counting for %d,%d space - midline: %d,%d\n", spacew, spaceh, midx, midy)
	debug := ""

	for {
		select {
		case bot := <-botinfo:
			botscounted++
			if bot.ex == midx || bot.ey == midy { // If the bot is on the middle line, ignore it
				debug = fmt.Sprintf("SF Bot (%d, %d)-(%d,%d): e: (%d, %d)\n", bot.sx, bot.sy, bot.vx, bot.vy, bot.ex, bot.ey)
				continue
			}
			if bot.ex < midx && bot.ey < midy { // top left
				debug = fmt.Sprintf("TL Bot (%d, %d)-(%d,%d): e: (%d, %d)\n", bot.sx, bot.sy, bot.vx, bot.vy, bot.ex, bot.ey)
				qbots[0]++
			} else if bot.ex > midx && bot.ey < midy { // top right
				debug = fmt.Sprintf("TR Bot (%d, %d)-(%d,%d): e: (%d, %d)\n", bot.sx, bot.sy, bot.vx, bot.vy, bot.ex, bot.ey)
				qbots[1]++
			} else if bot.ex < midx && bot.ey > midy { // bottom left
				debug = fmt.Sprintf("BL Bot (%d, %d)-(%d,%d): e: (%d, %d)\n", bot.sx, bot.sy, bot.vx, bot.vy, bot.ex, bot.ey)
				qbots[2]++
			} else if bot.ex > midx && bot.ey > midy { // bottom right
				debug = fmt.Sprintf("BR Bot (%d, %d)-(%d,%d): e: (%d, %d)\n", bot.sx, bot.sy, bot.vx, bot.vy, bot.ex, bot.ey)
				qbots[3]++
			}
			if verbose {
				fmt.Println(debug)
			}

		case <-time.After(time.Second * 1):
			for _, q := range qbots {

				safety *= q
			}
			factor <- safety
			fmt.Println("Counted ", botscounted, " Bots. Safety: ", safety, " Quadrants: ", qbots)
			return
		}
	}
}

// Setup puzzle parameters - Full problem
const (
	bottimer   = 100
	spacew     = 101
	spaceh     = 103
	flyworkers = 5
	verbose    = false
)

// Setup puzzle parameters - Test problem
/*
const (
	bottimer   = 100
	spacew     = 11
	spaceh     = 7
	flyworkers = 5
	verbose    = false
)
*/

func puzzlea(inF string) int {
	botdesc := aoclib.ReadStringSlice(inF)
	bots := parseBots(botdesc)

	//	fmt.Println(len(bots), bots)

	// Setup channels
	botinfo := make(chan robo, 50)
	factor := make(chan int)

	// Setup sync wait group
	wg := sync.WaitGroup{}
	wg.Add(1 + flyworkers)

	// Start quadrant qualifier
	go workerQCounter(&wg, botinfo, factor)

	// split the bots into groups for each fly worker
	// Start bot flyers
	for i := 1; i <= flyworkers; i++ {
		if i == flyworkers {
			// Send the remainder to the last worker
			go workerFlyBots(&wg, botinfo, bots[((i-1)/len(bots))/flyworkers:])
		} else {
			go workerFlyBots(&wg, botinfo, bots[((i-1)/len(bots))/flyworkers:((i)/len(bots))/flyworkers])
		}
	}

	// Get Safety
	safety := <-factor

	// Sync up
	wg.Wait()

	return safety
}
