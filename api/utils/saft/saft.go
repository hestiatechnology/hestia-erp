package saft

import (
	"encoding/xml"
	"errors"
)

// UQ = Unique Constraint
// KR = Key Reference
var (
	ErrorUQAccountId                  = errors.New("saft: unique constraint violation - account id")
	ErrorUQCustomerId                 = errors.New("saft: unique constraint violation - customer id")
	ErrorUQSupplierId                 = errors.New("saft: unique constraint violation - supplier id")
	ErrorUQProductCode                = errors.New("saft: unique constraint violation - product code")
	ErrorUQJournalId                  = errors.New("saft: unique constraint violation - journal id")
	ErrorUQTransactionId              = errors.New("saft: unique constraint violation - transaction id")
	ErrorUQInvoiceNo                  = errors.New("saft: unique constraint violation - invoice no")
	ErrorUQDocumentNo                 = errors.New("saft: unique constraint violation - document no")
	ErrorUQWorkDocNo                  = errors.New("saft: unique constraint violation - work document no")
	ErrorKRGenLedgerEntriesSupplierID = errors.New("saft: key reference violation - general ledger entries supplier id")
	ErrorKRGenLedgerEntriesAccountID  = errors.New("saft: key reference violation - general ledger entries account id")
	ErrorKRGenLedgerEntriesCustomerID = errors.New("saft: key reference violation - general ledger entries customer id")
	ErrorKRInvoiceCustomerID          = errors.New("saft: key reference violation - invoice customer id")
	ErrorKRInvoiceProductCode         = errors.New("saft: key reference violation - invoice product code")
	ErrorKRStockMovementCustomerID    = errors.New("saft: key reference violation - stock movement customer id")
	ErrorKRStockMovementSupplierID    = errors.New("saft: key reference violation - stock movement supplier id")
	ErrorKRStockMovementProductCode   = errors.New("saft: key reference violation - stock movement product code")
	ErrorKRWorkDocumentCustomerID     = errors.New("saft: key reference violation - work document customer id")
	ErrorKRWorkDocumentProductCode    = errors.New("saft: key reference violation - work document product code")
)

func (a *AuditFile) Validate() error {
	// Validate the AuditFile
	if err := a.CheckConstraints(); err != nil {
		return err
	}
	return nil
}

func (a *AuditFile) CheckConstraints() error {
	// AccountIDConstraint
	accounts := make(map[SafptglaccountId]bool)
	for _, account := range a.MasterFiles.GeneralLedgerAccounts.Account {
		if _, ok := accounts[account.AccountId]; ok {
			return ErrorUQAccountId
		}
		accounts[account.AccountId] = true
	}

	// TODO: GroupingCodeConstraint
	// ????????? Unsure
	// Should just do a match, however PDF docs dont support this

	// CustomerIDConstraint
	customers := make(map[SafpttextTypeMandatoryMax30Car]bool)
	for _, customer := range a.MasterFiles.Customer {
		if _, ok := customers[customer.CustomerId]; ok {
			return ErrorUQCustomerId
		}
		customers[customer.CustomerId] = true
	}

	// SupplierIDConstraint
	suppliers := make(map[SafpttextTypeMandatoryMax30Car]bool)
	for _, supplier := range a.MasterFiles.Supplier {
		if _, ok := suppliers[supplier.SupplierId]; ok {
			return ErrorUQSupplierId
		}
		suppliers[supplier.SupplierId] = true
	}

	// ProductCodeConstraint
	products := make(map[SafpttextTypeMandatoryMax60Car]bool)
	for _, product := range a.MasterFiles.Product {
		if _, ok := products[product.ProductCode]; ok {
			return ErrorUQProductCode
		}
		products[product.ProductCode] = true
	}

	// GeneralLedgerEntriesDebitLineAccountIDConstraint
	for _, entry := range a.GeneralLedgerEntries.Journal {
		for _, line := range entry.Transaction {
			for _, debit := range line.Lines.DebitLine {
				if _, ok := accounts[debit.AccountId]; !ok {
					return ErrorKRGenLedgerEntriesAccountID
				}
			}
		}
	}

	// GeneralLedgerEntriesCreditLineAccountIDConstraint
	for _, entry := range a.GeneralLedgerEntries.Journal {
		for _, line := range entry.Transaction {
			for _, credit := range line.Lines.CreditLine {
				if _, ok := accounts[credit.AccountId]; !ok {
					return ErrorKRGenLedgerEntriesAccountID
				}
			}
		}
	}

	// GeneralLedgerEntriesCustomerIDConstraint
	for _, entry := range a.GeneralLedgerEntries.Journal {
		for _, line := range entry.Transaction {
			if line.CustomerId != nil && *line.CustomerId != "" {
				if _, ok := customers[*line.CustomerId]; !ok {
					return ErrorKRGenLedgerEntriesCustomerID
				}
			}
		}
	}

	// GeneralLedgerEntriesJournalIdConstraint
	journals := make(map[SafptjournalId]bool)
	for _, entry := range a.GeneralLedgerEntries.Journal {
		if _, ok := journals[entry.JournalId]; ok {
			return ErrorUQJournalId
		}
		journals[entry.JournalId] = true
	}

	// GeneralLedgerEntriesSupplierIDConstraint
	for _, entry := range a.GeneralLedgerEntries.Journal {
		for _, line := range entry.Transaction {
			if line.SupplierId != nil && *line.SupplierId != "" {
				if _, ok := suppliers[*line.SupplierId]; !ok {
					return ErrorKRGenLedgerEntriesSupplierID
				}
			}
		}
	}

	// GeneralLedgerEntriesTransactionIdConstraint
	transactions := make(map[SafpttransactionId]bool)
	for _, entry := range a.GeneralLedgerEntries.Journal {
		for _, line := range entry.Transaction {
			if _, ok := transactions[line.TransactionId]; ok {
				return ErrorUQTransactionId
			}
			transactions[line.TransactionId] = true
		}
	}

	// InvoiceNoConstraint
	invoices := make(map[string]bool)
	for _, invoice := range a.SourceDocuments.SalesInvoices.Invoice {
		if _, ok := invoices[invoice.InvoiceNo]; ok {
			return ErrorUQInvoiceNo
		}
		invoices[invoice.InvoiceNo] = true
	}

	// InvoiceCustomerIDConstraint
	for _, invoice := range a.SourceDocuments.SalesInvoices.Invoice {
		if _, ok := customers[invoice.CustomerId]; !ok {
			return ErrorKRInvoiceCustomerID
		}
	}

	// InvoiceProductCodeConstraint
	for _, invoice := range a.SourceDocuments.SalesInvoices.Invoice {
		for _, line := range invoice.Line {
			if _, ok := products[line.ProductCode]; !ok {
				return ErrorKRInvoiceProductCode
			}
		}
	}

	// DocumentNumberConstraint
	documents := make(map[string]bool)
	for _, stock := range a.SourceDocuments.MovementOfGoods.StockMovement {
		if _, ok := documents[stock.DocumentNumber]; !ok {
			return ErrorUQDocumentNo
		}
		documents[stock.DocumentNumber] = true
	}

	// StockMovementCustomerIDConstraint
	for _, stock := range a.SourceDocuments.MovementOfGoods.StockMovement {
		if _, ok := customers[*stock.CustomerId]; !ok {
			return ErrorKRStockMovementCustomerID
		}
	}

	// StockMovementSupplierIDConstraint
	for _, stock := range a.SourceDocuments.MovementOfGoods.StockMovement {
		if _, ok := suppliers[*stock.SupplierId]; !ok {
			return ErrorKRStockMovementSupplierID
		}
	}

	// StockMovementProductCodeConstraint
	for _, stock := range a.SourceDocuments.MovementOfGoods.StockMovement {
		for _, line := range stock.Line {
			if _, ok := products[line.ProductCode]; !ok {
				return ErrorKRStockMovementProductCode
			}
		}
	}

	// WorkDocumentDocumentNumberConstraint
	workDocs := make(map[string]bool)
	for _, workDoc := range a.SourceDocuments.WorkingDocuments.WorkDocument {
		if _, ok := workDocs[workDoc.DocumentNumber]; !ok {
			return ErrorUQWorkDocNo
		}
		workDocs[workDoc.DocumentNumber] = true
	}

	// WorkDocumentDocumentCustomerIDConstraint
	for _, workDoc := range a.SourceDocuments.WorkingDocuments.WorkDocument {
		if _, ok := customers[workDoc.CustomerId]; !ok {
			return ErrorKRWorkDocumentCustomerID
		}
	}

	// WorkDocumentDocumentProductCodeConstraint
	for _, workDoc := range a.SourceDocuments.WorkingDocuments.WorkDocument {
		for _, line := range workDoc.Line {
			if _, ok := products[line.ProductCode]; !ok {
				return ErrorKRWorkDocumentProductCode
			}
		}
	}

	return nil
}

func (a *AuditFile) ToXML() (string, error) {
	// Convert the AuditFile to XML
	if err := a.Validate(); err != nil {
		return "", err
	}

	out, err := xml.MarshalIndent(a, "", "    ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(out), nil
}
