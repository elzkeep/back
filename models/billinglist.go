package models

import (
    //"zkeep/config"
    
    "zkeep/models/billinglist"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    
    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Billinglist struct {
            
    Id                int64 `json:"id"`         
    Price                int `json:"price"`         
    Status                billinglist.Status `json:"status"`         
    Giro                billinglist.Giro `json:"giro"`         
    Billdate                string `json:"billdate"`         
    Month                string `json:"month"`         
    Endmonth                string `json:"endmonth"`         
    Period                int `json:"period"`         
    Company                int64 `json:"company"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"`         
    Buildingname                string `json:"buildingname"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"`         
    Companyname                string `json:"companyname"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type BillinglistManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Billinglist) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewBillinglistManager(conn interface{}) *BillinglistManager {
    var item BillinglistManager

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

func (p *BillinglistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BillinglistManager) SetIndex(index string) {
    p.Index = index
}

func (p *BillinglistManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *BillinglistManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *BillinglistManager) GetQuery() string {
    ret := ""

    str := "select bi_id, bi_price, bi_status, bi_giro, bi_billdate, bi_month, bi_endmonth, bi_period, bi_company, bi_building, bi_date, bi_buildingname, bi_billingname, bi_billingtel, bi_billingemail, bi_companyname from billinglist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BillinglistManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from billinglist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BillinglistManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate billinglist_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *BillinglistManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from billinglist_vw where bi_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *BillinglistManager) DeleteWhere(args []interface{}) error {
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
                query += " and bi_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bi_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from billinglist_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *BillinglistManager) IncreasePrice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinglist_vw set bi_price = bi_price + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinglistManager) IncreasePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinglist_vw set bi_period = bi_period + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinglistManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinglist_vw set bi_company = bi_company + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinglistManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinglist_vw set bi_building = bi_building + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *BillinglistManager) GetIdentity() int64 {
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

func (p *Billinglist) InitExtra() {
    p.Extra = map[string]interface{}{
            "status":     billinglist.GetStatus(p.Status),
            "giro":     billinglist.GetGiro(p.Giro),

    }
}

func (p *BillinglistManager) ReadRow(rows *sql.Rows) *Billinglist {
    var item Billinglist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Price, &item.Status, &item.Giro, &item.Billdate, &item.Month, &item.Endmonth, &item.Period, &item.Company, &item.Building, &item.Date, &item.Buildingname, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Companyname)
        
        
        
        
        
        
        
        
        if item.Billdate == "0000-00-00" || item.Billdate == "1000-01-01" {
            item.Billdate = ""
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

func (p *BillinglistManager) ReadRows(rows *sql.Rows) []Billinglist {
    var items []Billinglist

    for rows.Next() {
        var item Billinglist
        
    
        err := rows.Scan(&item.Id, &item.Price, &item.Status, &item.Giro, &item.Billdate, &item.Month, &item.Endmonth, &item.Period, &item.Company, &item.Building, &item.Date, &item.Buildingname, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Companyname)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        if item.Billdate == "0000-00-00" || item.Billdate == "1000-01-01" {
            item.Billdate = ""
        }
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        
        
        
        
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *BillinglistManager) Get(id int64) *Billinglist {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and bi_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *BillinglistManager) Count(args []interface{}) int {
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
                query += " and bi_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bi_" + item.Column + " " + item.Compare + " ?"
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

func (p *BillinglistManager) FindAll() []Billinglist {
    return p.Find(nil)
}

func (p *BillinglistManager) Find(args []interface{}) []Billinglist {
    if p.Conn == nil && p.Tx == nil {
        var items []Billinglist
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
                query += " and bi_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bi_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "bi_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "bi_" + orderby
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
            orderby = "bi_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "bi_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Billinglist
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *BillinglistManager) Sum(args []interface{}) *Billinglist {
    if p.Conn == nil && p.Tx == nil {
        var item Billinglist
        return &item
    }

    var params []interface{}

    
    query := "select sum(bi_price) from billinglist_vw"

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
                query += " and bi_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bi_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bi_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "bi_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "bi_" + orderby
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
            orderby = "bi_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "bi_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Billinglist
    
    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        rows.Scan(&item.Price)        
    }

    return &item        
}
