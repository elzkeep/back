package models

import (
    //"zkeep/config"
    
    "zkeep/models/customer"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Customer struct {
            
    Id                int64 `json:"id"`         
    Type                customer.Type `json:"type"`         
    Checkdate                int `json:"checkdate"`         
    Managername                string `json:"managername"`         
    Managertel                string `json:"managertel"`         
    Manageremail                string `json:"manageremail"`         
    Contractstartdate                string `json:"contractstartdate"`         
    Contractenddate                string `json:"contractenddate"`         
    Contractprice                int `json:"contractprice"`         
    Contractday                int `json:"contractday"`         
    Billingdate                int `json:"billingdate"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"`         
    Status                int `json:"status"`         
    Company                int64 `json:"company"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type CustomerManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Customer) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewCustomerManager(conn interface{}) *CustomerManager {
    var item CustomerManager

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

func (p *CustomerManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CustomerManager) SetIndex(index string) {
    p.Index = index
}

func (p *CustomerManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CustomerManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CustomerManager) GetQuery() string {
    ret := ""

    str := "select cu_id, cu_type, cu_checkdate, cu_managername, cu_managertel, cu_manageremail, cu_contractstartdate, cu_contractenddate, cu_contractprice, cu_contractday, cu_billingdate, cu_billingname, cu_billingtel, cu_billingemail, cu_status, cu_company, cu_building, cu_date from customer_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CustomerManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from customer_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CustomerManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate customer_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *CustomerManager) Insert(item *Customer) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into customer_tb (cu_id, cu_type, cu_checkdate, cu_managername, cu_managertel, cu_manageremail, cu_contractstartdate, cu_contractenddate, cu_contractprice, cu_contractday, cu_billingdate, cu_billingname, cu_billingtel, cu_billingemail, cu_status, cu_company, cu_building, cu_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Contractday, item.Billingdate, item.Billingname, item.Billingtel, item.Billingemail, item.Status, item.Company, item.Building, item.Date)
    } else {
        query = "insert into customer_tb (cu_type, cu_checkdate, cu_managername, cu_managertel, cu_manageremail, cu_contractstartdate, cu_contractenddate, cu_contractprice, cu_contractday, cu_billingdate, cu_billingname, cu_billingtel, cu_billingemail, cu_status, cu_company, cu_building, cu_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Contractday, item.Billingdate, item.Billingname, item.Billingtel, item.Billingemail, item.Status, item.Company, item.Building, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *CustomerManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from customer_tb where cu_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *CustomerManager) DeleteWhere(args []interface{}) error {
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
                query += " and cu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cu_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from customer_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *CustomerManager) Update(item *Customer) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update customer_tb set cu_type = ?, cu_checkdate = ?, cu_managername = ?, cu_managertel = ?, cu_manageremail = ?, cu_contractstartdate = ?, cu_contractenddate = ?, cu_contractprice = ?, cu_contractday = ?, cu_billingdate = ?, cu_billingname = ?, cu_billingtel = ?, cu_billingemail = ?, cu_status = ?, cu_company = ?, cu_building = ?, cu_date = ? where cu_id = ?"
	_, err := p.Exec(query , item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Contractday, item.Billingdate, item.Billingname, item.Billingtel, item.Billingemail, item.Status, item.Company, item.Building, item.Date, item.Id)
    
        
    return err
}


func (p *CustomerManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_type = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateCheckdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_checkdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManagername(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_managername = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManagertel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_managertel = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManageremail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_manageremail = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractstartdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractstartdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractenddate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractenddate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractprice = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractday(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractday = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingname(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingname = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingtel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingtel = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingemail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingemail = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_status = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_company = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_building = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *CustomerManager) IncreaseCheckdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_checkdate = cu_checkdate + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractprice = cu_contractprice + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseContractday(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractday = cu_contractday + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseBillingdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingdate = cu_billingdate + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_status = cu_status + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_company = cu_company + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_building = cu_building + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *CustomerManager) GetIdentity() int64 {
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

func (p *Customer) InitExtra() {
    p.Extra = map[string]interface{}{
            "type":     customer.GetType(p.Type),

    }
}

func (p *CustomerManager) ReadRow(rows *sql.Rows) *Customer {
    var item Customer
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Type, &item.Checkdate, &item.Managername, &item.Managertel, &item.Manageremail, &item.Contractstartdate, &item.Contractenddate, &item.Contractprice, &item.Contractday, &item.Billingdate, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Status, &item.Company, &item.Building, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" {
            item.Contractstartdate = ""
        }
        
        if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" {
            item.Contractenddate = ""
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

func (p *CustomerManager) ReadRows(rows *sql.Rows) []Customer {
    var items []Customer

    for rows.Next() {
        var item Customer
        
    
        err := rows.Scan(&item.Id, &item.Type, &item.Checkdate, &item.Managername, &item.Managertel, &item.Manageremail, &item.Contractstartdate, &item.Contractenddate, &item.Contractprice, &item.Contractday, &item.Billingdate, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Status, &item.Company, &item.Building, &item.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" {
            item.Contractstartdate = ""
        }
        if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" {
            item.Contractenddate = ""
        }
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *CustomerManager) Get(id int64) *Customer {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and cu_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CustomerManager) Count(args []interface{}) int {
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
                query += " and cu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cu_" + item.Column + " " + item.Compare + " ?"
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

func (p *CustomerManager) FindAll() []Customer {
    return p.Find(nil)
}

func (p *CustomerManager) Find(args []interface{}) []Customer {
    if p.Conn == nil && p.Tx == nil {
        var items []Customer
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
                query += " and cu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cu_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "cu_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "cu_" + orderby
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
            orderby = "cu_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "cu_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Customer
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




