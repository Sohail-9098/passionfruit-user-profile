package models

type Profile struct {
	ID          string      `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	Bio         string      `json:"bio" bson:"bio"`
	Photos      string      `json:"photos" bson:"photos"`
	Preferences Preferences `json:"preferences" bson:"preferences"`
}

type Preferences struct {
	AgeRange  [2]int   `json:"age_range" bson:"age_range"`
	Location  string   `json:"location" bson:"location"`
	Interests []string `json:"interests" bson:"interests"`
}
