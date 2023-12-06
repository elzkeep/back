package models

import (
    //"aoi/config"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    
    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Statisticsfaction struct {
            
    Faction                int `json:"faction"`         
    Rank1                int `json:"rank1"`         
    Rank2                int `json:"rank2"`         
    Rank3                int `json:"rank3"`         
    Rank4                int `json:"rank4"`         
    Rank5                int `json:"rank5"`         
    Count                int64 `json:"count"`         
    Avg                int64 `json:"avg"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type StatisticsfactionManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Statisticsfaction) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewStatisticsfactionManager(conn interface{}) *StatisticsfactionManager {
    var item StatisticsfactionManager

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

func (p *StatisticsfactionManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *StatisticsfactionManager) SetIndex(index string) {
    p.Index = index
}

func (p *StatisticsfactionManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *StatisticsfactionManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *StatisticsfactionManager) GetQuery() string {
    ret := ""

    str := "select gu_faction, gu_rank1, gu_rank2, gu_rank3, gu_rank4, gu_rank5, gu_count, gu_avg from statisticsfaction_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *StatisticsfactionManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from statisticsfaction_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *StatisticsfactionManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate statisticsfaction_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *StatisticsfactionManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from statisticsfaction_vw where gu_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *StatisticsfactionManager) DeleteWhere(args []interface{}) error {
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

    query = "delete from statisticsfaction_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *StatisticsfactionManager) IncreaseFaction(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_faction = gu_faction + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseRank1(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_rank1 = gu_rank1 + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseRank2(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_rank2 = gu_rank2 + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseRank3(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_rank3 = gu_rank3 + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseRank4(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_rank4 = gu_rank4 + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseRank5(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_rank5 = gu_rank5 + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseCount(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_count = gu_count + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *StatisticsfactionManager) IncreaseAvg(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update statisticsfaction_vw set gu_avg = gu_avg + ? where gu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *StatisticsfactionManager) GetIdentity() int64 {
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

func (p *Statisticsfaction) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *StatisticsfactionManager) ReadRow(rows *sql.Rows) *Statisticsfaction {
    var item Statisticsfaction
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Faction, &item.Rank1, &item.Rank2, &item.Rank3, &item.Rank4, &item.Rank5, &item.Count, &item.Avg)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *StatisticsfactionManager) ReadRows(rows *sql.Rows) []Statisticsfaction {
    var items []Statisticsfaction

    for rows.Next() {
        var item Statisticsfaction
        
    
        err := rows.Scan(&item.Faction, &item.Rank1, &item.Rank2, &item.Rank3, &item.Rank4, &item.Rank5, &item.Count, &item.Avg)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        
        
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *StatisticsfactionManager) Get(id int64) *Statisticsfaction {
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

func (p *StatisticsfactionManager) Count(args []interface{}) int {
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

func (p *StatisticsfactionManager) FindAll() []Statisticsfaction {
    return p.Find(nil)
}

func (p *StatisticsfactionManager) Find(args []interface{}) []Statisticsfaction {
    if p.Conn == nil && p.Tx == nil {
        var items []Statisticsfaction
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
        case Base:
             item := v

             baseQuery = item.Query
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

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Statisticsfaction
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *StatisticsfactionManager) Sum(args []interface{}) *Statisticsfaction {
    if p.Conn == nil && p.Tx == nil {
        var item Statisticsfaction
        return &item
    }

    var params []interface{}

    
    query := "select count from statisticsfaction_vw"

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
                query += " and gu_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
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

    var item Statisticsfaction
    
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