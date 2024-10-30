package models

type OpenPosition struct {
	Symbol     string  `json:"symbol"`
	Quantity   int     `json:"quantity"`
	EntryPrice float64 `json:"entryPrice"`
}

type User struct {
	Id                string         `json:"id"`
	UserName          string         `json:"userName"`
	Balance           float64        `json:"balance"`
	AvailMargin       float64        `json:"availMargin"`
	TradingPermission []string       `json:"tradingPermission"`
	OpenPosition      []OpenPosition `json:"openPositioin"`
	RiskTolerance     string         `json:"riskTolerence"`
	PositionLimit     int            `json:"positionLimit"`
	ExposureLimit     int            `json:"exposureLimit"`
}


// In memory users
var Users = []User{
	{
		Id:                "user001",
		UserName:          "user1",
		Balance:           1000.0,
		TradingPermission: []string{"AAPL", "GOOG", "MSFT"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "AAPL",
				Quantity:   50,
				EntryPrice: 160.00,
			},
		},
		RiskTolerance: "medium",
		PositionLimit: 1000,
		ExposureLimit: 50000,
	},
	{
		Id:                "user002",
		UserName:          "user2",
		Balance:           2000.0,
		TradingPermission: []string{"AMZN", "TSLA", "FB"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "TSLA",
				Quantity:   20,
				EntryPrice: 620.00,
			},
		},
		RiskTolerance: "high",
		PositionLimit: 2000,
		ExposureLimit: 10000,
	},
	{
		Id:                "user003",
		UserName:          "user3",
		Balance:           1500.0,
		TradingPermission: []string{"NFLX", "NVDA", "INTC"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "NFLX",
				Quantity:   10,
				EntryPrice: 500.00,
			},
		},
		RiskTolerance: "low",
		PositionLimit: 500,
		ExposureLimit: 3000,
	},
	{
		Id:                "user004",
		UserName:          "user4",
		Balance:           2500.0,
		TradingPermission: []string{"AMD", "AAPL", "NVDA"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "AMD",
				Quantity:   30,
				EntryPrice: 90.00,
			},
		},
		RiskTolerance: "medium",
		PositionLimit: 1500,
		ExposureLimit: 7000,
	},
	{
		Id:                "user005",
		UserName:          "user5",
		Balance:           1800.0,
		TradingPermission: []string{"GOOG", "TSLA", "MSFT"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "GOOG",
				Quantity:   5,
				EntryPrice: 2500.00,
			},
		},
		RiskTolerance: "high",
		PositionLimit: 1000,
		ExposureLimit: 6000,
	},
	{
		Id:                "user006",
		UserName:          "user6",
		Balance:           1200.0,
		TradingPermission: []string{"FB", "NFLX", "AMZN"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "FB",
				Quantity:   40,
				EntryPrice: 300.00,
			},
		},
		RiskTolerance: "medium",
		PositionLimit: 1200,
		ExposureLimit: 5000,
	},
	{
		Id:                "user007",
		UserName:          "user7",
		Balance:           3000.0,
		TradingPermission: []string{"NVDA", "INTC", "AMD"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "NVDA",
				Quantity:   15,
				EntryPrice: 700.00,
			},
		},
		RiskTolerance: "low",
		PositionLimit: 700,
		ExposureLimit: 4000,
	},
	{
		Id:                "user008",
		UserName:          "user8",
		Balance:           5000.0,
		TradingPermission: []string{"AAPL", "GOOG", "MSFT"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "MSFT",
				Quantity:   25,
				EntryPrice: 280.00,
			},
		},
		RiskTolerance: "high",
		PositionLimit: 2500,
		ExposureLimit: 12000,
	},
	{
		Id:                "user009",
		UserName:          "user9",
		Balance:           2200.0,
		TradingPermission: []string{"AMZN", "TSLA", "FB"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "AMZN",
				Quantity:   8,
				EntryPrice: 3500.00,
			},
		},
		RiskTolerance: "medium",
		PositionLimit: 1000,
		ExposureLimit: 8000,
	},
	{
		Id:                "user010",
		UserName:          "user10",
		Balance:           1100.0,
		TradingPermission: []string{"NFLX", "NVDA", "INTC"},
		OpenPosition: []OpenPosition{
			{
				Symbol:     "INTC",
				Quantity:   60,
				EntryPrice: 60.00,
			},
		},
		RiskTolerance: "low",
		PositionLimit: 600,
		ExposureLimit: 4000,
	},
}
