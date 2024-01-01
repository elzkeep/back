package models

import (
    //"zkeep/config"
    
    "zkeep/models/report"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Report struct {
            
    Id                int64 `json:"id"`         
    Title                string `json:"title"`         
    Period                int `json:"period"`         
    Number                int `json:"number"`         
    Checkdate                string `json:"checkdate"`         
    Checktime                string `json:"checktime"`         
    Content                string `json:"content"`         
    Status                report.Status `json:"status"`         
    Company                int64 `json:"company"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type ReportManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Report) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewReportManager(conn interface{}) *ReportManager {
    var item ReportManager

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

func (p *ReportManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ReportManager) SetIndex(index string) {
    p.Index = index
}

func (p *ReportManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *ReportManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *ReportManager) GetQuery() string {
    ret := ""

    str := "select r_id, r_title, r_period, r_number, r_checkdate, r_checktime, r_content, r_status, r_company, r_date, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_buildingname, c_buildingcompanyno, c_buildingceo, c_buildingaddress, c_buildingaddressetc, c_type, c_checkdate, c_managername, c_managertel, c_manageremail, c_contractstartdate, c_contractenddate, c_contractprice, c_billingdate, c_billingname, c_billingtel, c_billingemail, c_status, c_date from report_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and r_company = c_id "
    

    return ret;
}

func (p *ReportManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from report_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and r_company = c_id "
    

    return ret;
}

func (p *ReportManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate report_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *ReportManager) Insert(item *Report) error {
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
        query = "insert into report_tb (r_id, r_title, r_period, r_number, r_checkdate, r_checktime, r_content, r_status, r_company, r_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Title, item.Period, item.Number, item.Checkdate, item.Checktime, item.Content, item.Status, item.Company, item.Date)
    } else {
        query = "insert into report_tb (r_title, r_period, r_number, r_checkdate, r_checktime, r_content, r_status, r_company, r_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Title, item.Period, item.Number, item.Checkdate, item.Checktime, item.Content, item.Status, item.Company, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *ReportManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from report_tb where r_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *ReportManager) DeleteWhere(args []interface{}) error {
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
                query += " and r_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from report_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *ReportManager) Update(item *Report) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update report_tb set r_title = ?, r_period = ?, r_number = ?, r_checkdate = ?, r_checktime = ?, r_content = ?, r_status = ?, r_company = ?, r_date = ? where r_id = ?"
	_, err := p.Exec(query , item.Title, item.Period, item.Number, item.Checkdate, item.Checktime, item.Content, item.Status, item.Company, item.Date, item.Id)
    
        
    return err
}


func (p *ReportManager) UpdateTitle(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_title = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdatePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_period = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdateNumber(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_number = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdateCheckdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_checkdate = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdateChecktime(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_checktime = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdateContent(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_content = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_status = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_company = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *ReportManager) IncreasePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_period = r_period + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) IncreaseNumber(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_number = r_number + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update report_tb set r_company = r_company + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *ReportManager) GetIdentity() int64 {
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

func (p *Report) InitExtra() {
    p.Extra = map[string]interface{}{
            "status":     report.GetStatus(p.Status),

    }
}

func (p *ReportManager) ReadRow(rows *sql.Rows) *Report {
    var item Report
    var err error

    var _company Company
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Title, &item.Period, &item.Number, &item.Checkdate, &item.Checktime, &item.Content, &item.Status, &item.Company, &item.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Address, &_company.Addressetc, &_company.Buildingname, &_company.Buildingcompanyno, &_company.Buildingceo, &_company.Buildingaddress, &_company.Buildingaddressetc, &_company.Type, &_company.Checkdate, &_company.Managername, &_company.Managertel, &_company.Manageremail, &_company.Contractstartdate, &_company.Contractenddate, &_company.Contractprice, &_company.Billingdate, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Status, &_company.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *ReportManager) ReadRows(rows *sql.Rows) []Report {
    var items []Report

    for rows.Next() {
        var item Report
        var _company Company
            
    
        err := rows.Scan(&item.Id, &item.Title, &item.Period, &item.Number, &item.Checkdate, &item.Checktime, &item.Content, &item.Status, &item.Company, &item.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Address, &_company.Addressetc, &_company.Buildingname, &_company.Buildingcompanyno, &_company.Buildingceo, &_company.Buildingaddress, &_company.Buildingaddressetc, &_company.Type, &_company.Checkdate, &_company.Managername, &_company.Managertel, &_company.Manageremail, &_company.Contractstartdate, &_company.Contractenddate, &_company.Contractprice, &_company.Billingdate, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Status, &_company.Date)
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

func (p *ReportManager) Get(id int64) *Report {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and r_id = ?"

    
    query += " and r_company = c_id"
    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ReportManager) Count(args []interface{}) int {
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
                query += " and r_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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

func (p *ReportManager) FindAll() []Report {
    return p.Find(nil)
}

func (p *ReportManager) Find(args []interface{}) []Report {
    if p.Conn == nil && p.Tx == nil {
        var items []Report
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
                query += " and r_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and r_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and r_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "r_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "r_" + orderby
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
            orderby = "r_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "r_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Report
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




