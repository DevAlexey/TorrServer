package settings

var (
	tdb      *TDB
	Path     string
	ReadOnly bool
	PlaylistDir string
)

func InitSets(readOnly bool) {
	ReadOnly = readOnly
	tdb = NewTDB()
	loadBTSets()
	Migrate()
}

func CloseDB() {
	tdb.CloseDB()
}
