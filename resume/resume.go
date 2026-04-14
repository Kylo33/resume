package resume

import (
	"fmt"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
)

type Resume struct {
	Name      string
	Email     string
	Phone     string
	Location  string
	Website   string
	Github    string
	Skills    []Skill
	Education []Education
	Work      []Job
	Projects  []Project
}

type Skill struct {
	Category string
	Items    []string
}

type Education struct {
	Institutiton string
	Location     string
	StartDate    YearMonthTime `yaml:"start_date"`
	EndDate      YearMonthTime `yaml:"end_date"`
	Degree       string
	Gpa          string
	Extra        []string
}

type Job struct {
	Title     string
	Company   string
	Location  string
	StartDate YearMonthTime `yaml:"start_date"`
	EndDate   YearMonthTime `yaml:"end_date"`
}

type Project struct {
	Name  string
	Url   string
	Tags  []string
	Date  YearMonthTime
	Extra []string
}

func Parse(src []byte) Resume {
	r := Resume{}
	yaml.Unmarshal(src, &r)
	return r
}

type YearMonthTime time.Time

func (t *YearMonthTime) UnmarshalYAML(dat []byte) error {
	dateString := strings.TrimSpace(string(dat))
	time, err := time.Parse("2006-01", dateString)
	if err != nil {
		fmt.Printf("Error parsing %v: %v\n", dateString, err)
	}

	*t = YearMonthTime(time)
	return nil
}

func (t *YearMonthTime) String() string {
	return time.Time(*t).Format("Jan 2006")
}
