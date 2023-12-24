package models

type Socials struct {
	Twitter   string `json: "twitter" bson: "twitter"`
	Instagram string `json: "instagram" bson: "instagram"`
	Snapchat  string `json: "snapchat" bson: "snapchat"`
}
