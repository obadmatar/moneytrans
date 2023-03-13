package account

type Repositoy interface{
	List() ([]*Account, error)

	Get(id string) (*Account, error)

	Update(account *Account) error
}