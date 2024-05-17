package api

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/billing"
	"zkeep/models/user"

	"github.com/dustin/go-humanize"
	"github.com/signintech/gopdf"
)

type DownloadController struct {
	controllers.Controller
}

func (c *DownloadController) File(id int64) {
	/*
		conn := c.NewConnection()

		manager := models.NewFileManager(conn)
		item := manager.Get(id)

		fullFilename := fmt.Sprintf("%v/%v", config.UploadPath, item.Filename)
		c.Download(fullFilename, item.Originalfilename)
	*/
}

func (c *DownloadController) Giro(ids []int64) {
	conn := c.NewConnection()

	log.Println("Print==================")
	log.Println(ids)

	user := c.Session

	companyManager := models.NewCompanyManager(conn)
	billingManager := models.NewBillingManager(conn)
	customerManager := models.NewCustomerManager(conn)

	items := billingManager.Find([]interface{}{
		models.Where{Column: "id", Value: ids, Compare: "in"},
	})

	for _, v := range items {
		v.Giro = billing.GiroComplete
		billingManager.Update(&v)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	err := pdf.AddTTFFont("noto", "./fonts/noto.ttf")
	if err != nil {
		log.Println("error")
		log.Print(err.Error())
		return
	}

	err = pdf.AddTTFFont("ocr", "./fonts/OCR-B1.ttf")
	if err != nil {
		log.Println("error")
		log.Print(err.Error())
		return
	}

	my := companyManager.Get(user.Company)
	today := global.GetDate(time.Now())

	for _, v := range items {
		company := v.Extra["company"].(models.Company)
		building := v.Extra["building"].(models.Building)

		pdf.SetFont("noto", "", 14)

		vat := int(v.Price / 11)
		onlyPrice := vat * 10

		pdf.AddPage()
		pdf.SetXY(2.8*90, 2.8*80)
		pdf.Cell(nil, fmt.Sprintf("%v %v", company.Address, company.Addressetc))
		pdf.SetXY(2.8*90, 2.8*90)
		pdf.Cell(nil, fmt.Sprintf("%v 귀하", company.Name))

		//pdf.SetXY(2.8*70, 2.8*140)
		//pdf.Cell(nil, "입금 계좌")

		//pdf.SetXY(2.8*70, 2.8*150)
		//pdf.Cell(nil, my.Bankname)

		//pdf.SetXY(2.8*70, 2.8*158)
		//pdf.Cell(nil, my.Bankno)

		pdf.SetXY(2.8*155, 2.8*223)
		pdf.Cell(nil, humanize.Comma(int64(v.Price)))

		pdf.SetFont("noto", "", 10)

		pdf.SetXY(2.8*14, 2.8*140)
		pdf.Cell(nil, humanize.Comma(int64(v.Price)))

		pdf.SetXY(2.8*12, 2.8*150)
		pdf.Cell(nil, humanize.Comma(int64(onlyPrice)))

		pdf.SetXY(2.8*39, 2.8*150)
		pdf.Cell(nil, humanize.Comma(int64(vat)))

		pdf.SetXY(2.8*24, 2.8*154)
		pdf.Cell(nil, fmt.Sprintf("%v", company.Id))

		pdf.SetXY(2.8*12, 2.8*158)
		pdf.Cell(nil, company.Name)

		pdf.SetXY(2.8*24, 2.8*164)
		pdf.Cell(nil, today)

		pdf.SetXY(2.8*73, 2.8*254)
		pdf.Cell(nil, fmt.Sprintf("%v", company.Id))

		pdf.SetXY(2.8*73, 2.8*260)
		pdf.Cell(nil, company.Name)

		pdf.SetXY(2.8*73, 2.8*266)
		pdf.Cell(nil, company.Ceo)

		pdf.SetXY(2.8*73, 2.8*272)
		temp := strings.Split(v.Billdate, "-")
		month := global.Atoi(temp[1])

		if v.Period == 1 {
			billdate := fmt.Sprintf("%v", month)
			pdf.Cell(nil, billdate)
		} else {
			billdate := fmt.Sprintf("%v ~ %v", month, month+v.Period-1)
			pdf.Cell(nil, billdate)
		}

		pdf.SetFont("orc", "", 12)

		price := v.Price
		sum := 0

		muls := []int{7, 3, 1}
		mulsPos := 0
		for i := 1; i <= 7; i++ {
			remain := price % 10

			sum += remain * muls[mulsPos]

			mulsPos++
			if mulsPos == 3 {
				mulsPos = 0
			}

			price -= remain
			price /= 10

			if price == 0 {
				break
			}
		}

		digit := 0
		if sum < 10 {
			digit = 10 - sum
		} else {
			sum = sum % 10
			if sum > 0 {
				digit = 10 - sum
			}
		}

		strPrice := global.Itoa(v.Price)
		spaces := strings.Repeat(" ", 10-(len(strPrice)+1))

		customer := customerManager.GetByCompanyBuilding(user.Company, building.Id)
		log.Println(user.Company, building.Id)
		log.Println("customer number", customer.Number)
		companyNo := 1000000000 + int64(user.Company)*100000 + int64(customer.Number)

		sum = 0

		muls = []int{2, 1}
		mulsPos = 0
		for i := 1; i <= 7; i++ {
			remain := price % 10

			temp := remain * muls[mulsPos]

			if temp > 10 {
				sum += temp%10 + 1
			} else {
				sum += temp
			}

			mulsPos++
			if mulsPos == 2 {
				mulsPos = 0
			}

			price -= remain
			price /= 10

			if price == 0 {
				break
			}
		}

		digit2 := 0
		if sum < 10 {
			digit2 = 10 - sum
		} else {
			sum = sum % 10
			if sum > 0 {
				digit2 = 10 - sum
			}
		}

		strCompanyNo := fmt.Sprintf("%v", companyNo)
		spaces2 := strings.Repeat(" ", 20-(len(strCompanyNo)+1))

		str := fmt.Sprintf("<%v+ %v+%v%v+ %v+%v%v< <11<", my.Giro, spaces2, strCompanyNo, digit2, spaces, strPrice, digit)
		pdf.SetXY(2.8*73, 2.8*237)
		pdf.Cell(nil, str)
	}

	fullFilename := global.GetTempFilename()
	log.Println("fullFilename", fullFilename)

	pdf.WritePdf(fullFilename)

	c.Download(fullFilename, "hello.pdf")
	os.Remove(fullFilename)
}

func (c *DownloadController) Company() {
	conn := c.NewConnection()

	user := c.Session
	companylistManager := models.NewCompanylistManager(conn)
	items := companylistManager.Find([]interface{}{
		models.Where{Column: "company", Value: user.Company, Compare: "="},
		models.Ordering("c_name"),
	})

	header := []string{"번호", "사업자명", "대표자", "사업자번호", "주소", "보유 건물수", "계약총액", "등록일"}
	width := []int{25, 70, 25, 40, 100, 20, 30, 40}
	align := []string{"C", "L", "L", "L", "L", "R", "R", "C"}
	excel := global.NewExcel("고객 현황", header, width, align)
	excel.SetHeight(30)

	for _, v := range items {
		excel.CellInt64(v.Id)
		excel.Cell(v.Name)
		excel.Cell(v.Ceo)
		excel.Cell(v.Companyno)
		excel.Cell(fmt.Sprintf("%v %v", v.Address, v.Addressetc))
		excel.CellInt64(v.Buildingcount)
		excel.CellInt(v.Contractprice)
		excel.Cell(v.Date)
	}

	fullFilename := excel.Save("")
	log.Println("filename", fullFilename)

	c.Download(fullFilename, "company.xlsx")
	//os.Remove(fullFilename)
}

func (c *DownloadController) User() {
	conn := c.NewConnection()

	session := c.Session

	departmentManager := models.NewDepartmentManager(conn)
	departments := departmentManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
	})

	userlistManager := models.NewUserlistManager(conn)
	items := userlistManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Ordering("u_name"),
	})

	header := []string{"팀", "로그인아이디", "이름", "이메일", "연락처", "주소", "권한", "상태", "점수", "등록일"}
	width := []int{30, 20, 15, 50, 30, 80, 15, 15, 15, 30}
	align := []string{"L", "L", "L", "L", "L", "L", "L", "C", "R", "C"}
	excel := global.NewExcel("소속회원 현황", header, width, align)
	excel.SetHeight(30)

	for _, v := range items {
		str := ""
		for _, v2 := range departments {
			if v.Department == v2.Id {
				str = v.Name
				break
			}
		}

		excel.Cell(str)
		excel.Cell(v.Loginid)
		excel.Cell(v.Name)
		excel.Cell(v.Email)
		excel.Cell(v.Tel)
		excel.Cell(fmt.Sprintf("%v %v", v.Address, v.Addressetc))
		excel.Cell(user.GetLevel(user.Level(v.Level)))
		excel.Cell(user.GetStatus(user.Status(v.Status)))

		if v.Totalscore == 0 {
			excel.Cell(fmt.Sprintf("%v / %v", 0, v.Score))
		} else {
			excel.Cell(fmt.Sprintf("%v / %v", global.ToFixed(float64(v.Totalscore), 1), v.Score))
		}

		excel.Cell(v.Date)
	}

	fullFilename := excel.Save("")
	log.Println("filename", fullFilename)

	c.Download(fullFilename, "user.xlsx")
	//os.Remove(fullFilename)
}

func (c *DownloadController) CompanyExample() {
	c.Download("./doc/company.xlsx", "company.xlsx")
}

func (c *DownloadController) CustomerExample() {
	c.Download("./doc/customer.xlsx", "customer.xlsx")
}

func (c *DownloadController) UserExample() {
	c.Download("./doc/user.xlsx", "user.xlsx")
}
