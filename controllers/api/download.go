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

	session := c.Session

	companyManager := models.NewCompanyManager(conn)
	billingManager := models.NewBillingManager(conn)
	customerManager := models.NewCustomerManager(conn)
	userManager := models.NewUserManager(conn)

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

	my := companyManager.Get(session.Company)
	today := global.GetDate(time.Now())

	yRatio := 2.9

	for _, v := range items {
		building := v.Extra["building"].(models.Building)
		company := v.Extra["company"].(models.Company)

		customer := customerManager.GetByCompanyBuilding(session.Company, building.Id)

		billdate := ""
		temp := strings.Split(v.Billdate, "-")
		year := global.Atoi(temp[0])
		month := global.Atoi(temp[1])

		//lastday := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

		if v.Period == 1 {
			billdate = fmt.Sprintf("%v/%v", year, month)
		} else {
			billdate = fmt.Sprintf("%v/%v~%v", year, month, month+v.Period-1)
		}

		pdf.SetFont("noto", "", 13)

		vat := int(v.Price / 11)
		onlyPrice := vat * 10

		pdf.AddPage()

		zip := building.Postzip
		address := building.Postaddress
		if address == "" {
			zip = building.Zip
			address = building.Address
		}

		postname := building.Postname
		if postname == "" {
			postname = customer.Billingname
		}

		marginY := -1.4
		pdf.SetXY(yRatio*80, 2.8*80+marginY)
		pdf.Cell(nil, address)

		pdf.SetXY(yRatio*80, 2.8*90+marginY)
		if customer.Managername != "" {
			pdf.Cell(nil, fmt.Sprintf("%v (%v) 귀하", postname, customer.Managername))
		} else {
			pdf.Cell(nil, fmt.Sprintf("%v 귀하", postname))
		}

		pdf.SetXY(yRatio*160, 2.8*100+marginY)
		pdf.Cell(nil, zip)

		pdf.SetFont("noto", "", 9)

		pdf.SetXY(yRatio*19, 2.8*110+marginY)
		pdf.Cell(nil, building.Name)

		if customer.User > 0 {
			user := userManager.Get(customer.User)

			if user != nil {
				pdf.SetXY(yRatio*170, 2.8*110+marginY)
				pdf.Cell(nil, fmt.Sprintf("(%v %v)", user.Id, user.Name))
			}
		}

		pdf.SetFont("noto", "", 13)

		//pdf.RectFromLowerLeftWithStyle(yRatio*(177-50), 2.8*(225+8)+marginY, 2.8*50, 2.8*8, "DF")

		pdf.SetXY(yRatio*(177-50), 2.8*(225)+marginY)
		pdf.CellWithOption(&gopdf.Rect{W: 2.8 * 50, H: 2.8 * 8}, humanize.Comma(int64(v.Price)), gopdf.CellOption{
			Align:  gopdf.Right | gopdf.Top,
			Border: 0,
			Float:  gopdf.Right,
		})

		//pdf.SetXY(yRatio*177, 2.8*225+marginY)
		//pdf.Cell(nil, humanize.Comma(int64(v.Price)))

		// 좌측

		pdf.SetFont("noto", "", 10)

		//pdf.SetXY(yRatio*24, 2.8*143+marginY)
		//pdf.Cell(nil, humanize.Comma(int64(v.Price)))

		//pdf.RectFromLowerLeftWithStyle(yRatio*(24-14), 2.8*(143+4)+marginY, 2.8*27, 2.8*4, "D")
		pdf.SetXY(yRatio*(24-14), 2.8*(143)+marginY)
		pdf.CellWithOption(&gopdf.Rect{W: 2.8 * 27, H: 2.8 * 4}, humanize.Comma(int64(v.Price)), gopdf.CellOption{
			Align:  gopdf.Right | gopdf.Top,
			Border: 0,
			Float:  gopdf.Right,
		})

		pdf.SetFont("noto", "", 8)

		pdf.SetXY(yRatio*39, 2.8*143+marginY)
		pdf.Cell(nil, billdate)

		// 중앙 컨텐츠
		pdf.SetFont("noto", "", 8)

		contents := strings.Split(my.Content, "\n")
		for i, content := range contents {
			pdf.SetXY(yRatio*75, 2.8*float64(140+i*6)+marginY)
			pdf.Cell(nil, content)
		}

		pdf.SetFont("noto", "", 10)

		pdf.SetXY(yRatio*25, 2.8*148+marginY)
		pdf.Cell(nil, fmt.Sprintf("%v", customer.Number))

		pdf.SetXY(yRatio*17, 2.8*152.5+marginY)
		pdf.Cell(nil, humanize.Comma(int64(onlyPrice)))

		pdf.SetXY(yRatio*44, 2.8*152.5+marginY)
		pdf.Cell(nil, humanize.Comma(int64(vat)))

		pdf.SetXY(yRatio*23.5, 2.8*157.5+marginY)
		pdf.Cell(nil, fmt.Sprintf("%v", company.Companyno))

		pdf.SetXY(yRatio*12, 2.8*162+marginY)
		pdf.Cell(nil, company.Name)

		pdf.SetXY(yRatio*23.5, 2.8*171.5+marginY)
		pdf.Cell(nil, today)

		// 하단

		pdf.SetXY(yRatio*73, 2.8*256.5+marginY)
		pdf.Cell(nil, fmt.Sprintf("%v", customer.Number))

		pdf.SetXY(yRatio*140, 2.8*256.5+marginY)
		pdf.Cell(nil, fmt.Sprintf("%v", month))

		pdf.SetXY(yRatio*156, 2.8*256.5+marginY)
		pdf.Cell(nil, fmt.Sprintf("%v", customer.Collectday))

		pdf.SetXY(yRatio*73, 2.8*263+marginY)
		pdf.Cell(nil, company.Name)

		pdf.SetXY(yRatio*73, 2.8*269.5+marginY)
		pdf.Cell(nil, company.Ceo)

		pdf.SetXY(yRatio*73, 2.8*276+marginY)
		pdf.Cell(nil, billdate)

		// OCR

		pdf.SetFont("ocr", "", 12)

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

		//companyNo := 1000000000 + int64(user.Company)*100000 + int64(customer.Number)
		companyNo := 1000000000 + v.Id

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

		str := fmt.Sprintf("<%v+%v+%v%v+ %v+%v%v< <11<", my.Giro, spaces2, strCompanyNo, digit2, spaces, strPrice, digit)
		pdf.SetXY(yRatio*68, 2.8*241.5+marginY)
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
	excel := global.NewExcel("고객 현황", "", 12, header, width, align)
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
	excel := global.NewExcel("소속회원 현황", "", 12, header, width, align)
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
	os.Remove(fullFilename)
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

func (c *DownloadController) All(category int) {
	log.Println("======================================")
	log.Println("typeid", category)
	log.Println("======================================")

	nodata := false
	if category > 2 {
		nodata = true
		category -= 3
	}
	conn := c.NewConnection()

	session := c.Session

	departmentManager := models.NewDepartmentManager(conn)
	userlistManager := models.NewUserlistManager(conn)
	customerManager := models.NewCustomerManager(conn)
	licenseManager := models.NewLicenseManager(conn)

	departments := departmentManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
	})

	items := customerManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Ordering("cu_number,cu_id"),
	})

	users := userlistManager.Find([]interface{}{
		models.Where{Column: "company", Value: session.Company, Compare: "="},
		models.Ordering("u_name"),
	})

	excel := global.New()

	if category != 2 {
		contracttypes := []string{"", "안전관리", "유지보수", "안전관리+유지보수"}

		header := []string{"고객코드", "고객명", "사업자번호", "대표자", "기본주소", "상세주소", "연락처", "이메일",
			"점검건물명", "사업자번호", "대표자", "우편번호", "기본주소", "상세주소",
			"업태", "종목", "건물용도",
			"계약형태", "계약용량", "점검자", "영업자", "계약일자", "계약만료일",
			"관할구청", "한전 고객번호", "안전공사 고객번호", "정기점검 주기", "최종검사일",

			"(관리)담당자", "(관리)담당자 연락처", "(관리)담당자 Email",
			"(계약)담당자", "(계약)담당자 연락처", "(계약)책임자 Email",
			"우편수령지 우편번호", "우편수령지", "우편수령지 수신자", "우편수령지 전화번호", "Fax No.",
			"대행수수료", "부가세", "계산서 발행일", "청구방법", "수금일", "특이사항"}
		width := []int{10, 30, 15, 50, 40, 40, 20, 30,
			30, 20, 15, 12, 50, 40,
			20, 20, 20,
			20, 20, 20, 20, 20, 20,
			20, 15, 15, 15, 20,

			20, 20, 50,
			20, 20, 50,
			20, 50, 15, 20, 20,
			15, 15, 20, 15, 20, 50}
		align := []string{"L", "L", "L", "L", "L", "L", "L", "L",
			"L", "L", "L", "L", "L", "L",
			"L", "L", "L",
			"C", "R", "L", "L", "C", "C",
			"L", "L", "L", "L", "C",

			"L", "L", "L",
			"L", "L", "L",
			"L", "L", "L", "L", "L",
			"R", "R", "C", "C", "C", "L"}
		excel.NewSheet("고객 현황", header, width, align)
		excel.SetHeight(24)

		excel.InsertRow(1, 1)
		excel.SetRowHeight(1, 24)
		excel.MergeCell("A", 1, "A", 2)
		excel.MergeCell("B", 1, "H", 1)
		excel.MergeCell("I", 1, "AS", 1)
		excel.SetHeaderStyle("A", 1, 10)
		excel.SetHeaderStyle("B", 1, 10)
		excel.SetHeaderStyle("I", 1, 10)
		excel.SetHeaderStyle("AS", 1, 10)
		excel.SetCellValue("A", 1, "고객코드")
		excel.SetCellValue("B", 1, "고객정보")
		excel.SetCellValue("I", 1, "점건건물정보")

		excel.Rows++

		if nodata == false {
			for _, v := range items {
				company := v.Extra["company"].(models.Company)
				building := v.Extra["building"].(models.Building)
				excel.CellInt(v.Number)
				excel.Cell(company.Name)
				excel.Cell(company.Companyno)
				excel.Cell(company.Ceo)
				excel.Cell(company.Address)
				excel.Cell(company.Addressetc)
				excel.Cell(company.Tel)
				excel.Cell(company.Email)

				excel.Cell(building.Name)
				excel.Cell(building.Companyno)
				excel.Cell(building.Ceo)
				excel.Cell(building.Zip)
				excel.Cell(building.Address)
				excel.Cell(building.Addressetc)

				excel.Cell(building.Businesscondition)
				excel.Cell(building.Businessitem)

				excel.Cell(building.Usage)

				excel.Cell(contracttypes[v.Contracttype])
				excel.Cell(humanize.FormatFloat("#,###.#", float64(building.Totalweight)))

				username := ""
				for _, user := range users {
					if user.Id == v.User {
						username = user.Name
						break
					}
				}
				excel.Cell(username)

				saileusername := ""
				for _, user := range users {
					if user.Id == v.Salesuser {
						saileusername = user.Name
						break
					}
				}
				excel.Cell(saileusername)

				excel.Cell(v.Contractstartdate)
				excel.Cell(v.Contractenddate)
				excel.Cell(building.District)
				excel.Cell(v.Kepconumber)
				excel.Cell(v.Kesconumber)

				excel.Cell(v.Periodic)
				excel.Cell(v.Lastdate)

				excel.Cell(v.Managername)
				excel.Cell(v.Managertel)
				excel.Cell(v.Manageremail)
				excel.Cell(v.Billingname)
				excel.Cell(v.Billingtel)
				excel.Cell(v.Billingemail)

				excel.Cell(building.Postzip)
				excel.Cell(building.Postaddress)
				excel.Cell(building.Postname)
				excel.Cell(building.Posttel)

				excel.Cell(v.Fax)
				excel.Cell(fmt.Sprintf("%v", v.Contractprice))
				excel.Cell(fmt.Sprintf("%v", v.Contractvat))

				excel.Cell(fmt.Sprintf("%v일", v.Billingdate))
				if v.Billingtype == 1 {
					excel.Cell("지로")
				} else {
					excel.Cell("계산서")
				}

				month := ""
				if v.Collectmonth == 1 {
					month = "매월"
				} else {
					month = "익월"
				}

				excel.Cell(fmt.Sprintf("%v %v일", month, v.Collectday))

				excel.Cell(v.Remark)
			}
		}
	}

	if category != 1 {
		header := []string{"팀", "로그인아이디", "이름", "이메일", "연락처", "주소", "권한", "상태", "점수", "입사일", "기술자격", "등록번호", "기술자격등급", "기술자격취득일자"}
		width := []int{30, 20, 15, 50, 30, 80, 15, 15, 15, 30, 50, 50, 20, 30}
		align := []string{"L", "L", "L", "L", "L", "L", "L", "C", "R", "C", "L", "L", "C", "C"}
		excel.NewSheet("소속회원", header, width, align)
		excel.SetHeight(24)

		if nodata == false {
			for _, v := range users {
				departmentName := ""

				for _, department := range departments {
					if department.Id == v.Department {
						departmentName = department.Name
						break
					}
				}

				excel.Cell(departmentName)
				excel.Cell(v.Loginid)
				excel.Cell(v.Name)
				excel.Cell(v.Email)
				excel.Cell(v.Tel)
				excel.Cell(fmt.Sprintf("%v %v", v.Address, v.Addressetc))
				excel.Cell(user.GetLevel(user.Level(v.Level)))
				excel.Cell(user.GetStatus(user.Status(v.Status)))
				excel.Cell(fmt.Sprintf("%v", v.Score))
				excel.Cell(v.Joindate)

				licenses := licenseManager.FindByUser(session.Id)
				if len(licenses) == 0 {
					excel.Cell("")
					excel.Cell("")
					excel.Cell("")
					excel.Cell("")
				} else {
					category := ""
					level := ""
					no := ""
					date := ""
					for i, license := range licenses {
						if i > 0 {
							category += "\n"
							level += "\n"
							no += "\n"
							date += "\n"
						}
						category += license.Extra["licensecategory"].(models.Licensecategory).Name
						level += license.Extra["licenselevel"].(models.Licenselevel).Name
						no += license.Number
						date += license.Takingdate
					}

					excel.Cell(category)
					excel.Cell(no)
					excel.Cell(level)
					excel.Cell(date)
				}
			}
		}
	}

	fullFilename := excel.Save("")
	log.Println("filename", fullFilename)

	c.Download(fullFilename, "user.xlsx")
	os.Remove(fullFilename)
}
