package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	ContextKey       string
	LoginCredentials struct {
		Email    string `json:"email" example:"test@mail.com"`
		Password string `json:"password" example:"abc123"`
	}
)

const (
	PhotographerNotVerifiedStatus = "NOT_VERIFIED"
	PhotographerPendingStatus     = "PENDING"
	PhotographerVerifiedStatus    = "VERIFIED"
	PhotographerRejectedStatus    = "REJECTED"
)

const (
	MALE   = "MALE"
	FEMALE = "FEMALE"
	OTHER  = "OTHER"
)

type User struct {
	bun.BaseModel      `bun:"table:users,alias:u"`
	Id                 uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Username           string    `bun:"username,type:varchar" json:"username"`
	Email              string    `bun:"email,type:varchar" json:"email"`
	Provider           *string   `bun:"provider,type:varchar" json:"provider"`
	Password           *string   `bun:"password,type:varchar" json:"-"`
	LoggedOut          bool      `bun:"logged_out,type:boolean" json:"logged_out"`
	ProfilePictureKey  *string   `bun:"profile_picture_key,type:varchar" json:"profile_picture_key"`
	Firstname          string    `bun:"firstname,type:varchar" json:"firstname"`
	Lastname           string    `bun:"lastname,type:varchar" json:"lastname"`
	VerificationStatus string    `bun:"verification_status,type:varchar" json:"verification_status"`
	IsAdmin            bool      `bun:"is_admin,type:boolean" json:"is_admin"`
	About              *string   `bun:"about,type:varchar" json:"about"`
	Address            *string   `bun:"address,type:varchar" json:"address"`
	PhoneNumber        *string   `bun:"phone_number,type:varchar" json:"phone_number"`
	Gender             *string   `bun:"gender,type:varchar" json:"gender"`
}

type UserInput struct {
	Email     string  `json:"email" example:"test@mail.com"`
	Password  *string `json:"password" example:"root"`
	Firstname string  `json:"firstname" example:"test"`
	Lastname  string  `json:"lastname" example:"test"`
}

type UserUpdateInput struct {
	Email       *string `json:"email" example:"test@mail.com"`
	Password    *string `json:"password" example:"root"`
	PhoneNumber *string `json:"phone_number" example:"096198923"`
	Firstname   *string `json:"firstname" example:"test"`
	Lastname    *string `json:"lastname" example:"test"`
	Gender      *string `json:"gender" example:"Male"`
	About       *string `json:"about" example:"Hello"`
	Username    *string `json:"username" example:"test"`
	Address     *string `json:"address" example:"Bangkok"`
}

type VerificationTicket struct {
	bun.BaseModel         `bun:"table:verification_ticket,alias:vrf_ticket"`
	Id                    uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	UserId                uuid.UUID `bun:"user_id,type:uuid" json:"-"`
	User                  User      `bun:"-" json:"user"`
	IdCardNumber          string    `bun:"id_card_number,type:varchar" json:"id_card_number"`
	IdCardPictureKey      string    `bun:"id_card_picture_key,type:varchar" json:"-"`
	IdCardPictureURL      string    `bun:"-" json:"id_card_picture_url"`
	AdditionalDescription *string   `bun:"additional_desc,type:varchar" json:"additional_desc"`
	CreatedAt             time.Time `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	DueDate               time.Time `bun:"due_date,type:timestamptz,default:now()" json:"due_date"`
}

type VerificationTicketInput struct {
	IdCardNumber          string                `form:"id_card_number" binding:"required"`
	IdCardPicture         *multipart.FileHeader `form:"id_card_picture" binding:"required"`
	AdditionalDescription *string               `form:"addition_desc"`
}

type Gallery struct {
	bun.BaseModel  `bun:"table:galleries,alias:galleries"`
	Id             uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	PhotographerId uuid.UUID `bun:"photographer_id,type:uuid" json:"photographer_id"`
	Location       string    `bun:"location,type:varchar" json:"location"`
	Name           string    `bun:"name,type:varchar" json:"name"`
	AvgRating      *float32  `bun:"avg_rating,type:real" json:"avg_rating"`
	Price          int       `bun:"price,type:integer" json:"price"`
	Hours          int       `bun:"hours,type:integer" json:"hours"`
	Description    *string   `bun:"description,type:varchar" json:"description"`
	DeliveryTime   int       `bun:"delivery_time,type:integer" json:"delivery_time"`
	Included       []string  `bun:",array" json:"included"`
}

type GalleryInput struct {
	Name         *string  `bun:"name,type:varchar" json:"name"`
	Location     *string  `bun:"name,type:varchar" json:"location"`
	Price        *int     `bun:"price,type:integer" json:"price"`
	Hours        *int     `bun:"hours,type:integer" json:"hours"`
	Description  *string  `bun:"description,type:varchar" json:"description"`
	DeliveryTime *int     `bun:"delivery_time,type:integer" json:"delivery_time"`
	Included     []string `bun:",array" json:"included"`
}

const (
	BookingDraftStatus                 = "DRAFT"
	BookingPaidStatus                  = "USER_PAID"
	BookingCancelledStatus             = "CANCELLED"
	BookingCustomerReqCancelStatus     = "C_REQ_CANCEL"
	BookingPhotographerReqCancelStatus = "P_REQ_CANCEL"
	BookingCompletedStatus             = "COMPLETED"
	BookingPaidOutStatus               = "PAID_OUT"
	BookingRefundReqStatus             = "REQ_REFUND"
)

type BookingProposal struct {
	CustomerId      uuid.UUID `json:"customer_id"`
	RoomId          uuid.UUID `bun:"room_id,type:uuid" json:"room_id"`
	NegotiatedPrice *int      `json:"negotiated_price"`
	StartTime       time.Time `bun:"start_time,type:timestamptz" json:"start_time"`
	EndTime         time.Time `bun:"end_time,type:timestamptz" json:"end_time"`
}

type Booking struct {
	bun.BaseModel `bun:"table:bookings,alias:bookings"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CustomerId    uuid.UUID `bun:"customer_id,type:uuid" json:"customer_id"`
	RoomId        uuid.UUID `bun:"room_id,type:uuid" json:"-"`
	Room          Room      `bun:"-" json:"room"`
	ResultedPrice int       `bun:"resulted_price,type:integer" json:"resulted_price"`
	StartTime     time.Time `bun:"start_time,type:timestamptz" json:"start_time"`
	EndTime       time.Time `bun:"end_time,type:timestamptz" json:"end_time"`
	Status        string    `bun:"status,type:varchar" json:"status"`
	CreatedAt     time.Time `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
}

type SearchFilter struct {
	PhotographerId                  *string `binding:"omitempty,uuid" form:"photographer_id"`
	MatchedConditionPhotographerIds []uuid.UUID
	GalleryName                     *string `form:"gallery_name"`
	PhotographerName                *string `form:"photographer_name"`
	Location                        *string `form:"location"`
	MinPrice                        *int    `form:"min_price"`
	MaxPrice                        *int    `form:"max_price"`
}

type Room struct {
	bun.BaseModel `bun:"table:rooms,alias:rooms"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	GalleryId     uuid.UUID  `bun:"gallery_id,type:uuid" json:"-"`
	Gallery       Gallery    `bun:"-" json:"gallery"`
	OtherUsers    []*User    `bun:"-" json:"other_users,omitempty"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type UserRoomLookup struct {
	bun.BaseModel `bun:"table:user_room_lookup,alias:urlookup"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	UserId        uuid.UUID  `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID  `bun:"room_id,type:uuid" json:"room_id"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type Conversation struct {
	bun.BaseModel `bun:"table:conversations,alias:convs"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Text          string     `bun:"text,type:varchar" json:"text"`
	UserId        uuid.UUID  `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID  `bun:"room_id,type:uuid" json:"room_id"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type RoomMemberInput struct {
	MemberIds []uuid.UUID `binding:"required" json:"member_ids"`
	GalleryId uuid.UUID   `json:"gallery_id"`
}

type Photo struct {
	bun.BaseModel `bun:"table:photos,alias:photos"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	GalleryId     uuid.UUID `bun:"gallery_id,type:uuid" json:"gallery_id"`
	PhotoKey      string    `bun:"photo_key,type:varchar" json:"photo_key"`
}

type ReviewInput struct {
	BookingId  *uuid.UUID `json:"booking_id"`
	Rating     *int       `json:"rating"`
	ReviewText *string    `json:"review_text"`
}

type Review struct {
	bun.BaseModel `bun:"table:reviews,alias:reviews"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CustomerId    uuid.UUID `bun:"customer_id,type:uuid" json:"-"`
	Customer      User      `bun:"-" json:"customer"`
	BookingId     uuid.UUID `bun:"booking_id,type:uuid" json:"-"`
	Booking       Booking   `bun:"-" json:"booking"`
	Rating        int       `bun:"rating,type:integer" json:"rating"`
	ReviewText    *string   `bun:"review_text,type:varchar" json:"review_text"`
}

const (
	IssueOpenStatus   = "OPEN"
	IssueClosedStatus = "CLOSED"
)

const (
	IssueRefundSubject    = "REFUND"
	IssueTechnicalSubject = "TECHNICAL"
)

type Issue struct {
	bun.BaseModel `bun:"table:issues,alias:issues"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ReporterId    uuid.UUID  `bun:"reporter_id,type:uuid" json:"-"`
	Reporter      User       `bun:"-" json:"reporter"`
	BookingId     *uuid.UUID `bun:"booking_id,type:uuid" json:"-"`
	Booking       *Booking   `bun:"-" json:"booking"`
	Status        string     `bun:"status,type:varchar" json:"status"`
	Subject       string     `bun:"subject,type:varchar" json:"subject"`
	DueDate       time.Time  `bun:"due_date,type:timestamptz,default:now()" json:"due_date"`
	Description   string     `bun:"description,type:varchar" json:"description"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
}

type IssueInput struct {
	Description *string `json:"description"`
}

type IssueFilter struct {
	ReporterId *string    `binding:"omitempty,uuid" form:"reporter_id"`
	Status     *string    `form:"status"`
	DueDate    *time.Time `form:"due_date" time_format:"2006-01-02" time_utc:"7"`
	CreatedAt  *time.Time `form:"created_at" time_format:"2006-01-02" time_utc:"7"`
	Subject    *string    `form:"subject"`
}

type IssueHeaderMetadata struct {
	PendingTickets  int `json:"pending_tickets"`
	TicketsToday    int `json:"tickets_today"`
	TicketsDueToday int `json:"tickets_due_today"`
	ClosedTickets   int `json:"closed_tickets"`
}
