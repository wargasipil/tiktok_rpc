package repo

type BotAuthData struct {
	ID    uint `gorm:"primarykey"`
	Email string
}
