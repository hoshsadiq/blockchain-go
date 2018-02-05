package block

type Transaction struct {
    Sender    string  `json:"sender"`
    Recipient string  `json:"recipient"`
    Amount    float32 `json:"amount"`
}

func NewTransaction(sender string, recipient string, amount float32) *Transaction {
    return &Transaction{
        Sender:    sender,
        Recipient: recipient,
        Amount:    amount,
    }
}

func (tx *Transaction) GetSender() string {
    return tx.Sender
}

func (tx *Transaction) GetRecipient() string {
    return tx.Recipient
}

func (tx *Transaction) GetAmount() float32 {
    return tx.Amount
}
