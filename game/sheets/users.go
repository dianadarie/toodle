package sheets

import (
	"context"
	"google.golang.org/api/sheets/v4"
)

func (s *GoogleSheets) AddUserToSheet(username, password string) error {
	ctx := context.Background()

	authSheet, err := s.GetSheetNameByID(sheetAuthenticationID)
	if err != nil {
		return err
	}

	row := &sheets.ValueRange{
		Values: [][]interface{}{{username, password, "nope"}},
	}

	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"
	res, err := s.Server.Spreadsheets.Values.Append(spreadsheetID, authSheet, row).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		return err
	}

	return nil
}

func (s *GoogleSheets) IsUser(username, password string) (bool, error) {
	authSheet, err := s.GetSheetNameByID(sheetAuthenticationID)
	if err != nil {
		return false, err
	}

	readRange := authSheet + sheetRange
	res, err := s.Server.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return false, err
	}

	for _, row := range res.Values {
		if row[0] == username && row[1] == password && row[2] == "yes" {
			return true, nil
		}
	}
	return false, nil
}
