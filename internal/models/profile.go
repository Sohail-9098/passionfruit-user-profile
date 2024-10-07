package models

type Profile struct {
	ID               string   `json:"id" bson:"_id,omitempty"`
	Name             string   `json:"name" bson:"name" validation:"required"`
	Bio              string   `json:"bio" bson:"bio"`
	Gender           string   `json:"gender" bson:"gender"`
	Age              int      `jsosn:"age" bson:"age"`
	ProfilePicUrl    string   `json:"profile_pic_url" bson:"profile_pic_url" validation:"required"`
	AdditionalPhotos []string `json:"photos" bson:"photos" validation:"url"`
}
