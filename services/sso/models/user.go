package models

import (
	//"github.com/google/uuid"
	"time"
)

type User struct {
	Id         int    `json:"USR_ID,omitempty"`
	Name       string `json:"USR_NAME,omitempty"`
	Password   int    `json:"USR_PASS,omitempty"`
	FlagMaster bool   `json:"USR_FLAGMASTER,omitempty"`
	Prefix     string `json:"USR_PREF,omitempty"`
	//Description string    `json:"USR_DESCR,omitempty"`
	//Type        int       `json:"USR_TYPE,omitempty"`
	//SID         string    `json:"USR_SID,omitempty"`
	//Type1C      bool      `json:"USR_1CTYPE,omitempty"`
	Active    bool      `json:"USR_ACTIVE,omitempty"`
	CreatedAt time.Time `json:"USR_CREATEDT,omitempty"`
	UpdatedAt time.Time `json:"USR_UPDDT,omitempty"`
}
