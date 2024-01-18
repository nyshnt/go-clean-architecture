package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the equivalent Go struct for the MongoDB UsersSchema
type User struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	FirstName          string             `bson:"firstName,omitempty"`
	LastName           string             `bson:"lastName,omitempty"`
	Username           string             `bson:"username,omitempty"`
	Email              string             `bson:"email,omitempty"`
	EmailVerified      bool               `bson:"emailVerified,omitempty"`
	Salt               string             `bson:"salt,omitempty"`
	Hash               string             `bson:"hash,omitempty"`
	Country            string             `bson:"country,omitempty"`
	DOB                time.Time          `bson:"dob,omitempty"`
	Gender             string             `bson:"gender,omitempty"`
	ProfileImgUrl      string             `bson:"profileImgUrl,omitempty"`
	CreatedDate        time.Time          `bson:"createdDate,omitempty"`
	Role               primitive.ObjectID `bson:"role,omitempty" json:"-"` // Exclude from JSON
	DeviceID           string             `bson:"deviceId,omitempty"`
	ReferralCode       string             `bson:"referralCode,omitempty"`
	ReferrerUserID     primitive.ObjectID `bson:"referrerUserId,omitempty"`
	ReferrerCeroPoints int                `bson:"referrerCeroPoints,omitempty"`
	SocialID           string             `bson:"socialId,omitempty"`
	SocialType         string             `bson:"socialType,omitempty"`
	Lang               string             `bson:"lang,omitempty"`
	TimeDiff           int                `bson:"timeDiff,omitempty"`
	CarbonEmission     float64            `bson:"carbonEmission,omitempty"`
	OffsetTarget       float64            `bson:"offsetTarget,omitempty"`
	SignupType         string             `bson:"signupType,omitempty"`
	LastActiveDtm      time.Time          `bson:"lastActiveDtm,omitempty"`
	ProfileCreated     bool               `bson:"profileCreated,omitempty"`
	Deleted            bool               `bson:"deleted,omitempty"`
}
