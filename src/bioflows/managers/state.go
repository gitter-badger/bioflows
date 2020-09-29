package managers


type StateManager interface {
	Setup(map[string]interface{}) error
	GetStateByID(string) (interface{},error)
	SetStateByID(string,interface{}) error
}






