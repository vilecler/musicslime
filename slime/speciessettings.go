package slime

import (
	types "gitlab.utc.fr/vilecler/musicslime/types"
)

type SpeciesSettings struct {
	MoveSpeed float64
	TurnSpeed float64

	SensorAngleDegrees   float64
	SensorOffsetDistance float64
	SensorSize           int
	Colour               types.Colour
}

func NewSpeciesSettings(moveSpeed float64, turnSpeed float64, sensorAngleDegrees float64, sensorOffsetDistance float64, sensorSize int, colour types.Colour) *SpeciesSettings {
	return &SpeciesSettings{moveSpeed, turnSpeed, sensorAngleDegrees, sensorOffsetDistance, sensorSize, colour}
}

func GetSpeciesSettingsByID(id int) *SpeciesSettings {
	switch id {
	case 1:
		var colour types.Colour
		colour[0] = 0.10588236
		colour[1] = 0.6240185
		colour[2] = 0.83137256
		colour[3] = 1
		return NewSpeciesSettings(50, -90, 112, 30, 1, colour)
	case 2:
		var colour types.Colour
		colour[0] = 0.47160125
		colour[1] = 0.8301887
		colour[2] = 0.105731584
		colour[3] = 1
		return NewSpeciesSettings(50, 50, 70, 30, 1, colour)
	}

	var colour types.Colour
	colour[0] = 0.53608304
	colour[1] = 0.062745094
	colour[2] = 1
	colour[3] = 0
	return NewSpeciesSettings(70, 30, 60, 20, 1, colour)
}
