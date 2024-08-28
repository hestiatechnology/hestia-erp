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
	ErrorKRGenLedgerEntriesAccountID  = errors.New("saft: key reference violation - general ledger entries account id")
	ErrorKRGenLedgerEntriesCustomerID = errors.New("saft: key reference violation - general ledger entries customer id")
	ErrorUQJournalId                  = errors.New("saft: unique constraint violation - journal id")
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
