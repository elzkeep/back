package models

import (
    //"zkeep/config"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    
    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Billinguserlist struct {
            
    Id                int64 `json:"id"`         
    Title                string `json:"title"`         
    Price                int `json:"price"`         
    Depositprice                int `json:"depositprice"`         
    Vat                int `json:"vat"`         
    Status                int `json:"status"`         
    Giro                int `json:"giro"`         
    Billdate                string `json:"billdate"`         
    Month                string `json:"month"`         
    Endmonth                string `json:"endmonth"`         
    Period                int `json:"period"`         
    Billingtype                int `json:"billingtype"`         
    Remark                string `json:"remark"`         
    Company                int64 `json:"company"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"`         
    Buildingname                string `json:"buildingname"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type BillinguserlistManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Billinguserlist) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewBillinguserlistManager(conn interface{}) *BillinguserlistManager {
    var item BillinguserlistManager

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

func (p *BillinguserlistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BillinguserlistManager) SetIndex(index string) {
    p.Index = index
}

func (p *BillinguserlistManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *BillinguserlistManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *BillinguserlistManager) GetQuery() string {
    ret := ""

    str := "select bi_id, bi_title, bi_price, bi_depositprice, bi_vat, bi_status, bi_giro, bi_billdate, bi_month, bi_endmonth, bi_period, bi_billingtype, bi_remark, bi_company, bi_building, bi_date, bi_buildingname, bi_billingname, bi_billingtel, bi_billingemail from billinguserlist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BillinguserlistManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from billinguserlist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BillinguserlistManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate billinguserlist_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *BillinguserlistManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from billinguserlist_vw where bi_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *BillinguserlistManager) DeleteWhere(args []interface{}) error {
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

    query = "delete from billinguserlist_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *BillinguserlistManager) IncreasePrice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_price = bi_price + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseDepositprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_depositprice = bi_depositprice + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseVat(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_vat = bi_vat + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_status = bi_status + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseGiro(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_giro = bi_giro + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreasePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_period = bi_period + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseBillingtype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_billingtype = bi_billingtype + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_company = bi_company + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillinguserlistManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billinguserlist_vw set bi_building = bi_building + ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *BillinguserlistManager) GetIdentity() int64 {
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

func (p *Billinguserlist) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *BillinguserlistManager) ReadRow(rows *sql.Rows) *Billinguserlist {
    var item Billinguserlist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Title, &item.Price, &item.Depositprice, &item.Vat, &item.Status, &item.Giro, &item.Billdate, &item.Month, &item.Endmonth, &item.Period, &item.Billingtype, &item.Remark, &item.Company, &item.Building, &item.Date, &item.Buildingname, &item.Billingname, &item.Billingtel, &item.Billingemail)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *BillinguserlistManager) ReadRows(rows *sql.Rows) []Billinguserlist {
    var items []Billinguserlist

    for rows.Next() {
        var item Billinguserlist
        
    
        err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Depositprice, &item.Vat, &item.Status, &item.Giro, &item.Billdate, &item.Month, &item.Endmonth, &item.Period, &item.Billingtype, &item.Remark, &item.Company, &item.Building, &item.Date, &item.Buildingname, &item.Billingname, &item.Billingtel, &item.Billingemail)
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

func (p *BillinguserlistManager) Get(id int64) *Billinguserlist {
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

func (p *BillinguserlistManager) Count(args []interface{}) int {
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

func (p *BillinguserlistManager) FindAll() []Billinguserlist {
    return p.Find(nil)
}

func (p *BillinguserlistManager) Find(args []interface{}) []Billinguserlist {
    if p.Conn == nil && p.Tx == nil {
        var items []Billinguserlist
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
                query += " and bi_" + item.Column
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
        var items []Billinguserlist
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *BillinguserlistManager) Sum(args []interface{}) *Billinguserlist {
    if p.Conn == nil && p.Tx == nil {
        var item Billinguserlist
        return &item
    }

    var params []interface{}

    
    query := "select sum(bi_price) from billinguserlist_vw"

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

    var item Billinguserlist
    
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
