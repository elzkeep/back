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

type Webfaq struct {
            
    Id                int64 `json:"id"`         
    Category                int `json:"category"`         
    Title                string `json:"title"`         
    Content                string `json:"content"`         
    Order                int `json:"order"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type WebfaqManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Webfaq) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewWebfaqManager(conn interface{}) *WebfaqManager {
    var item WebfaqManager

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

func (p *WebfaqManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *WebfaqManager) SetIndex(index string) {
    p.Index = index
}

func (p *WebfaqManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *WebfaqManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *WebfaqManager) GetQuery() string {
    ret := ""

    str := "select wf_id, wf_category, wf_title, wf_content, wf_order, wf_date from webfaq_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *WebfaqManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from webfaq_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *WebfaqManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate webfaq_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *WebfaqManager) Insert(item *Webfaq) error {
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
        query = "insert into webfaq_tb (wf_id, wf_category, wf_title, wf_content, wf_order, wf_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Category, item.Title, item.Content, item.Order, item.Date)
    } else {
        query = "insert into webfaq_tb (wf_category, wf_title, wf_content, wf_order, wf_date) values (?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Category, item.Title, item.Content, item.Order, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *WebfaqManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from webfaq_tb where wf_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *WebfaqManager) DeleteWhere(args []interface{}) error {
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
                query += " and wf_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and wf_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and wf_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from webfaq_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *WebfaqManager) Update(item *Webfaq) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update webfaq_tb set wf_category = ?, wf_title = ?, wf_content = ?, wf_order = ?, wf_date = ? where wf_id = ?"
	_, err := p.Exec(query , item.Category, item.Title, item.Content, item.Order, item.Date, item.Id)
    
        
    return err
}


func (p *WebfaqManager) UpdateCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webfaq_tb set wf_category = ? where wf_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebfaqManager) UpdateTitle(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webfaq_tb set wf_title = ? where wf_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebfaqManager) UpdateContent(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webfaq_tb set wf_content = ? where wf_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebfaqManager) UpdateOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webfaq_tb set wf_order = ? where wf_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *WebfaqManager) IncreaseCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webfaq_tb set wf_category = wf_category + ? where wf_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *WebfaqManager) IncreaseOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update webfaq_tb set wf_order = wf_order + ? where wf_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *WebfaqManager) GetIdentity() int64 {
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

func (p *Webfaq) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *WebfaqManager) ReadRow(rows *sql.Rows) *Webfaq {
    var item Webfaq
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Category, &item.Title, &item.Content, &item.Order, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
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

func (p *WebfaqManager) ReadRows(rows *sql.Rows) []Webfaq {
    var items []Webfaq

    for rows.Next() {
        var item Webfaq
        
    
        err := rows.Scan(&item.Id, &item.Category, &item.Title, &item.Content, &item.Order, &item.Date)
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

func (p *WebfaqManager) Get(id int64) *Webfaq {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and wf_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *WebfaqManager) Count(args []interface{}) int {
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
                query += " and wf_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and wf_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and wf_" + item.Column + " " + item.Compare + " ?"
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

func (p *WebfaqManager) FindAll() []Webfaq {
    return p.Find(nil)
}

func (p *WebfaqManager) Find(args []interface{}) []Webfaq {
    if p.Conn == nil && p.Tx == nil {
        var items []Webfaq
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

            if strings.Contains(item.Column, "_") {
                query += " and " + item.Column
            } else {
                query += " and wf_" + item.Column
            }
            
            if item.Compare == "in" {
                query += " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " " + item.Compare + " ?"

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
            orderby = "wf_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "wf_" + orderby
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
            orderby = "wf_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "wf_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Webfaq
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




