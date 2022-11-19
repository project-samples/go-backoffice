package test

import "time"

type Test struct {
	TestId        string     `mapstructure:"test_id" json:"testId,omitempty" gorm:"column:testid;primary_key" bson:"_id,omitempty" dynamodbav:"testId,omitempty" firestore:"testId,omitempty" avro:"testId" validate:"omitempty,max=40"`
	Title         *string    `mapstructure:"title" json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" avro:"title" validate:"required,max=255"`
	EffectiveDate *time.Time `mapstructure:"effective_date" json:"effectiveDate,omitempty" gorm:"column:effectivedate" bson:"effectiveDate,omitempty" dynamodbav:"effectiveDate,omitempty" firestore:"effectiveDate,omitempty" avro:"effectiveDate" validate:"required"`
	Questions     []string   `mapstructure:"questions" json:"questions,omitempty" gorm:"column:questions" bson:"questions,omitempty" dynamodbav:"questions,omitempty" firestore:"questions,omitempty" avro:"questions" validate:"required,max=500"`
	Tags          []string   `json:"tags,omitempty" gorm:"column:tags" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	CategoryId    *string    `mapstructure:"category_id" json:"categoryId,omitempty" gorm:"column:categoryid" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty" avro:"categoryId" validate:"required,max=45"`
	CreatedBy     *string    `json:"createdBy,omitempty" gorm:"column:createdBy" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"createdAt,omitempty"`
	UpdatedBy     *string    `json:"updatedBy,omitempty" gorm:"column:updatedBy" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"updatedAt,omitempty"`
}
