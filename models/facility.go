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

type Facility struct {
            
    Id                int64 `json:"id"`         
    Category                int `json:"category"`         
    Parent                int64 `json:"parent"`         
    Value1                string `json:"value1"`         
    Value2                string `json:"value2"`         
    Value3                string `json:"value3"`         
    Value4                string `json:"value4"`         
    Value5                string `json:"value5"`         
    Value6                string `json:"value6"`         
    Value7                string `json:"value7"`         
    Value8                string `json:"value8"`         
    Value9                string `json:"value9"`         
    Value10                string `json:"value10"`         
    Value11                string `json:"value11"`         
    Value12                string `json:"value12"`         
    Value13                string `json:"value13"`         
    Value14                string `json:"value14"`         
    Value15                string `json:"value15"`         
    Value16                string `json:"value16"`         
    Value17                string `json:"value17"`         
    Value18                string `json:"value18"`         
    Value19                string `json:"value19"`         
    Value20                string `json:"value20"`         
    Content                string `json:"content"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type FacilityManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Facility) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewFacilityManager(conn interface{}) *FacilityManager {
    var item FacilityManager

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

func (p *FacilityManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *FacilityManager) SetIndex(index string) {
    p.Index = index
}

func (p *FacilityManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *FacilityManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *FacilityManager) GetQuery() string {
    ret := ""

    str := "select f_id, f_category, f_parent, f_value1, f_value2, f_value3, f_value4, f_value5, f_value6, f_value7, f_value8, f_value9, f_value10, f_value11, f_value12, f_value13, f_value14, f_value15, f_value16, f_value17, f_value18, f_value19, f_value20, f_content, f_building, f_date from facility_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *FacilityManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from facility_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *FacilityManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate facility_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *FacilityManager) Insert(item *Facility) error {
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
        query = "insert into facility_tb (f_id, f_category, f_parent, f_value1, f_value2, f_value3, f_value4, f_value5, f_value6, f_value7, f_value8, f_value9, f_value10, f_value11, f_value12, f_value13, f_value14, f_value15, f_value16, f_value17, f_value18, f_value19, f_value20, f_content, f_building, f_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Category, item.Parent, item.Value1, item.Value2, item.Value3, item.Value4, item.Value5, item.Value6, item.Value7, item.Value8, item.Value9, item.Value10, item.Value11, item.Value12, item.Value13, item.Value14, item.Value15, item.Value16, item.Value17, item.Value18, item.Value19, item.Value20, item.Content, item.Building, item.Date)
    } else {
        query = "insert into facility_tb (f_category, f_parent, f_value1, f_value2, f_value3, f_value4, f_value5, f_value6, f_value7, f_value8, f_value9, f_value10, f_value11, f_value12, f_value13, f_value14, f_value15, f_value16, f_value17, f_value18, f_value19, f_value20, f_content, f_building, f_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Category, item.Parent, item.Value1, item.Value2, item.Value3, item.Value4, item.Value5, item.Value6, item.Value7, item.Value8, item.Value9, item.Value10, item.Value11, item.Value12, item.Value13, item.Value14, item.Value15, item.Value16, item.Value17, item.Value18, item.Value19, item.Value20, item.Content, item.Building, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *FacilityManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from facility_tb where f_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *FacilityManager) DeleteWhere(args []interface{}) error {
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
                query += " and f_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and f_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and f_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from facility_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *FacilityManager) Update(item *Facility) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update facility_tb set f_category = ?, f_parent = ?, f_value1 = ?, f_value2 = ?, f_value3 = ?, f_value4 = ?, f_value5 = ?, f_value6 = ?, f_value7 = ?, f_value8 = ?, f_value9 = ?, f_value10 = ?, f_value11 = ?, f_value12 = ?, f_value13 = ?, f_value14 = ?, f_value15 = ?, f_value16 = ?, f_value17 = ?, f_value18 = ?, f_value19 = ?, f_value20 = ?, f_content = ?, f_building = ?, f_date = ? where f_id = ?"
	_, err := p.Exec(query , item.Category, item.Parent, item.Value1, item.Value2, item.Value3, item.Value4, item.Value5, item.Value6, item.Value7, item.Value8, item.Value9, item.Value10, item.Value11, item.Value12, item.Value13, item.Value14, item.Value15, item.Value16, item.Value17, item.Value18, item.Value19, item.Value20, item.Content, item.Building, item.Date, item.Id)
    
        
    return err
}


func (p *FacilityManager) UpdateCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_category = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateParent(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_parent = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue1(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value1 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue2(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value2 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue3(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value3 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue4(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value4 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue5(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value5 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue6(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value6 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue7(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value7 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue8(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value8 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue9(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value9 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue10(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value10 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue11(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value11 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue12(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value12 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue13(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value13 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue14(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value14 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue15(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value15 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue16(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value16 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue17(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value17 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue18(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value18 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue19(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value19 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateValue20(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_value20 = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateContent(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_content = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) UpdateBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_building = ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *FacilityManager) IncreaseCategory(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_category = f_category + ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) IncreaseParent(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_parent = f_parent + ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *FacilityManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update facility_tb set f_building = f_building + ? where f_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *FacilityManager) GetIdentity() int64 {
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

func (p *Facility) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *FacilityManager) ReadRow(rows *sql.Rows) *Facility {
    var item Facility
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Category, &item.Parent, &item.Value1, &item.Value2, &item.Value3, &item.Value4, &item.Value5, &item.Value6, &item.Value7, &item.Value8, &item.Value9, &item.Value10, &item.Value11, &item.Value12, &item.Value13, &item.Value14, &item.Value15, &item.Value16, &item.Value17, &item.Value18, &item.Value19, &item.Value20, &item.Content, &item.Building, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *FacilityManager) ReadRows(rows *sql.Rows) []Facility {
    var items []Facility

    for rows.Next() {
        var item Facility
        
    
        err := rows.Scan(&item.Id, &item.Category, &item.Parent, &item.Value1, &item.Value2, &item.Value3, &item.Value4, &item.Value5, &item.Value6, &item.Value7, &item.Value8, &item.Value9, &item.Value10, &item.Value11, &item.Value12, &item.Value13, &item.Value14, &item.Value15, &item.Value16, &item.Value17, &item.Value18, &item.Value19, &item.Value20, &item.Content, &item.Building, &item.Date)
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

func (p *FacilityManager) Get(id int64) *Facility {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and f_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *FacilityManager) Count(args []interface{}) int {
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
                query += " and f_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and f_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and f_" + item.Column + " " + item.Compare + " ?"
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

func (p *FacilityManager) FindAll() []Facility {
    return p.Find(nil)
}

func (p *FacilityManager) Find(args []interface{}) []Facility {
    if p.Conn == nil && p.Tx == nil {
        var items []Facility
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
                query += " and f_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and f_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and f_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "f_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "f_" + orderby
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
            orderby = "f_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "f_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Facility
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *FacilityManager) DeleteByBuildingCategory(building int64, category int) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from facility_tb where f_building = ? and f_category = ?"
    _, err := p.Exec(query, building, category)

    return err
}



