package model

type User struct {
	Username      string
	GamesPlayed   int
	GamesWon      int
	TotalAttempts int
}

func (u *User) RecordGame(win bool, attempts int) {
	u.GamesPlayed++
	if win {
		u.GamesWon++
	}
	u.TotalAttempts += attempts
}

func (u *User) AverageAttempts() float64 {
	if u.GamesPlayed == 0 {
		return 0
	}
	return float64(u.TotalAttempts) / float64(u.GamesPlayed)
}
