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
    Month                string `json:"month"`         
    Endmonth                string `json:"endmonth"`         
    Period                int `json:"period"`         
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
    log.Println(query)
    log.Println(params)    
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *BillingManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    log.Println(query)
    log.Println(params)    
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *BillingManager) GetQuery() string {
    ret := ""

    str := "select bi_id, bi_price, bi_status, bi_giro, bi_billdate, bi_month, bi_endmonth, bi_period, bi_company, bi_building, bi_date, b_id, b_name, b_companyno, b_ceo, b_zip, b_address, b_addressetc, b_contractvolumn, b_receivevolumn, b_generatevolumn, b_sunlightvolumn, b_volttype, b_weight, b_totalweight, b_checkcount, b_receivevolt, b_generatevolt, b_periodic, b_businesscondition, b_businessitem, b_usage, b_district, b_score, b_status, b_company, b_date, c_id, c_name, c_companyno, c_ceo, c_tel, c_email, c_address, c_addressetc, c_type, c_billingname, c_billingtel, c_billingemail, c_bankname, c_bankno, c_businesscondition, c_businessitem, c_giro, c_content, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3, c_x4, c_y4, c_x5, c_y5, c_x6, c_y6, c_x7, c_y7, c_x8, c_y8, c_x9, c_y9, c_x10, c_y10, c_x11, c_y11, c_x12, c_y12, c_x13, c_y13, c_status, c_date from billing_tb, building_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and bi_building = b_id "
    
    ret += "and b_company = c_id "
    

    return ret;
}

func (p *BillingManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from billing_tb, building_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and bi_building = b_id "    
    
    ret += "and b_company = c_id "    
    

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
        query = "insert into billing_tb (bi_id, bi_price, bi_status, bi_giro, bi_billdate, bi_month, bi_endmonth, bi_period, bi_company, bi_building, bi_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Price, item.Status, item.Giro, item.Billdate, item.Month, item.Endmonth, item.Period, item.Company, item.Building, item.Date)
    } else {
        query = "insert into billing_tb (bi_price, bi_status, bi_giro, bi_billdate, bi_month, bi_endmonth, bi_period, bi_company, bi_building, bi_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Price, item.Status, item.Giro, item.Billdate, item.Month, item.Endmonth, item.Period, item.Company, item.Building, item.Date)
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

	query := "update billing_tb set bi_price = ?, bi_status = ?, bi_giro = ?, bi_billdate = ?, bi_month = ?, bi_endmonth = ?, bi_period = ?, bi_company = ?, bi_building = ?, bi_date = ? where bi_id = ?"
	_, err := p.Exec(query , item.Price, item.Status, item.Giro, item.Billdate, item.Month, item.Endmonth, item.Period, item.Company, item.Building, item.Date, item.Id)
    
        
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

func (p *BillingManager) UpdateMonth(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_month = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdateEndmonth(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_endmonth = ? where bi_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BillingManager) UpdatePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_period = ? where bi_id = ?"
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

func (p *BillingManager) IncreasePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update billing_tb set bi_period = bi_period + ? where bi_id = ?"
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
    var _company Company
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Price, &item.Status, &item.Giro, &item.Billdate, &item.Month, &item.Endmonth, &item.Period, &item.Company, &item.Building, &item.Date, &_building.Id, &_building.Name, &_building.Companyno, &_building.Ceo, &_building.Zip, &_building.Address, &_building.Addressetc, &_building.Contractvolumn, &_building.Receivevolumn, &_building.Generatevolumn, &_building.Sunlightvolumn, &_building.Volttype, &_building.Weight, &_building.Totalweight, &_building.Checkcount, &_building.Receivevolt, &_building.Generatevolt, &_building.Periodic, &_building.Businesscondition, &_building.Businessitem, &_building.Usage, &_building.District, &_building.Score, &_building.Status, &_building.Company, &_building.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Tel, &_company.Email, &_company.Address, &_company.Addressetc, &_company.Type, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Bankname, &_company.Bankno, &_company.Businesscondition, &_company.Businessitem, &_company.Giro, &_company.Content, &_company.X1, &_company.Y1, &_company.X2, &_company.Y2, &_company.X3, &_company.Y3, &_company.X4, &_company.Y4, &_company.X5, &_company.Y5, &_company.X6, &_company.Y6, &_company.X7, &_company.Y7, &_company.X8, &_company.Y8, &_company.X9, &_company.Y9, &_company.X10, &_company.Y10, &_company.X11, &_company.Y11, &_company.X12, &_company.Y12, &_company.X13, &_company.Y13, &_company.Status, &_company.Date)
        
        
        
        
        
        
        
        
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
_company.InitExtra()
        item.AddExtra("company",  _company)

        return &item
    }
}

func (p *BillingManager) ReadRows(rows *sql.Rows) []Billing {
    var items []Billing

    for rows.Next() {
        var item Billing
        var _building Building
            var _company Company
            
    
        err := rows.Scan(&item.Id, &item.Price, &item.Status, &item.Giro, &item.Billdate, &item.Month, &item.Endmonth, &item.Period, &item.Company, &item.Building, &item.Date, &_building.Id, &_building.Name, &_building.Companyno, &_building.Ceo, &_building.Zip, &_building.Address, &_building.Addressetc, &_building.Contractvolumn, &_building.Receivevolumn, &_building.Generatevolumn, &_building.Sunlightvolumn, &_building.Volttype, &_building.Weight, &_building.Totalweight, &_building.Checkcount, &_building.Receivevolt, &_building.Generatevolt, &_building.Periodic, &_building.Businesscondition, &_building.Businessitem, &_building.Usage, &_building.District, &_building.Score, &_building.Status, &_building.Company, &_building.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Tel, &_company.Email, &_company.Address, &_company.Addressetc, &_company.Type, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Bankname, &_company.Bankno, &_company.Businesscondition, &_company.Businessitem, &_company.Giro, &_company.Content, &_company.X1, &_company.Y1, &_company.X2, &_company.Y2, &_company.X3, &_company.Y3, &_company.X4, &_company.Y4, &_company.X5, &_company.Y5, &_company.X6, &_company.Y6, &_company.X7, &_company.Y7, &_company.X8, &_company.Y8, &_company.X9, &_company.Y9, &_company.X10, &_company.Y10, &_company.X11, &_company.Y11, &_company.X12, &_company.Y12, &_company.X13, &_company.Y13, &_company.Status, &_company.Date)
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
_company.InitExtra()
        item.AddExtra("company",  _company)

        items = append(items, item)
    }


     return items
}

func (p *BillingManager) Get(id int64) *Billing {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and bi_id = ?"

    
    query += " and bi_building = b_id "    
    
    query += " and b_company = c_id "    
    
    
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

    log.Println(baseQuery + query)
    log.Println(params)
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

    
    query := "select sum(bi_price) from billing_tb"

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
