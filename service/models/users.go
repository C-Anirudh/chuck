package models

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/C-Anirudh/chuck/service/hash"
	"github.com/C-Anirudh/chuck/service/rand"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	// imported for the effects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is a custom error we return when a resource we are looking for is not present in the db
	ErrNotFound = errors.New("models: resource not found")

	// ErrIDInvalid is a custom error we return when the id of user we want to delete is invalid
	ErrIDInvalid = errors.New("models: ID provided was invalid")

	// ErrPasswordIncorrect is a custom error we return when the user enters an invalid password in login page
	ErrPasswordIncorrect = errors.New("models: incorrect password provided")

	// ErrEmailRequired is a custom error we return when the email address is not provided when creating a user
	ErrEmailRequired = errors.New("models: email address is required")

	// ErrEmailInvalid is a custom error we return when the email address provided doesn't match requirements
	ErrEmailInvalid = errors.New("models: email address is not valid")

	// ErrEmailTaken is a custom error we return when create or update is called with an email address that is already in use
	ErrEmailTaken = errors.New("models: email address is already taken")

	// ErrPasswordTooShort is a custom error we return when the password set at account creation is too short
	ErrPasswordTooShort = errors.New("models: password must be atleast 8 characters long")

	// ErrPasswordRequired is a custom error we return when user tries to create an account without setting a password
	ErrPasswordRequired = errors.New("models: password is required")

	// ErrNameRequired is a custom error we return when user tries to create an account without setting a name
	ErrNameRequired = errors.New("models: Name is required")

	// ErrRememberRequired is a custom error we return when create or update is attempted without a user remember token hash
	ErrRememberRequired = errors.New("models: remember token is required")

	// ErrRememberTooShort is a custom error we return when the remember token is not 32 bytes long
	ErrRememberTooShort = errors.New("models: remember token must be atleast 32 bytes")
)

const (
	hmacSecretKey = "secret-hmac-key"
	userPwPepper  = "secret-random-string"
)

// User is the database model for our customer
type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

// UserDB is used to interact with the users database
type UserDB interface {
	// Methods for querying for single users
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	// Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	// Used to close a DB connection
	Close() error

	// Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

// UserService is a set of methods used to manipulate and work with the user model
type UserService interface {
	Authenticate(email, password string) (*User, error)
	UserDB
}

type userService struct {
	UserDB
}

type userGorm struct {
	db *gorm.DB
}

type userValidator struct {
	UserDB
	hmac       hash.HMAC
	emailRegex *regexp.Regexp
}

type userValFn func(*User) error

func newUserValidator(udb UserDB, hmac hash.HMAC) *userValidator {
	return &userValidator{
		UserDB:     udb,
		hmac:       hmac,
		emailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
	}
}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &userGorm{
		db: db,
	}, nil
}

// NewUserService is an abstraction layer providing us a connection with the db
func NewUserService(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(hmacSecretKey)
	uv := newUserValidator(ug, hmac)
	return &userService{
		UserDB: uv,
		/*
			uv implements all the methods in the UserDB interface, so we are able to assign it to the UserDB type
			in userService struct.
		*/
	}, nil
	/*
		We are able to return a userService struct type for UserService interface type as the userService struct type
		implements all the methods in the UserService interface.
	*/
}

// Authenticate is used to vet users
func (us *userService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrPasswordIncorrect
	default:
		return nil, err
	}
}

/*
	********************************
	********************************
	Start of functions related to db
	********************************
	********************************
*/

// ByID is used to search a user by ID from the db
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ByEmail is used to search a user by email from the db
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ByRemember is used to search a user by remember token from the db
func (ug *userGorm) ByRemember(rememberHash string) (*User, error) {
	var user User
	err := first(ug.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DestructiveReset drops the user table and rebuilds it
func (ug *userGorm) DestructiveReset() error {
	if err := ug.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return ug.AutoMigrate()
}

// Close is a function that is used to close the connection with the db
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

// Create is used to add a new user
func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(user).Error
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return nil
}

// Update is used to update user data in the db
func (ug *userGorm) Update(user *User) error {
	return ug.db.Save(user).Error
}

// Delete is used to delete a user from the db
func (ug *userGorm) Delete(id uint) error {
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

// AutoMigrate is used to automatically migrate the relations in the db
func (ug *userGorm) AutoMigrate() error {
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

/*
	***********************************************
	***********************************************
	Start of validation and normalization functions
	***********************************************
	***********************************************
*/

// Validation code for ByRemember
func (uv *userValidator) ByRemember(token string) (*User, error) {
	user := User{
		Remember: token,
	}
	if err := runUserValFns(&user, uv.hmacRemember); err != nil {
		return nil, err
	}
	return uv.UserDB.ByRemember(user.RememberHash)
}

// Validation code for Create
func (uv *userValidator) Create(user *User) error {
	log.Println("In Create Validator")
	err := runUserValFns(user,
		uv.requireName,
		uv.requireEmail,
		uv.passwordRequired,
		uv.passwordMinLength,
		uv.bcryptPassword,
		uv.passwordHashRequired,
		uv.setRememberIfUnset,
		uv.rememberMinBytes,
		uv.hmacRemember,
		uv.rememberHashRequired,
		uv.normalizeEmail,
		uv.emailFormat,
		uv.emailIsAvail)
	if err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

// Validation code for Update
func (uv *userValidator) Update(user *User) error {
	err := runUserValFns(user,
		uv.passwordMinLength,
		uv.bcryptPassword,
		uv.passwordHashRequired,
		uv.rememberMinBytes,
		uv.hmacRemember,
		uv.rememberHashRequired,
		uv.requireEmail,
		uv.normalizeEmail,
		uv.emailFormat,
		uv.emailIsAvail)
	if err != nil {
		return err
	}
	return uv.UserDB.Update(user)
}

// Validation code for Delete
func (uv *userValidator) Delete(id uint) error {
	var user User
	user.ID = id
	err := runUserValFns(&user, uv.idGreaterThan(0))
	if err != nil {
		return err
	}
	return uv.UserDB.Delete(id)
}

func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	err := runUserValFns(&user, uv.normalizeEmail)
	if err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user.Email)
}

// Normalization code for emails
func (uv *userValidator) normalizeEmail(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

func (uv *userValidator) requireEmail(user *User) error {
	if user.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

func (uv *userValidator) emailFormat(user *User) error {
	if user.Email == "" {
		return nil
	}
	if !uv.emailRegex.MatchString(user.Email) {
		return ErrEmailInvalid
	}
	return nil
}

func (uv *userValidator) emailIsAvail(user *User) error {
	existing, err := uv.ByEmail(user.Email)
	if err == ErrNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	if user.ID != existing.ID {
		return ErrEmailTaken
	}
	return nil
}

func (uv *userValidator) requireName(user *User) error {
	if user.Name == "" {
		return ErrNameRequired
	}
	return nil
}

func (uv *userValidator) passwordMinLength(user *User) error {
	if len(user.Password) < 8 {
		return ErrPasswordTooShort
	}
	return nil
}

func (uv *userValidator) passwordRequired(user *User) error {
	if user.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func (uv *userValidator) passwordHashRequired(user *User) error {
	if user.PasswordHash == "" {
		return ErrPasswordRequired
	}
	return nil
}

func (uv *userValidator) rememberHashRequired(user *User) error {
	if user.RememberHash == "" {
		return ErrRememberRequired
	}
	return nil
}

func (uv *userValidator) rememberMinBytes(user *User) error {
	if user.Remember == "" {
		return nil
	}
	n, err := rand.NBytes(user.Remember)
	if err != nil {
		return err
	}
	if n < 32 {
		return ErrRememberTooShort
	}
	return nil
}

/*
	******************************
	******************************
	Re-usable validation functions
	******************************
	******************************
*/

func (uv *userValidator) bcryptPassword(user *User) error {
	if user.Password == "" { // to check whether the password has been updated
		return nil
	}
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return nil
}

func (uv *userValidator) hmacRemember(user *User) error {
	if user.Remember == "" {
		return nil
	}
	user.RememberHash = uv.hmac.Hash(user.Remember)
	return nil
}

func (uv *userValidator) setRememberIfUnset(user *User) error {
	if user.Remember != "" {
		return nil
	}
	token, err := rand.RememberToken()
	if err != nil {
		return err
	}
	user.Remember = token
	return nil
}

func (uv *userValidator) idGreaterThan(n uint) userValFn {
	return userValFn(func(user *User) error {
		if user.ID <= n {
			return ErrIDInvalid
		}
		return nil
	})
}

func runUserValFns(user *User, fns ...userValFn) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

// simple statements to check whether the interface is implemented properly
// will give error in compile time, if not implemented properly
var _ UserDB = &userGorm{}
var _ UserService = &userService{}
