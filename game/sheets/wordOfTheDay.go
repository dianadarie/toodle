package sheets

import (
	"context"
	"google.golang.org/api/sheets/v4"
	"time"
)

func (s *GoogleSheets) GetWordGivenDay() (string, error) {
	wordOfDaySheet, err := s.GetSheetNameByID(sheetWordOfTheDayID)
	if err != nil {
		return "", err
	}

	readRange := wordOfDaySheet + sheetRange
	res, err := s.Server.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return "", err
	}

	day := time.Now().Format(TimeFormat)
	for _, row := range res.Values {
		if row[0] == day {
			val, _ := row[1].(string)
			return val, nil
		}
	}

	return "", nil
}

func (s *GoogleSheets) SetWordGivenDay(word string) error {
	ctx := context.Background()

	wordOfDaySheet, err := s.GetSheetNameByID(sheetWordOfTheDayID)
	if err != nil {
		return err
	}

	row := &sheets.ValueRange{
		Values: [][]interface{}{{time.Now().Format(TimeFormat), word}},
	}

	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"
	res, err := s.Server.Spreadsheets.Values.Append(spreadsheetID, wordOfDaySheet, row).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		return err
	}

	return nil
}
