package constants

type DataCollectionStages string

var (
	DataCollectionStageSingle   DataCollectionStages = "SINGLE"
	DataCollectionStageMultiple DataCollectionStages = "MULTIPLE"
)

type Direction string

var (
	North Direction = "NORTH"
	South Direction = "SOUTH"
	East  Direction = "EAST"
	West  Direction = "WEST"
)
