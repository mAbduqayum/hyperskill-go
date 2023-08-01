//package _9_University_admission

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Applicant struct {
	id               int
	fullName         string
	exams            []float64
	admissionScore   float64
	departments      []string
	departmentsScore []float64
	acceptedWave     int
}

var Departments = [5]string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}

type Admission struct {
	applicants         []Applicant
	acceptedApplicants []int
	departments        map[string][]Applicant
	studentsPerDep     int
}

func (a Admission) fillApplicants() {
	for wave := 0; wave < 3; wave++ {
		for _, department := range Departments {
			applicants := a.peekApplicants(department, wave)
			for _, applicant := range applicants {
				if len(a.departments[department]) >= a.studentsPerDep {
					break
				}
				if contains(a.acceptedApplicants, applicant.id) {
					continue
				}
				a.acceptedApplicants = append(a.acceptedApplicants, applicant.id)
				applicant.acceptedWave = wave
				a.departments[department] = append(a.departments[department], applicant)
			}
		}
	}
	for _, department := range Departments {
		sort.Slice(a.departments[department], func(i, j int) bool {
			applicantI, applicantJ := a.departments[department][i], a.departments[department][j]
			scoreI := applicantI.departmentsScore[applicantI.acceptedWave]
			scoreJ := applicantJ.departmentsScore[applicantJ.acceptedWave]
			if scoreI != scoreJ {
				return scoreI > scoreJ
			}
			return applicantI.fullName < applicantJ.fullName
		})
	}
}

func (a Admission) peekApplicants(department string, wave int) []Applicant {
	rez := make([]Applicant, 0)
	applicants := a.applicants
	for _, applicant := range applicants {
		if applicant.departments[wave] == department {
			rez = append(rez, applicant)
		}
	}
	sort.Slice(rez, func(i, j int) bool {
		applicantI, applicantJ := rez[i], rez[j]
		scoreI := applicantI.departmentsScore[wave]
		scoreJ := applicantJ.departmentsScore[wave]
		if scoreI == scoreJ {
			return applicantI.fullName < applicantJ.fullName
		}
		return scoreI > scoreJ
	})
	return rez
}

func (a Admission) results() {
	for _, department := range Departments {
		fileName := strings.ToLower(department) + ".txt"
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		applicants := a.departments[department]
		for _, applicant := range applicants {
			score := applicant.departmentsScore[applicant.acceptedWave]
			scoreF := fmt.Sprintf("%.1f", score)
			fmt.Fprintln(file, applicant.fullName, scoreF)
		}
	}
}

func NewAdmission(applicants []Applicant, n int) Admission {
	departments := make(map[string][]Applicant)
	for _, departmentName := range Departments {
		departments[departmentName] = []Applicant{}
	}
	return Admission{
		applicants:         applicants,
		departments:        departments,
		studentsPerDep:     n,
		acceptedApplicants: make([]int, 0),
	}
}

func main() {
	n, _ := strconv.Atoi(input())
	students := make([]Applicant, 0)
	for i, line := range readFile() {
		data := strings.Split(line, " ")
		fullName := data[0] + " " + data[1]
		exams := make([]float64, 0)
		for i := 0; i < 4; i++ {
			avgD, _ := strconv.ParseFloat(data[2+i], 64)
			exams = append(exams, avgD)
		}
		admissionScore, _ := strconv.ParseFloat(data[6], 64)
		departments := []string{
			data[7],
			data[8],
			data[9],
		}
		departmentsScore := scores(exams, departments, admissionScore)
		newStudent := Applicant{
			id:               i,
			fullName:         fullName,
			exams:            exams,
			admissionScore:   admissionScore,
			departments:      departments,
			departmentsScore: departmentsScore,
			acceptedWave:     -1,
		}
		students = append(students, newStudent)
	}
	admission := NewAdmission(students, n)
	admission.fillApplicants()
	admission.results()
}

func readFile() []string {
	file, err := os.Open("applicants.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	rez := make([]string, 0)
	for scanner.Scan() {
		rez = append(rez, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rez
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	return func() string {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		return line
	}()
}

func scores(exams []float64, departments []string, admissionScore float64) []float64 {
	rez := make([]float64, 0)
	var newRez float64
	for _, department := range departments {
		switch department {
		case "Physics":
			newRez = (exams[0] + exams[2]) / 2
		case "Chemistry":
			newRez = exams[1]
		case "Mathematics":
			newRez = exams[2]
		case "Engineering":
			newRez = (exams[2] + exams[3]) / 2
		case "Biotech":
			newRez = (exams[0] + exams[1]) / 2
		}
		newRez = math.Max(newRez, admissionScore)
		rez = append(rez, newRez)
	}
	return rez
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
