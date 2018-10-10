package benchs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var g_orm *gorm.DB

func init() {
	st := NewSuite("gorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GormInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GormInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GormUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GormRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GormReadSlice)

		db, _ := gorm.Open("mysql", ORM_SOURCE)
		g_orm = db.Table("model")
		//defer db.Close()
	}
}

func GormInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		g_orm.Create(m)
		if g_orm.NewRecord(m) {
			fmt.Println("error")
			b.FailNow()
		}
	}
}

func GormInsertMulti(b *B) {
	panic(fmt.Errorf("Not support multi insert"))
	var ms []*Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.InsertMulti(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		g_orm.Create(m)
	})

	for i := 0; i < b.N; i++ {

		if err := g_orm.Update(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		g_orm.Create(m)
	})

	for i := 0; i < b.N; i++ {
		if err := g_orm.Select(m).Error; err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GormReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.Id = 0
			if err := g_orm.Create(m).Error; err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := g_orm.Where("id > ?", 0).Limit(100).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
		fmt.Println(models)
	}
}
