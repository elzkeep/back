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

type Reportlist struct {
            
    Id                int64 `json:"id"`         
    Title                string `json:"title"`         
    Period                int `json:"period"`         
    Number                int `json:"number"`         
    Checkdate                string `json:"checkdate"`         
    Checktime                string `json:"checktime"`         
    Content                string `json:"content"`         
    Image                string `json:"image"`         
    Sign1                string `json:"sign1"`         
    Sign2                string `json:"sign2"`         
    Status                int `json:"status"`         
    Company                int64 `json:"company"`         
    User                int64 `json:"user"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type ReportlistManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Reportlist) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewReportlistManager(conn interface{}) *ReportlistManager {
    var item ReportlistManager

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

func (p *ReportlistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ReportlistManager) SetIndex(index string) {
    p.Index = index
}

func (p *ReportlistManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	log.Println(query)
	log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *ReportlistManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	log.Println(query)
	log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *ReportlistManager) GetQuery() string {
    ret := ""

    str := "select r_id, r_title, r_period, r_number, r_checkdate, r_checktime, r_content, r_image, r_sign1, r_sign2, r_status, r_company, r_user, r_building, r_date from reportlist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *ReportlistManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from reportlist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *ReportlistManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate reportlist_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *ReportlistManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from reportlist_vw where r_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *ReportlistManager) DeleteWhere(args []interface{}) error {
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

    query = "delete from reportlist_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *ReportlistManager) IncreasePeriod(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update reportlist_vw set r_period = r_period + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportlistManager) IncreaseNumber(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update reportlist_vw set r_number = r_number + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportlistManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update reportlist_vw set r_status = r_status + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportlistManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update reportlist_vw set r_company = r_company + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportlistManager) IncreaseUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update reportlist_vw set r_user = r_user + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ReportlistManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update reportlist_vw set r_building = r_building + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *ReportlistManager) GetIdentity() int64 {
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

func (p *Reportlist) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *ReportlistManager) ReadRow(rows *sql.Rows) *Reportlist {
    var item Reportlist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Title, &item.Period, &item.Number, &item.Checkdate, &item.Checktime, &item.Content, &item.Image, &item.Sign1, &item.Sign2, &item.Status, &item.Company, &item.User, &item.Building, &item.Date)
        
        
        
        
        
        
        
        
        if item.Checkdate == "0000-00-00" || item.Checkdate == "1000-01-01" {
            item.Checkdate = ""
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

func (p *ReportlistManager) ReadRows(rows *sql.Rows) []Reportlist {
    var items []Reportlist

    for rows.Next() {
        var item Reportlist
        
    
        err := rows.Scan(&item.Id, &item.Title, &item.Period, &item.Number, &item.Checkdate, &item.Checktime, &item.Content, &item.Image, &item.Sign1, &item.Sign2, &item.Status, &item.Company, &item.User, &item.Building, &item.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        if item.Checkdate == "0000-00-00" || item.Checkdate == "1000-01-01" {
            item.Checkdate = ""
        }
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *ReportlistManager) Get(id int64) *Reportlist {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and r_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ReportlistManager) Count(args []interface{}) int {
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

func (p *ReportlistManager) FindAll() []Reportlist {
    return p.Find(nil)
}

func (p *ReportlistManager) Find(args []interface{}) []Reportlist {
    if p.Conn == nil && p.Tx == nil {
        var items []Reportlist
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
        var items []Reportlist
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




