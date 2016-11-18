package freeze

import (
	"math/rand"
	"time"
)

func getRandom(duration int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(duration)
}

// GetRandomElement returns a random element from a given slice
func GetRandomElement(tab []string) string {
	return tab[getRandom(len(tab))]
}

// Sleep sleeps for a random duration from 0 to 'duration' seconds
func Sleep(duration int) {
	random := getRandom(duration)
	time.Sleep(time.Second * time.Duration(random))
}

// SleepMinMax sleeps for a random duration from 'min' to 'max' seconds
func SleepMinMax(min, max int) {
	if max < min {
		temp := min
		min = max
		max = temp
	} else if max == min {
		max = min + 1
	}
	random := getRandom(max - min + 1)
	time.Sleep(time.Second * time.Duration(min+random))
}

func maybe(chance, totalChance int) bool {
	random := getRandom(totalChance)
	if random <= chance-1 {
		return true
	}
	return false
}

// MaybeSleepMinMax sleeps for a random duration from 'min' to 'max' seconds
// with a chance of 'chance' / 'totalChance'
func MaybeSleepMinMax(chance, totalChance, min, max int) {
	if maybe(chance, totalChance) {
		SleepMinMax(min, max)
	}
}
