package view

import "github.com/rlaaudgjs5638/dsk/client"

type ViewModel interface {
	InteractWithClient() client.ClientModel
}
