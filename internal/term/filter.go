package term

import "github.com/core-go/search"

type TermFilter struct {
	*search.Filter
	Id            string            `json:"id" gorm:"primary_key;column:id" bson:"_id" dynamodbav:"id,omitempty" firestore:"-" match:"equal"`
	Title         string            `json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Name          string            `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Description   string            `json:"description,omitempty" gorm:"column:description" bson:"description" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	EffectiveDate *search.TimeRange `mapstructure:"effective_date" json:"effectiveDate,omitempty" gorm:"column:effectivedate" bson:"effectiveDate,omitempty" dynamodbav:"effectiveDate,omitempty" firestore:"effectiveDate,omitempty" avro:"effectiveDate" validate:"required"`
	PublishedAt   *search.TimeRange `json:"publishedAt,omitempty" gorm:"column:publishedat" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	Type          string            `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required"`
	Tags          []string          `json:"tags,omitempty" gorm:"column:tags" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	Status        string            `json:"status,omitempty" gorm:"column:status" bson:"status" dynamodbav:"status,omitempty" firestore:"status,omitempty"`
}
