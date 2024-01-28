package repo

import "gorm.io/gorm"

type Account struct {
	Email          string             `json:"email" gorm:"primaryKey;autoIncrement:false"`
	TargetPlanTask []*PlanCreatorTask `gorm:"foreignKey:Email"`
}

type AccountRepo struct {
	Db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{
		Db: db,
	}
}

func (repo *AccountRepo) Get(email string) (*Account, error) {
	var akun Account
	return &akun, repo.Db.Where(&Account{Email: email}).First(&akun).Error
}

func (repo *AccountRepo) Delete(email string) error {
	return repo.Db.Delete(&Account{Email: email}).Error
}

func (repo *AccountRepo) ListAccount() ([]*Account, error) {
	var hasil []*Account

	err := repo.Db.Model(&Account{}).Find(&hasil).Error
	return hasil, err
}
