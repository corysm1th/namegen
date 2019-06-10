package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// Case represents the case style for the name
type Case uint

const (
	// Snake looks like "snake_case"
	Snake = iota
	// Camel looks like "camelCase"
	// TODO: Camel
	// Pascal looks like "PascalCase"
	// TODO: Pascal
)

// Things returns a random name for a group of things nobody would ever want
func Things(c Case) string {
	names := []string{"pogo_sticks", "pogo_balls", "ET_fingers", "lawn_darts", "balzaks", "moon_shoes", "3_wheelers", "socker_boppers", "slinkies", "magic_mitts", "bopits", "slap_wraps", "blue_blockers"}
	colors := []string{"red", "pink", "orange", "yellow", "green", "blue", "purple", "magenta"}
	adjectives := []string{"exploding", "rubberized", "electrocuting", "vengeful", "malicious", "firery", "flaming", "flogging", "pulsating", "pulverizing", "sonic", "radioactive", "melting", "transient", "tickling", "teleporting"}

	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	nameIndex := gen.Intn(len(names))
	name := names[nameIndex]

	colorIndex := gen.Intn(len(colors))
	color := colors[colorIndex]

	adjIndex := gen.Intn(len(adjectives))
	adj := adjectives[adjIndex]

	time.Sleep(time.Duration(nameIndex) * time.Nanosecond)
	return fmt.Sprintf("%s_%s_%s", color, adj, name)
}
