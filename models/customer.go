package models

import (
    //"zkeep/config"
    
    "zkeep/models/customer"
    
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    log "github.com/sirupsen/logrus"    
    _ "github.com/go-sql-driver/mysql"

    
)

type Customer struct {
            
    Id                int64 `json:"id"`         
    Number                int `json:"number"`         
    Kepconumber                string `json:"kepconumber"`         
    Kesconumber                string `json:"kesconumber"`         
    Type                customer.Type `json:"type"`         
    Checkdate                int `json:"checkdate"`         
    Managername                string `json:"managername"`         
    Managertel                string `json:"managertel"`         
    Manageremail                string `json:"manageremail"`         
    Contractstartdate                string `json:"contractstartdate"`         
    Contractenddate                string `json:"contractenddate"`         
    Contractprice                int `json:"contractprice"`         
    Contractvat                int `json:"contractvat"`         
    Contractday                int `json:"contractday"`         
    Contracttype                int `json:"contracttype"`         
    Billingdate                int `json:"billingdate"`         
    Billingtype                int `json:"billingtype"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Collectmonth                int `json:"collectmonth"`         
    Collectday                int `json:"collectday"`         
    Manager                string `json:"manager"`         
    Tel                string `json:"tel"`         
    Fax                string `json:"fax"`         
    Periodic                string `json:"periodic"`         
    Lastdate                string `json:"lastdate"`         
    Remark                string `json:"remark"`         
    Status                int `json:"status"`         
    Salesuser                int64 `json:"salesuser"`         
    User                int64 `json:"user"`         
    Company                int64 `json:"company"`         
    Building                int64 `json:"building"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type CustomerManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Customer) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewCustomerManager(conn interface{}) *CustomerManager {
    var item CustomerManager

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

func (p *CustomerManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *CustomerManager) SetIndex(index string) {
    p.Index = index
}

func (p *CustomerManager) Exec(query string, params ...interface{}) (sql.Result, error) {
    log.Println(query)
    log.Println(params)    
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *CustomerManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
    log.Println(query)
    log.Println(params)    
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *CustomerManager) GetQuery() string {
    ret := ""

    str := "select cu_id, cu_number, cu_kepconumber, cu_kesconumber, cu_type, cu_checkdate, cu_managername, cu_managertel, cu_manageremail, cu_contractstartdate, cu_contractenddate, cu_contractprice, cu_contractvat, cu_contractday, cu_contracttype, cu_billingdate, cu_billingtype, cu_billingname, cu_billingtel, cu_billingemail, cu_address, cu_addressetc, cu_collectmonth, cu_collectday, cu_manager, cu_tel, cu_fax, cu_periodic, cu_lastdate, cu_remark, cu_status, cu_salesuser, cu_user, cu_company, cu_building, cu_date, b_id, b_name, b_companyno, b_ceo, b_zip, b_address, b_addressetc, b_contractvolumn, b_receivevolumn, b_generatevolumn, b_sunlightvolumn, b_volttype, b_weight, b_totalweight, b_checkcount, b_receivevolt, b_generatevolt, b_periodic, b_businesscondition, b_businessitem, b_usage, b_district, b_score, b_status, b_company, b_date, c_id, c_name, c_companyno, c_ceo, c_tel, c_email, c_address, c_addressetc, c_type, c_billingname, c_billingtel, c_billingemail, c_bankname, c_bankno, c_businesscondition, c_businessitem, c_giro, c_content, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3, c_x4, c_y4, c_x5, c_y5, c_x6, c_y6, c_x7, c_y7, c_x8, c_y8, c_x9, c_y9, c_x10, c_y10, c_x11, c_y11, c_x12, c_y12, c_x13, c_y13, c_status, c_date from customer_tb, building_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and cu_building = b_id "
    
    ret += "and b_company = c_id "
    

    return ret;
}

func (p *CustomerManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from customer_tb, building_tb, company_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    
    ret += "and cu_building = b_id "    
    
    ret += "and b_company = c_id "    
    

    return ret;
}

func (p *CustomerManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate customer_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *CustomerManager) Insert(item *Customer) error {
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
    if item.Lastdate == "" {
       item.Lastdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into customer_tb (cu_id, cu_number, cu_kepconumber, cu_kesconumber, cu_type, cu_checkdate, cu_managername, cu_managertel, cu_manageremail, cu_contractstartdate, cu_contractenddate, cu_contractprice, cu_contractvat, cu_contractday, cu_contracttype, cu_billingdate, cu_billingtype, cu_billingname, cu_billingtel, cu_billingemail, cu_address, cu_addressetc, cu_collectmonth, cu_collectday, cu_manager, cu_tel, cu_fax, cu_periodic, cu_lastdate, cu_remark, cu_status, cu_salesuser, cu_user, cu_company, cu_building, cu_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Number, item.Kepconumber, item.Kesconumber, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Contractvat, item.Contractday, item.Contracttype, item.Billingdate, item.Billingtype, item.Billingname, item.Billingtel, item.Billingemail, item.Address, item.Addressetc, item.Collectmonth, item.Collectday, item.Manager, item.Tel, item.Fax, item.Periodic, item.Lastdate, item.Remark, item.Status, item.Salesuser, item.User, item.Company, item.Building, item.Date)
    } else {
        query = "insert into customer_tb (cu_number, cu_kepconumber, cu_kesconumber, cu_type, cu_checkdate, cu_managername, cu_managertel, cu_manageremail, cu_contractstartdate, cu_contractenddate, cu_contractprice, cu_contractvat, cu_contractday, cu_contracttype, cu_billingdate, cu_billingtype, cu_billingname, cu_billingtel, cu_billingemail, cu_address, cu_addressetc, cu_collectmonth, cu_collectday, cu_manager, cu_tel, cu_fax, cu_periodic, cu_lastdate, cu_remark, cu_status, cu_salesuser, cu_user, cu_company, cu_building, cu_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Number, item.Kepconumber, item.Kesconumber, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Contractvat, item.Contractday, item.Contracttype, item.Billingdate, item.Billingtype, item.Billingname, item.Billingtel, item.Billingemail, item.Address, item.Addressetc, item.Collectmonth, item.Collectday, item.Manager, item.Tel, item.Fax, item.Periodic, item.Lastdate, item.Remark, item.Status, item.Salesuser, item.User, item.Company, item.Building, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *CustomerManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from customer_tb where cu_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *CustomerManager) DeleteWhere(args []interface{}) error {
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
                query += " and cu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cu_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from customer_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *CustomerManager) Update(item *Customer) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Contractstartdate == "" {
       item.Contractstartdate = "1000-01-01"
    }
    if item.Contractenddate == "" {
       item.Contractenddate = "1000-01-01"
    }
    if item.Lastdate == "" {
       item.Lastdate = "1000-01-01"
    }
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update customer_tb set cu_number = ?, cu_kepconumber = ?, cu_kesconumber = ?, cu_type = ?, cu_checkdate = ?, cu_managername = ?, cu_managertel = ?, cu_manageremail = ?, cu_contractstartdate = ?, cu_contractenddate = ?, cu_contractprice = ?, cu_contractvat = ?, cu_contractday = ?, cu_contracttype = ?, cu_billingdate = ?, cu_billingtype = ?, cu_billingname = ?, cu_billingtel = ?, cu_billingemail = ?, cu_address = ?, cu_addressetc = ?, cu_collectmonth = ?, cu_collectday = ?, cu_manager = ?, cu_tel = ?, cu_fax = ?, cu_periodic = ?, cu_lastdate = ?, cu_remark = ?, cu_status = ?, cu_salesuser = ?, cu_user = ?, cu_company = ?, cu_building = ?, cu_date = ? where cu_id = ?"
	_, err := p.Exec(query , item.Number, item.Kepconumber, item.Kesconumber, item.Type, item.Checkdate, item.Managername, item.Managertel, item.Manageremail, item.Contractstartdate, item.Contractenddate, item.Contractprice, item.Contractvat, item.Contractday, item.Contracttype, item.Billingdate, item.Billingtype, item.Billingname, item.Billingtel, item.Billingemail, item.Address, item.Addressetc, item.Collectmonth, item.Collectday, item.Manager, item.Tel, item.Fax, item.Periodic, item.Lastdate, item.Remark, item.Status, item.Salesuser, item.User, item.Company, item.Building, item.Date, item.Id)
    
        
    return err
}


func (p *CustomerManager) UpdateNumber(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_number = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateKepconumber(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_kepconumber = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateKesconumber(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_kesconumber = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_type = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateCheckdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_checkdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManagername(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_managername = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManagertel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_managertel = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManageremail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_manageremail = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractstartdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractstartdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractenddate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractenddate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractprice = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractvat(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractvat = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContractday(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractday = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateContracttype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contracttype = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingtype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingtype = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingname(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingname = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingtel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingtel = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBillingemail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingemail = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateAddress(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_address = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateAddressetc(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_addressetc = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateCollectmonth(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_collectmonth = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateCollectday(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_collectday = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateManager(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_manager = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateTel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_tel = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateFax(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_fax = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdatePeriodic(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_periodic = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateLastdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_lastdate = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateRemark(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_remark = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_status = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateSalesuser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_salesuser = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_user = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_company = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) UpdateBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_building = ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *CustomerManager) IncreaseNumber(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_number = cu_number + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseCheckdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_checkdate = cu_checkdate + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseContractprice(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractprice = cu_contractprice + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseContractvat(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractvat = cu_contractvat + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseContractday(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contractday = cu_contractday + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseContracttype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_contracttype = cu_contracttype + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseBillingdate(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingdate = cu_billingdate + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseBillingtype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_billingtype = cu_billingtype + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseCollectmonth(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_collectmonth = cu_collectmonth + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseCollectday(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_collectday = cu_collectday + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_status = cu_status + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseSalesuser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_salesuser = cu_salesuser + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseUser(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_user = cu_user + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_company = cu_company + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CustomerManager) IncreaseBuilding(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update customer_tb set cu_building = cu_building + ? where cu_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *CustomerManager) GetIdentity() int64 {
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

func (p *Customer) InitExtra() {
    p.Extra = map[string]interface{}{
            "type":     customer.GetType(p.Type),

    }
}

func (p *CustomerManager) ReadRow(rows *sql.Rows) *Customer {
    var item Customer
    var err error

    var _building Building
    var _company Company
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Number, &item.Kepconumber, &item.Kesconumber, &item.Type, &item.Checkdate, &item.Managername, &item.Managertel, &item.Manageremail, &item.Contractstartdate, &item.Contractenddate, &item.Contractprice, &item.Contractvat, &item.Contractday, &item.Contracttype, &item.Billingdate, &item.Billingtype, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Address, &item.Addressetc, &item.Collectmonth, &item.Collectday, &item.Manager, &item.Tel, &item.Fax, &item.Periodic, &item.Lastdate, &item.Remark, &item.Status, &item.Salesuser, &item.User, &item.Company, &item.Building, &item.Date, &_building.Id, &_building.Name, &_building.Companyno, &_building.Ceo, &_building.Zip, &_building.Address, &_building.Addressetc, &_building.Contractvolumn, &_building.Receivevolumn, &_building.Generatevolumn, &_building.Sunlightvolumn, &_building.Volttype, &_building.Weight, &_building.Totalweight, &_building.Checkcount, &_building.Receivevolt, &_building.Generatevolt, &_building.Periodic, &_building.Businesscondition, &_building.Businessitem, &_building.Usage, &_building.District, &_building.Score, &_building.Status, &_building.Company, &_building.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Tel, &_company.Email, &_company.Address, &_company.Addressetc, &_company.Type, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Bankname, &_company.Bankno, &_company.Businesscondition, &_company.Businessitem, &_company.Giro, &_company.Content, &_company.X1, &_company.Y1, &_company.X2, &_company.Y2, &_company.X3, &_company.Y3, &_company.X4, &_company.Y4, &_company.X5, &_company.Y5, &_company.X6, &_company.Y6, &_company.X7, &_company.Y7, &_company.X8, &_company.Y8, &_company.X9, &_company.Y9, &_company.X10, &_company.Y10, &_company.X11, &_company.Y11, &_company.X12, &_company.Y12, &_company.X13, &_company.Y13, &_company.Status, &_company.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Contractstartdate == "0000-00-00" || item.Contractstartdate == "1000-01-01" {
            item.Contractstartdate = ""
        }
        
        if item.Contractenddate == "0000-00-00" || item.Contractenddate == "1000-01-01" {
            item.Contractenddate = ""
        }
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Lastdate == "0000-00-00" || item.Lastdate == "1000-01-01" {
            item.Lastdate = ""
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
        _building.InitExtra()
        item.AddExtra("building",  _building)
_company.InitExtra()
        item.AddExtra("company",  _company)

        return &item
    }
}

func (p *CustomerManager) ReadRows(rows *sql.Rows) []Customer {
    var items []Customer

    for rows.Next() {
        var item Customer
        var _building Building
            var _company Company
            
    
        err := rows.Scan(&item.Id, &item.Number, &item.Kepconumber, &item.Kesconumber, &item.Type, &item.Checkdate, &item.Managername, &item.Managertel, &item.Manageremail, &item.Contractstartdate, &item.Contractenddate, &item.Contractprice, &item.Contractvat, &item.Contractday, &item.Contracttype, &item.Billingdate, &item.Billingtype, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Address, &item.Addressetc, &item.Collectmonth, &item.Collectday, &item.Manager, &item.Tel, &item.Fax, &item.Periodic, &item.Lastdate, &item.Remark, &item.Status, &item.Salesuser, &item.User, &item.Company, &item.Building, &item.Date, &_building.Id, &_building.Name, &_building.Companyno, &_building.Ceo, &_building.Zip, &_building.Address, &_building.Addressetc, &_building.Contractvolumn, &_building.Receivevolumn, &_building.Generatevolumn, &_building.Sunlightvolumn, &_building.Volttype, &_building.Weight, &_building.Totalweight, &_building.Checkcount, &_building.Receivevolt, &_building.Generatevolt, &_building.Periodic, &_building.Businesscondition, &_building.Businessitem, &_building.Usage, &_building.District, &_building.Score, &_building.Status, &_building.Company, &_building.Date, &_company.Id, &_company.Name, &_company.Companyno, &_company.Ceo, &_company.Tel, &_company.Email, &_company.Address, &_company.Addressetc, &_company.Type, &_company.Billingname, &_company.Billingtel, &_company.Billingemail, &_company.Bankname, &_company.Bankno, &_company.Businesscondition, &_company.Businessitem, &_company.Giro, &_company.Content, &_company.X1, &_company.Y1, &_company.X2, &_company.Y2, &_company.X3, &_company.Y3, &_company.X4, &_company.Y4, &_company.X5, &_company.Y5, &_company.X6, &_company.Y6, &_company.X7, &_company.Y7, &_company.X8, &_company.Y8, &_company.X9, &_company.Y9, &_company.X10, &_company.Y10, &_company.X11, &_company.Y11, &_company.X12, &_company.Y12, &_company.X13, &_company.Y13, &_company.Status, &_company.Date)
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
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Lastdate == "0000-00-00" || item.Lastdate == "1000-01-01" {
            item.Lastdate = ""
        }
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" {
            item.Date = ""
        }
        
        item.InitExtra()        
        _building.InitExtra()
        item.AddExtra("building",  _building)
_company.InitExtra()
        item.AddExtra("company",  _company)

        items = append(items, item)
    }


     return items
}

func (p *CustomerManager) Get(id int64) *Customer {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and cu_id = ?"

    
    query += " and cu_building = b_id "    
    
    query += " and b_company = c_id "    
    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *CustomerManager) Count(args []interface{}) int {
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
                query += " and cu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cu_" + item.Column + " " + item.Compare + " ?"
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

func (p *CustomerManager) FindAll() []Customer {
    return p.Find(nil)
}

func (p *CustomerManager) Find(args []interface{}) []Customer {
    if p.Conn == nil && p.Tx == nil {
        var items []Customer
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
                query += " and cu_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and cu_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and cu_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "cu_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "cu_" + orderby
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
            orderby = "cu_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "cu_" + orderby
            }
        }
        query += " order by " + orderby
    }

    log.Println(baseQuery + query)
    log.Println(params)
    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Customer
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *CustomerManager) CountByCompanyBuilding(company int64, building int64, args ...interface{}) int {
    rets := make([]interface{}, 0)
    rets = append(rets, args...)
    
    if company != 0 { 
        rets = append(rets, Where{Column:"company", Value:company, Compare:"="})
     }
    if building != 0 { 
        rets = append(rets, Where{Column:"building", Value:building, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *CustomerManager) GetByCompanyBuilding(company int64, building int64, args ...interface{}) *Customer {
    if company != 0 {
        args = append(args, Where{Column:"company", Value:company, Compare:"="})        
    }
    if building != 0 {
        args = append(args, Where{Column:"building", Value:building, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}

func (p *CustomerManager) DeleteByCompanyBuilding(company int64, building int64) error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from customer_tb where cu_company = ? and cu_building = ?"
    _, err := p.Exec(query, company, building)

    return err
}



