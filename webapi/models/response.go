package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Response struct {
	// ID           string      `json:"id"`
	StatusCode   int         `json:"status_code,omitempty"`
	IsSuccessful bool        `json:"is_successful,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	OnTime       time.Time   `json:"on_time" default:"now"`
}

type Device struct {
	ID        string    `gorm:"size:21;not null;unique;index;" json:"id,omitempty"`
	Name      string    `gorm:"size:50;unique;" json:"name,omitempty"`
	OnPin     int64     `gorm:"size:5;unique;" json:"on_pin"`
	AlertOn   int64     `json:"alert_on"`
	IsActive  bool      `gorm:"type:boolean;not null" json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Device) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type LineToken struct {
	ID          string    `gorm:"size:21;not null;unique;index;" json:"id,omitempty"`
	Token       string    `gorm:"index;not null;unique;" json:"token"`
	Description string    `gorm:"size:255;not null" json:"description"`
	IsActive    bool      `gorm:"type:boolean;not null" json:"is_active"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *LineToken) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Notification struct {
	ID          string    `gorm:"size:21;not null;unique;index;" json:"id,omitempty"`
	DeviceID    string    `gorm:"size:21;not null;unique;index;" json:"device_id"`
	LineTokenID string    `gorm:"size:21;not null;unique;index;" json:"line_token_id"`
	IsAccept    bool      `gorm:"type:boolean;not null" json:"is_accept"`
	IsActive    bool      `gorm:"type:boolean;not null" json:"is_active"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
	Device      Device    `gorm:"ForeignKey:DeviceID;References:ID" json:"device"`
	LineToken   LineToken `gorm:"ForeignKey:LineTokenID;References:ID" json:"line_token"`
}

func (obj *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type TempData struct {
	ID          string    `gorm:"size:21;not null;unique;index;" json:"id,omitempty"`
	DeviceID    string    `gorm:"not null;size:21;" json:"device_id"`
	OnDateTime  time.Time `json:"on_date_time" default:"now"`
	Temp        float64   `gorm:"type:float;not null" json:"temp"`
	Humidity    float64   `gorm:"type:float;not null" json:"humidity"`
	Description string    `gorm:"size:255;not null" json:"description"`
	IsActive    bool      `gorm:"type:boolean;not null" json:"is_active"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
	Device      Device    `gorm:"ForeignKey:DeviceID;References:ID" json:"device"`
}

func (obj *TempData) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
