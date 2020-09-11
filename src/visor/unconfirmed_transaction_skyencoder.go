// Code generated by github.com/SkycoinProject/skyencoder. DO NOT EDIT.
package visor

import (
	"errors"
	"math"

	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/cipher/encoder"
	"github.com/SkycoinProject/skycoin/src/coin"
)

// encodeSizeUnconfirmedTransaction computes the size of an encoded object of type UnconfirmedTransaction
func encodeSizeUnconfirmedTransaction(obj *UnconfirmedTransaction) uint64 {
	i0 := uint64(0)

	// obj.Transaction.Length
	i0 += 4

	// obj.Transaction.Type
	i0++

	// obj.Transaction.InnerHash
	i0 += 32

	// obj.Transaction.Sigs
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 65

		i0 += uint64(len(obj.Transaction.Sigs)) * i1
	}

	// obj.Transaction.In
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 32

		i0 += uint64(len(obj.Transaction.In)) * i1
	}

	// obj.Transaction.Out
	i0 += 4
	{
		i1 := uint64(0)

		// x.Address.Version
		i1++

		// x.Address.Key
		i1 += 20

		// x.Coins
		i1 += 8

		// x.Hours
		i1 += 8

		i0 += uint64(len(obj.Transaction.Out)) * i1
	}

	i0 += 512

	// obj.Received
	i0 += 8

	// obj.Checked
	i0 += 8

	// obj.Announced
	i0 += 8

	// obj.IsValid
	i0++

	return i0
}

// encodeUnconfirmedTransaction encodes an object of type UnconfirmedTransaction to a buffer allocated to the exact size
// required to encode the object.
func encodeUnconfirmedTransaction(obj *UnconfirmedTransaction) ([]byte, error) {
	n := encodeSizeUnconfirmedTransaction(obj)
	buf := make([]byte, n)

	if err := encodeUnconfirmedTransactionToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeUnconfirmedTransactionToBuffer encodes an object of type UnconfirmedTransaction to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeUnconfirmedTransactionToBuffer(buf []byte, obj *UnconfirmedTransaction) error {
	if uint64(len(buf)) < encodeSizeUnconfirmedTransaction(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Transaction.Length
	e.Uint32(obj.Transaction.Length)

	// obj.Transaction.Type
	e.Uint8(obj.Transaction.Type)

	// obj.Transaction.InnerHash
	e.CopyBytes(obj.Transaction.InnerHash[:])

	// obj.Transaction.Sigs maxlen check
	if len(obj.Transaction.Sigs) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Transaction.Sigs length check
	if uint64(len(obj.Transaction.Sigs)) > math.MaxUint32 {
		return errors.New("obj.Transaction.Sigs length exceeds math.MaxUint32")
	}

	// obj.Transaction.Sigs length
	e.Uint32(uint32(len(obj.Transaction.Sigs)))

	// obj.Transaction.Sigs
	for _, x := range obj.Transaction.Sigs {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Transaction.In maxlen check
	if len(obj.Transaction.In) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Transaction.In length check
	if uint64(len(obj.Transaction.In)) > math.MaxUint32 {
		return errors.New("obj.Transaction.In length exceeds math.MaxUint32")
	}

	// obj.Transaction.In length
	e.Uint32(uint32(len(obj.Transaction.In)))

	// obj.Transaction.In
	for _, x := range obj.Transaction.In {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Transaction.Out maxlen check
	if len(obj.Transaction.Out) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Transaction.Out length check
	if uint64(len(obj.Transaction.Out)) > math.MaxUint32 {
		return errors.New("obj.Transaction.Out length exceeds math.MaxUint32")
	}

	// obj.Transaction.Out length
	e.Uint32(uint32(len(obj.Transaction.Out)))

	// obj.Transaction.Out
	for _, x := range obj.Transaction.Out {

		// x.Address.Version
		e.Uint8(x.Address.Version)

		// x.Address.Key
		e.CopyBytes(x.Address.Key[:])

		// x.Coins
		e.Uint64(x.Coins)

		// x.Hours
		e.Uint64(x.Hours)

	}

	e.CopyBytes(obj.Transaction.Tweet[:])

	// obj.Received
	e.Int64(obj.Received)

	// obj.Checked
	e.Int64(obj.Checked)

	// obj.Announced
	e.Int64(obj.Announced)

	// obj.IsValid
	e.Int8(obj.IsValid)

	return nil
}

// decodeUnconfirmedTransaction decodes an object of type UnconfirmedTransaction from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeUnconfirmedTransaction(buf []byte, obj *UnconfirmedTransaction) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Transaction.Length
		i, err := d.Uint32()
		if err != nil {
			return 0, err
		}
		obj.Transaction.Length = i
	}

	{
		// obj.Transaction.Type
		i, err := d.Uint8()
		if err != nil {
			return 0, err
		}
		obj.Transaction.Type = i
	}

	{
		// obj.Transaction.InnerHash
		if len(d.Buffer) < len(obj.Transaction.InnerHash) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Transaction.InnerHash[:], d.Buffer[:len(obj.Transaction.InnerHash)])
		d.Buffer = d.Buffer[len(obj.Transaction.InnerHash):]
	}

	{
		// obj.Transaction.Sigs

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Transaction.Sigs = make([]cipher.Sig, length)

			for z2 := range obj.Transaction.Sigs {
				{
					// obj.Transaction.Sigs[z2]
					if len(d.Buffer) < len(obj.Transaction.Sigs[z2]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Transaction.Sigs[z2][:], d.Buffer[:len(obj.Transaction.Sigs[z2])])
					d.Buffer = d.Buffer[len(obj.Transaction.Sigs[z2]):]
				}

			}
		}
	}

	{
		// obj.Transaction.In

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Transaction.In = make([]cipher.SHA256, length)

			for z2 := range obj.Transaction.In {
				{
					// obj.Transaction.In[z2]
					if len(d.Buffer) < len(obj.Transaction.In[z2]) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Transaction.In[z2][:], d.Buffer[:len(obj.Transaction.In[z2])])
					d.Buffer = d.Buffer[len(obj.Transaction.In[z2]):]
				}

			}
		}
	}

	{
		// obj.Transaction.Out

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Transaction.Out = make([]coin.TransactionOutput, length)

			for z2 := range obj.Transaction.Out {
				{
					// obj.Transaction.Out[z2].Address.Version
					i, err := d.Uint8()
					if err != nil {
						return 0, err
					}
					obj.Transaction.Out[z2].Address.Version = i
				}

				{
					// obj.Transaction.Out[z2].Address.Key
					if len(d.Buffer) < len(obj.Transaction.Out[z2].Address.Key) {
						return 0, encoder.ErrBufferUnderflow
					}
					copy(obj.Transaction.Out[z2].Address.Key[:], d.Buffer[:len(obj.Transaction.Out[z2].Address.Key)])
					d.Buffer = d.Buffer[len(obj.Transaction.Out[z2].Address.Key):]
				}

				{
					// obj.Transaction.Out[z2].Coins
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.Transaction.Out[z2].Coins = i
				}

				{
					// obj.Transaction.Out[z2].Hours
					i, err := d.Uint64()
					if err != nil {
						return 0, err
					}
					obj.Transaction.Out[z2].Hours = i
				}

			}
		}
	}

	copy(obj.Transaction.Tweet[:], d.Buffer[:512])
	d.Buffer = d.Buffer[512:]

	{
		// obj.Received
		i, err := d.Int64()
		if err != nil {
			return 0, err
		}
		obj.Received = i
	}

	{
		// obj.Checked
		i, err := d.Int64()
		if err != nil {
			return 0, err
		}
		obj.Checked = i
	}

	{
		// obj.Announced
		i, err := d.Int64()
		if err != nil {
			return 0, err
		}
		obj.Announced = i
	}

	{
		// obj.IsValid
		i, err := d.Int8()
		if err != nil {
			return 0, err
		}
		obj.IsValid = i
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeUnconfirmedTransactionExact decodes an object of type UnconfirmedTransaction from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeUnconfirmedTransactionExact(buf []byte, obj *UnconfirmedTransaction) error {
	if n, err := decodeUnconfirmedTransaction(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
