package product

import "github.com/core-go/search"

type ProductFilter struct {
	*search.Filter
	ProductId     string     `json:"productId,omitempty" gorm:"column:productId;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	ProductName   string            `json:"productName,omitempty" gorm:"column:productName" bson:"productName,omitempty" dynamodbav:"productName,omitempty" firestore:"productName,omitempty" validate:"required,max=50"`
	Title         string            `json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	Description   string            `json:"description,omitempty" gorm:"column:description" bson:"description" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	EffectiveDate *search.TimeRange `mapstructure:"effective_date" json:"effectiveDate,omitempty" gorm:"column:effectivedate" bson:"effectiveDate,omitempty" dynamodbav:"effectiveDate,omitempty" firestore:"effectiveDate,omitempty" avro:"effectiveDate" validate:"required"`
	PublishedAt   *search.TimeRange `json:"publishedAt,omitempty" gorm:"column:publishedat" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	Type          string            `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required"`
	Tags          []string          `json:"tags,omitempty" gorm:"column:tags" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	Status        string            `json:"status,omitempty" gorm:"column:status" bson:"status" dynamodbav:"status,omitempty" firestore:"status,omitempty"`
}
