package Task_3

import (
	"errors"

	"gorm.io/gorm"
)

type Student struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(100)" json:"name"`
	Age   uint   `gorm:"type:int" json:"age"`
	Grade string `gorm:"type:varchar(50)" json:"grade"`
}

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

// InsertStudent 插入新学生记录
func (r *StudentRepository) InsertStudent(student *Student) error {
	return r.db.Create(student).Error
}

// FindStudentsByAge 根据年龄查询学生
func (r *StudentRepository) FindStudentsByAge(minAge uint) ([]Student, error) {
	var students []Student
	err := r.db.Where("age > ?", minAge).Find(&students).Error
	return students, err
}

// UpdateStudentGradeByName 根据姓名更新学生成绩
func (r *StudentRepository) UpdateStudentGradeByName(name, newGrade string) error {
	return r.db.Model(&Student{}).Where("name = ?", name).Update("grade", newGrade).Error
}

// DeleteStudentsByAge 根据年龄删除学生
func (r *StudentRepository) DeleteStudentsByAge(maxAge int) error {
	return r.db.Where("age < ?", maxAge).Delete(&Student{}).Error
}

type Account struct {
	ID      uint
	Balance int64
}

type Transaction struct {
	ID            uint  `gorm:"primaryKey"`
	FromAccountID uint  `gorm:"not null"`
	ToAccountID   uint  `gorm:"not null"`
	Amount        int64 `gorm:"not null"`
}

type TransferService struct {
	db *gorm.DB
}

func NewTransferService(db *gorm.DB) *TransferService {
	return &TransferService{db: db}
}

func (s *TransferService) Transfer(fromAccountID, toAccountID uint, amount int64) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var fromAccount Account
		if err := tx.Where("id = ?", fromAccountID).First(&fromAccount).Error; err != nil {
			return errors.New("转出账户不存在")
		}

		if fromAccount.Balance < amount {
			return errors.New("账户余额不足")
		}

		var toAccount Account
		if err := tx.Where("id = ?", toAccountID).First(&toAccount).Error; err != nil {
			return errors.New("转入账户不存在")
		}

		// 扣除转出账户余额
		fromAccount.Balance -= amount
		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}

		// 增加转入账户余额
		toAccount.Balance += amount
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}

		// 记录转账交易
		transaction := Transaction{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})
}
