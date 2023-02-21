package ctrl

func (a *App) AddHooks() {
	a.PB.OnRecordBeforeCreateRequest().Add(a.Views.HWR.ValidateUser)
	a.PB.OnRecordsListRequest().Add(a.Views.Timetable.CheckStatus)
}
