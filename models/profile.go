package models

type Profile struct {
	Name     string  `json: "name" bson: "name"`
	Username string  `json: "username" bson: "username"`
	Bio      string  `json: "bio" bson: "bio"`
	Pronouns string  `json: "pronouns" bson: "pronouns"`
	Links    Socials `json: "links" bson: "links"`
}
