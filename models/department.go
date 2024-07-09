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

type Department struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Status                int `json:"status"`         
    Order                int `json:"order"`         
    Parent                int64 `json:"parent"`         
    Company                int64 `json:"company"`         
    Master                int64 `json:"master"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type DepartmentManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Department) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewDepartmentManager(conn interface{}) *DepartmentManager {
    var item DepartmentManager

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

func (p *DepartmentManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *DepartmentManager) SetIndex(index string) {
    p.Index = index
}

func (p *DepartmentManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *DepartmentManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *DepartmentManager) GetQuery() string {
    ret := ""

    str := "select de_id, de_name, de_status, de_order, de_parent, de_company, de_master, de_date from department_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *DepartmentManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from department_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *DepartmentManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate department_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *DepartmentManager) Insert(item *Department) error {
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
        query = "insert into department_tb (de_id, de_name, de_status, de_order, de_parent, de_company, de_master, de_date) values (?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Name, item.Status, item.Order, item.Parent, item.Company, item.Master, item.Date)
    } else {
        query = "insert into department_tb (de_name, de_status, de_order, de_parent, de_company, de_master, de_date) values (?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Name, item.Status, item.Order, item.Parent, item.Company, item.Master, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *DepartmentManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from department_tb where de_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *DepartmentManager) DeleteWhere(args []interface{}) error {
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
                query += " and de_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and de_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and de_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from department_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *DepartmentManager) Update(item *Department) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update department_tb set de_name = ?, de_status = ?, de_order = ?, de_parent = ?, de_company = ?, de_master = ?, de_date = ? where de_id = ?"
	_, err := p.Exec(query , item.Name, item.Status, item.Order, item.Parent, item.Company, item.Master, item.Date, item.Id)
    
        
    return err
}


func (p *DepartmentManager) UpdateName(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_name = ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_status = ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) UpdateOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_order = ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) UpdateParent(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_parent = ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_company = ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) UpdateMaster(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_master = ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *DepartmentManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_status = de_status + ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) IncreaseOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_order = de_order + ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) IncreaseParent(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_parent = de_parent + ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_company = de_company + ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *DepartmentManager) IncreaseMaster(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update department_tb set de_master = de_master + ? where de_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *DepartmentManager) GetIdentity() int64 {
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

func (p *Department) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *DepartmentManager) ReadRow(rows *sql.Rows) *Department {
    var item Department
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Status, &item.Order, &item.Parent, &item.Company, &item.Master, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *DepartmentManager) ReadRows(rows *sql.Rows) []Department {
    var items []Department

    for rows.Next() {
        var item Department
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Status, &item.Order, &item.Parent, &item.Company, &item.Master, &item.Date)
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

func (p *DepartmentManager) Get(id int64) *Department {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and de_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *DepartmentManager) Count(args []interface{}) int {
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
                query += " and de_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and de_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and de_" + item.Column + " " + item.Compare + " ?"
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

func (p *DepartmentManager) FindAll() []Department {
    return p.Find(nil)
}

func (p *DepartmentManager) Find(args []interface{}) []Department {
    if p.Conn == nil && p.Tx == nil {
        var items []Department
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
                query += " and de_" + item.Column
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
            orderby = "de_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "de_" + orderby
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
            orderby = "de_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "de_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Department
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *DepartmentManager) GetByCompanyName(company int64, name string, args ...interface{}) *Department {
    if company != 0 {
        args = append(args, Where{Column:"company", Value:company, Compare:"="})        
    }
    if name != "" {
        args = append(args, Where{Column:"name", Value:name, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}



