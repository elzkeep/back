package controllers

import (
	"path"
	"sync"
	"zkeep/config"
	"zkeep/global"
	"zkeep/models"

	"database/sql"
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/CloudyKit/jet/v3"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type Controller struct {
	Context    *fiber.Ctx
	Vars       jet.VarMap
	Result     map[string]interface{}
	Connection *sql.DB
	Current    string
	Code       int
	Session    *models.User

	Date string

	Page     int
	Pagesize int
	Mutex    sync.Mutex
}

func NewController(g *fiber.Ctx) *Controller {
	var ctl Controller
	ctl.Init(g)
	return &ctl
}

func (c *Controller) Init(g *fiber.Ctx) {
	c.Mutex = sync.Mutex{}
	c.Context = g
	c.Vars = make(jet.VarMap)
	c.Result = make(map[string]interface{})
	c.Result["code"] = "ok"
	c.Connection = c.NewConnection()
	c.Code = http.StatusOK

	user := g.Locals("user")

	if user != nil {
		c.Session = user.(*models.User)
		c.Session.Passwd = ""
	} else {
		c.Session = nil
	}

	t := time.Now()
	c.Date = global.GetDate(t)

	c.Set("_t", t.UnixNano())
}

func (c *Controller) Lock() {
	c.Mutex.Lock()
}

func (c *Controller) Unlock() {
	c.Mutex.Unlock()
}

func (c *Controller) Error(err string) {
	c.Set("code", "error")
	c.Set("error", err)
}

func (c *Controller) Set(name string, value interface{}) {
	c.Result[name] = value
}

func (c *Controller) SetArray(value map[string]interface{}) {
	for k, v := range value {
		c.Result[k] = v
	}
}

func (c *Controller) GetArrayComma(name string) []string {
	value := c.Get(name)

	return strings.Split(value, ",")
}

func (c *Controller) GetArrayCommai(name string) []int {
	value := c.Get(name)

	var items []int

	if value == "" {
		return items
	}

	values := strings.Split(value, ",")
	for _, item := range values {
		items = append(items, global.Atoi(item))
	}

	return items
}

func (c *Controller) Get(name string) string {
	return c.Query(name)
}

func (c *Controller) GetStartdate(name string) string {
	date := c.Get(name)

	if date != "" {
		date += ":00"
	}

	return date
}

func (c *Controller) GetEnddate(name string) string {
	date := c.Get(name)

	if date != "" {
		date += ":59"
	}

	return date
}

func (c *Controller) Geti(name string) int {
	return c.Queryi(name)
}

func (c *Controller) Geti64(name string) int64 {
	return c.Queryi64(name)
}

func (c *Controller) Getf(name string) float64 {
	return c.Queryf(name)
}

func (c *Controller) Geti64Array(name string) []int64 {
	str := c.Get(name)

	ret := make([]int64, 0)

	if str == "" {
		return ret
	}

	items := strings.Split(str, ",")

	for _, v := range items {
		ret = append(ret, global.Atol(strings.TrimSpace(v)))
	}

	return ret
}

func (c *Controller) DefaultGet(name string, defaultValue string) string {
	return c.DefaultQuery(name, defaultValue)
}

func (c *Controller) DefaultGeti(name string, defaultValue int) int {
	return c.DefaultQueryi(name, defaultValue)
}

func (c *Controller) DefaultGeti64(name string, defaultValue int64) int64 {
	return c.DefaultQueryi64(name, defaultValue)
}

func (c *Controller) Query(name string) string {
	return c.Context.Query(name)
}

func (c *Controller) Queryi(name string) int {
	value, _ := strconv.Atoi(c.Context.Query(name))
	return value
}

func (c *Controller) Queryi64(name string) int64 {
	value, _ := strconv.ParseInt(c.Context.Query(name), 10, 64)
	return value
}

func (c *Controller) Queryf(name string) float64 {
	value, _ := strconv.ParseFloat(c.Context.Query(name), 64)
	return value
}

func (c *Controller) DefaultQuery(name string, defaultValue string) string {
	value := c.Context.Query(name)

	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultQueryi(name string, defaultValue int) int {
	value, _ := strconv.Atoi(c.Context.Query(name))

	if value == 0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultQueryi64(name string, defaultValue int64) int64 {
	value, _ := strconv.ParseInt(c.Context.Query(name), 10, 64)

	if value == 0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) DefaultQueryf(name string, defaultValue float64) float64 {
	value, _ := strconv.ParseFloat(c.Context.Query(name), 64)

	if value == 0.0 {
		return defaultValue
	} else {
		return value
	}
}

func (c *Controller) Download(filename string, downloadFilename string) {
	filesize, _ := os.Stat(filename)
	c.Context.Append("Content-Type", "application/octet-stream")
	c.Context.Append("Content-Length", fmt.Sprintf("%v", filesize))
	c.Context.Append("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\";", downloadFilename))
	c.Context.Append("Content-Transfer-Encoding", "binary")
	c.Context.Append("Pragma", "no-cache")
	c.Context.Append("Expires", "0")

	c.Context.Download(filename)
}

func (c *Controller) NewConnection() *sql.DB {
	if c.Connection != nil {
		return c.Connection
	}

	c.Connection = models.NewConnection()
	return c.Connection
}

func (c *Controller) Close() {
	if c.Connection != nil {
		c.Connection.Close()
		c.Connection = nil
	}
}

func (c *Controller) Bind(obj interface{}) error {
	return c.Context.BodyParser(obj)
}

func (c *Controller) Paging(page int, totalRows int, pageSize int) {
	blockSize := 5

	totalPage := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	totalBlock := int(math.Ceil(float64(totalPage) / float64(blockSize)))
	currentBlock := int(math.Ceil(float64(page) / float64(blockSize)))

	startPage := (currentBlock-1)*blockSize + 1
	endPage := currentBlock * blockSize
	if endPage > totalPage {
		endPage = totalPage
	}

	s := make([]int, endPage-startPage+1)
	for i := range s {
		s[i] = startPage + i
	}

	c.Set("pages", s)
	c.Set("page", page)
	c.Set("blockSize", blockSize)
	c.Set("totalPage", totalPage)
	c.Set("totalBlock", totalBlock)
	c.Set("currentBlock", currentBlock)
}

func (c *Controller) GetUpload(uploadPath string, name string) (string, string) {
	log.Println("controller getupload")
	file, err := c.Context.FormFile(name)

	if err != nil {
		log.Println(err)
		return "", ""
	}

	t := time.Now()
	u2 := uuid.NewV4()

	filename := fmt.Sprintf("%04d%02d%02d%02d%02d%02d_%v%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), strings.Replace(u2.String(), "-", "", -1), filepath.Ext(file.Filename))
	if uploadPath != "" {
		filename = path.Join(uploadPath, filename)
	}

	fullFilename := path.Join(config.UploadPath, filename)

	log.Println("path : ", config.UploadPath, fullFilename)

	c.Context.SaveFile(file, fullFilename)

	return file.Filename, filename
}

func (c *Controller) IsAdmin() bool {
	return false
}
