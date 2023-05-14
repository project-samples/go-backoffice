package ticket

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Ticket struct {
	Id          string       `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"_id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty" avro:"id" validate:"omitempty,max=40"`
	Title       *string      `mapstructure:"title" json:"title,omitempty" gorm:"column:title" bson:"title,omitempty" dynamodbav:"title,omitempty" firestore:"title,omitempty" avro:"title" validate:"required,max=255"`
	Body        *string      `mapstructure:"body" json:"body,omitempty" gorm:"column:body" bson:"body,omitempty" dynamodbav:"body,omitempty" firestore:"body,omitempty" avro:"body" validate:"required,max=2000"`
	CategoryId  *string      `mapstructure:"category_id" json:"categoryId,omitempty" gorm:"column:categoryid" bson:"categoryId,omitempty" dynamodbav:"categoryId,omitempty" firestore:"categoryId,omitempty" avro:"categoryId" validate:"required,max=45"`
	Requestor   *string      `mapstructure:"requestor" json:"requestor,omitempty" gorm:"column:requestor" bson:"requestor,omitempty" dynamodbav:"requestor,omitempty" firestore:"requestor,omitempty" avro:"requestor" validate:"required,max=40"`
	RequestedAt *time.Time   `mapstructure:"requested_at" json:"requestedAt,omitempty" gorm:"column:requested_at" bson:"requestedAt,omitempty" dynamodbav:"requestedAt,omitempty" firestore:"requestedAt,omitempty" avro:"requestedAt" validate:"required"`
	Approver    *string      `mapstructure:"approver" json:"approver,omitempty" gorm:"column:approver" bson:"approver,omitempty" dynamodbav:"approver,omitempty" firestore:"approver,omitempty" avro:"approver" validate:"required,max=40"`
	ApprovedAt  *time.Time   `mapstructure:"approved_at" json:"approvedAt,omitempty" gorm:"column:approved_at" bson:"ApprovedAt,omitempty" dynamodbav:"approvedAt,omitempty" firestore:"approvedAt,omitempty" avro:"approvedAt" validate:"required"`
	Assignee    *string      `mapstructure:"assignee" json:"assignee,omitempty" gorm:"column:assignee" bson:"assignee,omitempty" dynamodbav:"assignee,omitempty" firestore:"assignee,omitempty" avro:"assignee" validate:"required,max=40"`
	CompletedAt *time.Time   `mapstructure:"completed_at" json:"completedAt,omitempty" gorm:"column:completed_at" bson:"CompletedAt,omitempty" dynamodbav:"completedAt,omitempty" firestore:"completedAt,omitempty" avro:"completedAt" validate:"required"`
	Status      *string      `mapstructure:"status" json:"status,omitempty" gorm:"column:status" bson:"status,omitempty" dynamodbav:"status,omitempty" firestore:"status,omitempty" avro:"status" validate:"required,max=1"`
	Attachments []Attachment `mapstructure:"attachments" json:"attachments,omitempty" gorm:"column:attachments" bson:"attachments,omitempty" dynamodbav:"attachments,omitempty" firestore:"attachments,omitempty" avro:"attachments"`
}

type Attachment struct {
	OriginalFileName string `json:"originalFileName,omitempty" gorm:"column:original_filename" bson:"original_filename,omitempty" dynamodbav:"original_filename,omitempty" firestore:"original_filename,omitempty"`
	FileName         string `json:"fileName,omitempty" gorm:"column:fileName" bson:"fileName,omitempty" dynamodbav:"fileName,omitempty" firestore:"fileName,omitempty"`
	Url              string `json:"url,omitempty" gorm:"column:url" bson:"url,omitempty" dynamodbav:"url,omitempty" firestore:"url,omitempty" avro:"url" validate:"required"`
	Type             string `json:"type,omitempty" gorm:"column:type" bson:"type,omitempty" dynamodbav:"type,omitempty" firestore:"type,omitempty"`
}

func (a Attachment) Value() (driver.Value, error) {
	return json.Marshal(a)
}
func (a *Attachment) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
