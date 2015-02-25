package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAddPairing(t *testing.T) {
	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TLVType_Method, TLVType_Method_PairingAdd)
	tlv8.SetByte(TLVType_SequenceNumber, 0x01)
	tlv8.SetString(TLVType_Username, "Unit Test")
	tlv8.SetBytes(TLVType_PublicKey, []byte{0x01, 0x02})

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	tlv8_out, err := controller.Handle(tlv8)
	assert.Nil(t, err)
	assert.NotNil(t, tlv8_out)
	assert.Equal(t, tlv8_out.GetByte(TLVType_SequenceNumber), byte(0x2))
}

func TestDeletePairing(t *testing.T) {
	username := "Unit Test"
	client := db.NewClient(username, []byte{0x01, 0x02})
	database, _ := db.NewDatabase(os.TempDir())
	database.SaveClient(client)

	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TLVType_Method, TLVType_Method_PairingDelete)
	tlv8.SetByte(TLVType_SequenceNumber, 0x01)
	tlv8.SetString(TLVType_Username, username)

	controller := NewPairingController(database)

	tlv8_out, err := controller.Handle(tlv8)
	assert.Nil(t, err)
	assert.NotNil(t, tlv8_out)
	assert.Equal(t, tlv8_out.GetByte(TLVType_SequenceNumber), byte(0x2))

	saved_client := database.ClientWithName(username)
	assert.Nil(t, saved_client)
}
