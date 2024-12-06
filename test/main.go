package main

import "fmt"

// توزيع النمل عبر المسارات
func DistributeAnts(paths [][]int, ants int) map[int][]int {
	antDistribution := make(map[int][]int) // تخزين المسار لكل نملة

	for i := 0; i < ants; i++ {
		bestPath := 0 // افتراض أن المسار الأول هو الأفضل
		minSteps := len(paths[0]) + len(antDistribution[0])

		// اختيار أفضل مسار
		for j := 1; j < len(paths); j++ {
			steps := len(paths[j]) + len(antDistribution[j])
			if steps < minSteps {
				bestPath = j
				minSteps = steps
			}
		}

		// إضافة النملة إلى المسار الأفضل
		antDistribution[bestPath] = append(antDistribution[bestPath], i+1)
	}

	return antDistribution
}

func main() {
	paths := [][]int{
		{0, 1, 4},
		{0, 2, 3, 4},
		{0, 3, 4},
	}

	ants := 6
	distribution := DistributeAnts(paths, ants)

	fmt.Println(distribution)
}
