package main

type Student struct {
	UserName string
	Sex      int
	Score    float32
	Grade    string
}

func NewStudent(username string, sex int, score float32, grade string) (stu *Student) {
	stu = &Student{
		UserName: username,
		Sex:      sex,
		Score:    score,
		Grade:    grade,
	}
	return
}
