package models

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateSucccess(t *testing.T) {
	user := new(User)
	user.Id = 123
	user.Name = "Jiten"
	user.Email = "Jiten@Outlook.Com"
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	err := user.Validate()

	if err != nil {
		t.Fail()
	}
}

func TestValidateFailure(t *testing.T) {
	user := new(User)
	user.Id = 123
	user.Email = "Jiten@Outlook.Com"
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	err := user.Validate()

	if err == nil {
		t.Fail()
	}
}

func TestValidateAll(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{
			"Valid User", User{Id: 100, Name: "Jiten", Email: "JItenP@Outlook.com", Status: "active", LastModified: time.Now().Unix()}, nil,
		},
		{
			"Missing Name", User{Id: 100, Email: "JItenP@Outlook.com", Status: "active", LastModified: time.Now().Unix()}, errors.New("invalid name"),
		}, {
			"Missing Email", User{Id: 100, Name: "Jiten", Status: "active", LastModified: time.Now().Unix()}, errors.New("invalid email"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualErr := tt.user.Validate() // actual result
			if tt.wantErr != nil {
				assert.EqualError(t, actualErr, tt.wantErr.Error())
			} else {
				assert.NoError(t, actualErr)
			}

			// if tt.wantErr != nil {
			// 	if actualErr.Error() == tt.wantErr.Error() {
			// 		t.Fail()
			// 	}
			// } else {
			// 	if actualErr != nil {
			// 		t.Fail()
			// 	}
			// }
		})
	}

}
