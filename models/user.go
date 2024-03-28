package models

import (
    //"zkeep/config"
    
    "zkeep/models/user"
    
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
    Loginid                string `json:"loginid"`         
    Passwd                string `json:"passwd"`         
    Name                string `json:"name"`         
    Email                string `json:"email"`         
    Tel                string `json:"tel"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Joindate                string `json:"joindate"`         
    Careeryear                int `json:"careeryear"`         
    Careermonth                int `json:"careermonth"`         
    Level                user.Level `json:"level"`         
    Score                Double `json:"score"`         
    Approval                user.Approval `json:"approval"`         
    Status                user.Status `json:"status"`         
    Company                int64 `json:"company"`         
    Department                int64 `json:"department"`         
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

    str := "select u_id, u_loginid, u_passwd, u_name, u_email, u_tel, u_address, u_addressetc, u_joindate, u_careeryear, u_careermonth, u_level, u_score, u_approval, u_status, u_company, u_department, u_date from user_tb "

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

    
    if item.Joindate == "" {
       item.Joindate = "1000-01-01 00:00:00"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into user_tb (u_id, u_loginid, u_passwd, u_name, u_email, u_tel, u_address, u_addressetc, u_joindate, u_careeryear, u_careermonth, u_level, u_score, u_approval, u_status, u_company, u_department, u_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Loginid, item.Passwd, item.Name, item.Email, item.Tel, item.Address, item.Addressetc, item.Joindate, item.Careeryear, item.Careermonth, item.Level, item.Score, item.Approval, item.Status, item.Company, item.Department, item.Date)
    } else {
        query = "insert into user_tb (u_loginid, u_passwd, u_name, u_email, u_tel, u_address, u_addressetc, u_joindate, u_careeryear, u_careermonth, u_level, u_score, u_approval, u_status, u_company, u_department, u_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Loginid, item.Passwd, item.Name, item.Email, item.Tel, item.Address, item.Addressetc, item.Joindate, item.Careeryear, item.Careermonth, item.Level, item.Score, item.Approval, item.Status, item.Company, item.Department, item.Date)
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
    
    
    if item.Joindate == "" {
       item.Joindate = "1000-01-01 00:00:00"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update user_tb set u_loginid = ?, u_passwd = ?, u_name = ?, u_email = ?, u_tel = ?, u_address = ?, u_addressetc = ?, u_joindate = ?, u_careeryear = ?, u_careermonth = ?, u_level = ?, u_score = ?, u_approval = ?, u_status = ?, u_company = ?, u_department = ?, u_date = ? where u_id = ?"
	_, err := p.Exec(query , item.Loginid, item.Passwd, item.Name, item.Email, item.Tel, item.Address, item.Addressetc, item.Joindate, item.Careeryear, item.Careermonth, item.Level, item.Score, item.Approval, item.Status, item.Company, item.Department, item.Date, item.Id)
    
        
    return err
}


func (p *UserManager) UpdateLoginid(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_loginid = ? where u_id = ?"
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

func (p *UserManager) UpdateEmail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_email = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateTel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_tel = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateAddress(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_address = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateAddressetc(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_addressetc = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateJoindate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_joindate = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateCareeryear(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_careeryear = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateCareermonth(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_careermonth = ? where u_id = ?"
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

func (p *UserManager) UpdateScore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_score = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateApproval(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_approval = ? where u_id = ?"
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

func (p *UserManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_company = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) UpdateDepartment(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_department = ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *UserManager) IncreaseCareeryear(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_careeryear = u_careeryear + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) IncreaseCareermonth(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_careermonth = u_careermonth + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) IncreaseScore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_score = u_score + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_company = u_company + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserManager) IncreaseDepartment(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update user_tb set u_department = u_department + ? where u_id = ?"
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
            "approval":     user.GetApproval(p.Approval),

    }
}

func (p *UserManager) ReadRow(rows *sql.Rows) *User {
    var item User
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Email, &item.Tel, &item.Address, &item.Addressetc, &item.Joindate, &item.Careeryear, &item.Careermonth, &item.Level, &item.Score, &item.Approval, &item.Status, &item.Company, &item.Department, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Joindate == "0000-00-00 00:00:00" || item.Joindate == "1000-01-01 00:00:00" {
            item.Joindate = ""
        }
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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
        
    
        err := rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Email, &item.Tel, &item.Address, &item.Addressetc, &item.Joindate, &item.Careeryear, &item.Careermonth, &item.Level, &item.Score, &item.Approval, &item.Status, &item.Company, &item.Department, &item.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        
        
        if item.Joindate == "0000-00-00 00:00:00" || item.Joindate == "1000-01-01 00:00:00" {
            item.Joindate = ""
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


func (p *UserManager) GetByLoginid(loginid string, args ...interface{}) *User {
    if loginid != "" {
        args = append(args, Where{Column:"loginid", Value:loginid, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}

func (p *UserManager) CountByLoginid(loginid string, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if loginid != "" { 
        rets = append(rets, Where{Column:"loginid", Value:loginid, Compare:"="})
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

func (p *UserManager) CountByCompany(company int64, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if company != 0 { 
        rets = append(rets, Where{Column:"company", Value:company, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *UserManager) GetByCompanyName(company int64, name string, args ...interface{}) *User {
    if company != 0 {
        args = append(args, Where{Column:"company", Value:company, Compare:"="})        
    }
    if name != "" {
        args = append(args, Where{Column:"name", Value:name, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}


func (p *UserManager) Sum(args []interface{}) *User {
    if p.Conn == nil && p.Tx == nil {
        var item User
        return &item
    }

    var params []interface{}

    
    query := "select sum(u_score) from user_tb"

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
        
        rows.Scan(&item.Score)        
    }

    return &item        
}
