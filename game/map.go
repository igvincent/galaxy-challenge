package game

import (
	"errors"
	"github.com/magleff/galaxy-challenge/common"
	"github.com/magleff/galaxy-challenge/dto"
	"math"
)

type Map struct {
	Planets     []dto.StatusPlanet
	Fleets      []dto.StatusFleet
	DistanceMap map[uint16]map[uint16]float64
	Turn        uint16
	Initialized bool
}

func (m *Map) InitDistanceMap() error {
	if m.Planets == nil {
		return errors.New("The planets have not been initialized.")
	}

	m.DistanceMap = make(map[uint16]map[uint16]float64)

	for _, planet := range m.Planets {
		m.DistanceMap[planet.ID] = initDistanceMapForPlanet(planet, m.Planets)
	}

	return nil
}

func (m *Map) Update(status dto.Status) {
	m.Planets = status.Planets
	m.Fleets = status.Fleets
	m.Turn = status.Config.Turn
}

func (m Map) MapMoveFleet(playerID uint16, moveFleet dto.MoveFleet) dto.StatusFleet {
	return dto.StatusFleet{
		OwnerID:  playerID,
		SourceID: moveFleet.SourceID,
		TargetID: moveFleet.TargetID,
		Units:    moveFleet.Units,
		Left:     m.computeTurnsLeft(moveFleet.SourceID, moveFleet.TargetID),
	}
}

func (m Map) computeTurnsLeft(sourceID uint16, targetID uint16) uint16 {
	return uint16(math.Floor(m.DistanceMap[sourceID][targetID] / common.DISTANCE_PER_TURN))
}

func initDistanceMapForPlanet(planet dto.StatusPlanet, otherPlanets []dto.StatusPlanet) map[uint16]float64 {
	planetDistanceMap := make(map[uint16]float64)

	for _, otherPlanet := range otherPlanets {
		planetDistanceMap[otherPlanet.ID] = computeDistance(planet, otherPlanet)
	}

	return planetDistanceMap
}

func computeDistance(p1 dto.StatusPlanet, p2 dto.StatusPlanet) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
