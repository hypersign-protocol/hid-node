package types

import "github.com/multiformats/go-multibase"

func (v VerificationMethod) GetPublicKey() ([]byte, error) {
	if len(v.PublicKeyMultibase) > 0 {
		_, key, err := multibase.Decode(v.PublicKeyMultibase)
		if err != nil {
			return nil, ErrInvalidPublicKey.Wrapf("Cannot decode verification method '%s' public key", v.Id)
		}
		return key, nil
	}

	return nil, ErrInvalidPublicKey.Wrapf("verification method '%s' public key not found", v.Id)
}

// TODO: Implement the following without the field Creator
// func accAddrByKeyRef(keyring keyring.Keyring, keyRef string) (sdk.AccAddress, error) {
// 	// Firstly check if the keyref is a key name of a key registered in a keyring
// 	info, err := keyring.Key(keyRef)

// 	if err == nil {
// 		return info.GetAddress(), nil
// 	}

// 	if !sdkerr.IsOf(err, sdkerr.ErrIO, sdkerr.ErrKeyNotFound) {
// 		return nil, err
// 	}

// 	// Fallback: convert keyref to address
// 	return sdk.AccAddressFromBech32(keyRef)
// }

// func (msg *MsgCreateDID) GetSigners() []sdk.AccAddress {
// 	ctx := client
// 	signerAccAddr, err := accAddrByKeyRef(ctx.Keyring, ctx.From)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return []sdk.AccAddress{signerAccAddr}
// }
