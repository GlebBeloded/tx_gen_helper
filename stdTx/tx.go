package stdTx

type Value struct {
	Msg        []Msg       `json:"msg"`
	Fee        Fee         `json:"fee"`
	Signatures []Signature `json:"signatures"`
	Memo       string      `json:"memo"`
}

type Tx struct {
	Msg        []Msg       `json:"msg"`
	Fee        Fee         `json:"fee"`
	Signatures []Signature `json:"signatures"`
	Memo       string      `json:"memo"`
}

type Msg interface{}

type Fee struct {
	Amount []Coin `json:"amount"`
	Gas    int    `json:"gas",String`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount int    `json:"amount",String`
}

type Signature struct {
	Pubkey    PubKey `json:"pub_key"`
	Signature string `json:"signature"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewTx(message Msg) Tx {
	return Tx{
		Msg:        []Msg{message},
		Fee:        Fee{Amount: []Coin{}, Gas: 200000},
		Signatures: nil,
	}
}
