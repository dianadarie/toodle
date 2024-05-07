package sheets

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
	"path/filepath"
)

type GoogleSheets struct {
	Server *sheets.Service
}

type UserCredentials struct {
	Username   string
	Password   string
	IsApproved string
}

func NewSheetsClient() (GoogleSheets, error) {
	ctx := context.Background()

	pwd, _ := os.Getwd()
	credBytes, err := os.ReadFile(filepath.Join(pwd, sheetsCredsPath))
	if err != nil {
		return GoogleSheets{}, err
	}

	config, err := google.JWTConfigFromJSON(credBytes, sheetsURL)
	if err != nil {
		return GoogleSheets{}, err
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return GoogleSheets{}, err
	}

	return GoogleSheets{Server: srv}, nil
}

func (s *GoogleSheets) GetSheetNameByID(ID int64) (string, error) {
	r, err := s.Server.Spreadsheets.Get(spreadsheetID).Fields("sheets(properties(sheetId,title))").Do()
	if err != nil {
		return "", err
	}

	sheetName := ""
	for _, v := range r.Sheets {
		prop := v.Properties
		sheetID := prop.SheetId
		if sheetID == ID {
			sheetName = prop.Title
			break
		}
	}

	return sheetName, nil
}

func (s *GoogleSheets) GetAllSheetsIDs() ([]int64, error) {
	var sheetIDs []int64
	r, err := s.Server.Spreadsheets.Get(spreadsheetID).Fields("sheets(properties(sheetId))").Do()
	if err != nil {
		return sheetIDs, err
	}

	for _, v := range r.Sheets {
		sheetIDs = append(sheetIDs, v.Properties.SheetId)
	}

	return sheetIDs, nil
}
