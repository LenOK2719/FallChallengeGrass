package main

import "fmt"
import "os"

func main() {
	var width, height int
	fmt.Scan(&width, &height)

	for {
		var myRobots []robot
		var enemyRobots []robot
		var myMatter, oppMatter int
		fmt.Scan(&myMatter, &oppMatter)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				// owner: 1 = me, 0 = foe, -1 = neutral
				var scrapAmount, owner, units, recycler, canBuild, canSpawn, inRangeOfRecycler int
				fmt.Scan(&scrapAmount, &owner, &units, &recycler, &canBuild, &canSpawn, &inRangeOfRecycler)
				if owner == 1 && units > 0 {
					for r := 0; r < units; r++ {
						myRobots = append(myRobots, robot{x: i, y: j})
					}
				} else if owner == 0 && units > 0 {
					for r := 0; r < units; r++ {
						enemyRobots = append(enemyRobots, robot{x: i, y: j})
					}
				}
			}
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// fmt.Fprintln(os.Stderr, myRobots[0].x)
		fmt.Fprintln(os.Stderr, len(myRobots))
		for i := 0; i < len(myRobots); i++ {
			fmt.Printf("MOVE 1 %v %v %v %v;", myRobots[i].y, myRobots[i].x, enemyRobots[len(enemyRobots)-1].y-1, enemyRobots[len(enemyRobots)-1].x-1) // Write action to stdout
		}
		fmt.Printf("SPAWN 1 %v %v;", myRobots[0].y, myRobots[0].x)
		fmt.Printf("\n")
		// fmt.Printf
	}
}

type robot struct {
	x      int
	y      int
	isMine bool
}
