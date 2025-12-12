package regional

type Region string

const (
	// Global
	RegionISO Region = "iso" // YYYY-MM-DD HH:MM:SS
	RegionUS  Region = "us"  // MM/DD/YYYY HH:MM AM/PM
	RegionEU  Region = "eu"  // DD/MM/YYYY HH:MM
	RegionCA  Region = "ca"  // YYYY-MM-DD

	// ASEAN
	RegionID Region = "id" // DD MonthName YYYY (Indonesia)
	RegionTH Region = "th" // DD/MM/YYYY BE (Thailand)
	RegionVN Region = "vn" // DD/MM/YYYY (Vietnam)
	RegionMY Region = "my" // DD/MM/YYYY (Malaysia)
	RegionSG Region = "sg" // DD/MM/YYYY (Singapore)
	RegionPH Region = "ph" // MM/DD/YYYY (Philippines)

	// East Asia
	RegionJP Region = "jp" // YYYY/MM/DD
	RegionKR Region = "kr" // YYYY.MM.DD
	RegionCN Region = "cn" // YYYY-MM-DD
)

const (
	LangEN = "en"
	LangID = "id"
	LangTH = "th"
)
