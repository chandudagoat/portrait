package models

type Profile struct {
	ID       string  `json: "id" bson: "_id"`
	Name     string  `json: "name" bson: "name"`
	Bio      string  `json: "bio" bson: "bio"`
	Pronouns string  `json: "pronouns" bson: "pronouns"`
	Links    Socials `json: "links" bson: "links"`
}
