package models

type Transaction struct {
	UUID 	  uint64  `json:"_id"`
	Origin    Wallet  `json:"origin"`
	Target 	  Wallet  `json:"target"`
	Cash	  float32 `json:"cash"`
	Message   string  `json:"message"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func NewTransaction(transaction Transaction) (bool, error) {
	con := Connect()
	defer con.Close()
	tx, err := con.Begin()
	if err != nil {
		return false, err
	}
	sql := "update wallet set balance = (balance + $1) where public_key = $2"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(transaction.Cash, transaction.Target.PublicKey)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	sql = "update wallet set balance = (balance - $1) where public_key = $2"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(transaction.Cash, transaction.Origin.PublicKey)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	sql = "insert into transactions (origin, target, cash, message) values ($1, $2, $3, $4)"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(transaction.Origin.PublicKey, transaction.Target.PublicKey,
		transaction.Cash, transaction.Message)
		if err != nil {
			tx.Rollback()
			return false, err
		}	
	}
	return true, tx.Commit()
}

func GetTransactions() ([]Transaction, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from transactions"
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var transactions []Transaction
	for rs.Next() {
		var transaction Transaction
		err := rs.Scan(&transaction.UUID, &transaction.Origin.PublicKey,
		&transaction.Target.PublicKey, &transaction.Cash, &transaction.Message,
		&transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}