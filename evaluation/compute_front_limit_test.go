package evaluation

import (
	"testing"
)

func TestComputeFrontLimit(t *testing.T) {
	// Arrange
	testCases := []struct {
		Planets         []*ResultPlanet
		MaxFrontPlanets int
		Expected        int16
	}{
		{
			[]*ResultPlanet{
				{DistanceToEnemy: 3},
			},
			0,
			0,
		},
		{
			[]*ResultPlanet{},
			1,
			0,
		},
		{
			[]*ResultPlanet{
				{DistanceToEnemy: 3}, {DistanceToEnemy: 5},
			},
			1,
			3,
		},
		{
			[]*ResultPlanet{
				{DistanceToEnemy: 3}, {DistanceToEnemy: 5},
			},
			5,
			5,
		},
	}

	for i, testCase := range testCases {
		// Act
		actual := ComputeFrontLimit(testCase.Planets, testCase.MaxFrontPlanets)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeFrontLimit(%d): expected %f, was %f", i, testCase.Expected, actual)
		}
	}
}
