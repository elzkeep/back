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

type Companylicense struct {
            
    Id                int64 `json:"id"`         
    Number                string `json:"number"`         
    Takingdate                string `json:"takingdate"`         
    Educationdate                string `json:"educationdate"`         
    Educationinstitution                string `json:"educationinstitution"`         
    Specialeducationdate                string `json:"specialeducationdate"`         
    Specialeducationinstitution                string `json:"specialeducationinstitution"`         
    Company                int64 `json:"company"`         
    Licensecategory                int64 `json:"licensecategory"`         
    Licenselevel                int64 `json:"licenselevel"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type CompanylicenseManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Companylicense) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewCompanylicenseManager(conn interface{}) *CompanylicenseManager {
    var item CompanylicenseManager

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

func (p *CompanylicenseManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CompanylicenseManager) SetIndex(index string) {
    p.Index = index
}

func (p *CompanylicenseManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CompanylicenseManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CompanylicenseManager) GetQuery() string {
    ret := ""

    str := "select l_id, l_number, l_takingdate, l_educationdate, l_educationinstitution, l_specialeducationdate, l_specialeducationinstitution, l_company, l_licensecategory, l_licenselevel, l_date from companylicense_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CompanylicenseManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from companylicense_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CompanylicenseManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate companylicense_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *CompanylicenseManager) Insert(item *Companylicense) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
    if item.Takingdate == "" {
       item.Takingdate = "1000-01-01"
    }
    if item.Educationdate == "" {
       item.Educationdate = "1000-01-01"
    }
    if item.Specialeducationdate == "" {
       item.Specialeducationdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into companylicense_tb (l_id, l_number, l_takingdate, l_educationdate, l_educationinstitution, l_specialeducationdate, l_specialeducationinstitution, l_company, l_licensecategory, l_licenselevel, l_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Number, item.Takingdate, item.Educationdate, item.Educationinstitution, item.Specialeducationdate, item.Specialeducationinstitution, item.Company, item.Licensecategory, item.Licenselevel, item.Date)
    } else {
        query = "insert into companylicense_tb (l_number, l_takingdate, l_educationdate, l_educationinstitution, l_specialeducationdate, l_specialeducationinstitution, l_company, l_licensecategory, l_licenselevel, l_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Number, item.Takingdate, item.Educationdate, item.Educationinstitution, item.Specialeducationdate, item.Specialeducationinstitution, item.Company, item.Licensecategory, item.Licenselevel, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *CompanylicenseManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from companylicense_tb where l_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *CompanylicenseManager) DeleteWhere(args []interface{}) error {
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
                query += " and l_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and l_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and l_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from companylicense_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *CompanylicenseManager) Update(item *Companylicense) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Takingdate == "" {
       item.Takingdate = "1000-01-01"
    }
    if item.Educationdate == "" {
       item.Educationdate = "1000-01-01"
    }
    if item.Specialeducationdate == "" {
       item.Specialeducationdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update companylicense_tb set l_number = ?, l_takingdate = ?, l_educationdate = ?, l_educationinstitution = ?, l_specialeducationdate = ?, l_specialeducationinstitution = ?, l_company = ?, l_licensecategory = ?, l_licenselevel = ?, l_date = ? where l_id = ?"
	_, err := p.Exec(query , item.Number, item.Takingdate, item.Educationdate, item.Educationinstitution, item.Specialeducationdate, item.Specialeducationinstitution, item.Company, item.Licensecategory, item.Licenselevel, item.Date, item.Id)
    
        
    return err
}


func (p *CompanylicenseManager) UpdateNumber(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_number = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateTakingdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_takingdate = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateEducationdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_educationdate = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateEducationinstitution(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_educationinstitution = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateSpecialeducationdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_specialeducationdate = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateSpecialeducationinstitution(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_specialeducationinstitution = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_company = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateLicensecategory(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_licensecategory = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) UpdateLicenselevel(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_licenselevel = ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *CompanylicenseManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_company = l_company + ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) IncreaseLicensecategory(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_licensecategory = l_licensecategory + ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanylicenseManager) IncreaseLicenselevel(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update companylicense_tb set l_licenselevel = l_licenselevel + ? where l_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *CompanylicenseManager) GetIdentity() int64 {
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

func (p *Companylicense) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *CompanylicenseManager) ReadRow(rows *sql.Rows) *Companylicense {
    var item Companylicense
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Number, &item.Takingdate, &item.Educationdate, &item.Educationinstitution, &item.Specialeducationdate, &item.Specialeducationinstitution, &item.Company, &item.Licensecategory, &item.Licenselevel, &item.Date)
        
        
        
        
        if item.Takingdate == "0000-00-00" || item.Takingdate == "1000-01-01" {
            item.Takingdate = ""
        }
        
        if item.Educationdate == "0000-00-00" || item.Educationdate == "1000-01-01" {
            item.Educationdate = ""
        }
        
        
        
        if item.Specialeducationdate == "0000-00-00" || item.Specialeducationdate == "1000-01-01" {
            item.Specialeducationdate = ""
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

func (p *CompanylicenseManager) ReadRows(rows *sql.Rows) []Companylicense {
    var items []Companylicense

    for rows.Next() {
        var item Companylicense
        
    
        err := rows.Scan(&item.Id, &item.Number, &item.Takingdate, &item.Educationdate, &item.Educationinstitution, &item.Specialeducationdate, &item.Specialeducationinstitution, &item.Company, &item.Licensecategory, &item.Licenselevel, &item.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        if item.Takingdate == "0000-00-00" || item.Takingdate == "1000-01-01" {
            item.Takingdate = ""
        }
        if item.Educationdate == "0000-00-00" || item.Educationdate == "1000-01-01" {
            item.Educationdate = ""
        }
        
        if item.Specialeducationdate == "0000-00-00" || item.Specialeducationdate == "1000-01-01" {
            item.Specialeducationdate = ""
        }
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *CompanylicenseManager) Get(id int64) *Companylicense {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and l_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CompanylicenseManager) Count(args []interface{}) int {
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
                query += " and l_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and l_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and l_" + item.Column + " " + item.Compare + " ?"
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

func (p *CompanylicenseManager) FindAll() []Companylicense {
    return p.Find(nil)
}

func (p *CompanylicenseManager) Find(args []interface{}) []Companylicense {
    if p.Conn == nil && p.Tx == nil {
        var items []Companylicense
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
                query += " and l_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and l_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and l_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "l_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "l_" + orderby
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
            orderby = "l_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "l_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Companylicense
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




