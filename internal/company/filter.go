package company

import . "github.com/core-go/search"

type CompanyFilter struct {
	*Filter
	CompanyId   string   `json:"companyId,omitempty" gorm:"column:companyid;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	CompanyName string   `json:"companyName,omitempty" gorm:"column:companyname" bson:"companyName,omitempty" dynamodbav:"companyName,omitempty" firestore:"companyName,omitempty" validate:"required,max=80" q:"true"`
	Email       string   `json:"email,omitempty" gorm:"column:email" bson:"email,omitempty" dynamodbav:"email,omitempty" firestore:"email,omitempty" validate:"email,max=100" q:"prefix"`
	DisplayName string   `json:"displayName,omitempty" gorm:"column:displayname" bson:"displayName,omitempty" dynamodbav:"displayName,omitempty" firestore:"displayName,omitempty" validate:"max=100" q:"true"`
	Status      []string `json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" match:"equal" validate:"required,max=1,code"`
}
