package models

import (
    //"aoi/config"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Gametile struct {
            
    Id                int64 `json:"id"`         
    Type                int `json:"type"`         
    Number                int `json:"number"`         
    Order                int `json:"order"`         
    Game                int64 `json:"game"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type GametileManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Gametile) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewGametileManager(conn interface{}) *GametileManager {
    var item GametileManager

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

func (p *GametileManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *GametileManager) SetIndex(index string) {
    p.Index = index
}

func (p *GametileManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *GametileManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *GametileManager) GetQuery() string {
    ret := ""

    str := "select gt_id, gt_type, gt_number, gt_order, gt_game, gt_date from gametile_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GametileManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from gametile_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GametileManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate gametile_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *GametileManager) Insert(item *Gametile) error {
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
        query = "insert into gametile_tb (gt_id, gt_type, gt_number, gt_order, gt_game, gt_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Type, item.Number, item.Order, item.Game, item.Date)
    } else {
        query = "insert into gametile_tb (gt_type, gt_number, gt_order, gt_game, gt_date) values (?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Type, item.Number, item.Order, item.Game, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}
func (p *GametileManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from gametile_tb where gt_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}
func (p *GametileManager) Update(item *Gametile) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update gametile_tb set gt_type = ?, gt_number = ?, gt_order = ?, gt_game = ?, gt_date = ? where gt_id = ?"
	_, err := p.Exec(query , item.Type, item.Number, item.Order, item.Game, item.Date, item.Id)
    
        
    return err
}


func (p *GametileManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gametile_tb set gt_type = ? where gt_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GametileManager) UpdateNumber(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gametile_tb set gt_number = ? where gt_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GametileManager) UpdateOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gametile_tb set gt_order = ? where gt_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GametileManager) UpdateGame(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gametile_tb set gt_game = ? where gt_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *GametileManager) GetIdentity() int64 {
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

func (p *Gametile) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *GametileManager) ReadRow(rows *sql.Rows) *Gametile {
    var item Gametile
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Type, &item.Number, &item.Order, &item.Game, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
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

func (p *GametileManager) ReadRows(rows *sql.Rows) []Gametile {
    var items []Gametile

    for rows.Next() {
        var item Gametile
        
    
        err := rows.Scan(&item.Id, &item.Type, &item.Number, &item.Order, &item.Game, &item.Date)
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

func (p *GametileManager) Get(id int64) *Gametile {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and gt_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *GametileManager) Count(args []interface{}) int {
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
                query += " and gt_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gt_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gt_" + item.Column + " " + item.Compare + " ?"
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

func (p *GametileManager) Find(args []interface{}) []Gametile {
    if p.Conn == nil && p.Tx == nil {
        var items []Gametile
        return items
    }

    var params []interface{}
    query := p.GetQuery()

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
                query += " and gt_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gt_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gt_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "gt_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "gt_" + orderby
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
            orderby = "gt_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "gt_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Gametile
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




