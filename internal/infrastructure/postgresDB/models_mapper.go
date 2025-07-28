package postgres

import (
	d "tg-bot/internal/domain/entities"
)

func toDBcl(cl *d.Client) *Client {
	c := &Client{
		Firstn: cl.Firstn,
		Lastn:  cl.Lastn,
		NID:    cl.NID,
	}

	return c
}

func fromDBcl(dbcl *Client) *d.Client {
	c := &d.Client{
		Firstn: dbcl.Firstn,
		Lastn:  dbcl.Lastn,
		NID:    dbcl.NID,
	}

	return c
}

func toDBacc(acc *d.Account) *Account {

	a := &Account{
		AID:     acc.AID,
		Balance: acc.Balance,
		NID:     acc.NID,
	}

	return a
}

func fromDBacc(dbacc *Account) *d.Account {

	a := &d.Account{
		AID:     dbacc.AID,
		Balance: dbacc.Balance,
		NID:     dbacc.NID,
	}

	return a
}

func toDBtr(tr *d.Transaction) *Transaction {
	t := &Transaction{
		AID: tr.AID,
		Sum: tr.Sum,
		TID: tr.TID,
	}

	return t
}

func fromDBtr(dbtr *Transaction) *d.Transaction {
	t := &d.Transaction{
		AID: dbtr.AID,
		Sum: dbtr.Sum,
		TID: dbtr.TID,
	}

	return t
}

func toDBfile(file *d.File) *File {
	f := &File{
		Filename: file.Filename,
		Hash:     file.Hash,
		AID:      file.AID,
		Bytes:    file.Bytes,
	}

	return f
}

func fromDBfile(dbfile *File) *d.File {
	f := &d.File{
		Filename: dbfile.Filename,
		Hash:     dbfile.Hash,
		AID:      dbfile.AID,
		Bytes:    dbfile.Bytes,
	}

	return f
}
