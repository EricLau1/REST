package models

import (
	"api/utils"
)

type Wallet struct {
	PublicKey string `json:"public_key"`
	User 	User	 `json:"user"`		
	Balance float32  `json:"balance"`
}

func (w *Wallet) generatePublicKey() {
	w.PublicKey = utils.Md5(w.User.Nickname + w.User.Password)
}

func GetWallets() ([]Wallet, error) {
	con := Connect()
	defer con.Close()
	sql := `
		select u.uuid, u.nickname, u.email, u.password, u.status, u.created_at, u.updated_at,
		w.public_key, w.usr, w.balance
		from wallet as w 
		inner join users as u on u.uuid = w.usr order by w.usr asc
	`
	rs, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var wallets []Wallet
	for rs.Next() {
		var wallet Wallet
		err := rs.Scan(&wallet.User.UUID, 
		&wallet.User.Nickname, &wallet.User.Email, &wallet.User.Password, &wallet.User.Status,
		&wallet.User.CreatedAt, &wallet.User.UpdatedAt,
		&wallet.PublicKey, &wallet.User.UUID, &wallet.Balance)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}
	return wallets, nil
}

func GetWalletByPublicKey(publicKey string) (Wallet, error) {
	con := Connect()
	defer con.Close()
	sql := `
		select u.uuid, u.nickname, u.email, u.password, u.status, u.created_at, u.updated_at,
		w.public_key, w.usr, w.balance
		from wallet as w 
		inner join users as u on u.uuid = w.usr
		where w.public_key = $1
	`
	rs, err := con.Query(sql, publicKey)
	if err != nil {
		return Wallet{}, err
	}
	defer rs.Close()
	var wallet Wallet
	if rs.Next() {
		err := rs.Scan(&wallet.User.UUID, 
		&wallet.User.Nickname, &wallet.User.Email, &wallet.User.Password, &wallet.User.Status,
		&wallet.User.CreatedAt, &wallet.User.UpdatedAt,
		&wallet.PublicKey, &wallet.User.UUID, &wallet.Balance)
		if err != nil {
			return Wallet{}, err
		}
	}
	return wallet, nil
}

func UpdateWallet(wallet Wallet) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "update wallet set balance = $1 where public_key = $2"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(wallet.Balance, wallet.PublicKey)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func AddCashWallet(wallet Wallet) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "update wallet set balance = (balance + $1) where public_key = $2"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(wallet.Balance, wallet.PublicKey)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}
