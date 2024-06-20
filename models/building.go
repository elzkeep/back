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

type Building struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Companyno                string `json:"companyno"`         
    Ceo                string `json:"ceo"`         
    Zip                string `json:"zip"`         
    Address                string `json:"address"`         
    Addressetc                string `json:"addressetc"`         
    Postzip                string `json:"postzip"`         
    Postaddress                string `json:"postaddress"`         
    Postaddressetc                string `json:"postaddressetc"`         
    Postname                string `json:"postname"`         
    Posttel                string `json:"posttel"`         
    Cmsnumber                string `json:"cmsnumber"`         
    Cmsbank                string `json:"cmsbank"`         
    Cmsaccountno                string `json:"cmsaccountno"`         
    Cmsconfirm                string `json:"cmsconfirm"`         
    Cmsstartdate                string `json:"cmsstartdate"`         
    Cmsenddate                string `json:"cmsenddate"`         
    Contractvolumn                Double `json:"contractvolumn"`         
    Receivevolumn                Double `json:"receivevolumn"`         
    Generatevolumn                Double `json:"generatevolumn"`         
    Sunlightvolumn                Double `json:"sunlightvolumn"`         
    Volttype                int `json:"volttype"`         
    Weight                Double `json:"weight"`         
    Totalweight                Double `json:"totalweight"`         
    Checkcount                int `json:"checkcount"`         
    Receivevolt                int `json:"receivevolt"`         
    Generatevolt                int `json:"generatevolt"`         
    Periodic                int `json:"periodic"`         
    Businesscondition                string `json:"businesscondition"`         
    Businessitem                string `json:"businessitem"`         
    Usage                string `json:"usage"`         
    District                string `json:"district"`         
    Check                int `json:"check"`         
    Checkpost                int `json:"checkpost"`         
    Score                Double `json:"score"`         
    Status                int `json:"status"`         
    Company                int64 `json:"company"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]interface{} `json:"extra"`
}


type BuildingManager struct {
    Conn    *sql.DB
    Tx    *sql.Tx    
    Result  *sql.Result
    Index   string
}

func (c *Building) AddExtra(key string, value interface{}) {    
	c.Extra[key] = value     
}

func NewBuildingManager(conn interface{}) *BuildingManager {
    var item BuildingManager

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

func (p *BuildingManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *BuildingManager) SetIndex(index string) {
    p.Index = index
}

func (p *BuildingManager) Exec(query string, params ...interface{}) (sql.Result, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Exec(query, params...)
    } else {
       return p.Tx.Exec(query, params...)    
    }
}

func (p *BuildingManager) Query(query string, params ...interface{}) (*sql.Rows, error) {
	//log.Println(query)
	//log.Println(params)
    if p.Conn != nil {
       return p.Conn.Query(query, params...)
    } else {
       return p.Tx.Query(query + " FOR UPDATE", params...)    
    }
}

func (p *BuildingManager) GetQuery() string {
    ret := ""

    str := "select b_id, b_name, b_companyno, b_ceo, b_zip, b_address, b_addressetc, b_postzip, b_postaddress, b_postaddressetc, b_postname, b_posttel, b_cmsnumber, b_cmsbank, b_cmsaccountno, b_cmsconfirm, b_cmsstartdate, b_cmsenddate, b_contractvolumn, b_receivevolumn, b_generatevolumn, b_sunlightvolumn, b_volttype, b_weight, b_totalweight, b_checkcount, b_receivevolt, b_generatevolt, b_periodic, b_businesscondition, b_businessitem, b_usage, b_district, b_check, b_checkpost, b_score, b_status, b_company, b_date from building_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BuildingManager) GetQuerySelect() string {
    ret := ""
    
    str := "select count(*) from building_tb "

    if p.Index == "" {
        ret = str
    } else {
        ret = str + " use index(" + p.Index + ") "
    }

    ret += "where 1=1 "
    

    return ret;
}

func (p *BuildingManager) Truncate() error {
     if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    query := "truncate building_tb "
    _, err := p.Exec(query)

    if err != nil {
        log.Println(err)
    }

    return nil
}

func (p *BuildingManager) Insert(item *Building) error {
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
        query = "insert into building_tb (b_id, b_name, b_companyno, b_ceo, b_zip, b_address, b_addressetc, b_postzip, b_postaddress, b_postaddressetc, b_postname, b_posttel, b_cmsnumber, b_cmsbank, b_cmsaccountno, b_cmsconfirm, b_cmsstartdate, b_cmsenddate, b_contractvolumn, b_receivevolumn, b_generatevolumn, b_sunlightvolumn, b_volttype, b_weight, b_totalweight, b_checkcount, b_receivevolt, b_generatevolt, b_periodic, b_businesscondition, b_businessitem, b_usage, b_district, b_check, b_checkpost, b_score, b_status, b_company, b_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Id, item.Name, item.Companyno, item.Ceo, item.Zip, item.Address, item.Addressetc, item.Postzip, item.Postaddress, item.Postaddressetc, item.Postname, item.Posttel, item.Cmsnumber, item.Cmsbank, item.Cmsaccountno, item.Cmsconfirm, item.Cmsstartdate, item.Cmsenddate, item.Contractvolumn, item.Receivevolumn, item.Generatevolumn, item.Sunlightvolumn, item.Volttype, item.Weight, item.Totalweight, item.Checkcount, item.Receivevolt, item.Generatevolt, item.Periodic, item.Businesscondition, item.Businessitem, item.Usage, item.District, item.Check, item.Checkpost, item.Score, item.Status, item.Company, item.Date)
    } else {
        query = "insert into building_tb (b_name, b_companyno, b_ceo, b_zip, b_address, b_addressetc, b_postzip, b_postaddress, b_postaddressetc, b_postname, b_posttel, b_cmsnumber, b_cmsbank, b_cmsaccountno, b_cmsconfirm, b_cmsstartdate, b_cmsenddate, b_contractvolumn, b_receivevolumn, b_generatevolumn, b_sunlightvolumn, b_volttype, b_weight, b_totalweight, b_checkcount, b_receivevolt, b_generatevolt, b_periodic, b_businesscondition, b_businessitem, b_usage, b_district, b_check, b_checkpost, b_score, b_status, b_company, b_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        res, err = p.Exec(query , item.Name, item.Companyno, item.Ceo, item.Zip, item.Address, item.Addressetc, item.Postzip, item.Postaddress, item.Postaddressetc, item.Postname, item.Posttel, item.Cmsnumber, item.Cmsbank, item.Cmsaccountno, item.Cmsconfirm, item.Cmsstartdate, item.Cmsenddate, item.Contractvolumn, item.Receivevolumn, item.Generatevolumn, item.Sunlightvolumn, item.Volttype, item.Weight, item.Totalweight, item.Checkcount, item.Receivevolt, item.Generatevolt, item.Periodic, item.Businesscondition, item.Businessitem, item.Usage, item.District, item.Check, item.Checkpost, item.Score, item.Status, item.Company, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        log.Println(err)
        p.Result = nil
    }

    return err
}

func (p *BuildingManager) Delete(id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

    query := "delete from building_tb where b_id = ?"
    _, err := p.Exec(query, id)

    
    return err
}

func (p *BuildingManager) DeleteWhere(args []interface{}) error {
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
                query += " and b_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and b_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and b_" + item.Column + " " + item.Compare + " ?"
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

    query = "delete from building_tb where " + query[5:]
    _, err := p.Exec(query, params...)

    
    return err
}

func (p *BuildingManager) Update(item *Building) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }
    
    
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }

	query := "update building_tb set b_name = ?, b_companyno = ?, b_ceo = ?, b_zip = ?, b_address = ?, b_addressetc = ?, b_postzip = ?, b_postaddress = ?, b_postaddressetc = ?, b_postname = ?, b_posttel = ?, b_cmsnumber = ?, b_cmsbank = ?, b_cmsaccountno = ?, b_cmsconfirm = ?, b_cmsstartdate = ?, b_cmsenddate = ?, b_contractvolumn = ?, b_receivevolumn = ?, b_generatevolumn = ?, b_sunlightvolumn = ?, b_volttype = ?, b_weight = ?, b_totalweight = ?, b_checkcount = ?, b_receivevolt = ?, b_generatevolt = ?, b_periodic = ?, b_businesscondition = ?, b_businessitem = ?, b_usage = ?, b_district = ?, b_check = ?, b_checkpost = ?, b_score = ?, b_status = ?, b_company = ?, b_date = ? where b_id = ?"
	_, err := p.Exec(query , item.Name, item.Companyno, item.Ceo, item.Zip, item.Address, item.Addressetc, item.Postzip, item.Postaddress, item.Postaddressetc, item.Postname, item.Posttel, item.Cmsnumber, item.Cmsbank, item.Cmsaccountno, item.Cmsconfirm, item.Cmsstartdate, item.Cmsenddate, item.Contractvolumn, item.Receivevolumn, item.Generatevolumn, item.Sunlightvolumn, item.Volttype, item.Weight, item.Totalweight, item.Checkcount, item.Receivevolt, item.Generatevolt, item.Periodic, item.Businesscondition, item.Businessitem, item.Usage, item.District, item.Check, item.Checkpost, item.Score, item.Status, item.Company, item.Date, item.Id)
    
        
    return err
}


func (p *BuildingManager) UpdateName(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_name = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCompanyno(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_companyno = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCeo(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_ceo = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateZip(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_zip = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateAddress(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_address = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateAddressetc(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_addressetc = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdatePostzip(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_postzip = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdatePostaddress(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_postaddress = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdatePostaddressetc(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_postaddressetc = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdatePostname(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_postname = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdatePosttel(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_posttel = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCmsnumber(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_cmsnumber = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCmsbank(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_cmsbank = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCmsaccountno(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_cmsaccountno = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCmsconfirm(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_cmsconfirm = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCmsstartdate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_cmsstartdate = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCmsenddate(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_cmsenddate = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateContractvolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_contractvolumn = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateReceivevolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_receivevolumn = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateGeneratevolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_generatevolumn = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateSunlightvolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_sunlightvolumn = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateVolttype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_volttype = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateWeight(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_weight = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateTotalweight(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_totalweight = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCheckcount(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_checkcount = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateReceivevolt(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_receivevolt = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateGeneratevolt(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_generatevolt = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdatePeriodic(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_periodic = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateBusinesscondition(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_businesscondition = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateBusinessitem(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_businessitem = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateUsage(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_usage = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateDistrict(value string, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_district = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCheck(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_check = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCheckpost(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_checkpost = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateScore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_score = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_status = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) UpdateCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_company = ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}



func (p *BuildingManager) IncreaseContractvolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_contractvolumn = b_contractvolumn + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseReceivevolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_receivevolumn = b_receivevolumn + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseGeneratevolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_generatevolumn = b_generatevolumn + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseSunlightvolumn(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_sunlightvolumn = b_sunlightvolumn + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseVolttype(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_volttype = b_volttype + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseWeight(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_weight = b_weight + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseTotalweight(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_totalweight = b_totalweight + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseCheckcount(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_checkcount = b_checkcount + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseReceivevolt(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_receivevolt = b_receivevolt + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseGeneratevolt(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_generatevolt = b_generatevolt + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreasePeriodic(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_periodic = b_periodic + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseCheck(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_check = b_check + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseCheckpost(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_checkpost = b_checkpost + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseScore(value Double, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_score = b_score + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseStatus(value int, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_status = b_status + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}

func (p *BuildingManager) IncreaseCompany(value int64, id int64) error {
    if p.Conn == nil && p.Tx == nil {
        return errors.New("Connection Error")
    }

	query := "update building_tb set b_company = b_company + ? where b_id = ?"
	_, err := p.Exec(query, value, id)

    return err
}


func (p *BuildingManager) GetIdentity() int64 {
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

func (p *Building) InitExtra() {
    p.Extra = map[string]interface{}{

    }
}

func (p *BuildingManager) ReadRow(rows *sql.Rows) *Building {
    var item Building
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Zip, &item.Address, &item.Addressetc, &item.Postzip, &item.Postaddress, &item.Postaddressetc, &item.Postname, &item.Posttel, &item.Cmsnumber, &item.Cmsbank, &item.Cmsaccountno, &item.Cmsconfirm, &item.Cmsstartdate, &item.Cmsenddate, &item.Contractvolumn, &item.Receivevolumn, &item.Generatevolumn, &item.Sunlightvolumn, &item.Volttype, &item.Weight, &item.Totalweight, &item.Checkcount, &item.Receivevolt, &item.Generatevolt, &item.Periodic, &item.Businesscondition, &item.Businessitem, &item.Usage, &item.District, &item.Check, &item.Checkpost, &item.Score, &item.Status, &item.Company, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *BuildingManager) ReadRows(rows *sql.Rows) []Building {
    var items []Building

    for rows.Next() {
        var item Building
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Companyno, &item.Ceo, &item.Zip, &item.Address, &item.Addressetc, &item.Postzip, &item.Postaddress, &item.Postaddressetc, &item.Postname, &item.Posttel, &item.Cmsnumber, &item.Cmsbank, &item.Cmsaccountno, &item.Cmsconfirm, &item.Cmsstartdate, &item.Cmsenddate, &item.Contractvolumn, &item.Receivevolumn, &item.Generatevolumn, &item.Sunlightvolumn, &item.Volttype, &item.Weight, &item.Totalweight, &item.Checkcount, &item.Receivevolt, &item.Generatevolt, &item.Periodic, &item.Businesscondition, &item.Businessitem, &item.Usage, &item.District, &item.Check, &item.Checkpost, &item.Score, &item.Status, &item.Company, &item.Date)
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

func (p *BuildingManager) Get(id int64) *Building {
    if p.Conn == nil && p.Tx == nil {
        return nil
    }

    query := p.GetQuery() + " and b_id = ?"

    
    
    rows, err := p.Query(query, id)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *BuildingManager) Count(args []interface{}) int {
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
                query += " and b_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and b_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and b_" + item.Column + " " + item.Compare + " ?"
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

func (p *BuildingManager) FindAll() []Building {
    return p.Find(nil)
}

func (p *BuildingManager) Find(args []interface{}) []Building {
    if p.Conn == nil && p.Tx == nil {
        var items []Building
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
                query += " and b_" + item.Column + " in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and b_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and b_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "b_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "b_" + orderby
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
            orderby = "b_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "b_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(baseQuery + query, params...)

    if err != nil {
        log.Printf("query error : %v, %v\n", err, query)
        var items []Building
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *BuildingManager) GetByCompanyName(company int64, name string, args ...interface{}) *Building {
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


func (p *BuildingManager) Sum(args []interface{}) *Building {
    if p.Conn == nil && p.Tx == nil {
        var item Building
        return &item
    }

    var params []interface{}

    
    query := "select sum(b_score) from building_tb"

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
                query += " and b_id in (" + strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]") + ")"
            } else if item.Compare == "between" {
                query += " and b_" + item.Column + " between ? and ?"

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                query += " and b_" + item.Column + " " + item.Compare + " ?"
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
            orderby = "b_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                orderby = "b_" + orderby
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
            orderby = "b_id"
        } else {
            if !strings.Contains(orderby, "_") {
                orderby = "b_" + orderby
            }
        }
        query += " order by " + orderby
    }

    rows, err := p.Query(query, params...)

    var item Building
    
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
