package empire

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Vehicle struct {
	Id                   int         `json:"id,omitempty"`
	Name                 null.String `json:"name,omitempty"`
	Model                null.String `json:"model,omitempty"`
	Manufacturer         null.String `json:"manufacturer,omitempty"`
	CostInCredits        null.String `json:"cost_in_credits,omitempty"`
	Length               null.String `json:"length,omitempty"`
	MaxAtmospheringSpeed null.String `json:"max_atmosphering_speed,omitempty"`
	Crew                 null.String `json:"crew,omitempty"`
	Passengers           null.String `json:"passengers,omitempty"`
	CargoCapacity        null.String `json:"cargo_capacity,omitempty"`
	Consumables          null.String `json:"consumables,omitempty"`
	VehicleClass         null.String `json:"vehicle_class,omitempty"`
	Pilots               []Person    `json:"pilots,omitempty"`
	Films                []Film      `json:"films,omitempty"`
	Created              null.String `json:"created,omitempty"`
	Edited               null.String `json:"edited,omitempty"`
}

func GetVehicle(id int) Vehicle {
	var v Vehicle
	GetHttp(fmt.Sprintf("/vehicles/%d", id), &v)
	return v
}
