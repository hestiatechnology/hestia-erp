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
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var colStyle = &props.Cell{
	//BackgroundColor: &props.Color{Red: 80, Green: 80, Blue: 80},
	BorderType:      border.Bottom,
	BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 0},
	LineStyle:       linestyle.Solid,
	BorderThickness: 0.25,
}

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
		WithDefaultFont(&props.Font{
			Family: fontfamily.Helvetica,
		}).
		WithPageNumber().
		Build()

	// darkGrayColor := getDarkGrayColor()
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

	transactions := getTransactions()

	for _, transaction := range transactions {
		m.AddAutoRow(transaction...)
	}

	return m
}

func getTransactions() [][]core.Col {
	textStyle := props.Text{
		/* 		Top:   1, */
		Size:  8,
		Align: align.Center,
	}

	textStyleProduct := textStyle
	textStyleProduct.Align = align.Left
	textStyleProduct.Left = 1

	contents := getContents()

	var cols [][]core.Col
	for _, content := range contents {
		logger.DebugLogger.Println(content)

		col := []core.Col{
			text.NewCol(0, ""),
			text.NewCol(2, content[0]+"\n ", textStyle),
			text.NewCol(5, content[2]+"\n ", textStyleProduct),
			text.NewCol(1, "kilo", textStyle),
			text.NewCol(1, content[3]+"\n ", textStyle),
			text.NewCol(1, content[3]+"\n ", textStyle),
			text.NewCol(1, "23%", textStyle),
			text.NewCol(1, content[3]+"€ ", textStyle),
		}

		cols = append(cols, col)
	}
	return cols
}

func getPageHeader(img []byte, ext extension.Type) []core.Row {
	var rows []core.Row
	rows = append(rows, row.New(20).Add(

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
			text.New("Original", props.Text{
				Top:   8,
				Style: fontstyle.Bold,
				Align: align.Right,
				Size:  9,
			}),
			text.New("27/06/2004", props.Text{
				Top:   13,
				Style: fontstyle.Normal,
				Align: align.Right,
				Size:  7,
			}),
		),
	))

	rows = append(rows, text.NewRow(4, "Hestia Technology & Daniel Pereira Unipessoal Lda", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Size:  10,
		Align: align.Left,
	}))

	rows = append(rows, text.NewRow(4, "Rua de Rio Covo 62", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Left,
	}))

	rows = append(rows, text.NewRow(4, "4755-466 Rio Covo Santa Eugénia", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Left,
	}))

	rows = append(rows, text.NewRow(4, "Portugal", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Left,
	}))

	rows = append(rows, text.NewRow(4, "Tel: 55 024 12345-1234", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Left,
	}))

	rows = append(rows, text.NewRow(4, "Cliente ABC XYZ Unipessoal Lda", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Size:  10,
		Align: align.Right,
	}))

	rows = append(rows, text.NewRow(4, "Rua de Rio Covo 62", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Right,
	}))

	rows = append(rows, text.NewRow(4, "4755-466 Rio Covo Santa Eugénia", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Right,
	}))

	rows = append(rows, text.NewRow(4, "Portugal", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Right,
	}))
	// r =

	rows = append(rows, text.NewRow(4, "Tel: 55 024 12345-1234", props.Text{
		Top:   3,
		Style: fontstyle.Normal,
		Size:  8,
		Align: align.Right,
	}))

	rows = append(rows, row.New(7))
	//rows = append(rows, r)
	rows = append(rows, text.NewRow(6, "Transactions", props.Text{
		Top:   1.5,
		Size:  9,
		Style: fontstyle.Bold,
		Align: align.Center,
		Color: &props.WhiteColor,
	}).WithStyle(&props.Cell{BackgroundColor: getDarkGrayColor()}),
	)

	rows = append(rows, row.New(5).Add(
		//col.New(3),
		text.NewCol(2, "Código", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
		// text.NewCol(2, "Cod. Barras", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		text.NewCol(5, "Artigo", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
		text.NewCol(1, "Uni", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
		text.NewCol(1, "Quant.", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
		text.NewCol(1, "€/Uni", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
		text.NewCol(1, "IVA", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
		text.NewCol(1, "Total", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).WithStyle(colStyle),
	))
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

/*
	 func getGrayColor() *props.Color {
		return &props.Color{
			Red:   200,
			Green: 200,
			Blue:  200,
		}
	}
*/
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
		{"P029029", "1234567890123", "Aspirador Robô Xiaomi ", "12", "R$ 4,00"},
		{"P030592", "1234567890123", "Coluna Xiaomi Mi Smart ", "4", "R$ 90,00"},
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
	}
}
