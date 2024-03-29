// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id             int    `json:"id"             ` //
	FullName       string `json:"fullName"       ` //
	Email          string `json:"email"          ` //
	HashedPassword string `json:"hashedPassword" ` //
	IsActive       bool   `json:"isActive"       ` //
	IsSuperuser    bool   `json:"isSuperuser"    ` //
}
