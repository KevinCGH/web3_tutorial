package Task_3

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTask3(t *testing.T) {

	Convey("SQL 语句练习", t, func() {
		// "file::memory:?cache=shared"
		db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		So(err, ShouldBeNil)

		err = db.AutoMigrate(&Student{}, &Account{}, &Transaction{})
		So(err, ShouldBeNil)
		//if err != nil {
		//	t.Fatalf("Failed to migrate database: %v", err)
		//}

		Convey("基本 CRUD 操作", func() {
			db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Student{})
			repo := NewStudentRepository(db)
			// 向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
			err := repo.InsertStudent(&Student{Name: "张三", Age: 20, Grade: "三年级"})
			err = repo.InsertStudent(&Student{Name: "李四", Age: 19, Grade: "二年级"})
			err = repo.InsertStudent(&Student{Name: "王五", Age: 14, Grade: "一年级"})
			So(err, ShouldBeNil)

			var student Student
			db.First(&student, "name=?", "张三")
			So(student.Name, ShouldEqual, "张三")
			So(student.Age, ShouldEqual, 20)
			So(student.Grade, ShouldEqual, "三年级")

			// 查询 students 表中所有年龄大于 18 岁的学生信息
			students, err := repo.FindStudentsByAge(18)
			So(err, ShouldBeNil)
			So(len(students), ShouldEqual, 2)

			// 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
			err = repo.UpdateStudentGradeByName("张三", "四年级")
			So(err, ShouldBeNil)
			db.First(&student, "name=?", "张三")
			So(student.Grade, ShouldEqual, "四年级")

			// 删除 students 表中年龄小于 15 岁的学生记录
			err = repo.DeleteStudentsByAge(15)
			So(err, ShouldBeNil)
			var count int64
			db.Model(&Student{}).Count(&count)
			So(count, ShouldEqual, 2)
		})

		Convey("事务语句", func() {
			db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Account{})
			db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Transaction{})
			transferService := NewTransferService(db)

			accountA := Account{ID: 1, Balance: 500}
			accountB := Account{ID: 2, Balance: 300}

			err = db.Create(&accountA).Error
			So(err, ShouldBeNil)
			err = db.Create(&accountB).Error
			So(err, ShouldBeNil)

			Convey("正常转账测试", func() {
				err := transferService.Transfer(1, 2, 100)
				So(err, ShouldBeNil)

				// 验证账户余额
				var updatedAccountA Account
				db.Where("id = ?", 1).First(&updatedAccountA)
				So(updatedAccountA.Balance, ShouldEqual, 400)

				var updatedAccountB Account
				db.Where("id = ?", 1).First(&updatedAccountB)
				So(updatedAccountB.Balance, ShouldEqual, 400)

				// 验证交易记录
				var transaction Transaction
				err = db.Where("from_account_id = ? AND to_account_id = ?", 1, 2).First(&transaction).Error
				So(err, ShouldBeNil)
				So(transaction.Amount, ShouldEqual, 100)
			})

			Convey("余额不足转账测试", func() {
				err := transferService.Transfer(1, 2, 1000)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "账户余额不足")

				// 验证账户余额未发生变化
				var validAccountA Account
				db.Where("id = ?", 1).First(&validAccountA)
				So(validAccountA.Balance, ShouldEqual, 500)

				var validAccountB Account
				db.Where("id = ?", 2).First(&validAccountB)
				So(validAccountB.Balance, ShouldEqual, 300)
			})

			Convey("账户不存在测试", func() {
				err := transferService.Transfer(999, 2, 100)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "转出账户不存在")

				err = transferService.Transfer(1, 999, 100)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "转入账户不存在")

			})
		})
	})

	Convey("Sqlx 入门", t, func() {
		db, err := sqlx.Open("sqlite3", "file::memory:?cache=shared")
		So(err, ShouldBeNil)

		Convey("查询", func() {
			repo := NewEmployeeRepository(db)
			err = repo.CreateTable()
			So(err, ShouldBeNil)
			err := repo.InsertEmployee(&Employee{
				Department: "技术部",
				Name:       "张三",
				Salary:     5000,
			})
			if err != nil {
				fmt.Println("插入员工数据失败")
			}

			Convey("查询技术部员工", func() {
				employees, err := repo.FindEmployeesByDepartment("技术部")
				So(err, ShouldBeNil)
				So(employees, ShouldNotBeNil)
				for _, emp := range employees {
					So(emp.Department, ShouldEqual, "技术部")
				}
			})

			Convey("查询工资最高的员工", func() {
				employee, err := repo.FindHighestPaidEmployee()
				So(err, ShouldBeNil)
				So(employee, ShouldNotBeNil)
				So(employee.Salary, ShouldBeGreaterThanOrEqualTo, 0)
			})
		})

		Convey("实现类型安全映射", func() {
			repo := NewBookRepository(db)
			err := repo.CreateTable()
			So(err, ShouldBeNil)
			err = repo.InsertBook(&Book{
				Title:  "Go语言基础",
				Author: "Go语言",
				Price:  69.9,
			})
			if err != nil {
				fmt.Println("插入书本数据失败")
			}

			Convey("查询价格大于 50元的书本", func() {
				books, err := repo.FindBooksByPrice(50.0)
				So(err, ShouldBeNil)
				So(books, ShouldNotBeNil)
				for _, book := range books {
					So(book.Price, ShouldBeGreaterThan, 50)
				}
			})
		})
	})

	Convey("进阶 GORM", t, func() {
		// 初始化
		db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
		So(err, ShouldBeNil)
		db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{}, &Post{}, &Comment{})

		Convey("测试 Blog 系统", func() {
			// 创建测试用户
			Convey("关联查询", func() {
				user := User{Name: "Alice", Email: "alice@example.com"}
				db.Create(&user)
				service := &BlogService{DB: db}
				post1 := Post{Title: "Post 1", Content: "Content 1", UserID: user.ID}
				post2 := Post{Title: "Post 2", Content: "Content 2", UserID: user.ID}
				db.Create(&post1)
				db.Create(&post2)

				// 创建评论
				comment1 := Comment{Content: "Comment 1", PostID: post1.ID, UserID: user.ID}
				comment2 := Comment{Content: "Comment 2", PostID: post1.ID, UserID: user.ID}
				comment3 := Comment{Content: "Comment 3", PostID: post1.ID, UserID: user.ID}
				db.Create(&comment1)
				db.Create(&comment2)
				db.Create(&comment3)

				// 测试查询用户的文章和评论过
				posts, err := service.GetUserPostsAndComments(user.ID)
				So(err, ShouldBeNil)
				So(len(posts), ShouldEqual, 2)
				for _, post := range posts {
					if post.ID == post1.ID {
						So(len(post.Comments), ShouldEqual, 3)
					} else if post.ID == post2.ID {
						So(len(post.Comments), ShouldEqual, 0)
					}
				}

				// 测试查询评论最多的文章
				post, err := service.GetMostCommentedPost()
				So(err, ShouldBeNil)
				So(post.ID, ShouldEqual, post1.ID)
				So(post.CommentCount, ShouldEqual, 3)

			})

			Convey("测试文章创建和钩子函数", func() {
				user := &User{Name: "Test User 1", Email: "test1@example.com"}
				db.Create(&user)

				// 新用户文章数为 0
				So(user.PostCount, ShouldEqual, 0)
				post := &Post{
					Title:   "Test Post",
					Content: "This is a test post",
					UserID:  user.ID,
				}
				err = db.Create(&post).Error
				So(err, ShouldBeNil)

				// 重新加载用户信息，检查钩子函数是否更新了文章数
				var updatedUser User
				db.First(&updatedUser, user.ID)
				So(updatedUser.PostCount, ShouldEqual, 1)
			})

			Convey("测试评论创建和钩子函数", func() {
				user := &User{Name: "Test User 2", Email: "test2@example.com"}
				db.Create(&user)

				// 新用户文章数为 0
				So(user.PostCount, ShouldEqual, 0)
				post := Post{
					Title:   "Test Post 2",
					Content: "This is a test post",
					UserID:  user.ID,
				}
				db.Create(&post)
				So(err, ShouldBeNil)

				// 创建评论
				comment1 := Comment{Content: "Comment 1", PostID: post.ID, UserID: user.ID}
				comment2 := Comment{Content: "Comment 2", PostID: post.ID, UserID: user.ID}

				db.Create(&comment1)
				db.Create(&comment2)

				db.Delete(&comment1)
				// 重新加载用户信息，检查钩子函数是否更新了文章数
				var updatedPost Post
				db.First(&updatedPost, post.ID)
				So(updatedPost.CommentCount, ShouldEqual, 1)
			})
		})
	})
}
