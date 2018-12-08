package models

import (
	"encoding/json"
	"time"
    "errors"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)
/**
 *  Status ID's and their meanings
 *  1 - Request Accepted
 *  2 - Pending response
 *  3 - Request Withdrawn/Canceled *(by sender)
 *  4 - Request Denied *(by receiver)
 */
type Request struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Status      int       `json:"status" db:"status"`
	SenderID    uuid.UUID `json:"senderid" db:"senderid"`
	ReceiverID  uuid.UUID `json:"receiverid" db:"receiverid"`
	Topic       string    `json:"topic" db:"topic"`
}

// String is not required by pop and may be deleted
func (r Request) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Requests is not required by pop and may be deleted
type Requests []Request

// String is not required by pop and may be deleted
func (r Requests) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Request) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: r.Status, Name: "Status"},
        &validators.StringIsPresent{Field: r.Topic, Name: "Topic"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Request) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Request) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (req *Request) CreateRequest(tx *pop.Connection) (*validate.Errors, error){
    return tx.ValidateAndCreate(req)
}

func CreateNewRequestData(sender *User, receiver *User, topic string) (*Request, error){
    if (sender.ID == receiver.ID){
        return nil, errors.New("Cannot send a request to yourself")
    }
    if (sender.IsTutor == receiver.IsTutor){
        return nil, errors.New("Cannot Send a request to a user of the same type")
    }
    var request *Request

    request = &Request{
        Status:     3,
        SenderID:   sender.ID,
        ReceiverID: receiver.ID,
        Topic:      topic,
    }

    return request, nil
}


func GetRequestsSent(user *User, tx *pop.Connection) (Requests, error){
	query := tx.RawQuery("SELECT * FROM requests WHERE senderid = ?", user.GoogleID)

    var requests Requests
	if err := query.All(requests); err != nil {
		return nil, err
	} else {
		return requests, nil
	}

}


func GetRequestsReceived(user *User, tx *pop.Connection) (Requests, error){
	query := tx.RawQuery("SELECT * FROM requests WHERE receiverid = ?", user.GoogleID)

    var requests Requests
	if err := query.All(requests); err != nil {
		return nil, err
	} else {
		return requests, nil
	}

}


