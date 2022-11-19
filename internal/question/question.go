package question

import "time"

type Question struct {
	Id         string     `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty" avro:"id" validate:"omitempty,max=40"`
	Title      *string    `mapstructure:"title" json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" avro:"title" validate:"required,max=255"`
	Body       *string    `mapstructure:"body" json:"body,omitempty" gorm:"column:body" bson:"body,omitempty" dynamodbav:"body,omitempty" firestore:"body,omitempty" avro:"body" validate:"required,max=2000"`
	Mixed      bool       `mapstructure:"mixed" json:"mixed,omitempty" gorm:"column:mixed" bson:"mixed,omitempty" dynamodbav:"mixed,omitempty" firestore:"mixed,omitempty" avro:"mixed"`
	Options    []string   `mapstructure:"options" json:"options,omitempty" gorm:"column:options" bson:"options,omitempty" dynamodbav:"options,omitempty" firestore:"options,omitempty" avro:"options" validate:"required,max=45"`
	Answers    []int32    `mapstructure:"answers" json:"answers,omitempty" gorm:"column:answers" bson:"answers,omitempty" dynamodbav:"answers,omitempty" firestore:"answers,omitempty" avro:"answers" validate:"required"`
	Tags       []string   `json:"tags,omitempty" gorm:"column:tags" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	CategoryId *string    `mapstructure:"category_id" json:"categoryId,omitempty" gorm:"column:categoryid" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty" avro:"categoryId" validate:"required,max=45"`
	CreatedBy  *string    `json:"createdBy,omitempty" gorm:"column:createdBy" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"createdAt,omitempty"`
	UpdatedBy  *string    `json:"updatedBy,omitempty" gorm:"column:updatedBy" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"updatedAt,omitempty"`
}

type Answer struct {
	Body    string `mapstructure:"body" json:"body,omitempty" gorm:"column:body" bson:"body,omitempty" dynamodbav:"body,omitempty" firestore:"body,omitempty" avro:"body" validate:"required,max=2000"`
	Correct *bool  `mapstructure:"correct" json:"correct,omitempty" gorm:"column:correct" bson:"correct,omitempty" dynamodbav:"correct,omitempty" firestore:"correct,omitempty" avro:"correct"`
}
