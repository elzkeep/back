package models

import (
    //"zkeep/config"
    
    "zkeep/models/item"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Item struct {
            
    Id                int64 `json:"id"`         
    Title                string `json:"title"`         
    Type                item.Type `json:"type"`         
    Value1                int `json:"value1"`         
    Value2                int `json:"value2"`         
    Value3                int `json:"value3"`         
    Value4                int `json:"value4"`         
    Value5                int `json:"value5"`         
    Value6                int `json:"value6"`         
    Value7                int `json:"value7"`         
    Value8                int `json:"value8"`         
    Value                int `json:"value"`         
    Unit                string `json:"unit"`         
    Status                item.Status `json:"status"`         
    Reason                int `json:"reason"`         
    Reasontext                string `json:"reasontext"`         
    Action                int `json:"action"`         
    Actiontext                string `json:"actiontext"`         
    Image                string `json:"image"`         
    Order                int `json:"order"`         
    Data                int64 `json:"data"`         
    Report                int64 `json:"report"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type ItemManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Item) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewItemManager(conn interface{}) *ItemManager {
    var item ItemManager

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

func (p *ItemManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *ItemManager) SetIndex(index string) {
    p.Index = index
}

func (p *ItemManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *ItemManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *ItemManager) GetQuery() string {
    ret := ""

    str := "select i_id, i_title, i_type, i_value1, i_value2, i_value3, i_value4, i_value5, i_value6, i_value7, i_value8, i_value, i_unit, i_status, i_reason, i_reasontext, i_action, i_actiontext, i_image, i_order, i_data, i_report, i_date from item_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *ItemManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from item_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *ItemManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate item_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *ItemManager) Insert(item *Item) error {
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
        query = "insert into item_tb (i_id, i_title, i_type, i_value1, i_value2, i_value3, i_value4, i_value5, i_value6, i_value7, i_value8, i_value, i_unit, i_status, i_reason, i_reasontext, i_action, i_actiontext, i_image, i_order, i_data, i_report, i_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Title, item.Type, item.Value1, item.Value2, item.Value3, item.Value4, item.Value5, item.Value6, item.Value7, item.Value8, item.Value, item.Unit, item.Status, item.Reason, item.Reasontext, item.Action, item.Actiontext, item.Image, item.Order, item.Data, item.Report, item.Date)
    } else {
        query = "insert into item_tb (i_title, i_type, i_value1, i_value2, i_value3, i_value4, i_value5, i_value6, i_value7, i_value8, i_value, i_unit, i_status, i_reason, i_reasontext, i_action, i_actiontext, i_image, i_order, i_data, i_report, i_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Title, item.Type, item.Value1, item.Value2, item.Value3, item.Value4, item.Value5, item.Value6, item.Value7, item.Value8, item.Value, item.Unit, item.Status, item.Reason, item.Reasontext, item.Action, item.Actiontext, item.Image, item.Order, item.Data, item.Report, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *ItemManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from item_tb where i_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *ItemManager) DeleteWhere(args []interface{}) error {
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
                query += " and i_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and i_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and i_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from item_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *ItemManager) Update(item *Item) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update item_tb set i_title = ?, i_type = ?, i_value1 = ?, i_value2 = ?, i_value3 = ?, i_value4 = ?, i_value5 = ?, i_value6 = ?, i_value7 = ?, i_value8 = ?, i_value = ?, i_unit = ?, i_status = ?, i_reason = ?, i_reasontext = ?, i_action = ?, i_actiontext = ?, i_image = ?, i_order = ?, i_data = ?, i_report = ?, i_date = ? where i_id = ?"
	_, err := p.Exec(query , item.Title, item.Type, item.Value1, item.Value2, item.Value3, item.Value4, item.Value5, item.Value6, item.Value7, item.Value8, item.Value, item.Unit, item.Status, item.Reason, item.Reasontext, item.Action, item.Actiontext, item.Image, item.Order, item.Data, item.Report, item.Date, item.Id)
    
        
    return err
}


func (p *ItemManager) UpdateTitle(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_title = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_type = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue1(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value1 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue2(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value2 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue3(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value3 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue4(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value4 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue5(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value5 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue6(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value6 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue7(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value7 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue8(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value8 = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateValue(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateUnit(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_unit = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_status = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateReason(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_reason = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateReasontext(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_reasontext = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateAction(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_action = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateActiontext(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_actiontext = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateImage(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_image = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_order = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateData(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_data = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) UpdateReport(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_report = ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *ItemManager) IncreaseValue1(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value1 = i_value1 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue2(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value2 = i_value2 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue3(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value3 = i_value3 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue4(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value4 = i_value4 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue5(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value5 = i_value5 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue6(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value6 = i_value6 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue7(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value7 = i_value7 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue8(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value8 = i_value8 + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseValue(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_value = i_value + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseReason(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_reason = i_reason + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseAction(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_action = i_action + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseOrder(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_order = i_order + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseData(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_data = i_data + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *ItemManager) IncreaseReport(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update item_tb set i_report = i_report + ? where i_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *ItemManager) GetIdentity() int64 {
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

func (p *Item) InitExtra() {
    p.Extra = map[string]interface{}{
            "type":     item.GetType(p.Type),
            "status":     item.GetStatus(p.Status),

    }
}

func (p *ItemManager) ReadRow(rows *sql.Rows) *Item {
    var item Item
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Title, &item.Type, &item.Value1, &item.Value2, &item.Value3, &item.Value4, &item.Value5, &item.Value6, &item.Value7, &item.Value8, &item.Value, &item.Unit, &item.Status, &item.Reason, &item.Reasontext, &item.Action, &item.Actiontext, &item.Image, &item.Order, &item.Data, &item.Report, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *ItemManager) ReadRows(rows *sql.Rows) []Item {
    var items []Item

    for rows.Next() {
        var item Item
        
    
        err := rows.Scan(&item.Id, &item.Title, &item.Type, &item.Value1, &item.Value2, &item.Value3, &item.Value4, &item.Value5, &item.Value6, &item.Value7, &item.Value8, &item.Value, &item.Unit, &item.Status, &item.Reason, &item.Reasontext, &item.Action, &item.Actiontext, &item.Image, &item.Order, &item.Data, &item.Report, &item.Date)
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

func (p *ItemManager) Get(id int64) *Item {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and i_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *ItemManager) Count(args []interface{}) int {
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
                query += " and i_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and i_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and i_" + item.Column + " " + item.Compare + " ?"
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

func (p *ItemManager) FindAll() []Item {
    return p.Find(nil)
}

func (p *ItemManager) Find(args []interface{}) []Item {
    if p.Conn == nil && p.Tx == nil {
        var items []Item
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
                query += " and i_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and i_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and i_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "i_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "i_" + orderby
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
            orderby = "i_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "i_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Item
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




