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

type Webjoin struct {
            
    Id                int64 `json:"id"`         
    Category                int `json:"category"`         
    Name                string `json:"name"`         
    Manager                string `json:"manager"`         
    Tel                string `json:"tel"`         
    Email                string `json:"email"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type WebjoinManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Webjoin) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewWebjoinManager(conn interface{}) *WebjoinManager {
    var item WebjoinManager

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

func (p *WebjoinManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *WebjoinManager) SetIndex(index string) {
    p.Index = index
}

func (p *WebjoinManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *WebjoinManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *WebjoinManager) GetQuery() string {
    ret := ""

    str := "select wj_id, wj_category, wj_name, wj_manager, wj_tel, wj_email, wj_date from webjoin_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *WebjoinManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from webjoin_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *WebjoinManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate webjoin_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *WebjoinManager) Insert(item *Webjoin) error {
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
        query = "insert into webjoin_tb (wj_id, wj_category, wj_name, wj_manager, wj_tel, wj_email, wj_date) values (?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Category, item.Name, item.Manager, item.Tel, item.Email, item.Date)
    } else {
        query = "insert into webjoin_tb (wj_category, wj_name, wj_manager, wj_tel, wj_email, wj_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Category, item.Name, item.Manager, item.Tel, item.Email, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *WebjoinManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from webjoin_tb where wj_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *WebjoinManager) DeleteWhere(args []interface{}) error {
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
                query += " and wj_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and wj_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and wj_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from webjoin_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *WebjoinManager) Update(item *Webjoin) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update webjoin_tb set wj_category = ?, wj_name = ?, wj_manager = ?, wj_tel = ?, wj_email = ?, wj_date = ? where wj_id = ?"
	_, err := p.Exec(query , item.Category, item.Name, item.Manager, item.Tel, item.Email, item.Date, item.Id)
    
        
    return err
}


func (p *WebjoinManager) UpdateCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webjoin_tb set wj_category = ? where wj_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebjoinManager) UpdateName(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webjoin_tb set wj_name = ? where wj_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebjoinManager) UpdateManager(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webjoin_tb set wj_manager = ? where wj_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebjoinManager) UpdateTel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webjoin_tb set wj_tel = ? where wj_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebjoinManager) UpdateEmail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webjoin_tb set wj_email = ? where wj_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *WebjoinManager) IncreaseCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webjoin_tb set wj_category = wj_category + ? where wj_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *WebjoinManager) GetIdentity() int64 {
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

func (p *Webjoin) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *WebjoinManager) ReadRow(rows *sql.Rows) *Webjoin {
    var item Webjoin
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Category, &item.Name, &item.Manager, &item.Tel, &item.Email, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *WebjoinManager) ReadRows(rows *sql.Rows) []Webjoin {
    var items []Webjoin

    for rows.Next() {
        var item Webjoin
        
    
        err := rows.Scan(&item.Id, &item.Category, &item.Name, &item.Manager, &item.Tel, &item.Email, &item.Date)
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

func (p *WebjoinManager) Get(id int64) *Webjoin {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and wj_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *WebjoinManager) Count(args []interface{}) int {
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
                query += " and wj_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and wj_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and wj_" + item.Column + " " + item.Compare + " ?"
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

func (p *WebjoinManager) FindAll() []Webjoin {
    return p.Find(nil)
}

func (p *WebjoinManager) Find(args []interface{}) []Webjoin {
    if p.Conn == nil && p.Tx == nil {
        var items []Webjoin
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
                query += " and wj_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and wj_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and wj_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "wj_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "wj_" + orderby
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
            orderby = "wj_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "wj_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Webjoin
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




