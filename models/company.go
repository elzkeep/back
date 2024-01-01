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

type Company struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Companyno                string `json:"companyno"`         
    Ceo                string `json:"ceo"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Buildingname                string `json:"buildingname"`         
    Buildingcompanyno                string `json:"buildingcompanyno"`         
    Buildingceo                string `json:"buildingceo"`         
    Buildingaddress                string `json:"buildingaddress"`         
    Buildingaddressetc                string `json:"buildingaddressetc"`         
    Type                int `json:"type"`         
    Checkdate                int `json:"checkdate"`         
    Managername                string `json:"managername"`         
    Managertel                string `json:"managertel"`         
    Manageremail                string `json:"manageremail"`         
    Contractstartdate                string `json:"contractstartdate"`         
    Contractenddate                string `json:"contractenddate"`         
    Contractprice                int `json:"contractprice"`         
    Billingdate                int `json:"billingdate"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"`         
    Status                int `json:"status"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type CompanyManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Company) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewCompanyManager(conn interface{}) *CompanyManager {
    var item CompanyManager

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

func (p *CompanyManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CompanyManager) SetIndex(index string) {
    p.Index = index
}

func (p *CompanyManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CompanyManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CompanyManager) GetQuery() string {
    ret := ""

    str := "select c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_buildingname, c_buildingcompanyno, c_buildingceo, c_buildingaddress, c_buildingaddressetc, c_type, c_checkdate, c_managername, c_managertel, c_manageremail, c_contractstartdate, c_contractenddate, c_contractprice, c_billingdate, c_billingname, c_billingtel, c_billingemail, c_status, c_date from company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CompanyManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *CompanyManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate company_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *CompanyManager) Insert(item *Company) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into company_tb (c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_buildingname, c_buildingcompanyno, c_buildingceo, c_buildingaddress, c_buildingaddressetc, c_type, c_checkdate, c_managername, c_managertel, c_manageremail, c_contractstartdate, c_contractenddate, c_contractprice, c_billingdate, c_billingname, c_billingtel, c_billingemail, c_status, c_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Name, item.Companyno, item.Ceo, item.Address, item.Addressetc, item.Buildingname, item.Buildingcompanyno, item.Buildingceo, item.Buildingaddress, item.Buildingaddressetc, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Billingdate, item.Billingname, item.Billingtel, item.Billingemail, item.Status, item.Date)
    } else {
        query = "insert into company_tb (c_name, c_companyno, c_ceo, c_address, c_addressetc, c_buildingname, c_buildingcompanyno, c_buildingceo, c_buildingaddress, c_buildingaddressetc, c_type, c_checkdate, c_managername, c_managertel, c_manageremail, c_contractstartdate, c_contractenddate, c_contractprice, c_billingdate, c_billingname, c_billingtel, c_billingemail, c_status, c_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Name, item.Companyno, item.Ceo, item.Address, item.Addressetc, item.Buildingname, item.Buildingcompanyno, item.Buildingceo, item.Buildingaddress, item.Buildingaddressetc, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Billingdate, item.Billingname, item.Billingtel, item.Billingemail, item.Status, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *CompanyManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from company_tb where c_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *CompanyManager) DeleteWhere(args []interface{}) error {
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

    query = "delete from company_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *CompanyManager) Update(item *Company) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update company_tb set c_name = ?, c_companyno = ?, c_ceo = ?, c_address = ?, c_addressetc = ?, c_buildingname = ?, c_buildingcompanyno = ?, c_buildingceo = ?, c_buildingaddress = ?, c_buildingaddressetc = ?, c_type = ?, c_checkdate = ?, c_managername = ?, c_managertel = ?, c_manageremail = ?, c_contractstartdate = ?, c_contractenddate = ?, c_contractprice = ?, c_billingdate = ?, c_billingname = ?, c_billingtel = ?, c_billingemail = ?, c_status = ?, c_date = ? where c_id = ?"
	_, err := p.Exec(query , item.Name, item.Companyno, item.Ceo, item.Address, item.Addressetc, item.Buildingname, item.Buildingcompanyno, item.Buildingceo, item.Buildingaddress, item.Buildingaddressetc, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Billingdate, item.Billingname, item.Billingtel, item.Billingemail, item.Status, item.Date, item.Id)
    
        
    return err
}


func (p *CompanyManager) UpdateName(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_name = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateCompanyno(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_companyno = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateCeo(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_ceo = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateAddress(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_address = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateAddressetc(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_addressetc = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBuildingname(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_buildingname = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBuildingcompanyno(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_buildingcompanyno = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBuildingceo(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_buildingceo = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBuildingaddress(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_buildingaddress = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBuildingaddressetc(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_buildingaddressetc = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_type = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateCheckdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_checkdate = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateManagername(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_managername = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateManagertel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_managertel = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateManageremail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_manageremail = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateContractstartdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_contractstartdate = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateContractenddate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_contractenddate = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_contractprice = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBillingdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_billingdate = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBillingname(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_billingname = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBillingtel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_billingtel = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBillingemail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_billingemail = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_status = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *CompanyManager) IncreaseType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_type = c_type + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseCheckdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_checkdate = c_checkdate + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_contractprice = c_contractprice + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseBillingdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_billingdate = c_billingdate + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_status = c_status + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *CompanyManager) GetIdentity() int64 {
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

func (p *Company) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *CompanyManager) ReadRow(rows *sql.Rows) *Company {
    var item Company
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Address, &item.Addressetc, &item.Buildingname, &item.Buildingcompanyno, &item.Buildingceo, &item.Buildingaddress, &item.Buildingaddressetc, &item.Type, &item.Checkdate, &item.Managername, &item.Managertel, &item.Manageremail, &item.Contractstartdate, &item.Contractenddate, &item.Contractprice, &item.Billingdate, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Status, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" {
            item.Contractstartdate = ""
        }
        
        if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" {
            item.Contractenddate = ""
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

func (p *CompanyManager) ReadRows(rows *sql.Rows) []Company {
    var items []Company

    for rows.Next() {
        var item Company
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Address, &item.Addressetc, &item.Buildingname, &item.Buildingcompanyno, &item.Buildingceo, &item.Buildingaddress, &item.Buildingaddressetc, &item.Type, &item.Checkdate, &item.Managername, &item.Managertel, &item.Manageremail, &item.Contractstartdate, &item.Contractenddate, &item.Contractprice, &item.Billingdate, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Status, &item.Date)
        if err != nil {
           log.Printf("ReadRows error : %v\n", err)
           break
        }

        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" {
            item.Contractstartdate = ""
        }
        if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" {
            item.Contractenddate = ""
        }
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *CompanyManager) Get(id int64) *Company {
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

func (p *CompanyManager) Count(args []interface{}) int {
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

func (p *CompanyManager) FindAll() []Company {
    return p.Find(nil)
}

func (p *CompanyManager) Find(args []interface{}) []Company {
    if p.Conn == nil && p.Tx == nil {
        var items []Company
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
        var items []Company
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}




