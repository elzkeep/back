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

type Customercompany struct {
            
    Company                int64 `json:"company"`         
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Companyno                string `json:"companyno"`         
    Ceo                string `json:"ceo"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Type                int `json:"type"`         
    Status                int `json:"status"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"`         
    Date                string `json:"date"`         
    Buildingcount                int64 `json:"buildingcount"`         
    Contractprice                int `json:"contractprice"`         
    Score                Double `json:"score"` 
    
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
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CustomercompanyManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CustomercompanyManager) GetQuery() string {
    ret := ""

    str := "select c_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_billingname, c_billingtel, c_billingemail, c_date, c_buildingcount, c_contractprice, c_score from customercompany_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CustomercompanyManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from customercompany_vw "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CustomercompanyManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate customercompany_vw "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}



func (p *CustomercompanyManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from customercompany_vw where c_id = ?"
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
                query += " and c_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and c_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and c_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from customercompany_vw where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}



func (p *CustomercompanyManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_vw set c_company = c_company + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) IncreaseType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_vw set c_type = c_type + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_vw set c_status = c_status + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) IncreaseBuildingcount(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_vw set c_buildingcount = c_buildingcount + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) IncreaseContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_vw set c_contractprice = c_contractprice + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomercompanyManager) IncreaseScore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customercompany_vw set c_score = c_score + ? where c_id = ?"
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

    

    if rows.Next() {
        err = rows.Scan(&item.Company, &item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Address, &item.Addressetc, &item.Type, &item.Status, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Date, &item.Buildingcount, &item.Contractprice, &item.Score)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *CustomercompanyManager) ReadRows(rows *sql.Rows) []Customercompany {
    var items []Customercompany

    for rows.Next() {
        var item Customercompany
        
    
        err := rows.Scan(&item.Company, &item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Address, &item.Addressetc, &item.Type, &item.Status, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Date, &item.Buildingcount, &item.Contractprice, &item.Score)
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

func (p *CustomercompanyManager) Get(id int64) *Customercompany {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and c_id = ?"

    
    
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
                query += " and c_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and c_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and c_" + item.Column + " " + item.Compare + " ?"
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
                query += " and c_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and c_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and c_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "c_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "c_" + orderby
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
            orderby = "c_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "c_" + orderby
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



func (p *CustomercompanyManager) Sum(args []interface{}) *Customercompany {
    if p.Conn == nil && p.Tx == nil {
        var item Customercompany
        return &item
    }

    var params []interface{}

    
    query := "select sum(c_score) from customercompany_vw"

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
                query += " and c_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and c_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and c_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "c_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "c_" + orderby
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
            orderby = "c_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "c_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Customercompany
    
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
