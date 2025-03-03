package repository

import (
	"backend/protos/gen/go/production_task"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var mockProducts = []*production_task.Product{
	{
		Id:                1,
		ManufacturingDate: timestamp.New(time.Date(2023, 9, 15, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product A",
		NumberFrame:       123,
		PartName:          "Part A",
		WeightSpKg:        10.5,
		WeightGpKg:        11.0,
	},
	{
		Id:                2,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product B",
		NumberFrame:       456,
		PartName:          "Part B",
		WeightSpKg:        20.5,
		WeightGpKg:        21.0,
	},
	{
		Id:                3,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product C",
		NumberFrame:       789,
		PartName:          "Part C",
		WeightSpKg:        30.5,
		WeightGpKg:        31.0,
	},
	{
		Id:                4,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product D",
		NumberFrame:       101,
		PartName:          "Part D",
		WeightSpKg:        40.5,
		WeightGpKg:        41.0,
	},
	{
		Id:                5,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product E",
		NumberFrame:       202,
		PartName:          "Part E",
		WeightSpKg:        50.5,
		WeightGpKg:        51.0,
	},
	// Добавляем еще 15 записей...
	{
		Id:                6,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product F",
		NumberFrame:       303,
		PartName:          "Part F",
		WeightSpKg:        60.5,
		WeightGpKg:        61.0,
	},
	{
		Id:                7,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product G",
		NumberFrame:       404,
		PartName:          "Part G",
		WeightSpKg:        70.5,
		WeightGpKg:        71.0,
	},
	{
		Id:                8,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 7, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product H",
		NumberFrame:       505,
		PartName:          "Part H",
		WeightSpKg:        80.5,
		WeightGpKg:        81.0,
	},
	{
		Id:                9,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product I",
		NumberFrame:       606,
		PartName:          "Part I",
		WeightSpKg:        90.5,
		WeightGpKg:        91.0,
	},
	{
		Id:                10,
		ManufacturingDate: timestamp.New(time.Date(2023, 10, 9, 0, 0, 0, 0, time.UTC)),
		Nomenclature:      "Product J",
		NumberFrame:       707,
		PartName:          "Part J",
		WeightSpKg:        100.5,
		WeightGpKg:        101.0,
	},
}
