package tokenmeta

import (
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/types"
)

type Instruction uint8

const (
	InstructionCreateMetadataAccount Instruction = iota
	InstructionUpdateMetadataAccount
	InstructionDeprecatedCreateMasterEdition
	InstructionDeprecatedMintNewEditionFromMasterEditionViaPrintingToken
	InstructionUpdatePrimarySaleHappenedViaToken
	InstructionDeprecatedSetReservationList
	InstructionDeprecatedCreateReservationList
	InstructionSignMetadata
	InstructionDeprecatedMintPrintingTokensViaToken
	InstructionDeprecatedMintPrintingTokens
	InstructionCreateMasterEdition
	InstructionMintNewEditionFromMasterEditionViaToken
	InstructionConvertMasterEditionV1ToV2
	InstructionMintNewEditionFromMasterEditionViaVaultProxy
	InstructionPuffMetadata
	InstructionUpdateMetadataAccountV2
	InstructionCreateMetadataAccountV2
	InstructionCreateMasterEditionV3
	InstructionVerifyCollection
	InstructionUtilize
	InstructionApproveUseAuthority
	InstructionRevokeUseAuthority
	InstructionUnverifyCollection
	InstructionApproveCollectionAuthority
	InstructionRevokeCollectionAuthority
	InstructionSetAndVerifyCollection
	InstructionFreezeDelegatedAccount
	InstructionThawDelegatedAccount
	InstructionRemoveCreatorVerification
	InstructionBurnANFT
	Instruction_
	InstructionUnverifySizedCollectionItem
	InstructionSetAndVerifySizedCollectionItem
	InstructionCreateMetadataAccountV3
	InstructionSetCollectionSize
	InstructionSetTokenStandard
)

type CreateMetadataAccountParam struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	MintData                Data
}

func CreateMetadataAccount(param CreateMetadataAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        Data
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccount,
		Data:        param.MintData,
		IsMutable:   param.IsMutable,
	})

	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type UpdateMetadataAccountParam struct {
	MetadataAccount     common.PublicKey
	UpdateAuthority     common.PublicKey
	Data                *Data
	NewUpdateAuthority  *common.PublicKey
	PrimarySaleHappened *bool
}

func UpdateMetadataAccount(param UpdateMetadataAccountParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction         Instruction
		Data                *Data
		NewUpdateAuthority  *common.PublicKey
		PrimarySaleHappened *bool
	}{
		Instruction:         InstructionUpdateMetadataAccount,
		Data:                param.Data,
		NewUpdateAuthority:  param.NewUpdateAuthority,
		PrimarySaleHappened: param.PrimarySaleHappened,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.MetadataAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMasterEditionParam struct {
	Edition         common.PublicKey
	Mint            common.PublicKey
	UpdateAuthority common.PublicKey
	MintAuthority   common.PublicKey
	Metadata        common.PublicKey
	Payer           common.PublicKey
	MaxSupply       *uint64
}

func CreateMasterEdition(param CreateMasterEditionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		MaxSupply   *uint64
	}{
		Instruction: InstructionCreateMasterEdition,
		MaxSupply:   param.MaxSupply,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type SignMetadataParam struct {
	Metadata common.PublicKey
	Creator  common.PublicKey
}

func SignMetadata(param SignMetadataParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSignMetadata,
	})
	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Creator,
				IsSigner:   true,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type MintNewEditionFromMasterEditionViaTokeParam struct {
	NewMetaData                common.PublicKey
	NewEdition                 common.PublicKey
	MasterEdition              common.PublicKey
	NewMint                    common.PublicKey
	EditionMark                common.PublicKey
	NewMintAuthority           common.PublicKey
	Payer                      common.PublicKey
	TokenAccountOwner          common.PublicKey
	TokenAccount               common.PublicKey
	NewMetadataUpdateAuthority common.PublicKey
	MasterMetadata             common.PublicKey
	Edition                    uint64
}

func MintNewEditionFromMasterEditionViaToken(param MintNewEditionFromMasterEditionViaTokeParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Edition     uint64
	}{
		Instruction: InstructionMintNewEditionFromMasterEditionViaToken,
		Edition:     param.Edition,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.NewMetaData,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewEdition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.MasterEdition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewMint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.EditionMark,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.NewMintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccountOwner,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.NewMetadataUpdateAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MasterMetadata,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMetadataAccountV2Param struct {
	Metadata                common.PublicKey
	Mint                    common.PublicKey
	MintAuthority           common.PublicKey
	Payer                   common.PublicKey
	UpdateAuthority         common.PublicKey
	UpdateAuthorityIsSigner bool
	IsMutable               bool
	Data                    DataV2
}

func CreateMetadataAccountV2(param CreateMetadataAccountV2Param) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		Data        DataV2
		IsMutable   bool
	}{
		Instruction: InstructionCreateMetadataAccountV2,
		Data:        param.Data,
		IsMutable:   param.IsMutable,
	})

	if err != nil {
		panic(err)
	}

	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   param.UpdateAuthorityIsSigner,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type CreateMasterEditionV3Param struct {
	Edition         common.PublicKey
	Mint            common.PublicKey
	UpdateAuthority common.PublicKey
	MintAuthority   common.PublicKey
	Metadata        common.PublicKey
	Payer           common.PublicKey
	MaxSupply       *uint64
}

func CreateMasterEditionV3(param CreateMasterEditionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
		MaxSupply   *uint64
	}{
		Instruction: InstructionCreateMasterEditionV3,
		MaxSupply:   param.MaxSupply,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Edition,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.MintAuthority,
				IsSigner:   true,
				IsWritable: false,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     common.TokenProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SystemProgramID,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     common.SysVarRentPubkey,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type VerifyCollectionParam struct {
	Payer                          common.PublicKey //payer
	Metadata                       common.PublicKey //Metadata account
	CollectionAuthority            common.PublicKey //Collection Update authority
	CollectionMint                 common.PublicKey //Mint of the Collection
	Collection                     common.PublicKey //Metadata Account of the Collection
	CollectionMasterEditionAccount common.PublicKey //MasterEdition2 Account of the Collection Token
}

func VerifyCollection(param VerifyCollectionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionVerifyCollection,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Collection,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type UnverifyCollectionParam struct {
	Metadata                       common.PublicKey
	CollectionAuthority            common.PublicKey
	CollectionMint                 common.PublicKey
	Collection                     common.PublicKey
	CollectionMasterEditionAccount common.PublicKey
	CollectionAuthorityRecord      common.PublicKey
}

func UnverifyCollection(param UnverifyCollectionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionUnverifyCollection,
	})
	if err != nil {
		panic(err)
	}
	return types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Collection,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionAuthorityRecord,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}
}

type SetAndVerifyCollectionParam struct {
	Payer                          common.PublicKey
	Metadata                       common.PublicKey
	CollectionAuthority            common.PublicKey
	UpdateAuthority                common.PublicKey
	CollectionMint                 common.PublicKey
	Collection                     common.PublicKey
	CollectionMasterEditionAccount common.PublicKey
	CollectionAuthorityRecord      common.PublicKey
}

var (
	EmptyPubKey = common.PublicKey{}
)

func SetAndVerifyCollection(param SetAndVerifyCollectionParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionSetAndVerifyCollection,
	})
	if err != nil {
		panic(err)
	}
	ix := types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.CollectionAuthority,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Payer,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.UpdateAuthority,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMint,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.Collection,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMasterEditionAccount,
				IsSigner:   false,
				IsWritable: false,
			},
		},
		Data: data,
	}

	if param.CollectionAuthorityRecord != EmptyPubKey {
		ix.Accounts = append(ix.Accounts, types.AccountMeta{
			PubKey:     param.CollectionAuthorityRecord,
			IsSigner:   false,
			IsWritable: false,
		})
	}

	return ix
}

type BurnANFTParam struct {
	Metadata             common.PublicKey
	Owner                common.PublicKey
	Mint                 common.PublicKey
	TokenAccount         common.PublicKey
	MasterEditionAccount common.PublicKey
	SplTokenProgram      common.PublicKey
	CollectionMetadata   common.PublicKey
}

func BurnANFT(param BurnANFTParam) types.Instruction {
	data, err := borsh.Serialize(struct {
		Instruction Instruction
	}{
		Instruction: InstructionBurnANFT,
	})
	if err != nil {
		panic(err)
	}
	ix := types.Instruction{
		ProgramID: common.MetaplexTokenMetaProgramID,
		Accounts: []types.AccountMeta{
			{
				PubKey:     param.Metadata,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.Owner,
				IsSigner:   true,
				IsWritable: true,
			},
			{
				PubKey:     param.Mint,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.TokenAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.MasterEditionAccount,
				IsSigner:   false,
				IsWritable: true,
			},
			{
				PubKey:     param.SplTokenProgram,
				IsSigner:   false,
				IsWritable: false,
			},
			{
				PubKey:     param.CollectionMetadata,
				IsSigner:   false,
				IsWritable: true,
			},
		},
		Data: data,
	}
	return ix
}
