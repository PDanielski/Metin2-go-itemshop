package wallet

//Currency is a currency that can be used in the shop
type Currency struct {
	Code     string
	Name     string
	IconLink string
}

func NewCurrency(code string, name string, iconLink string) *Currency {
	return &Currency{code, name, iconLink}
}

type CurrencyProvider interface {
	Provide() ([]*Currency, error)
}

type SimpleCurrencyProvider struct {
	currencies []*Currency
}

func (p *SimpleCurrencyProvider) Provide() ([]*Currency, error) {
	return p.currencies, nil
}

func NewSimpleCurrencyProvider(c []*Currency) CurrencyProvider {
	return &SimpleCurrencyProvider{c}
}

//Wallet manages the currencies of an account
type Wallet struct {
	accountID uint
	handler   *Handler
}

//New creates a new Wallet. How it manages the operations depends on the handler instance given
func New(accountID uint, handler *Handler) *Wallet {
	return &Wallet{accountID, handler}
}

//AccountID returns the account id
func (wallet *Wallet) AccountID() uint {
	return wallet.accountID
}

//Handler defines the behavior of the wallet
type Handler interface {
	Withdraw(currency *Currency, amount int) error
	Deposit(currency *Currency, amount int) error
	Balance(currency *Currency) (int, error)
}
