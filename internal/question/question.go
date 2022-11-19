package question

type Question struct {
    Id *string `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty" avro:"id" validate:"required,max=40"`
    Title *string `mapstructure:"title" json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" avro:"title" validate:"required,max=255"`
    Body *string `mapstructure:"body" json:"body,omitempty" gorm:"column:body" bson:"body,omitempty" dynamodbav:"body,omitempty" firestore:"body,omitempty" avro:"body" validate:"required,max=2000"`
    Options []string `mapstructure:"options" json:"options,omitempty" gorm:"column:options" bson:"options,omitempty" dynamodbav:"options,omitempty" firestore:"options,omitempty" avro:"options" validate:"required,max=45"`
    Answers []int32 `mapstructure:"answers" json:"answers,omitempty" gorm:"column:answers" bson:"answers,omitempty" dynamodbav:"answers,omitempty" firestore:"answers,omitempty" avro:"answers" validate:"required"`
    Tags []string  `json:"tags,omitempty" gorm:"column:tags" bson:"tags,omitempty" dynamodbav:"tags,omitempty" firestore:"tags,omitempty"`
    CategoryId *string `mapstructure:"category_id" json:"categoryId,omitempty" gorm:"column:categoryid" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty" avro:"categoryId" validate:"required,max=45"`
}
