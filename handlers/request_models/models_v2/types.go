package models_v2

type CalculatorAddParams struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Product struct {
	Id   int    `json:"id" xml:"Id" yaml:"id"`
	Name string `json:"name" xml:"Name" yaml:"name"`
}

type PrintJob struct {
	JobId int `json:"jobId" binding:"required,gte=10000"`
	Pages int `json:"pages" binding:"required,gte=1,lte=100"`
}

type PrintInvoiceJob struct {
	JobId     int    `json:"jobId" binding:"gte=0"`
	Format    string `json:"format" binding:"required"`
	InvoiceId int    `json:"invoiceId" binding:"required,gte=1"`
}
