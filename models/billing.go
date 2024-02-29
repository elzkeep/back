package models

import (
    //"zkeep/config"
    
    "zkeep/models/billing"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Billing struct {
            
    Id                int64 `json:"id"`         
    Price                int `json:"price"`         
    Status                billing.Status `json:"status"`         
    Giro                billing.Giro `json:"giro"`         
    Billdate                string `json:"billdate"`         
    Company                int64 `json:"company"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type BillingManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Billing) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewBillingManager(conn interface{}) *BillingManager {
    var item BillingManager

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

func (p *BillingManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BillingManager) SetIndex(index string) {
    p.Index = index
}

func (p *BillingManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *BillingManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *BillingManager) GetQuery() string {
    ret := ""

    str := "select bi_id, bi_price, bi_status, bi_giro, bi_billdate, bi_company, bi_building, bi_date, b_id, b_name, b_conpanyno, b_ceo, b_address, b_addressetc, b_score, b_status, b_company, b_date from billing_tb, building_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and bi_building = b_id "
    

    return ret;
}

func (p *BillingManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from billing_tb, building_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and bi_building = b_id "
    

    return ret;
}

func (p *BillingManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate billing_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *BillingManager) Insert(item *Billing) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
    if item.Billdate == "" {
       item.Billdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into billing_tb (bi_id, bi_price, bi_status, bi_giro, bi_billdate, bi_company, bi_building, bi_date) values (?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Price, item.Status, item.Giro, item.Billdate, item.Company, item.Building, item.Date)
    } else {
        query = "insert into billing_tb (bi_price, bi_status, bi_giro, bi_billdate, bi_company, bi_building, bi_date) values (?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Price, item.Status, item.Giro, item.Billdate, item.Company, item.Building, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *BillingManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from billing_tb where bi_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *BillingManager) DeleteWhere(args []interface{}) error {
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

    query = "delete from billing_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *BillingManager) Update(item *Billing) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Billdate == "" {
       item.Billdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update billing_tb set bi_price = ?, bi_status = ?, bi_giro = ?, bi_billdate = ?, bi_company = ?, bi_building = ?, bi_date = ? where bi_id = ?"
	_, err := p.Exec(query , item.Price, item.Status, item.Giro, item.Billdate, item.Company, item.Building, item.Date, item.Id)
    
        
    return err
}


func (p *BillingManager) UpdatePrice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_price = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_status = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdateGiro(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_giro = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdateBilldate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_billdate = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_company = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdateBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_building = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *BillingManager) IncreasePrice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_price = bi_price + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_company = bi_company + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_building = bi_building + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *BillingManager) GetIdentity() int64 {
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

func (p *Billing) InitExtra() {
    p.Extra = map[string]interface{}{
            "status":     billing.GetStatus(p.Status),
            "giro":     billing.GetGiro(p.Giro),

    }
}

func (p *BillingManager) ReadRow(rows *sql.Rows) *Billing {
    var item Billing
    var err error

    var _building Building
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Price, &item.Status, &item.Giro, &item.Billdate, &item.Company, &item.Building, &item.Date, &_building.Id, &_building.Name, &_building.Conpanyno, &_building.Ceo, &_building.Address, &_building.Addressetc, &_building.Score, &_building.Status, &_building.Company, &_building.Date)
        
        
        
        
        
        
        
        
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
        _building.InitExtra()
        item.AddExtra("building",  _building)

        return &item
    }
}

func (p *BillingManager) ReadRows(rows *sql.Rows) []Billing {
    var items []Billing

    for rows.Next() {
        var item Billing
        var _building Building
            
    
        err := rows.Scan(&item.Id, &item.Price, &item.Status, &item.Giro, &item.Billdate, &item.Company, &item.Building, &item.Date, &_building.Id, &_building.Name, &_building.Conpanyno, &_building.Ceo, &_building.Address, &_building.Addressetc, &_building.Score, &_building.Status, &_building.Company, &_building.Date)
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
        _building.InitExtra()
        item.AddExtra("building",  _building)

        items = append(items, item)
    }


     return items
}

func (p *BillingManager) Get(id int64) *Billing {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and bi_id = ?"

    
    query += " and bi_building = b_id"
    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *BillingManager) Count(args []interface{}) int {
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

func (p *BillingManager) FindAll() []Billing {
    return p.Find(nil)
}

func (p *BillingManager) Find(args []interface{}) []Billing {
    if p.Conn == nil && p.Tx == nil {
        var items []Billing
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
        var items []Billing
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *BillingManager) Sum(args []interface{}) *Billing {
    if p.Conn == nil && p.Tx == nil {
        var item Billing
        return &item
    }

    var params []interface{}

    
    query := "select price from billing_tb"

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

    var item Billing
    
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
