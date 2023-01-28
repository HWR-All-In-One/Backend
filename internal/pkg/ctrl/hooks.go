package ctrl

import (
	"github.com/HWR-All-In-One/Backend/internal/pkg/encrypt"
	"github.com/pocketbase/pocketbase/core"
)

func (a *App) encryptHwrPaswordRecord() {
	a.PB.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "users" {
			hwrPassword := e.Record.GetString("hwr_password")
			passwordHash := e.Record.GetString("passwordHash")
			enc, err := encrypt.AESEncrypt(passwordHash, hwrPassword)

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
