package models

type Profile struct {
	ID               string   `json:"id" bson:"_id,omitempty"`
	Name             string   `json:"name" bson:"name" binding:"required"`
	Bio              string   `json:"bio" bson:"bio" binding:"required"`
	Gender           string   `json:"gender" bson:"gender" binding:"required"`
	Age              int      `json:"age" bson:"age" binding:"required,gte=18"`
	ProfilePicUrl    string   `json:"profile_pic_url" bson:"profile_pic_url" binding:"required"`
	AdditionalPhotos []string `json:"additional_photos" bson:"additional_photos" binding:"dive,url"`
}
