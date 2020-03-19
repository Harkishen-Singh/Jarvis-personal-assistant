package controllers

type response struct {
	username string
	message  string
}

type statusCode struct {
	status string
}

type messageQueryBody struct {
	Head     string `json:"head"`
	Link     string `json:"link"`
	Desc     string `json:"desc"`
	DescLink string `json:"dlink"`
}

type reminderResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Result  []reminder `json:"result"`
}

type jsonResponseQuery struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Result  []messageQueryBody `json:"result"`
}

type jsonResponseWeather struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Result  weatherStr `json:"result"`
}

type weatherStr struct {
	Time        string `json:"time"`
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	DewPoint    string `json:"dew_point"`
	Humidity    string `json:"humidity"`
	Visibility  string `json:"visibility"`
	FeelsLike   string `json:"feels_like"`
}

type meaningStr struct {
	Meaning    string       `json:"meaning"`
	Example    string       `json:"example"`
	Submeaning []submeanStr `json:"submeaning"`
}

type submeanStr struct {
	Smean      string
	Subexample string
}

type jsonResponseMeaning struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Result  []meaningStr `json:"result"`
}
