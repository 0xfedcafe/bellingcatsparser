package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// All Groups:
// 		{
//			1302000987, 	Civilian Casualty
//			1204076368, 	Russian Firing Positions
//			744433847, 		Bombing Shelling Explosion
//			2613895895, 	Gunfire, Fighting, Battle
//			2159062771,		Civilian Infrastructure Damage April 2022
//			182598014,		Civilian Infrastructure Damage March 2022
//			3730511149,		Civilian Infrastructure Damage February 2022
//			1946403459,		Munitions
//			878638621,		Other
//			2758805904,		Russian Military Losses
//			3442810439,		Ukrainian Military Losses
//			3640522656,		Military Infrastructure Damage
//			4263226937,		Russian Allies Movements March 2022
//			1422527994,		Russian Allies Movements February 2022
//			2327490580		Russian Allies Movements January 2022
//		}
// Doesn't affect category

// All marker colors
// {
//		"#cc1b15", Civilian Casualty, Civilian Infrastructure Damage, Ukrainian Military Losses
//		"#444444", (Munitions?)
//		"#f18729", (Bombing, Shelling, Explosions), (Gunfire, Fighting, Battle)
//		"#ffcc00", Others
//		"#005e38" Russian Allies Movements/Losses
//	}

func Parse(fName string) (MapObject, error) {
	j, err := os.Open(fName)

	if err != nil {
		fmt.Println(err)
		return MapObject{}, err
	}

	b, err := ioutil.ReadAll(j)

	if err != nil {
		fmt.Println(err)
		return MapObject{}, err
	}
	var entities Entities

	err = json.Unmarshal(b, &entities)
	if err != nil {
		fmt.Println(err)
		return MapObject{}, err
	}

	//fmt.Println(entities[0])
	var res MapObject
	for _, k := range entities {

		var nEntity NormalisedEntity
		nEntity.Date = k.Properties.Title[0:10]
		nEntity.Title = k.Properties.Title[11:]
		nEntity.Latitude = k.Geometry.Coordinates[0]
		nEntity.Longitude = k.Geometry.Coordinates[1]
		nEntity.MediaURL = k.Properties.MediaURL

		if char := k.Properties.Description[33]; '0' <= char && char <= '9' {
			nEntity.Level = char - '0'
		} else {
			nEntity.Level = 6
		}
		switch group := k.Properties.Group; group {
		case 1302000987:
			res.Casualties = append(res.Casualties, nEntity)
		case 744433847:
			res.Shellings = append(res.Shellings, nEntity)
		case 2159062771, 182598014, 3730511149:
			res.InfraDamage = append(res.InfraDamage, nEntity)
		}
	}

	err = j.Close()
	if err != nil {
		fmt.Println(err)
		return MapObject{}, err
	}
	return res, nil
}
