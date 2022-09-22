package main

import (
	"fmt"
	"strings"
	"crypto/sha256"
)

type Student struct {
	rollNo   int
	name     string
	address  string
	subjects []string
}

func addStudent(rollNo int, name string, address string, sub []string) *Student {
	s := new(Student)
	s.rollNo = rollNo
	s.name = name
	s.address = address
	s.subjects = sub
	return s
}

type StudentList struct {
	sList []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, s1 []string) *Student {
	s := addStudent(rollno, name, address, s1)
	ls.sList = append(ls.sList, s)
	return s
}


// Write a function to print the details of all the students in student list

func (ls *StudentList) PrintStudentDetails() {
	counter := 1
	for _, s := range ls.sList {

		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 10), counter, strings.Repeat("=", 10))
		fmt.Println("Roll No: ", s.rollNo)
		fmt.Println("Name: ", s.name)
		fmt.Println("Address: ", s.address)
		fmt.Println("Subjects: ", s.subjects)
		counter++
	}
}


func (ls StudentList) String() string {
	var buffer bytes.Buffer
	for _, s := range ls.Sl {
		buffer.Reset()

		buffer.WriteString(fmt.Sprintf("Roll No: %d, Name: %s, Address: %s, Subjects: %s", s.rollNo, s.name, s.address, s.subjects))
		sum := sha256.Sum256([]byte(buffer.String()))
		fmt.Println(buffer.String())

		fmt.Printf("%x\n", sum)
		fmt.Println()
	}
	//fmt.Println(buffer.String())
	//Print new line
	//fmt.Println()
	return buffer.String()
}


// func (ls *StudentList) CalcHash() {
// 	counter := 1
// 	for _, s := range ls.sList {
// 		arr := []string{}
// 		arr = append{ s.rollNo}
// 		sum := sha256.Sum256([]byte(arr))
// 		fmt.Printf("%s\n", sum)
// 		// fmt.Printf("%s List %d %s \n", strings.Repeat("=", 10), counter, strings.Repeat("=", 10))
// 		// fmt.Println("Roll No: ", s.rollNo)
// 		// fmt.Println("Name: ", s.name)
// 		// fmt.Println("Address: ", s.address)
// 		// fmt.Println("Subjects: ", s.subjects)
// 		counter++
// 	}
// }


func main() {
	s1 := []string{"pf", "IS", "SMD"}
	s2 := []string{"pf", "IS", "SMD"}
	s3 := []string{"pf", "IS", "SMD"}

	ls := new(StudentList)

	ls.CreateStudent(1, "Shamshad", "PAK", s1)
	ls.CreateStudent(2, "Jawad", "Islamabad", s2)
	ls.CreateStudent(3, "Thomson", "planet", s3)

	ls.PrintStudentDetails()
}
