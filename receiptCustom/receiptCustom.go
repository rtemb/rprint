package receiptCustom


type PdfDocument struct {
	FileSetting *FileSetting
	ReceiptN	*ReceiptN

}

type ReceiptN struct {
	MPlaceName		*MPlaceName
	MPlaceAddress	*MPlaceAddress
	MPlaceINN		*MPlaceINN
	OperationType	*OperationType
}

type MPlaceName struct {
	Value		string `json:"Value"`
	LineSetting	*LineSetting
}

type MPlaceAddress struct {
	Value		string `json:"Value"`
	LineSetting	*LineSetting
}

type MPlaceINN struct {
	Value			string `json:"Value"`
	LineSetting	*LineSetting
}

type OperationType struct {
	Value		string `json:"Value"`
	LineSetting	*LineSetting
}

type LineSetting struct {
	Fsize	int		`json:"Fsize"`
	Fstyle	string	`json:"Fstyle"`
	PosX	int		`json:"PosX"`
	PosY	int		`json:"PosY"`
	Align	string	`json:"Align"`
}

type FileSetting struct {
	Format	string	`json:"Format"`
	ZeroX	int	`json:"ZeroX"`
	ZeroY	int	`json:"ZeroY"`
}