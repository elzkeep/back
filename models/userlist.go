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

type Userlist struct {
            
    Id                int64 `json:"id"`         
    Loginid                string `json:"loginid"`         
    Passwd                string `json:"passwd"`         
    Name                string `json:"name"`         
    Email                string `json:"email"`         
    Tel                string `json:"tel"`         
    Zip                string `json:"zip"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Joindate                string `json:"joindate"`         
    Careeryear                int `json:"careeryear"`         
    Careermonth                int `json:"careermonth"`         
    Level                int `json:"level"`         
    Score                Double `json:"score"`         
    Approval                int `json:"approval"`         
    Educationdate                string `json:"educationdate"`         
    Educationinstitution                string `json:"educationinstitution"`         
    Specialeducationdate                string `json:"specialeducationdate"`         
    Specialeducationinstitution                string `json:"specialeducationinstitution"`         
    Rejectreason                string `json:"rejectreason"`         
    Status                int `json:"status"`         
    Company                int64 `json:"company"`         
    Department                int64 `json:"department"`         
    Date                string `json:"date"`         
    Totalscore                Double `json:"totalscore"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type UserlistManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Userlist) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewUserlistManager(conn interface{}) *UserlistManager {
    var item UserlistManager

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

func (p *UserlistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *UserlistManager) SetIndex(index string) {
    p.Index = index
}

func (p *UserlistManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *UserlistManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *UserlistManager) GetQuery() string {
    ret := ""

    str := "select u_id, u_loginid, u_passwd, u_name, u_email, u_tel, u_zip, u_address, u_addressetc, u_joindate, u_careeryear, u_careermonth, u_level, u_score, u_approval, u_educationdate, u_educationinstitution, u_specialeducationdate, u_specialeducationinstitution, u_rejectreason, u_status, u_company, u_department, u_date, u_totalscore from userlist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *UserlistManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from userlist_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *UserlistManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate userlist_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *UserlistManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from userlist_vw where u_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *UserlistManager) DeleteWhere(args []interface{}) error {
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
                query += " and u_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from userlist_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *UserlistManager) IncreaseCareeryear(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_careeryear = u_careeryear + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseCareermonth(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_careermonth = u_careermonth + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseLevel(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_level = u_level + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseScore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_score = u_score + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseApproval(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_approval = u_approval + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_status = u_status + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_company = u_company + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseDepartment(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_department = u_department + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *UserlistManager) IncreaseTotalscore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update userlist_vw set u_totalscore = u_totalscore + ? where u_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *UserlistManager) GetIdentity() int64 {
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

func (p *Userlist) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *UserlistManager) ReadRow(rows *sql.Rows) *Userlist {
    var item Userlist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Email, &item.Tel, &item.Zip, &item.Address, &item.Addressetc, &item.Joindate, &item.Careeryear, &item.Careermonth, &item.Level, &item.Score, &item.Approval, &item.Educationdate, &item.Educationinstitution, &item.Specialeducationdate, &item.Specialeducationinstitution, &item.Rejectreason, &item.Status, &item.Company, &item.Department, &item.Date, &item.Totalscore)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Joindate == "0000-00-00" || item.Joindate == "1000-01-01" {
            item.Joindate = ""
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

func (p *UserlistManager) ReadRows(rows *sql.Rows) []Userlist {
    var items []Userlist

    for rows.Next() {
        var item Userlist
        
    
        err := rows.Scan(&item.Id, &item.Loginid, &item.Passwd, &item.Name, &item.Email, &item.Tel, &item.Zip, &item.Address, &item.Addressetc, &item.Joindate, &item.Careeryear, &item.Careermonth, &item.Level, &item.Score, &item.Approval, &item.Educationdate, &item.Educationinstitution, &item.Specialeducationdate, &item.Specialeducationinstitution, &item.Rejectreason, &item.Status, &item.Company, &item.Department, &item.Date, &item.Totalscore)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        
        
        if item.Joindate == "0000-00-00" || item.Joindate == "1000-01-01" {
            item.Joindate = ""
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

func (p *UserlistManager) Get(id int64) *Userlist {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and u_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *UserlistManager) Count(args []interface{}) int {
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
                query += " and u_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
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

func (p *UserlistManager) FindAll() []Userlist {
    return p.Find(nil)
}

func (p *UserlistManager) Find(args []interface{}) []Userlist {
    if p.Conn == nil && p.Tx == nil {
        var items []Userlist
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
                query += " and u_" + item.Column
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
            orderby = "u_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "u_" + orderby
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
            orderby = "u_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "u_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Userlist
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}



func (p *UserlistManager) Sum(args []interface{}) *Userlist {
    if p.Conn == nil && p.Tx == nil {
        var item Userlist
        return &item
    }

    var params []interface{}

    
    query := "select sum(u_score) from userlist_vw"

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
                query += " and u_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and u_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and u_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "u_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "u_" + orderby
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
            orderby = "u_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "u_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Userlist
    
    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return &item
    }

    defer rows.Close()

    if rows.Next() {
        
        rows.Scan(&item.Score)        
    }

    return &item        
}
