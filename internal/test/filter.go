package test

import . "github.com/core-go/search"

type TestFilter struct {
	*Filter
    TestId *string `mapstructure:"test_id" json:"testId,omitempty" gorm:"column:testid;primary_key" bson:"_id,omitempty" dynamodbav:"testId,omitempty" firestore:"testId,omitempty" avro:"testId" validate:"required,max=40"`
    Title *string `mapstructure:"title" json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" avro:"title" validate:"required,max=255"`
    EffectiveDate *TimeRange `mapstructure:"effective_date" json:"effectiveDate,omitempty" gorm:"column:effectivedate" bson:"effectiveDate,omitempty" dynamodbav:"effectiveDate,omitempty" firestore:"effectiveDate,omitempty" avro:"effectiveDate" validate:"required"`
    Questions *string `mapstructure:"questions" json:"questions,omitempty" gorm:"column:questions" bson:"questions,omitempty" dynamodbav:"questions,omitempty" firestore:"questions,omitempty" avro:"questions" validate:"required,max=500"`
    CategoryId *string `mapstructure:"category_id" json:"categoryId,omitempty" gorm:"column:categoryid" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty" avro:"categoryId" validate:"required,max=45"`
}
