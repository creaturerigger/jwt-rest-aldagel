package models

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/tokenization"
)

type User struct {
	gorm.Model
	Email    string  `gorm:"not null;unique" json:"email"`
	Password string  `gorm:"not null" json:"password"`
	FullName string  `gorm:"not null" json:"fullName"`
	ImageUrl string  `gorm:"not null" json:"imageUrl"`
	Bio      string  `gorm:"not null" json:"bio"`
	Friends  []*User `gorm:"many2many:user_friends" json:"friends"`
	Orders   []Order `gorm:"foreignkey:UserID" json:"orders"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email          string `json:"email"`
	StandardClaims *jwt.StandardClaims
}

var UserDB *gorm.DB

func AddFriend(user *User, friend *User) error {
	userTemp := &User{}
	friendTemp := &User{}
	db.Table("users").
		Where("email = ?", user.Email).
		First(&userTemp)
	if err := db.Table("users").
		Where("email = ?", friend.Email).
		First(&friendTemp).Error; err != nil {
		return &UserNotFoundError{friend.Email}
	}
	err := db.Model(&userTemp).
		Association("Friends").
		Append(friendTemp)
	db.Model(&friendTemp).
		Association("Friends").
		Append(userTemp)
	db.Save(&userTemp)
	db.Save(&friendTemp)
	return err
}

func GetFriends(user *User) (*[]User, error) {
	if err := db.Table("users").
		Where("email = ?", user.Email).
		First(&user).Error; err != nil {
		return nil, &UserNotFoundError{user.Email}
	}
	var friends *[]User
	err := db.Model(&user).
		Association("Friends").
		Find(&friends)
	return friends, err
}

func Register(user *User) (*User, error) {
	userEmail := db.Table("users").
		Where("email = ?", user.Email).
		First(&User{}).Error
	if userEmail == nil {
		return nil, &UserExistsError{user.Email}
	}
	return user, db.Create(&user).Error
}

func Login(creds *UserCredentials) (*Claims, error) {
	var user *User
	err := db.Table("users").
		Where("email = ?", creds.Email).
		First(&user).Error
	if err != nil {
		return nil, &UserNotFoundError{creds.Email}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return nil, &WrongPasswordError{creds.Email}
	}
	return &Claims{
		Email:          user.Email,
		StandardClaims: tokenization.GenerateClaims(user.Email, jwt.At(time.Now().Add(time.Hour*24*7))),
	}, nil
}

func GetUser(id string) (*User, error) {
	var user *User
	if err := db.Table("users").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, &UserNotFoundError{id}
	}
	return user, nil
}

func GetUserWithEmail(email string) (*User, error) {
	var user *User
	if err := db.Table("users").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, &UserNotFoundError{email}
	}
	return user, nil
}

func UserExistsWithEmail(email string) bool {
	var user *User
	err := db.Table("users").
		Where("email = ?", email).
		First(&user).Error
	return err == nil
}

func UserExistsWithID(id string) bool {
	var user *User
	err := db.Table("users").
		Where("id = ?", id).
		First(&user).Error
	return err == nil
}
