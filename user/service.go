package user

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

var (
	err                  = godotenv.Load()
	emailName            = os.Getenv("EMAIL_NAME")
	emailPass            = os.Getenv("EMAIL_PASS")
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = fmt.Sprintf("PT. Refactory Akselerasi <%s>", emailName)
	CONFIG_AUTH_EMAIL    = emailName
	CONFIG_AUTH_PASSWORD = emailPass
)

type Service interface {
	GetByID(userID string) (User, error)
	GetByEmail(email string) (User, error)
	Register(input UserRegister, uuid string, avatarPath string) (User, error)
	RegisterAdmin(input UserRegister, uuid string, avatarPath string) (User, error)
	GetByUsername(username string) (User, error)
	// sendMail(to []string, cc []string, subject, message string) error
	SendEMailConfirmation(email string, confirmationKey string)
	VerifiedEmailByUserID(userID string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByID(userID string) (User, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return user, err
	}

	if user.ID == "" || len(user.ID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}
func (s *service) GetByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == "" || len(user.ID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}

func (s *service) GetByUsername(username string) (User, error) {
	// fmt.Println("masuk service getbyusername")
	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	if user.ID == "" || len(user.ID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}

func (s *service) Register(input UserRegister, uuid string, avatarPath string) (User, error) {

	checkEmailUser, err := s.repository.FindByEmail(input.Email)

	if checkEmailUser.ID != "" || len(checkEmailUser.ID) > 1 {
		return User{}, errors.New("email has been registered")
	}

	checkUsernameUser, err := s.repository.FindByEmail(input.Username)

	if checkUsernameUser.ID != "" || len(checkUsernameUser.ID) > 1 {
		return User{}, errors.New("username has been registered")
	}

	if err != nil {
		return User{}, err
	}

	genPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	var newUser = User{
		ID:            uuid,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Photo:         avatarPath,
		Username:      input.Username,
		Email:         input.Email,
		VerifiedEmail: false,
		Password:      string(genPassword),
		Role:          "user",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	user, err := s.repository.Create(newUser)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) RegisterAdmin(input UserRegister, uuid string, avatarPath string) (User, error) {
	checkEmailUser, err := s.repository.FindByEmail(input.Email)

	if checkEmailUser.ID != "" || len(checkEmailUser.ID) > 1 {
		return User{}, errors.New("email has been registered")
	}

	checkUsernameUser, err := s.repository.FindByEmail(input.Username)

	if checkUsernameUser.ID != "" || len(checkUsernameUser.ID) > 1 {
		return User{}, errors.New("username has been registered")
	}

	if err != nil {
		return User{}, err
	}

	genPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	var newUser = User{
		ID:            uuid,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Photo:         avatarPath,
		Username:      input.Username,
		Email:         input.Email,
		VerifiedEmail: false,
		Password:      string(genPassword),
		Role:          "admin",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	user, err := s.repository.Create(newUser)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) VerifiedEmailByUserID(userID string) (User, error) {
	var dataUpdate = map[string]interface{}{}

	checkUser, err := s.repository.FindByID(userID)

	if err != nil {
		return checkUser, err
	}

	if checkUser.ID == "" || len(checkUser.ID) == 0 {
		return checkUser, errors.New("user id not found")
	}

	dataUpdate["verified_email"] = true
	dataUpdate["updated_at"] = time.Now()

	user, err := s.repository.UpdateByUserID(userID, dataUpdate)
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (s *service) sendMail(to []string, cc []string, subject, message string) error {
// 	body := "From: " + CONFIG_SENDER_NAME + "\n" +
// 		"To: " + strings.Join(to, ",") + "\n" +
// 		"Cc: " + strings.Join(cc, ",") + "\n" +
// 		"Subject: " + subject + "\n\n" +
// 		message

// 	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
// 	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

// 	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (s *service) SendEMailConfirmation(email string, confirmationKey string) {
	sendLink := fmt.Sprintf(`<p>Or you cna click link here for confirmation :</p><form action="http://localhost:8080/api/email_confirmation/%s" method="post"><button type="submit"> confirmation email </button></form>`, confirmationKey)
	sendBody := fmt.Sprintf("<h4>Thank you for registration to out application</h4><p>this is email confirmation key : %s </p> %s <p>have a nice day</p>", confirmationKey, sendLink)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetAddressHeader("Cc", "pratama.1208979@gmail.com", "pratama")
	mailer.SetHeader("Subject", "Email Confirmation")
	mailer.SetBody("text/html", sendBody)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
