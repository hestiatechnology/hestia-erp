/* package pdf
 */package main

import (
	"hestia/api/utils/logger"
	qrcode "hestia/api/utils/pdf/codes/at-qrcode"
	"os"

	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("billingv2.pdf") //.Save("billingv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("billingv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithLeftMargin(5).
		WithTopMargin(0).
		WithRightMargin(5).
		WithBottomMargin(6).
		WithPageNumber().
		Build()

	darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	img, err := os.ReadFile("FB_Cover.png")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	err = m.RegisterHeader(getPageHeader(img, extension.Png)...)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(getPageFooter())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(text.NewRow(10, "Invoice ABC123456789", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))

	m.AddRow(7,
		text.NewCol(3, "Transactions", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: darkGrayColor})

	m.AddRow(5,
		//col.New(3),
		text.NewCol(2, "Código", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		text.NewCol(2, "Cod. Barras", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		text.NewCol(3, "Artigo", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		text.NewCol(3, "Uni", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		text.NewCol(3, "Quant.", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
	)

	transactions := getTransactions()

	for i, transaction := range transactions {
		logger.DebugLogger.Println(transaction)
		if i%2 != 0 {
			gray := getGrayColor()
			m.AddRow(4, transaction...).WithStyle(&props.Cell{
				BackgroundColor: gray,
				BorderType:      border.Full,
				BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 0},
				LineStyle:       linestyle.Solid,
				BorderThickness: 0.25,
			})

		} else {
			m.AddRow(4, transaction...).WithStyle(&props.Cell{
				BorderType:      border.Full,
				BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 0},
				LineStyle:       linestyle.Solid,
				BorderThickness: 0.25,
			})
		}
	}

	/* m.AddRow(15,
		col.New(6).Add(
			code.NewBar("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			}),
			text.New("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  fontstyle.Bold,
				Size:   9,
				Align:  align.Center,
			}),
		),
		col.New(6),
	) */
	return m
}

func getTransactions() [][]core.Col {
	colStyle := &props.Cell{
		//BackgroundColor: &props.Color{Red: 80, Green: 80, Blue: 80},

	}
	/* 	rows := []core.Row{
		row.New(5).Add(
			//col.New(3),
			text.NewCol(1, "Código", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Cod. Barras", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Artigo", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Uni", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Quant.", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		),
	} */

	/* 	var contentsRow []core.Row */
	contents := getContents()
	//for i := 0; i < 8; i++ {
	//	contents = append(contents, contents...)
	//}

	var cols [][]core.Col
	for _, content := range contents {
		logger.DebugLogger.Println(content)
		col := []core.Col{
			text.NewCol(2, content[0]+"\n ", props.Text{Size: 8, Align: align.Center}).WithStyle(colStyle),
			text.NewCol(2, content[1]+"\n ", props.Text{Size: 8, Align: align.Center}).WithStyle(colStyle),
			text.NewCol(3, content[2]+"\n ", props.Text{Size: 8, Align: align.Center}).WithStyle(colStyle),
			text.NewCol(3, content[3]+"\n ", props.Text{Size: 8, Align: align.Center}).WithStyle(colStyle),
			text.NewCol(3, content[3]+"\n ", props.Text{Size: 8, Align: align.Center}).WithStyle(colStyle),
		}
		cols = append(cols, col)
	}

	/* 		if i%2 == 0 {
		gray := getGrayColor()
		r.WithStyle(&props.Cell{BackgroundColor: gray})
	} */

	/* contentsRow = append(contentsRow, r) */

	/* 	rows = append(rows, contentsRow...)

	   	rows = append(rows, row.New(20).Add(
	   		col.New(7),
	   		text.NewCol(2, "Total:", props.Text{
	   			Top:   5,
	   			Style: fontstyle.Bold,
	   			Size:  8,
	   			Align: align.Right,
	   		}),
	   		text.NewCol(3, "R$ 2.567,00", props.Text{
	   			Top:   5,
	   			Style: fontstyle.Bold,
	   			Size:  8,
	   			Align: align.Center,
	   		}),
	   	)) */
	return cols
}

func getPageHeader(img []byte, ext extension.Type) []core.Row {
	var rows []core.Row
	r := row.New(10).Add(

		col.New(3).Add(
			image.NewFromBytes(img, ext, props.Rect{
				Top:     2,
				Percent: 100,
			}),
		),

		col.New(3),
		col.New(6).Add(
			text.New("Fatura-Recibo FR 201298343/1", props.Text{
				Top:   3,
				Style: fontstyle.Bold,
				Size:  13,
				Align: align.Right,
			}),

			//code.NewQr("ABCD23EF-12345", props.Rect{}),
		),
	)
	rows = append(rows, r)

	r = text.NewRow(4, "Hestia Technology & Daniel Pereira Socios e filhos lololololololololololololololololo Unipessoal Lda", props.Text{
		Top:   2,
		Style: fontstyle.Bold,
		Size:  10,
		Align: align.Left,
	})

	rows = append(rows, r)

	r = text.NewRow(4, "1234 Street Name, City Name, Country Name", props.Text{
		Top:   2,
		Style: fontstyle.Bold,
		Size:  8,
		Align: align.Left,
	})

	rows = append(rows, r)

	r = text.NewRow(4, "Tel: 55 024 12345-1234", props.Text{
		Top:   2,
		Style: fontstyle.Bold,
		Size:  8,
		Align: align.Left,
	})

	rows = append(rows, r)

	r = row.New(10).Add(
		col.New(12).Add(
			text.New("Hestia Technology & Daniel Pereira Socios e filhos lololololololololololololololololo Unipessoal Lda", props.Text{
				Top:   2,
				Style: fontstyle.Bold,
				Size:  10,
				Align: align.Right,
			}),
		),
	)

	rows = append(rows, r)
	return rows
}

func getPageFooter() core.Row {
	return row.New(50).Add(
		col.New(9).Add(
			text.New("Tel: 55 024 12345-1234", props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
			text.New("www.mycompany.com", props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
		),
		col.New(4).Add(
			text.New("ATCUD: IUSDIUSFD-02389", props.Text{
				Style: fontstyle.Bold,
				Size:  8,
				Align: align.Center,
			}),
			image.NewFromBytes(qrcode.Qr(), extension.Jpeg, props.Rect{
				Percent: 80,
				Center:  true,
			}),
		),
	)
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

//func getRedColor() *props.Color {
//	return &props.Color{
//		Red:   150,
//		Green: 10,
//		Blue:  10,
//	}
//}

func getContents() [][]string {
	return [][]string{
		{"72347823746", "1234567890123", "Swamp", "12", "R$ 4,00"},
		/* {"23453675869", "1234567890123", "Sorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A PlaneswalkerSorin, A Planeswalker", "4", "R$ 90,00"}, */
		{"23453675869", "1234567890123", "Tassa", "4", "R$ 30,00"},
		{"23453675869", "1234567890123", "Skinrender", "4", "R$ 9,00"},
		{"23453675869", "1234567890123", "Island", "12", "R$ 4,00"},
		{"23453675869", "1234567890123", "Mountain", "12", "R$ 4,00"},
		{"23453675869", "1234567890123", "Plain", "12", "R$ 4,00"},
		{"23453675869", "1234567890123", "Black Lotus", "1", "R$ 1.000,00"},
		{"23453675869", "1234567890123", "Time Walk", "1", "R$ 1.000,00"},
		{"23453675869", "1234567890123", "Emberclave", "4", "R$ 44,00"},
		{"23453675869", "1234567890123", "Anax", "4", "R$ 32,00"},
		{"23453675869", "1234567890123", "Murderous Rider", "4", "R$ 22,00"},
		{"23453675869", "1234567890123", "dfsdjlsbfrhasdasdaaaaaa", "4", "R$ 2,00"},
		{"23453675869", "1234567890123", "Ajani's Pridemate", "4", "R$ 2,00"},
		{"23453675869", "1234567890123", "Renan, Chatuba", "4", "R$ 19,00"},
		{"23453675869", "1234567890123", "Tymarett", "4", "R$ 13,00"},
		{"23453675869", "1234567890123", "Doom Blade", "4", "R$ 5,00"},
		{"23453675869", "1234567890123", "Dark Lord", "3", "R$ 7,00"},
		{"23453675869", "1234567890123", "Memory of Thanatos", "3", "R$ 32,00"},
		{"23453675869", "1234567890123", "Poring", "4", "R$ 1,00"},
		{"23453675869", "1234567890123", "Deviling", "4", "R$ 99,00"},
		{"23453675869", "1234567890123", "Seiya", "4", "R$ 45,00"},
		{"23453675869", "1234567890123", "Harry Potter", "4", "R$ 62,00"},
		{"23453675869", "1234567890123", "Goku", "4", "R$ 77,00"},
		{"23453675869", "1234567890123", "Phreoni", "4", "R$ 22,00"},
		{"23453675869", "1234567890123", "Katheryn High Wizard", "4", "R$ 25,00"},
		{"23453675869", "1234567890123", "Lord Seyren", "4", "R$ 55,00"},
	}
}
