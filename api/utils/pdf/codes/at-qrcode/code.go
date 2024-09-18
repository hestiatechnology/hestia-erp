package qrcode

import (
	"github.com/hestiatechnology/autoridadetributaria/saft"
)

// QR code mandated by the Autoridade Tributária.
// Must be on all fiscal documents
// https://info.portaldasfinancas.gov.pt/pt/apoio_contribuinte/Novas_regras_faturacao/Documents/Especificacoes_Tecnicas_Codigo_QR.pdf
type ATQRCode struct {
	// A: Issuer NIF
	A saft.SafptportugueseVatNumber
	// B: Client NIF
	B saft.SafptportugueseVatNumber
	// C: Client Country
	C saft.CustomerCountry
	// D: Invoice type
	D string
	// E: Invoice status
	E string
	// F: Invoice date
	F string
	// G: Invoice total
	G string
	// H: Invoice hash
	H string
	// I1: Invoice total
	I1 string
	// N: Tax total
	N string
	// O: Tax exempt total
	O string
	// Q: Tax rate
	Q string
	// R: Invoice total
	R string
}
