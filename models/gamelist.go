package models

import (
    //"aoi/config"
    
    "aoi/models/gamelist"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    
    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Gamelist struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Count                int `json:"count"`         
    Join                int `json:"join"`         
    Map                int64 `json:"map"`         
    Illusionists                gamelist.Illusionists `json:"illusionists"`         
    Type                int `json:"type"`         
    Status                gamelist.Status `json:"status"`         
    Enddate                string `json:"enddate"`         
    User                int64 `json:"user"`         
    Date                string `json:"date"`         
    Gameuser                int64 `json:"gameuser"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type GamelistManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Gamelist) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewGamelistManager(conn interface{}) *GamelistManager {
    var item GamelistManager

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

func (p *GamelistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *GamelistManager) SetIndex(index string) {
    p.Index = index
}

func (p *GamelistManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *GamelistManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *GamelistManager) GetQuery() string {
    ret := ""

    str := "select g_id, g_name, g_count, g_join, g_map, g_illusionists, g_type, g_status, g_enddate, g_user, g_date, g_gameuser from gamelist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GamelistManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from gamelist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *GamelistManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate gamelist_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *GamelistManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from gamelist_vw where g_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *GamelistManager) DeleteWhere(args []interface{}) error {
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
                query += " and g_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and g_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and g_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from gamelist_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *GamelistManager) IncreaseCount(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamelist_vw set g_count = g_count + ? where g_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamelistManager) IncreaseJoin(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamelist_vw set g_join = g_join + ? where g_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamelistManager) IncreaseMap(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamelist_vw set g_map = g_map + ? where g_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamelistManager) IncreaseType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamelist_vw set g_type = g_type + ? where g_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamelistManager) IncreaseUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamelist_vw set g_user = g_user + ? where g_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *GamelistManager) IncreaseGameuser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update gamelist_vw set g_gameuser = g_gameuser + ? where g_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *GamelistManager) GetIdentity() int64 {
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

func (p *Gamelist) InitExtra() {
    p.Extra = map[string]interface{}{
            "status":     gamelist.GetStatus(p.Status),
            "illusionists":     gamelist.GetIllusionists(p.Illusionists),

    }
}

func (p *GamelistManager) ReadRow(rows *sql.Rows) *Gamelist {
    var item Gamelist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Count, &item.Join, &item.Map, &item.Illusionists, &item.Type, &item.Status, &item.Enddate, &item.User, &item.Date, &item.Gameuser)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Enddate == "0000-00-00 00:00:00" || item.Enddate == "1000-01-01 00:00:00" {
            item.Enddate = ""
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

func (p *GamelistManager) ReadRows(rows *sql.Rows) []Gamelist {
    var items []Gamelist

    for rows.Next() {
        var item Gamelist
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Count, &item.Join, &item.Map, &item.Illusionists, &item.Type, &item.Status, &item.Enddate, &item.User, &item.Date, &item.Gameuser)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        
        
        if item.Enddate == "0000-00-00 00:00:00" || item.Enddate == "1000-01-01 00:00:00" {
            item.Enddate = ""
        }
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *GamelistManager) Get(id int64) *Gamelist {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and g_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *GamelistManager) Count(args []interface{}) int {
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
                query += " and g_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and g_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and g_" + item.Column + " " + item.Compare + " ?"
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

func (p *GamelistManager) FindAll() []Gamelist {
    return p.Find(nil)
}

func (p *GamelistManager) Find(args []interface{}) []Gamelist {
    if p.Conn == nil && p.Tx == nil {
        var items []Gamelist
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
                query += " and g_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and g_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and g_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "g_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "g_" + orderby
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
            orderby = "g_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "g_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Gamelist
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *GamelistManager) Sum(args []interface{}) *Gamelist {
    if p.Conn == nil && p.Tx == nil {
        var item Gamelist
        return &item
    }

    var params []interface{}

    
    query := "select count from gamelist_vw"

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
                query += " and g_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and g_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and g_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "g_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "g_" + orderby
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
            orderby = "g_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "g_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Gamelist
    
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
