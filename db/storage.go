package db

import (
	"github.com/gohumble/cashu-feni/cashu"
	"github.com/gohumble/cashu-feni/lightning"
)

type MintStorage interface {
	GetUsedProofs() []cashu.Proof
	ProofsUsed([]string) []cashu.Proof
	StoreProof(proof cashu.Proof) error
	StorePromise(p cashu.Promise) error
	StoreScript(p cashu.P2SHScript) error
	StoreLightningInvoice(i lightning.Invoice) error
	GetLightningInvoice(hash string) (lightning.Invoice, error)
	UpdateLightningInvoice(hash string, issued bool) error
	Migrate(interface{}) error
}
