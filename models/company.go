package models

import (
    //"zkeep/config"
    
    "zkeep/models/company"
    
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
    Tel                string `json:"tel"`         
    Email                string `json:"email"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Type                company.Type `json:"type"`         
    Billingname                string `json:"billingname"`         
    Billingtel                string `json:"billingtel"`         
    Billingemail                string `json:"billingemail"`         
    Bankname                string `json:"bankname"`         
    Bankno                string `json:"bankno"`         
    Businesscondition                string `json:"businesscondition"`         
    Businessitem                string `json:"businessitem"`         
    Giro                string `json:"giro"`         
    Content                string `json:"content"`         
    X1                Double `json:"x1"`         
    Y1                Double `json:"y1"`         
    X2                Double `json:"x2"`         
    Y2                Double `json:"y2"`         
    X3                Double `json:"x3"`         
    Y3                Double `json:"y3"`         
    X4                Double `json:"x4"`         
    Y4                Double `json:"y4"`         
    X5                Double `json:"x5"`         
    Y5                Double `json:"y5"`         
    X6                Double `json:"x6"`         
    Y6                Double `json:"y6"`         
    X7                Double `json:"x7"`         
    Y7                Double `json:"y7"`         
    X8                Double `json:"x8"`         
    Y8                Double `json:"y8"`         
    X9                Double `json:"x9"`         
    Y9                Double `json:"y9"`         
    X10                Double `json:"x10"`         
    Y10                Double `json:"y10"`         
    X11                Double `json:"x11"`         
    Y11                Double `json:"y11"`         
    X12                Double `json:"x12"`         
    Y12                Double `json:"y12"`         
    X13                Double `json:"x13"`         
    Y13                Double `json:"y13"`         
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

    str := "select c_id, c_name, c_companyno, c_ceo, c_tel, c_email, c_address, c_addressetc, c_type, c_billingname, c_billingtel, c_billingemail, c_bankname, c_bankno, c_businesscondition, c_businessitem, c_giro, c_content, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3, c_x4, c_y4, c_x5, c_y5, c_x6, c_y6, c_x7, c_y7, c_x8, c_y8, c_x9, c_y9, c_x10, c_y10, c_x11, c_y11, c_x12, c_y12, c_x13, c_y13, c_status, c_date from company_tb "

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

    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        query = "insert into company_tb (c_id, c_name, c_companyno, c_ceo, c_tel, c_email, c_address, c_addressetc, c_type, c_billingname, c_billingtel, c_billingemail, c_bankname, c_bankno, c_businesscondition, c_businessitem, c_giro, c_content, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3, c_x4, c_y4, c_x5, c_y5, c_x6, c_y6, c_x7, c_y7, c_x8, c_y8, c_x9, c_y9, c_x10, c_y10, c_x11, c_y11, c_x12, c_y12, c_x13, c_y13, c_status, c_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Name, item.Companyno, item.Ceo, item.Tel, item.Email, item.Address, item.Addressetc, item.Type, item.Billingname, item.Billingtel, item.Billingemail, item.Bankname, item.Bankno, item.Businesscondition, item.Businessitem, item.Giro, item.Content, item.X1, item.Y1, item.X2, item.Y2, item.X3, item.Y3, item.X4, item.Y4, item.X5, item.Y5, item.X6, item.Y6, item.X7, item.Y7, item.X8, item.Y8, item.X9, item.Y9, item.X10, item.Y10, item.X11, item.Y11, item.X12, item.Y12, item.X13, item.Y13, item.Status, item.Date)
    } else {
        query = "insert into company_tb (c_name, c_companyno, c_ceo, c_tel, c_email, c_address, c_addressetc, c_type, c_billingname, c_billingtel, c_billingemail, c_bankname, c_bankno, c_businesscondition, c_businessitem, c_giro, c_content, c_x1, c_y1, c_x2, c_y2, c_x3, c_y3, c_x4, c_y4, c_x5, c_y5, c_x6, c_y6, c_x7, c_y7, c_x8, c_y8, c_x9, c_y9, c_x10, c_y10, c_x11, c_y11, c_x12, c_y12, c_x13, c_y13, c_status, c_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Name, item.Companyno, item.Ceo, item.Tel, item.Email, item.Address, item.Addressetc, item.Type, item.Billingname, item.Billingtel, item.Billingemail, item.Bankname, item.Bankno, item.Businesscondition, item.Businessitem, item.Giro, item.Content, item.X1, item.Y1, item.X2, item.Y2, item.X3, item.Y3, item.X4, item.Y4, item.X5, item.Y5, item.X6, item.Y6, item.X7, item.Y7, item.X8, item.Y8, item.X9, item.Y9, item.X10, item.Y10, item.X11, item.Y11, item.X12, item.Y12, item.X13, item.Y13, item.Status, item.Date)
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
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update company_tb set c_name = ?, c_companyno = ?, c_ceo = ?, c_tel = ?, c_email = ?, c_address = ?, c_addressetc = ?, c_type = ?, c_billingname = ?, c_billingtel = ?, c_billingemail = ?, c_bankname = ?, c_bankno = ?, c_businesscondition = ?, c_businessitem = ?, c_giro = ?, c_content = ?, c_x1 = ?, c_y1 = ?, c_x2 = ?, c_y2 = ?, c_x3 = ?, c_y3 = ?, c_x4 = ?, c_y4 = ?, c_x5 = ?, c_y5 = ?, c_x6 = ?, c_y6 = ?, c_x7 = ?, c_y7 = ?, c_x8 = ?, c_y8 = ?, c_x9 = ?, c_y9 = ?, c_x10 = ?, c_y10 = ?, c_x11 = ?, c_y11 = ?, c_x12 = ?, c_y12 = ?, c_x13 = ?, c_y13 = ?, c_status = ?, c_date = ? where c_id = ?"
	_, err := p.Exec(query , item.Name, item.Companyno, item.Ceo, item.Tel, item.Email, item.Address, item.Addressetc, item.Type, item.Billingname, item.Billingtel, item.Billingemail, item.Bankname, item.Bankno, item.Businesscondition, item.Businessitem, item.Giro, item.Content, item.X1, item.Y1, item.X2, item.Y2, item.X3, item.Y3, item.X4, item.Y4, item.X5, item.Y5, item.X6, item.Y6, item.X7, item.Y7, item.X8, item.Y8, item.X9, item.Y9, item.X10, item.Y10, item.X11, item.Y11, item.X12, item.Y12, item.X13, item.Y13, item.Status, item.Date, item.Id)
    
        
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

func (p *CompanyManager) UpdateTel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_tel = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateEmail(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_email = ? where c_id = ?"
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

func (p *CompanyManager) UpdateType(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_type = ? where c_id = ?"
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

func (p *CompanyManager) UpdateBankname(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_bankname = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBankno(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_bankno = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBusinesscondition(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_businesscondition = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateBusinessitem(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_businessitem = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateGiro(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_giro = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateContent(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_content = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX1(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x1 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY1(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y1 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX2(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x2 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY2(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y2 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX3(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x3 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY3(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y3 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX4(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x4 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY4(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y4 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX5(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x5 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY5(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y5 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX6(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x6 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY6(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y6 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX7(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x7 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY7(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y7 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX8(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x8 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY8(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y8 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX9(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x9 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY9(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y9 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX10(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x10 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY10(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y10 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX11(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x11 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY11(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y11 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX12(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x12 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY12(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y12 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateX13(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x13 = ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) UpdateY13(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y13 = ? where c_id = ?"
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



func (p *CompanyManager) IncreaseX1(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x1 = c_x1 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY1(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y1 = c_y1 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX2(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x2 = c_x2 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY2(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y2 = c_y2 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX3(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x3 = c_x3 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY3(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y3 = c_y3 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX4(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x4 = c_x4 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY4(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y4 = c_y4 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX5(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x5 = c_x5 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY5(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y5 = c_y5 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX6(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x6 = c_x6 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY6(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y6 = c_y6 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX7(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x7 = c_x7 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY7(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y7 = c_y7 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX8(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x8 = c_x8 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY8(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y8 = c_y8 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX9(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x9 = c_x9 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY9(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y9 = c_y9 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX10(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x10 = c_x10 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY10(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y10 = c_y10 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX11(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x11 = c_x11 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY11(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y11 = c_y11 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX12(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x12 = c_x12 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY12(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y12 = c_y12 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseX13(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_x13 = c_x13 + ? where c_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *CompanyManager) IncreaseY13(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update company_tb set c_y13 = c_y13 + ? where c_id = ?"
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
            "type":     company.GetType(p.Type),

    }
}

func (p *CompanyManager) ReadRow(rows *sql.Rows) *Company {
    var item Company
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Tel, &item.Email, &item.Address, &item.Addressetc, &item.Type, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Bankname, &item.Bankno, &item.Businesscondition, &item.Businessitem, &item.Giro, &item.Content, &item.X1, &item.Y1, &item.X2, &item.Y2, &item.X3, &item.Y3, &item.X4, &item.Y4, &item.X5, &item.Y5, &item.X6, &item.Y6, &item.X7, &item.Y7, &item.X8, &item.Y8, &item.X9, &item.Y9, &item.X10, &item.Y10, &item.X11, &item.Y11, &item.X12, &item.Y12, &item.X13, &item.Y13, &item.Status, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Tel, &item.Email, &item.Address, &item.Addressetc, &item.Type, &item.Billingname, &item.Billingtel, &item.Billingemail, &item.Bankname, &item.Bankno, &item.Businesscondition, &item.Businessitem, &item.Giro, &item.Content, &item.X1, &item.Y1, &item.X2, &item.Y2, &item.X3, &item.Y3, &item.X4, &item.Y4, &item.X5, &item.Y5, &item.X6, &item.Y6, &item.X7, &item.Y7, &item.X8, &item.Y8, &item.X9, &item.Y9, &item.X10, &item.Y10, &item.X11, &item.Y11, &item.X12, &item.Y12, &item.X13, &item.Y13, &item.Status, &item.Date)
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


func (p *CompanyManager) GetByCompanyno(companyno string, args ...interface{}) *Company {
    if companyno != "" {
        args = append(args, Where{Column:"companyno", Value:companyno, Compare:"="})        
    }
    
    items := p.Find(args)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}

func (p *CompanyManager) GetByName(name string, args ...interface{}) *Company {
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



