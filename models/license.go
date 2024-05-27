package models

import (
    //"zkeep/config"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type License struct {
            
    Id                int64 `json:"id"`         
    User                int64 `json:"user"`         
    Number                string `json:"number"`         
    Takingdate                string `json:"takingdate"`         
    Educationdate                string `json:"educationdate"`         
    Educationinstitution                string `json:"educationinstitution"`         
    Specialeducationdate                string `json:"specialeducationdate"`         
    Specialeducationinstitution                string `json:"specialeducationinstitution"`         
    Licensecategory                int64 `json:"licensecategory"`         
    Licenselevel                int64 `json:"licenselevel"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type LicenseManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *License) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewLicenseManager(conn interface{}) *LicenseManager {
    var item LicenseManager

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

func (p *LicenseManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *LicenseManager) SetIndex(index string) {
    p.Index = index
}

func (p *LicenseManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    log.Println(query)
    log.Println(params)    
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *LicenseManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    log.Println(query)
    log.Println(params)    
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *LicenseManager) GetQuery() string {
    ret := ""

    str := "select l_id, l_user, l_number, l_takingdate, l_educationdate, l_educationinstitution, l_specialeducationdate, l_specialeducationinstitution, l_licensecategory, l_licenselevel, l_date, lc_id, lc_name, lc_order, lc_date, ll_id, ll_name, ll_order, ll_date from license_tb, licensecategory_tb, licenselevel_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and l_licensecategory = lc_id "
    
    ret += "and l_licenselevel = ll_id "
    

    return ret;
}

func (p *LicenseManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from license_tb, licensecategory_tb, licenselevel_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and l_licensecategory = lc_id "    
    
    ret += "and l_licenselevel = ll_id "    
    

    return ret;
}

func (p *LicenseManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate license_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *LicenseManager) Insert(item *License) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
    if item.Takingdate == "" {
       item.Takingdate = "1000-01-01"
    }
    if item.Educationdate == "" {
       item.Educationdate = "1000-01-01"
    }
    if item.Specialeducationdate == "" {
       item.Specialeducationdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into license_tb (l_id, l_user, l_number, l_takingdate, l_educationdate, l_educationinstitution, l_specialeducationdate, l_specialeducationinstitution, l_licensecategory, l_licenselevel, l_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.User, item.Number, item.Takingdate, item.Educationdate, item.Educationinstitution, item.Specialeducationdate, item.Specialeducationinstitution, item.Licensecategory, item.Licenselevel, item.Date)
    } else {
        query = "insert into license_tb (l_user, l_number, l_takingdate, l_educationdate, l_educationinstitution, l_specialeducationdate, l_specialeducationinstitution, l_licensecategory, l_licenselevel, l_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.User, item.Number, item.Takingdate, item.Educationdate, item.Educationinstitution, item.Specialeducationdate, item.Specialeducationinstitution, item.Licensecategory, item.Licenselevel, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *LicenseManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from license_tb where l_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *LicenseManager) DeleteWhere(args []interface{}) error {
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
                query += " and l_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and l_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and l_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from license_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *LicenseManager) Update(item *License) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Takingdate == "" {
       item.Takingdate = "1000-01-01"
    }
    if item.Educationdate == "" {
       item.Educationdate = "1000-01-01"
    }
    if item.Specialeducationdate == "" {
       item.Specialeducationdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update license_tb set l_user = ?, l_number = ?, l_takingdate = ?, l_educationdate = ?, l_educationinstitution = ?, l_specialeducationdate = ?, l_specialeducationinstitution = ?, l_licensecategory = ?, l_licenselevel = ?, l_date = ? where l_id = ?"
	_, err := p.Exec(query , item.User, item.Number, item.Takingdate, item.Educationdate, item.Educationinstitution, item.Specialeducationdate, item.Specialeducationinstitution, item.Licensecategory, item.Licenselevel, item.Date, item.Id)
    
        
    return err
}


func (p *LicenseManager) UpdateUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_user = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateNumber(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_number = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateTakingdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_takingdate = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateEducationdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_educationdate = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateEducationinstitution(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_educationinstitution = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateSpecialeducationdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_specialeducationdate = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateSpecialeducationinstitution(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_specialeducationinstitution = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateLicensecategory(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_licensecategory = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) UpdateLicenselevel(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_licenselevel = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *LicenseManager) IncreaseUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_user = l_user + ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) IncreaseLicensecategory(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_licensecategory = l_licensecategory + ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *LicenseManager) IncreaseLicenselevel(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update license_tb set l_licenselevel = l_licenselevel + ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *LicenseManager) GetIdentity() int64 {
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

func (p *License) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *LicenseManager) ReadRow(rows *sql.Rows) *License {
    var item License
    var err error

    var _licensecategory Licensecategory
    var _licenselevel Licenselevel
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.User, &item.Number, &item.Takingdate, &item.Educationdate, &item.Educationinstitution, &item.Specialeducationdate, &item.Specialeducationinstitution, &item.Licensecategory, &item.Licenselevel, &item.Date, &_licensecategory.Id, &_licensecategory.Name, &_licensecategory.Order, &_licensecategory.Date, &_licenselevel.Id, &_licenselevel.Name, &_licenselevel.Order, &_licenselevel.Date)
        
        
        
        
        
        
        if item.Takingdate == "0000-00-00" || item.Takingdate == "1000-01-01" {
            item.Takingdate = ""
        }
        
        if item.Educationdate == "0000-00-00" || item.Educationdate == "1000-01-01" {
            item.Educationdate = ""
        }
        
        
        
        if item.Specialeducationdate == "0000-00-00" || item.Specialeducationdate == "1000-01-01" {
            item.Specialeducationdate = ""
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
        _licensecategory.InitExtra()
        item.AddExtra("licensecategory",  _licensecategory)
_licenselevel.InitExtra()
        item.AddExtra("licenselevel",  _licenselevel)

        return &item
    }
}

func (p *LicenseManager) ReadRows(rows *sql.Rows) []License {
    var items []License

    for rows.Next() {
        var item License
        var _licensecategory Licensecategory
            var _licenselevel Licenselevel
            
    
        err := rows.Scan(&item.Id, &item.User, &item.Number, &item.Takingdate, &item.Educationdate, &item.Educationinstitution, &item.Specialeducationdate, &item.Specialeducationinstitution, &item.Licensecategory, &item.Licenselevel, &item.Date, &_licensecategory.Id, &_licensecategory.Name, &_licensecategory.Order, &_licensecategory.Date, &_licenselevel.Id, &_licenselevel.Name, &_licenselevel.Order, &_licenselevel.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        if item.Takingdate == "0000-00-00" || item.Takingdate == "1000-01-01" {
            item.Takingdate = ""
        }
        if item.Educationdate == "0000-00-00" || item.Educationdate == "1000-01-01" {
            item.Educationdate = ""
        }
        
        if item.Specialeducationdate == "0000-00-00" || item.Specialeducationdate == "1000-01-01" {
            item.Specialeducationdate = ""
        }
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        _licensecategory.InitExtra()
        item.AddExtra("licensecategory",  _licensecategory)
_licenselevel.InitExtra()
        item.AddExtra("licenselevel",  _licenselevel)

        items = append(items, item)
    }


     return items
}

func (p *LicenseManager) Get(id int64) *License {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and l_id = ?"

    
    query += " and l_licensecategory = lc_id "    
    
    query += " and l_licenselevel = ll_id "    
    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *LicenseManager) Count(args []interface{}) int {
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
                query += " and l_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and l_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and l_" + item.Column + " " + item.Compare + " ?"
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

func (p *LicenseManager) FindAll() []License {
    return p.Find(nil)
}

func (p *LicenseManager) Find(args []interface{}) []License {
    if p.Conn == nil && p.Tx == nil {
        var items []License
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
                query += " and l_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and l_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and l_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "l_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "l_" + orderby
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
            orderby = "l_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "l_" + orderby
            }
        }
        query += " order by " + orderby
    }

    log.Println(baseQuery + query)
    log.Println(params)
    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []License
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *LicenseManager) GetByUserLicensecategory(user int64, licensecategory int64, args ...interface{}) *License {
    if user != 0 {
        args = append(args, Where{Column:"user", Value:user, Compare:"="})        
    }
    if licensecategory != 0 {
        args = append(args, Where{Column:"licensecategory", Value:licensecategory, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}

func (p *LicenseManager) DeleteByUser(user int64) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from license_tb where l_user = ?"
    _, err := p.Exec(query, user)

    return err
}

func (p *LicenseManager) FindByUser(user int64, args ...interface{}) []License {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)

    if user != 0 { 
        rets = append(rets, Where{Column:"user", Value:user, Compare:"="})
     }
    
    return p.Find(rets)
}



