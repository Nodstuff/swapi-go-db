package empire

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Starship struct {
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
	HyperdriveRating     null.String `json:"hyperdrive_rating,omitempty"`
	MGLT                 null.String `json:"MGLT,omitempty"`
	StarshipClass        null.String `json:"starship_class,omitempty"`
	Pilots               []Person    `json:"pilots,omitempty"`
	Films                []Film      `json:"films,omitempty"`
	Created              null.String `json:"created,omitempty"`
	Edited               null.String `json:"edited,omitempty"`
}

func GetStarship(id int) Starship {
	var s Starship
	GetHttp(fmt.Sprintf("/starships/%d", id), &s)
	return s
}
