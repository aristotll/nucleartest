package squirrel

import (
	sq "github.com/Masterminds/squirrel"
	"log"
)

// 感觉这玩意没什么卵用，还不如直接写 sql 来的实在

func test() {
	user := sq.
		Select("*").
		From("user").
		Join("email USING (email_id)")

	active := user.Where(sq.Eq{"delete_at": nil})

	sql, args, err := active.ToSql()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("args: ", args)
	log.Println("sql: ", sql)

	// Output:
	// args: []
	// sql: SELECT * FROM user JOIN email USING (email_id) WHERE delete_at IS NULL
}

func test1() {
	user := sq.
		Select("stu.id, stu.name, stu.age, teacher.id, teacher.name").
		From("stu").
		Join("teacher").
		JoinClause("on stu.teacher_id = teacher.id")

	sql, args, err := user.ToSql()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("args: ", args)
	log.Println("sql: ", sql)
}
