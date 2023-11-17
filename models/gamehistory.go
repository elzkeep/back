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

type Gamehistory struct {
            
    Id                int64 `json:"id"`         
    Round                int `json:"round"`         
    Command                string `json:"command"`         
    User                int64 `json:"user"`         
    Game                int64 `json:"game"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type GamehistoryManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Gamehistory) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewGamehistoryManager(conn interface{}) *GamehistoryManager {
    var item GamehistoryManager

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

func (p *GamehistoryManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *GamehistoryManager) SetIndex(index string) {
    p.Index = index
}

func (p *GamehistoryManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *GamehistoryManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *GamehistoryManager) GetQuery() string {
    ret := ""

    str := "select gh_id, gh_round, gh_command, gh_user, gh_game, gh_date from gamehistory_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GamehistoryManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from gamehistory_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GamehistoryManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate gamehistory_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *GamehistoryManager) Insert(item *Gamehistory) error {
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
        query = "insert into gamehistory_tb (gh_id, gh_round, gh_command, gh_user, gh_game, gh_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Round, item.Command, item.User, item.Game, item.Date)
    } else {
        query = "insert into gamehistory_tb (gh_round, gh_command, gh_user, gh_game, gh_date) values (?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Round, item.Command, item.User, item.Game, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}
func (p *GamehistoryManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from gamehistory_tb where gh_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}
func (p *GamehistoryManager) Update(item *Gamehistory) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update gamehistory_tb set gh_round = ?, gh_command = ?, gh_user = ?, gh_game = ?, gh_date = ? where gh_id = ?"
	_, err := p.Exec(query , item.Round, item.Command, item.User, item.Game, item.Date, item.Id)
    
        
    return err
}


func (p *GamehistoryManager) UpdateRound(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamehistory_tb set gh_round = ? where gh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamehistoryManager) UpdateCommand(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamehistory_tb set gh_command = ? where gh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamehistoryManager) UpdateUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamehistory_tb set gh_user = ? where gh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamehistoryManager) UpdateGame(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamehistory_tb set gh_game = ? where gh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *GamehistoryManager) GetIdentity() int64 {
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

func (p *Gamehistory) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *GamehistoryManager) ReadRow(rows *sql.Rows) *Gamehistory {
    var item Gamehistory
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Round, &item.Command, &item.User, &item.Game, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
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

func (p *GamehistoryManager) ReadRows(rows *sql.Rows) []Gamehistory {
    var items []Gamehistory

    for rows.Next() {
        var item Gamehistory
        
    
        err := rows.Scan(&item.Id, &item.Round, &item.Command, &item.User, &item.Game, &item.Date)
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

func (p *GamehistoryManager) Get(id int64) *Gamehistory {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and gh_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *GamehistoryManager) Count(args []interface{}) int {
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
                query += " and gh_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gh_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gh_" + item.Column + " " + item.Compare + " ?"
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

func (p *GamehistoryManager) Find(args []interface{}) []Gamehistory {
    if p.Conn == nil && p.Tx == nil {
        var items []Gamehistory
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
                query += " and gh_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and gh_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and gh_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "gh_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "gh_" + orderby
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
            orderby = "gh_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "gh_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Gamehistory
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




