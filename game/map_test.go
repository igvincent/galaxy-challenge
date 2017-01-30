package game

import (
	"github.com/magleff/galaxy-challenge/dto"
	"testing"
)

func TestInitDistanceMapEmptyPlanets(t *testing.T) {
	// Arrange
	gameMap := &Map{}

	// Act
	err := gameMap.InitDistanceMap()

	// Assert
	if err == nil {
		t.Error("TestInitDistanceMapEmptyPlanets: should trigger an error if the planets map is empty.")
	}
}

func TestInitDistanceMap(t *testing.T) {
	// Arrange
	var testCases = []struct {
		Input             []dto.StatusPlanet
		InputSourcePlanet uint16
		InputTargetPlanet uint16
		ExpectedDistance  float64
	}{
		{
			[]dto.StatusPlanet{
				{ID: 1, X: 0, Y: 0},
				{ID: 2, X: 15, Y: 0},
				{ID: 3, X: 5, Y: 0},
			},
			1, 3, 5,
		},
	}

	for index, testCase := range testCases {
		// Arrange
		gameMap := &Map{
			Planets: testCase.Input,
		}

		// Act
		gameMap.InitDistanceMap()

		// Assert
		distance := gameMap.DistanceMap[testCase.InputSourcePlanet][testCase.InputTargetPlanet]

		if testCase.ExpectedDistance != distance {
			t.Errorf("TestInitDistanceMap(%d): expected %s, actual %s", index,
				testCase.ExpectedDistance, distance)
		}
	}
}

func TestUpdateMap(t *testing.T) {
	// Arrange
	gameMap := &Map{}

	status := dto.Status{
		Planets: []dto.StatusPlanet{
			{}, {},
		},
		Fleets: []dto.StatusFleet{
			{}, {}, {},
		},
	}

	// Act
	gameMap.Update(status)

	// Assert
	if len(gameMap.Planets) != len(status.Planets) {
		t.Errorf("There should be %d planets after update.", len(status.Planets))
	}

	if len(gameMap.Fleets) != len(status.Fleets) {
		t.Errorf("There should be %d planets after update.", len(status.Fleets))
	}
}

func TestComputeTurnsLeft(t *testing.T) {
	// Arrange
	gameMap := &Map{
		Planets: []dto.StatusPlanet{
			{ID: 1, X: 0, Y: 0},
			{ID: 2, X: 40, Y: 0},
			{ID: 3, X: 45, Y: 0},
		},
	}

	gameMap.InitDistanceMap()

	testCases := []struct {
		SourceID uint16
		TargetID uint16
		Expected uint16
	}{
		{1, 2, 2},
		{1, 3, 2},
		{2, 3, 0},
	}

	for index, testCase := range testCases {
		// Act
		actual := gameMap.computeTurnsLeft(testCase.SourceID, testCase.TargetID)

		// Assert
		if actual != testCase.Expected {
			t.Errorf("TestComputeTurnsLeft(%d): expected %d, was %d", index, testCase.Expected, actual)
		}
	}
}
