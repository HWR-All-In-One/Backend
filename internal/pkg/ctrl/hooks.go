package ctrl

import (
	"errors"

	"github.com/HWR-All-In-One/Backend/internal/pkg/encrypt"
	"github.com/HWR-All-In-One/Backend/internal/pkg/hwr"
	"github.com/pocketbase/pocketbase/core"
)

func (a *App) encryptHwrPaswordRecord() {
	a.PB.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "users" {
			key := a.Safe.Get()

			password := e.Record.GetString("hwr_password")
			username := e.Record.GetString("hwr_email")
			isValid, err := hwr.ValidateUser(username, password)

			if err != nil {
				return err
			}

			if !isValid {
				return errors.New("user does not exist")
			}

			enc, err := encrypt.AESEncrypt(key, password)

			if err != nil {
				return err
			}
			e.Record.Set("hwr_password", enc)
		}

		return nil
	})
}

func (a *App) AddHooks() {
	a.encryptHwrPaswordRecord()
}
