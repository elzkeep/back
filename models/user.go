package models

import (
    //"aoi/config"
    
    "aoi/models/user"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type User struct {
            
    Id                int64 `json:"id"`         
    Email                string `json:"email"`         
    Passwd                string `json:"passwd"`         
    Name                string `json:"name"`         
    Level                user.Level `json:"level"`         
    Status                user.Status `json:"status"`         
    Elo                Double `json:"elo"`         
    Count                int `json:"count"`         
    Image                string `json:"image"`         
    Profile                string `json:"profile"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type UserManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *User) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewUserManager(conn interface{}) *UserManager {
    var item UserManager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        if v, ok := conn.(*sql.DB); ok {
            item.Conn = v
            item.Tx = nil
        } else {
            item.Tx = conn.(*sql.Tx)
            item.Conn = nil
        }
    }

    item.Index = ""

    return &item
}

func (p *UserManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *UserManager) SetIndex(index string) {
    p.Index = index
}

func (p *UserManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *UserManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *UserManager) GetQuery() string {
    ret := ""

    str := "select u_id, u_email, u_passwd, u_name, u_level, u_status, u_elo, u_count, u_image, u_profile, u_date from user_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *UserManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from user_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *UserManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate user_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *UserManager) Insert(item *User) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into user_tb (u_id, u_email, u_passwd, u_name, u_level, u_status, u_elo, u_count, u_image, u_profile, u_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Email, item.Passwd, item.Name, item.Level, item.Status, item.Elo, item.Count, item.Image, item.Profile, item.Date)
    } else {
        query = "insert into user_tb (u_email, u_passwd, u_name, u_level, u_status, u_elo, u_count, u_image, u_profile, u_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Email, item.Passwd, item.Name, item.Level, item.Status, item.Elo, item.Count, item.Image, item.Profile, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *UserManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from user_tb where u_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *UserManager) DeleteWhere(args []interface{}) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := ""
    var params []interface{}

    for _, arg := range args {
        switch v := arg.(type) {        
        case Where:
            item := v

            if item.Compare == "in" {
                query += " and u_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

             query += " and " + item.Query
        }        
    }

    query = "delete from user_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *UserManager) Update(item *User) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update user_tb set u_email = ?, u_passwd = ?, u_name = ?, u_level = ?, u_status = ?, u_elo = ?, u_count = ?, u_image = ?, u_profile = ?, u_date = ? where u_id = ?"
	_, err := p.Exec(query , item.Email, item.Passwd, item.Name, item.Level, item.Status, item.Elo, item.Count, item.Image, item.Profile, item.Date, item.Id)
    
        
    return err
}


func (p *UserManager) UpdateEmail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_email = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdatePasswd(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_passwd = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateName(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_name = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateLevel(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_level = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_status = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateElo(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_elo = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateCount(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_count = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateImage(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_image = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateProfile(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_profile = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *UserManager) IncreaseElo(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_elo = u_elo + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) IncreaseCount(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_count = u_count + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *UserManager) GetIdentity() int64 {
    if p.Result == nil && p.Tx == nil {
        return 0
    }

    id, err := (*p.Result).LastInsertId()

    if err != nil {
        return 0
    } else {
        return id
    }
}

func (p *User) InitExtra() {
    p.Extra = map[string]interface{}{
            "level":     user.GetLevel(p.Level),
            "status":     user.GetStatus(p.Status),

    }
}

func (p *UserManager) ReadRow(rows *sql.Rows) *User {
    var item User
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Email, &item.Passwd, &item.Name, &item.Level, &item.Status, &item.Elo, &item.Count, &item.Image, &item.Profile, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {

        item.InitExtra()
        
        return &item
    }
}

func (p *UserManager) ReadRows(rows *sql.Rows) []User {
    var items []User

    for rows.Next() {
        var item User
        
    
        err := rows.Scan(&item.Id, &item.Email, &item.Passwd, &item.Name, &item.Level, &item.Status, &item.Elo, &item.Count, &item.Image, &item.Profile, &item.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *UserManager) Get(id int64) *User {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and u_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *UserManager) Count(args []interface{}) int {
    if p.Conn == nil && p.Tx == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect()

    for _, arg := range args {
        switch v := arg.(type) {
        case Where:
            item := v

            if item.Compare == "in" {
                query += " and u_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }            
        case Custom:
             item := v

             query += " and " + item.Query
        }
    }

    rows, err := p.Query(query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return 0
    }

    defer rows.Close()

    if !rows.Next() {
        return 0
    }

    cnt := 0
    err = rows.Scan(&cnt)

    if err != nil {
        return 0
    } else {
        return cnt
    }
}

func (p *UserManager) FindAll() []User {
    return p.Find(nil)
}

func (p *UserManager) Find(args []interface{}) []User {
    if p.Conn == nil && p.Tx == nil {
        var items []User
        return items
    }

    var params []interface{}
    baseQuery := p.GetQuery()
    query := ""

    page := 0
    pagesize := 0
    orderby := ""
    
    for _, arg := range args {
        switch v := arg.(type) {
        case PagingType:
            item := v
            page = item.Page
            pagesize = item.Pagesize            
        case OrderingType:
            item := v
            orderby = item.Order
        case LimitType:
            item := v
            page = 1
            pagesize = item.Limit
        case OptionType:
            item := v
            if item.Limit > 0 {
                page = 1
                pagesize = item.Limit
            } else {
                page = item.Page
                pagesize = item.Pagesize                
            }
            orderby = item.Order
        case Where:
            item := v

            if item.Compare == "in" {
                query += " and u_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

             query += " and " + item.Query
        case Base:
             item := v

             baseQuery = item.Query
        }
    }
    
    startpage := (page - 1) * pagesize
    
    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "u_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "u_" + orderby
            }
            
        }
        query += " order by " + orderby
        //if config.Database == "mysql" {
            query += " limit ? offset ?"
            params = append(params, pagesize)
            params = append(params, startpage)
            /*
        } else if config.Database == "mssql" || config.Database == "sqlserver" {
            query += "OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
            params = append(params, startpage)
            params = append(params, pagesize)
        }
        */
    } else {
        if orderby == "" {
            orderby = "u_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "u_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []User
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *UserManager) GetByEmail(email string, args ...interface{}) *User {
    if email != "" {
        args = append(args, Where{Column:"email", Value:email, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}

func (p *UserManager) CountByEmail(email string, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if email != "" { 
        rets = append(rets, Where{Column:"email", Value:email, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *UserManager) FindByLevel(level user.Level, args ...interface{}) []User {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)

    if level != 0 { 
        rets = append(rets, Where{Column:"level", Value:level, Compare:"="})
     }
    
    return p.Find(rets)
}

func (p *UserManager) UpdateImageById(image string, id int64) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "update user_tb set u_image = ? where 1=1 and u_id = ?"
	_, err := p.Exec(query, image, id)

    return err    
}


func (p *UserManager) Sum(args []interface{}) *User {
    if p.Conn == nil && p.Tx == nil {
        var item User
        return &item
    }

    var params []interface{}

    
    query := "select count from user_tb"

    if p.Index != "" {
        query = query + " use index(" + p.Index + ") "
    }

    query += "where 1=1 "

    page := 0
    pagesize := 0
    orderby := ""
    
    for _, arg := range args {
        switch v := arg.(type) {
        case PagingType:
            item := v
            page = item.Page
            pagesize = item.Pagesize
        case OrderingType:
            item := v
            orderby = item.Order
        case LimitType:
            item := v
            page = 1
            pagesize = item.Limit
        case OptionType:
            item := v
            if item.Limit > 0 {
                page = 1
                pagesize = item.Limit
            } else {
                page = item.Page
                pagesize = item.Pagesize                
            }
            orderby = item.Order
        case Where:
            item := v

            if item.Compare == "in" {
                query += " and u_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

             query += " and " + item.Query
        }        
    }
    
    startpage := (page - 1) * pagesize
    
    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "u_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "u_" + orderby
            }
            
        }
        query += " order by " + orderby
        //if config.Database == "mysql" {
            query += " limit ? offset ?"
            params = append(params, pagesize)
            params = append(params, startpage)
            /*
        } else if config.Database == "mssql" || config.Database == "sqlserver" {
            query += "OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
            params = append(params, startpage)
            params = append(params, pagesize)
        }
        */
    } else {
        if orderby == "" {
            orderby = "u_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "u_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item User
    
    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        rows.Scan(&item.Count)        
    }

    return &item        
}
