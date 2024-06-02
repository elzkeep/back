package models

import (
    //"zkeep/config"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    
    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Calendarcompanylist struct {
            
    Id                int64 `json:"id"`         
    Company                int64 `json:"company"`         
    Month                string `json:"month"`         
    Day                string `json:"day"`         
    Status                int `json:"status"`         
    Count                int64 `json:"count"`         
    Checkdate                string `json:"checkdate"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type CalendarcompanylistManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Calendarcompanylist) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewCalendarcompanylistManager(conn interface{}) *CalendarcompanylistManager {
    var item CalendarcompanylistManager

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

func (p *CalendarcompanylistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CalendarcompanylistManager) SetIndex(index string) {
    p.Index = index
}

func (p *CalendarcompanylistManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CalendarcompanylistManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CalendarcompanylistManager) GetQuery() string {
    ret := ""

    str := "select r_id, r_company, r_month, r_day, r_status, r_count, r_checkdate from calendarcompanylist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CalendarcompanylistManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from calendarcompanylist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CalendarcompanylistManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate calendarcompanylist_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *CalendarcompanylistManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from calendarcompanylist_vw where r_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *CalendarcompanylistManager) DeleteWhere(args []interface{}) error {
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
                query += " and r_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from calendarcompanylist_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *CalendarcompanylistManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update calendarcompanylist_vw set r_company = r_company + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CalendarcompanylistManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update calendarcompanylist_vw set r_status = r_status + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CalendarcompanylistManager) IncreaseCount(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update calendarcompanylist_vw set r_count = r_count + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *CalendarcompanylistManager) GetIdentity() int64 {
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

func (p *Calendarcompanylist) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *CalendarcompanylistManager) ReadRow(rows *sql.Rows) *Calendarcompanylist {
    var item Calendarcompanylist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Company, &item.Month, &item.Day, &item.Status, &item.Count, &item.Checkdate)
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Checkdate == "0000-00-00 00:00:00" || item.Checkdate == "1000-01-01 00:00:00" {
            item.Checkdate = ""
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

func (p *CalendarcompanylistManager) ReadRows(rows *sql.Rows) []Calendarcompanylist {
    var items []Calendarcompanylist

    for rows.Next() {
        var item Calendarcompanylist
        
    
        err := rows.Scan(&item.Id, &item.Company, &item.Month, &item.Day, &item.Status, &item.Count, &item.Checkdate)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        if item.Checkdate == "0000-00-00 00:00:00" || item.Checkdate == "1000-01-01 00:00:00" {
            item.Checkdate = ""
        }
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *CalendarcompanylistManager) Get(id int64) *Calendarcompanylist {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and r_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CalendarcompanylistManager) Count(args []interface{}) int {
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
                query += " and r_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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

func (p *CalendarcompanylistManager) FindAll() []Calendarcompanylist {
    return p.Find(nil)
}

func (p *CalendarcompanylistManager) Find(args []interface{}) []Calendarcompanylist {
    if p.Conn == nil && p.Tx == nil {
        var items []Calendarcompanylist
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
                query += " and r_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "r_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "r_" + orderby
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
            orderby = "r_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "r_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Calendarcompanylist
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *CalendarcompanylistManager) Sum(args []interface{}) *Calendarcompanylist {
    if p.Conn == nil && p.Tx == nil {
        var item Calendarcompanylist
        return &item
    }

    var params []interface{}

    
    query := "select sum(r_count) from calendarcompanylist_vw"

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
                query += " and r_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "r_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "r_" + orderby
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
            orderby = "r_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "r_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Calendarcompanylist
    
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
