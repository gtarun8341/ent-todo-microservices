package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"ent-todo-microservices/user-service/ent"
	"ent-todo-microservices/user-service/ent/session"
	"ent-todo-microservices/user-service/ent/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsersENT(ctx context.Context,client *ent.Client) ([]*ent.User,error){
	users,err := client.User.Query().All(ctx)
	return users,err
}

func GetUserIdFromSessionENT(ctx context.Context,client *ent.Client,sessionId string) (string,error){

	session,err := client.Session.Query().Where(session.TokenEQ(sessionId)).WithUser().Only(ctx)
	log.Print(session.ID)

	if err != nil{
		return "",errors.New("unauthorized access")
	}
	if session.ExpiryAt < time.Now().Unix(){
		return "",errors.New("session expired login again")
	}
	return fmt.Sprintf("%d", session.Edges.User.ID), err
	// return session.Edges.User.ID,err
}

func RegisterUserENT(ctx context.Context,client *ent.Client,input ent.User) error {
log.Print(input)
	exist,err := client.User.Query().Where(user.EmailEQ(input.Email)).Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = client.User.
		Create().
		SetName(input.Name).
		SetEmail(input.Email).
		SetPassword(string(hashedPassword)).
		Save(ctx)
	return err
}

func LoginUserENT(ctx context.Context,client *ent.Client,input ent.User) ([]*ent.Session,error){
	userFromDb,err := client.User.Query().Where(user.EmailEQ(input.Email)).Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("email doesn't exist")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token := uuid.NewString()
	expiry := time.Now().Add(time.Hour * 1).Unix()

	session, err := client.Session.
		Create().
		SetUserID(userFromDb.ID).
		SetToken(token).
		SetExpiryAt(expiry).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return []*ent.Session{session}, nil
}
