package article

import "time"

type Article struct {
	Id          string     `json:"id" gorm:"primary_key;column:id" bson:"_id" dynamodbav:"id,omitempty" firestore:"-"`
	Title       string     `json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty"`
	AuthorId    string     `json:"authorId,omitempty" gorm:"column:authorid" bson:"authorId,omitempty" dynamodbav:"authorId,omitempty" firestore:"authorId,omitempty"`
	Name        string     `json:"name,omitempty" gorm:"column:name" bson:"name,omitempty" dynamodbav:"name,omitempty" firestore:"name,omitempty"`
	Description string     `json:"description,omitempty" gorm:"column:description" bson:"description" dynamodbav:"description,omitempty" firestore:"description,omitempty"`
	PublishedAt *time.Time `json:"publishedAt,omitempty" gorm:"column:publishedat" bson:"publishedAt,omitempty" dynamodbav:"publishedAt,omitempty" firestore:"publishedAt,omitempty"`
	Type        string     `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty" validate:"required"`
	Body        string     `json:"body,omitempty" gorm:"column:body" bson:"body,omitempty" dynamodbav:"body,omitempty" firestore:"body,omitempty" validate:"required,max=2000"`
	Tags        []string   `json:"tags,omitempty" gorm:"column:tags" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
	Status      *string    `json:"status,omitempty" gorm:"column:status" bson:"status" dynamodbav:"status,omitempty" firestore:"status,omitempty"`
	CreatedBy   *string    `json:"createdBy,omitempty" gorm:"column:createdBy" bson:"createdBy,omitempty" dynamodbav:"createdBy,omitempty" firestore:"createdBy,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt" bson:"createdAt,omitempty" dynamodbav:"createdAt,omitempty" firestore:"createdAt,omitempty"`
	UpdatedBy   *string    `json:"updatedBy,omitempty" gorm:"column:updatedBy" bson:"updatedBy,omitempty" dynamodbav:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt" bson:"updatedAt,omitempty" dynamodbav:"updatedAt,omitempty" firestore:"updatedAt,omitempty"`
}
