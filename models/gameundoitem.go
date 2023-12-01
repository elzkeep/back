package models

import (
    //"aoi/config"
    
    "aoi/models/gameundoitem"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Gameundoitem struct {
            
    Id                int64 `json:"id"`         
    Status                gameundoitem.Status `json:"status"`         
    Gameundo                int64 `json:"gameundo"`         
    Game                int64 `json:"game"`         
    User                int64 `json:"user"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type GameundoitemManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Gameundoitem) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewGameundoitemManager(conn interface{}) *GameundoitemManager {
    var item GameundoitemManager

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

func (p *GameundoitemManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *GameundoitemManager) SetIndex(index string) {
    p.Index = index
}

func (p *GameundoitemManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *GameundoitemManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *GameundoitemManager) GetQuery() string {
    ret := ""

    str := "select gi_id, gi_status, gi_gameundo, gi_game, gi_user, gi_date from gameundoitem_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GameundoitemManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from gameundoitem_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GameundoitemManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate gameundoitem_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *GameundoitemManager) Insert(item *Gameundoitem) error {
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
        query = "insert into gameundoitem_tb (gi_id, gi_status, gi_gameundo, gi_game, gi_user, gi_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Status, item.Gameundo, item.Game, item.User, item.Date)
    } else {
        query = "insert into gameundoitem_tb (gi_status, gi_gameundo, gi_game, gi_user, gi_date) values (?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Status, item.Gameundo, item.Game, item.User, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *GameundoitemManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from gameundoitem_tb where gi_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *GameundoitemManager) DeleteWhere(args []interface{}) error {
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
                query += " and gi_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gi_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from gameundoitem_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *GameundoitemManager) Update(item *Gameundoitem) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update gameundoitem_tb set gi_status = ?, gi_gameundo = ?, gi_game = ?, gi_user = ?, gi_date = ? where gi_id = ?"
	_, err := p.Exec(query , item.Status, item.Gameundo, item.Game, item.User, item.Date, item.Id)
    
        
    return err
}


func (p *GameundoitemManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_status = ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameundoitemManager) UpdateGameundo(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_gameundo = ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameundoitemManager) UpdateGame(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_game = ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameundoitemManager) UpdateUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_user = ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *GameundoitemManager) IncreaseGameundo(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_gameundo = gi_gameundo + ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameundoitemManager) IncreaseGame(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_game = gi_game + ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GameundoitemManager) IncreaseUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gameundoitem_tb set gi_user = gi_user + ? where gi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *GameundoitemManager) GetIdentity() int64 {
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

func (p *Gameundoitem) InitExtra() {
    p.Extra = map[string]interface{}{
            "status":     gameundoitem.GetStatus(p.Status),

    }
}

func (p *GameundoitemManager) ReadRow(rows *sql.Rows) *Gameundoitem {
    var item Gameundoitem
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Status, &item.Gameundo, &item.Game, &item.User, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
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

func (p *GameundoitemManager) ReadRows(rows *sql.Rows) []Gameundoitem {
    var items []Gameundoitem

    for rows.Next() {
        var item Gameundoitem
        
    
        err := rows.Scan(&item.Id, &item.Status, &item.Gameundo, &item.Game, &item.User, &item.Date)
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

func (p *GameundoitemManager) Get(id int64) *Gameundoitem {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and gi_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *GameundoitemManager) Count(args []interface{}) int {
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
                query += " and gi_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gi_" + item.Column + " " + item.Compare + " ?"
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

func (p *GameundoitemManager) FindAll() []Gameundoitem {
    return p.Find(nil)
}

func (p *GameundoitemManager) Find(args []interface{}) []Gameundoitem {
    if p.Conn == nil && p.Tx == nil {
        var items []Gameundoitem
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
                query += " and gi_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gi_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "gi_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "gi_" + orderby
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
            orderby = "gi_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "gi_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Gameundoitem
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *GameundoitemManager) FindByGame(game int64, args ...interface{}) []Gameundoitem {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)

    if game != 0 { 
        rets = append(rets, Where{Column:"game", Value:game, Compare:"="})
     }
    
    return p.Find(rets)
}

func (p *GameundoitemManager) CountByGame(game int64, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if game != 0 { 
        rets = append(rets, Where{Column:"game", Value:game, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *GameundoitemManager) DeleteByGame(game int64) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from gameundoitem_tb where gi_game = ?"
    _, err := p.Exec(query, game)

    return err
}



