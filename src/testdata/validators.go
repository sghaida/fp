package helpers_testData

type TTVldData struct {
	Name     string
	Data     interface{}
	HasError bool
}

func CreateVldData() []TTVldData {
	return []TTVldData{
		{
			Name:     "none empty string",
			Data:     "this is not empty",
			HasError: false,
		}, {
			Name: "struct with value",
			Data: struct {
				value string
			}{"have some value"},
			HasError: false,
		}, {
			Name:     "non empty struct by address",
			Data:     &struct{ value string }{value: "some data"},
			HasError: false,
		}, {
			Name:     "non empty map",
			Data:     map[interface{}]string{"somekey": "some data"},
			HasError: false,
		}, {
			Name:     "empty string",
			Data:     "",
			HasError: true,
		}, {
			Name:     "nil value",
			Data:     nil,
			HasError: true,
		}, {
			Name:     "empty struct",
			Data:     struct{}{},
			HasError: true,
		}, {
			Name:     "empty slice",
			Data:     []interface{}{},
			HasError: true,
		}, {
			Name:     "empty map",
			Data:     map[interface{}]string{},
			HasError: true,
		},
	}
}
