package model

func initDB() {
	if !db.Migrator().HasTable(&BookData{}) {
		db.Migrator().CreateTable(&BookData{})

		// 初始化数据
		initBookDataList := []*BookData{
			&BookData{Name: "The Moon and Sixpence", Author: "Maugham, W. Somerset", ISBN: 9781420951929, Published: "2016-01-05"},
			&BookData{Name: "The Old Man and the Sea", Author: "Hemingway, Ernest", ISBN: 9780684801223, Published: "1995-05-05"},
			&BookData{Name: "The Nightingale and the Rose", Author: "Wilde, Oscar", ISBN: 9798461534783, Published: "2021"},
			&BookData{Name: "Naissance Du Purgatoire", Author: "Le Jacques, Goff", ISBN: 9782070326440, Published: "1991-09-01"},
		}
		db.Create(initBookDataList)
	}

	if !db.Migrator().HasTable(&User{}) {
		db.Migrator().CreateTable(&User{})
	}
}
