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

type Customercompany struct {
            
    Id                int64 `json:"id"`         
    Company                int64 `json:"company"`         
    Customer                int64 `json:"customer"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type CustomercompanyManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Customercompany) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewCustomercompanyManager(conn interface{}) *CustomercompanyManager {
    var item CustomercompanyManager

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

func (p *CustomercompanyManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CustomercompanyManager) SetIndex(index string) {
    p.Index = index
}

func (p *CustomercompanyManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	log.Println(query)
	log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CustomercompanyManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	log.Println(query)
	log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CustomercompanyManager) GetQuery() string {
    ret := ""

    str := "select cc_id, cc_company, cc_customer, cc_date, c_id, c_name, c_companyno, c_ceo, c_tel, c_email, c_address, c_addressetc, c_type, c_billingname, c_billingtel, c_billingemail, c_bankname, c_bankno, c_businesscondition, c_businessitem, c_giro, c_egirologinid, c_egiropasswd, c_content, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3, c_x4, c_y4, c_x5, c_y5, c_x6, c_y6, c_x7, c_y7, c_x8, c_y8, c_x9, c_y9, c_x10, c_y10, c_x11, c_y11, c_x12, c_y12, c_x13, c_y13, c_status, c_date from customercompany_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and cc_customer = c_id "
    

    return ret;
}

func (p *CustomercompanyManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from customercompany_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and cc_customer = c_id "    
    

    return ret;
}

func (p *CustomercompanyManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate customercompany_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *CustomercompanyManager) Insert(item *Customercompany) error {
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
        query = "insert into customercompany_tb (cc_id, cc_company, cc_customer, cc_date) values (?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Company, item.Customer, item.Date)
    } else {
        query = "insert into customercompany_tb (cc_company, cc_customer, cc_date) values (?, ?, ?)"
        res, err = p.Exec(query , item.Company, item.Customer, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *CustomercompanyManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from customercompany_tb where cc_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *CustomercompanyManager) DeleteWhere(args []interface{}) error {
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
                query += " and cc_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cc_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cc_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from customercompany_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *CustomercompanyManager) Update(item *Customercompany) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update customercompany_tb set cc_company = ?, cc_customer = ?, cc_date = ? where cc_id = ?"
	_, err := p.Exec(query , item.Company, item.Customer, item.Date, item.Id)
    
        
    return err
}


func (p *CustomercompanyManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_tb set cc_company = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) UpdateCustomer(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_tb set cc_customer = ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *CustomercompanyManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_tb set cc_company = cc_company + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) IncreaseCustomer(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_tb set cc_customer = cc_customer + ? where cc_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *CustomercompanyManager) GetIdentity() int64 {
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

func (p *Customercompany) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *CustomercompanyManager) ReadRow(rows *sql.Rows) *Customercompany {
    var item Customercompany
    var err error

    var _company Company
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Company, &item.Customer, &item.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Tel, &_company.Email, &_company.Address, &_company.Addressetc, &_company.Type, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Bankname, &_company.Bankno, &_company.Businesscondition, &_company.Businessitem, &_company.Giro, &_company.Egirologinid, &_company.Egiropasswd, &_company.Content, &_company.X1, &_company.Y1, &_company.X2, &_company.Y2, &_company.X3, &_company.Y3, &_company.X4, &_company.Y4, &_company.X5, &_company.Y5, &_company.X6, &_company.Y6, &_company.X7, &_company.Y7, &_company.X8, &_company.Y8, &_company.X9, &_company.Y9, &_company.X10, &_company.Y10, &_company.X11, &_company.Y11, &_company.X12, &_company.Y12, &_company.X13, &_company.Y13, &_company.Status, &_company.Date)
        
        
        
        
        
        
        
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
        _company.InitExtra()
        item.AddExtra("company",  _company)

        return &item
    }
}

func (p *CustomercompanyManager) ReadRows(rows *sql.Rows) []Customercompany {
    var items []Customercompany

    for rows.Next() {
        var item Customercompany
        var _company Company
            
    
        err := rows.Scan(&item.Id, &item.Company, &item.Customer, &item.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Tel, &_company.Email, &_company.Address, &_company.Addressetc, &_company.Type, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Bankname, &_company.Bankno, &_company.Businesscondition, &_company.Businessitem, &_company.Giro, &_company.Egirologinid, &_company.Egiropasswd, &_company.Content, &_company.X1, &_company.Y1, &_company.X2, &_company.Y2, &_company.X3, &_company.Y3, &_company.X4, &_company.Y4, &_company.X5, &_company.Y5, &_company.X6, &_company.Y6, &_company.X7, &_company.Y7, &_company.X8, &_company.Y8, &_company.X9, &_company.Y9, &_company.X10, &_company.Y10, &_company.X11, &_company.Y11, &_company.X12, &_company.Y12, &_company.X13, &_company.Y13, &_company.Status, &_company.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        _company.InitExtra()
        item.AddExtra("company",  _company)

        items = append(items, item)
    }


     return items
}

func (p *CustomercompanyManager) Get(id int64) *Customercompany {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and cc_id = ?"

    
    query += " and cc_customer = c_id "    
    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CustomercompanyManager) Count(args []interface{}) int {
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
                query += " and cc_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cc_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cc_" + item.Column + " " + item.Compare + " ?"
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

func (p *CustomercompanyManager) FindAll() []Customercompany {
    return p.Find(nil)
}

func (p *CustomercompanyManager) Find(args []interface{}) []Customercompany {
    if p.Conn == nil && p.Tx == nil {
        var items []Customercompany
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
                query += " and cc_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cc_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cc_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "cc_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "cc_" + orderby
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
            orderby = "cc_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "cc_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Customercompany
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *CustomercompanyManager) GetByCompanyCustomer(company int64, customer int64, args ...interface{}) *Customercompany {
    if company != 0 {
        args = append(args, Where{Column:"company", Value:company, Compare:"="})        
    }
    if customer != 0 {
        args = append(args, Where{Column:"customer", Value:customer, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}



