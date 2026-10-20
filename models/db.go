package models

import (
	"fmt"
	"regexp"
	"strings"
	"zkeep/config"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type PagingType struct {
	Page     int
	Pagesize int
}

type OrderingType struct {
	Order string
}

type LimitType struct {
	Limit int
}

type OptionType struct {
	Page     int
	Pagesize int
	Order    string
	Limit    int
}

type Where struct {
	Column  string
	Value   interface{}
	Compare string
}

type Custom struct {
	Query string
}

type Base struct {
	Query string
}

func Paging(page int, pagesize int) PagingType {
	return PagingType{Page: page, Pagesize: pagesize}
}

func Ordering(order string) OrderingType {
	return OrderingType{Order: SafeOrder(order)}
}

var orderPattern = regexp.MustCompile(`^[A-Za-z0-9_.]+([ \t]+(?i:asc|desc))?$`)

// 정렬 문자열은 SQL에 그대로 붙으므로 "컬럼[ asc|desc], ..." 형식만 허용하고
// 그 외에는 무시(기본 정렬)함
func SafeOrder(order string) string {
	if strings.TrimSpace(order) == "" {
		return ""
	}

	for _, v := range strings.Split(order, ",") {
		if !orderPattern.MatchString(strings.TrimSpace(v)) {
			log.Println("invalid order :", order)
			return ""
		}
	}

	return order
}

// Custom 조건은 바인딩 파라미터를 지원하지 않으므로, 작은따옴표 문자열 안에 넣을 값은 반드시 이스케이프
func EscapeString(str string) string {
	str = strings.ReplaceAll(str, "\\", "\\\\")
	str = strings.ReplaceAll(str, "'", "''")

	return str
}

func Limit(limit int) LimitType {
	return LimitType{Limit: limit}
}

func GetConnection() *sql.DB {
	r1, err := sql.Open(config.Database.Type, config.Database.ConnectionString)
	if err != nil {
		log.Println("Database Connect Error")
		return nil
	}

	r1.SetMaxOpenConns(100)
	r1.SetMaxIdleConns(10)
	r1.SetConnMaxLifetime(5 * time.Minute)

	return r1
}

func NewConnection() *sql.DB {
	db := GetConnection()

	if db != nil {
		return db
	}

	time.Sleep(100 * time.Millisecond)

	db = GetConnection()

	if db != nil {
		return db
	}

	time.Sleep(500 * time.Millisecond)

	db = GetConnection()

	if db != nil {
		return db
	}

	time.Sleep(1 * time.Second)

	db = GetConnection()

	if db != nil {
		return db
	}

	time.Sleep(2 * time.Second)

	db = GetConnection()

	return db
}

func QueryArray(db *sql.DB, query string, items []interface{}) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	rows, err = db.Query(query, items...)
	return rows, err
}

func ExecArray(db *sql.DB, query string, items []interface{}) error {
	var err error

	_, err = db.Exec(query, items...)
	return err
}

func InitDate() string {
	return "1000-01-01 00:00:00"
}

type Double float64

func (c Double) MarshalJSON() ([]byte, error) {
	if float64(c) == float64(int(c)) {
		return []byte(fmt.Sprintf("%v.0", int64(c))), nil
	}

	return []byte(fmt.Sprintf("%v", float64(c))), nil
}
