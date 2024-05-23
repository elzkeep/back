package models

import (
    //"zkeep/config"
    
    "zkeep/models/data"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Data struct {
            
    Id                int64 `json:"id"`         
    Topcategory                int `json:"topcategory"`         
    Title                string `json:"title"`         
    Type                data.Type `json:"type"`         
    Category                int `json:"category"`         
    Order                int `json:"order"`         
    Report                int64 `json:"report"`         
    Company                int64 `json:"company"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type DataManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Data) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewDataManager(conn interface{}) *DataManager {
    var item DataManager

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

func (p *DataManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *DataManager) SetIndex(index string) {
    p.Index = index
}

func (p *DataManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *DataManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *DataManager) GetQuery() string {
    ret := ""

    str := "select d_id, d_topcategory, d_title, d_type, d_category, d_order, d_report, d_company, d_date from data_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *DataManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from data_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *DataManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate data_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *DataManager) Insert(item *Data) error {
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
        query = "insert into data_tb (d_id, d_topcategory, d_title, d_type, d_category, d_order, d_report, d_company, d_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Topcategory, item.Title, item.Type, item.Category, item.Order, item.Report, item.Company, item.Date)
    } else {
        query = "insert into data_tb (d_topcategory, d_title, d_type, d_category, d_order, d_report, d_company, d_date) values (?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Topcategory, item.Title, item.Type, item.Category, item.Order, item.Report, item.Company, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *DataManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from data_tb where d_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *DataManager) DeleteWhere(args []interface{}) error {
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
                query += " and d_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and d_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and d_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from data_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *DataManager) Update(item *Data) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update data_tb set d_topcategory = ?, d_title = ?, d_type = ?, d_category = ?, d_order = ?, d_report = ?, d_company = ?, d_date = ? where d_id = ?"
	_, err := p.Exec(query , item.Topcategory, item.Title, item.Type, item.Category, item.Order, item.Report, item.Company, item.Date, item.Id)
    
        
    return err
}


func (p *DataManager) UpdateTopcategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_topcategory = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) UpdateTitle(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_title = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_type = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) UpdateCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_category = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) UpdateOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_order = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) UpdateReport(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_report = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_company = ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *DataManager) IncreaseTopcategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_topcategory = d_topcategory + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) IncreaseCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_category = d_category + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) IncreaseOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_order = d_order + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) IncreaseReport(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_report = d_report + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DataManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update data_tb set d_company = d_company + ? where d_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *DataManager) GetIdentity() int64 {
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

func (p *Data) InitExtra() {
    p.Extra = map[string]interface{}{
            "type":     data.GetType(p.Type),

    }
}

func (p *DataManager) ReadRow(rows *sql.Rows) *Data {
    var item Data
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Topcategory, &item.Title, &item.Type, &item.Category, &item.Order, &item.Report, &item.Company, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *DataManager) ReadRows(rows *sql.Rows) []Data {
    var items []Data

    for rows.Next() {
        var item Data
        
    
        err := rows.Scan(&item.Id, &item.Topcategory, &item.Title, &item.Type, &item.Category, &item.Order, &item.Report, &item.Company, &item.Date)
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

func (p *DataManager) Get(id int64) *Data {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and d_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *DataManager) Count(args []interface{}) int {
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
                query += " and d_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and d_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and d_" + item.Column + " " + item.Compare + " ?"
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

    log.Println(query)
    log.Println(params)
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

func (p *DataManager) FindAll() []Data {
    return p.Find(nil)
}

func (p *DataManager) Find(args []interface{}) []Data {
    if p.Conn == nil && p.Tx == nil {
        var items []Data
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
                query += " and d_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and d_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and d_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "d_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "d_" + orderby
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
            orderby = "d_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "d_" + orderby
            }
        }
        query += " order by " + orderby
    }

    log.Println(baseQuery + query)
    log.Println(params)
    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Data
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *DataManager) DeleteByReportTopcategory(report int64, topcategory int) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from data_tb where d_report = ? and d_topcategory = ?"
    _, err := p.Exec(query, report, topcategory)

    return err
}



