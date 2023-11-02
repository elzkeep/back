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

type Gameuser struct {
            
    Id                int64 `json:"id"`         
    Order                int `json:"order"`         
    User                int64 `json:"user"`         
    Game                int64 `json:"game"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type GameuserManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Gameuser) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewGameuserManager(conn interface{}) *GameuserManager {
    var item GameuserManager

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

func (p *GameuserManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *GameuserManager) SetIndex(index string) {
    p.Index = index
}

func (p *GameuserManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *GameuserManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *GameuserManager) GetQuery() string {
    ret := ""

    str := "select gu_id, gu_order, gu_user, gu_game, gu_date from gameuser_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GameuserManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from gameuser_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GameuserManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate gameuser_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *GameuserManager) Insert(item *Gameuser) error {
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
        query = "insert into gameuser_tb (gu_id, gu_order, gu_user, gu_game, gu_date) values (?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Order, item.User, item.Game, item.Date)
    } else {
        query = "insert into gameuser_tb (gu_order, gu_user, gu_game, gu_date) values (?, ?, ?, ?)"
        res, err = p.Exec(query , item.Order, item.User, item.Game, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}
func (p *GameuserManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from gameuser_tb where gu_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}
func (p *GameuserManager) Update(item *Gameuser) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update gameuser_tb set gu_order = ?, gu_user = ?, gu_game = ?, gu_date = ? where gu_id = ?"
	_, err := p.Exec(query , item.Order, item.User, item.Game, item.Date, item.Id)
    
        
    return err
}


func (p *GameuserManager) UpdateOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameuser_tb set gu_order = ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameuserManager) UpdateUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameuser_tb set gu_user = ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameuserManager) UpdateGame(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameuser_tb set gu_game = ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *GameuserManager) GetIdentity() int64 {
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

func (p *Gameuser) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *GameuserManager) ReadRow(rows *sql.Rows) *Gameuser {
    var item Gameuser
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Order, &item.User, &item.Game, &item.Date)
        
        
        
        
        
        
        
        
        
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

func (p *GameuserManager) ReadRows(rows *sql.Rows) []Gameuser {
    var items []Gameuser

    for rows.Next() {
        var item Gameuser
        
    
        err := rows.Scan(&item.Id, &item.Order, &item.User, &item.Game, &item.Date)
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

func (p *GameuserManager) Get(id int64) *Gameuser {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and gu_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *GameuserManager) Count(args []interface{}) int {
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
                query += " and gu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gu_" + item.Column + " " + item.Compare + " ?"
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

func (p *GameuserManager) Find(args []interface{}) []Gameuser {
    if p.Conn == nil && p.Tx == nil {
        var items []Gameuser
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
                query += " and gu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gu_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "gu_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "gu_" + orderby
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
            orderby = "gu_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "gu_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Gameuser
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *GameuserManager) FindByGame(game int64, args ...interface{}) []Gameuser {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)

    if game != 0 { 
        rets = append(rets, Where{Column:"game", Value:game, Compare:"="})
     }
    
    return p.Find(rets)
}

func (p *GameuserManager) CountByGame(game int64, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if game != 0 { 
        rets = append(rets, Where{Column:"game", Value:game, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *GameuserManager) CountByGameUser(game int64, user int64, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if game != 0 { 
        rets = append(rets, Where{Column:"game", Value:game, Compare:"="})
     }
    if user != 0 { 
        rets = append(rets, Where{Column:"user", Value:user, Compare:"="})
     }
    
    return p.Count(rets)
}



