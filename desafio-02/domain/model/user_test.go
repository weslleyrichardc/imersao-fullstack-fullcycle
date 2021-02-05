package model_test

import (
	"testing"
	
	"github.com/weslleyrichardc/imersao/desafio-01/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "Weslley Richard"
	email := "weslleyrichardc@gmail.com"

	user, err := model.NewUser(name, email)
	
	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, error := model.NewUser("", "")

	require.NotNil(t, error)
}

