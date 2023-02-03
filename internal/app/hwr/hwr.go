package hwr

import (
	"errors"

	aes "github.com/HWR-All-In-One/Backend/internal/pkg/aes"
	"github.com/HWR-All-In-One/Backend/internal/pkg/hwr"
	"github.com/HWR-All-In-One/Backend/internal/pkg/safe"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type Environment struct {
	PB   *pocketbase.PocketBase
	Safe *safe.Safe
}

func (env *Environment) ValidateUser(e *core.RecordCreateEvent) error {
	if e.Record.Collection().Name == "users" {
		key := env.Safe.Get()
		password := e.Record.GetString("hwr_password")
		username := e.Record.GetString("hwr_email")
		isValid, err := hwr.ValidateUser(username, password)

		if err != nil {
			return err
		}

		if !isValid {
			return errors.New("hwr user does not exist")
		}

		enc, err := aes.Encrypt(key, password)

		if err != nil {
			return err
		}

		e.Record.Set("hwr_password", enc)
	}
	return nil
}
