package main

import "github.com/nitehawk/advent-of-code/aoclib"

func botTimeStep(bot robo) robo {
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
	return bot
}

func botRenderState(bots []robo) {
}

func botEggSearch(bots []robo) {
	for b := 0; b < len(bots); b++ {
		bots[b] = botTimeStep(bots[b])
	}
	botRenderState(bots)
}

func puzzleb(inF string) int {
	botdesc := aoclib.ReadStringSlice(inF)
	bots := parseBots(botdesc)

	botEggSearch(bots)

	return 0
}
