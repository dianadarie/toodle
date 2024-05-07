package sheets

type Solution struct {
	date         string
	username     string
	attempt1     string
	attempt2     string
	attempt3     string
	attempt4     string
	attempt5     string
	noOfAttempts string
}

func (s *GoogleSheets) GetUserSolution(username, date string) (string, error) {
	solution, err := s.GetSheetNameByID(sheetSolutionsID)
	if err != nil {
		return "", err
	}

	readRange := solution + sheetRange
	res, err := s.Server.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return "", err
	}

	for _, row := range res.Values {
		if row[0] == date && row[1] == solution {
			val, _ := row[1].(string)
			return val, nil
		}
	}

	return "", nil
}
