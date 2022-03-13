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
				AccountID: "account1",
				Score:     2},
			{
				AccountID: "account2",
				Score:     7},
			{
				AccountID: "account3",
				Score:     5},
			{
				AccountID: "account4",
				Score:     9},
			{
				AccountID: "account5",
				Score:     3},
		},
	}
	return tableRows, nil
}
