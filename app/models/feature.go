package models


import (
	"demo_gorm/configs/db"
	"time"
)

type FeatureConfig struct{
	Id uint8 `json:"id"`
	Feature_name string `json:"feature_name"`
	Enable bool `json:"enable"`
	UpdatedAt time.Time `json:"-"` 
	CreatedAt time.Time `json:"-"` 
}

type featureModel struct{}
var FeatureModel featureModel

func init() {
	FeatureModel = featureModel{}	
}

func(fm featureModel)All()[]FeatureConfig{
	var features []FeatureConfig
	db.MysqlConn.Find(&features)
	return features
}

func(fm featureModel)Create(args map[string]interface{})FeatureConfig{
	feature := FeatureConfig{	
								Feature_name: args["feature_name"].(string),
								Enable: args["enable"].(bool),
							}
	db.MysqlConn.Create(&feature)
	db.MysqlConn.Where(&FeatureConfig{Feature_name: feature.Feature_name}).First(&feature)
	return feature
}