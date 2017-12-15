package entities

type UserServiceProvider struct {}

var UserService = UserServiceProvider {}

func (*UserServiceProvider) Insert(user *User) error {
	tx, err := db.Begin()
	checkErr(err)

	dao := userDao{ tx }
	err := dao.Insert(user)
	checkErr(err)
	tx.Commit()
	return nil
}

func (*UserServiceProvider) QueryAll() []User {
	dao := userDao{db}
	return dao.QueryAll()
}

func (*UserServiceProvider) QueryByName(name string) User {
	dao := userDao{ db }
	user, err := dao.QueryByName(name)
	checkErr(err)
	return user
}

func (*UserServiceProvider)DeleteByName(name string) error {
	dao := userDao{ db }
	err := dao.DeleteByName(name)
	checkErr(err)
	return nil
}