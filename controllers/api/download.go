package api

import (
	"fmt"
	"log"
	"os"
	"time"
	"zkeep/controllers"
	"zkeep/global"
	"zkeep/models"
	"zkeep/models/billing"

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

	my := companyManager.Get(user.Company)
	today := global.GetDate(time.Now())

	for _, v := range items {
		company := v.Extra["company"].(models.Company)
		pdf.SetFont("noto", "", 14)

		vat := int(v.Price / 10)

		pdf.AddPage()
		pdf.SetXY(2.8*90, 2.8*80)
		pdf.Cell(nil, fmt.Sprintf("%v %v", company.Address, company.Addressetc))
		pdf.SetXY(2.8*90, 2.8*90)
		pdf.Cell(nil, fmt.Sprintf("%v 귀하", company.Name))

		pdf.SetXY(2.8*70, 2.8*140)
		pdf.Cell(nil, "입금 계좌")

		pdf.SetXY(2.8*70, 2.8*150)
		pdf.Cell(nil, my.Bankname)

		pdf.SetXY(2.8*70, 2.8*158)
		pdf.Cell(nil, my.Bankno)

		pdf.SetXY(2.8*155, 2.8*223)
		pdf.Cell(nil, humanize.Comma(int64(v.Price+vat)))

		pdf.SetFont("noto", "", 10)

		pdf.SetXY(2.8*14, 2.8*140)
		pdf.Cell(nil, humanize.Comma(int64(v.Price+vat)))

		pdf.SetXY(2.8*12, 2.8*150)
		pdf.Cell(nil, humanize.Comma(int64(v.Price)))

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
		pdf.Cell(nil, v.Billdate)
	}

	fullFilename := global.GetTempFilename()
	log.Println("fullFilename", fullFilename)

	pdf.WritePdf(fullFilename)

	c.Download(fullFilename, "hello.pdf")
	os.Remove(fullFilename)
}
