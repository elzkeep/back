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

type Billinghistory struct {
            
    Id                int64 `json:"id"`         
    Price                int `json:"price"`         
    Company                int64 `json:"company"`         
    Building                int64 `json:"building"`         
    Billing                int64 `json:"billing"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type BillinghistoryManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Billinghistory) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewBillinghistoryManager(conn interface{}) *BillinghistoryManager {
    var item BillinghistoryManager

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

func (p *BillinghistoryManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BillinghistoryManager) SetIndex(index string) {
    p.Index = index
}

func (p *BillinghistoryManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *BillinghistoryManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *BillinghistoryManager) GetQuery() string {
    ret := ""

    str := "select bh_id, bh_price, bh_company, bh_building, bh_billing, bh_date from billinghistory_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BillinghistoryManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from billinghistory_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BillinghistoryManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate billinghistory_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *BillinghistoryManager) Insert(item *Billinghistory) error {
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
        query = "insert into billinghistory_tb (bh_id, bh_price, bh_company, bh_building, bh_billing, bh_date) values (?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Price, item.Company, item.Building, item.Billing, item.Date)
    } else {
        query = "insert into billinghistory_tb (bh_price, bh_company, bh_building, bh_billing, bh_date) values (?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Price, item.Company, item.Building, item.Billing, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *BillinghistoryManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from billinghistory_tb where bh_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *BillinghistoryManager) DeleteWhere(args []interface{}) error {
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
                query += " and bh_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bh_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bh_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from billinghistory_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *BillinghistoryManager) Update(item *Billinghistory) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update billinghistory_tb set bh_price = ?, bh_company = ?, bh_building = ?, bh_billing = ?, bh_date = ? where bh_id = ?"
	_, err := p.Exec(query , item.Price, item.Company, item.Building, item.Billing, item.Date, item.Id)
    
        
    return err
}


func (p *BillinghistoryManager) UpdatePrice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_price = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinghistoryManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_company = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinghistoryManager) UpdateBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_building = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinghistoryManager) UpdateBilling(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_billing = ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *BillinghistoryManager) IncreasePrice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_price = bh_price + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinghistoryManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_company = bh_company + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinghistoryManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_building = bh_building + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinghistoryManager) IncreaseBilling(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinghistory_tb set bh_billing = bh_billing + ? where bh_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *BillinghistoryManager) GetIdentity() int64 {
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

func (p *Billinghistory) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *BillinghistoryManager) ReadRow(rows *sql.Rows) *Billinghistory {
    var item Billinghistory
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Price, &item.Company, &item.Building, &item.Billing, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
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

func (p *BillinghistoryManager) ReadRows(rows *sql.Rows) []Billinghistory {
    var items []Billinghistory

    for rows.Next() {
        var item Billinghistory
        
    
        err := rows.Scan(&item.Id, &item.Price, &item.Company, &item.Building, &item.Billing, &item.Date)
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

func (p *BillinghistoryManager) Get(id int64) *Billinghistory {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and bh_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *BillinghistoryManager) Count(args []interface{}) int {
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
                query += " and bh_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bh_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bh_" + item.Column + " " + item.Compare + " ?"
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

func (p *BillinghistoryManager) FindAll() []Billinghistory {
    return p.Find(nil)
}

func (p *BillinghistoryManager) Find(args []interface{}) []Billinghistory {
    if p.Conn == nil && p.Tx == nil {
        var items []Billinghistory
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
                query += " and bh_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bh_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bh_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "bh_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "bh_" + orderby
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
            orderby = "bh_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "bh_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Billinghistory
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *BillinghistoryManager) DeleteByBilling(billing int64) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from billinghistory_tb where bh_billing = ?"
    _, err := p.Exec(query, billing)

    return err
}


func (p *BillinghistoryManager) Sum(args []interface{}) *Billinghistory {
    if p.Conn == nil && p.Tx == nil {
        var item Billinghistory
        return &item
    }

    var params []interface{}

    
    query := "select sum(bh_price) from billinghistory_tb"

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
                query += " and bh_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and bh_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and bh_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "bh_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "bh_" + orderby
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
            orderby = "bh_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "bh_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Billinghistory
    
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
