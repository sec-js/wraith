package libwraith

// Constants for describing module types to the module tree
const (
	ModCommsChanTx modtype = iota
	ModCommsChanRx
	ModProtoLang
	ModProtoPart
)

type modtype int

// Every module must implement this interface to make sure
// it's meant to be used as a Wraith module
type GenericModule interface {
	WraithModuleInit(*Wraith)
}

type CommsChanTxModule interface {
	GenericModule
	StartTx()
	StopTx()
	TriggerTx(addr string, data []byte) bool
	CommsChanTxModule()
}

type CommsChanRxModule interface {
	GenericModule
	StartRx()
	StopRx()
	CommsChanRxModule()
}

type ProtoLangModule interface {
	GenericModule
	Encode(map[string]interface{}) ([]byte, error)
	Decode([]byte) (map[string]interface{}, error)
	Identify([]byte) bool
	ProtoLangModule()
}

type ProtoPartModule interface {
	GenericModule
	ProcessProtoPart(*Wraith, *HandlerKeyValueStore, interface{}) (string, map[string]interface{}, string)
	ProtoPartModule()
}
