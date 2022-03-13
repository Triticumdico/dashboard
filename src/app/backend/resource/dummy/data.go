package dummy

type TableRows struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	AccountID string `json:"account_id"`
	Score     int    `json:"score"`
}

func GetTableRows() (*TableRows, error) {
	tableRows := &TableRows{
		Rows: []Row{
			{
				AccountID: "test1",
				Score:     2},
			{
				AccountID: "test2",
				Score:     7},
		},
	}
	return tableRows, nil
}
