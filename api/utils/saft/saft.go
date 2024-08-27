package saft

import (
	"encoding/xml"
	"errors"
)

var (
	ErrorUniqueConstraintViolation = errors.New("unique constraint violation")
)

func (a *AuditFile) Validate() error {
	// Validate the AuditFile
	return nil
}

func (a *AuditFile) CheckUniqueConstraint() error {
	// Check for UQ in MasterFiles/GeneralLedgerAccounts/Account/AccountID
	accounts := make(map[SafptglaccountId]bool)
	for _, account := range a.MasterFiles.GeneralLedgerAccounts.Account {
		if _, ok := accounts[account.AccountId]; ok {
			return ErrorUniqueConstraintViolation
		}
		accounts[account.AccountId] = true
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
