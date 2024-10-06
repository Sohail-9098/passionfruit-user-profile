package models

type Profile struct {
	ID          string      `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name" validation:"required"`
	Bio         string      `json:"bio" bson:"bio"`
	PhotosUrl   []string    `json:"photos" bson:"photos" validation:"url"`
	Preferences Preferences `json:"preferences" bson:"preferences"`
}

type Preferences struct {
	AgeRange  [2]int   `json:"age_range" bson:"age_range"`
	Location  string   `json:"location" bson:"location"`
	Interests []string `json:"interests" bson:"interests"`
}
