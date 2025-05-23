package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type VerificationRelationship int

// A verification relationship expresses the relationship between the DID subject and a verification method.
// This enum is used to
// cfr. https://www.w3.org/TR/did-core/#verification-relationships
const (
	Authentication       = "authentication"       // https://www.w3.org/TR/did-core/#authentication
	AssertionMethod      = "assertionMethod"      // https://www.w3.org/TR/did-core/#assertion
	KeyAgreement         = "keyAgreement"         // https://www.w3.org/TR/did-core/#key-agreement
	CapabilityInvocation = "capabilityInvocation" // https://www.w3.org/TR/did-core/#capability-invocation
	CapabilityDelegation = "capabilityDelegation" // https://www.w3.org/TR/did-core/#capability-delegation
)

const (
	authentication VerificationRelationship = iota
	assertionMethod
	keyAgreement
	capabilityInvocation
	capabilityDelegation
)

// VerificationRelationships are the supported list of verification relationships
var VerificationRelationships = map[string]VerificationRelationship{
	Authentication:       authentication,
	AssertionMethod:      assertionMethod,
	KeyAgreement:         keyAgreement,
	CapabilityInvocation: capabilityInvocation,
	CapabilityDelegation: capabilityDelegation,
}

// getRelationships retrieve the pointer to the verification relationship
// if it exists, otherwise returns nil
func (didDoc *IidDocument) getRelationships(rel VerificationRelationship) *[]string {
	switch rel {
	case authentication:
		return &didDoc.Authentication
	case assertionMethod:
		return &didDoc.AssertionMethod
	case keyAgreement:
		return &didDoc.KeyAgreement
	case capabilityInvocation:
		return &didDoc.CapabilityInvocation
	case capabilityDelegation:
		return &didDoc.CapabilityDelegation
	default:
		return nil
	}
}

// parseRelationshipLabels parse relationships labels to a slice of VerificationRelationship
// making sure that the relationsips are not repeated
func parseRelationshipLabels(relNames ...string) (vrs []VerificationRelationship, err error) {
	names := distinct(relNames)
	vrs = make([]VerificationRelationship, len(names))
	for i, vrn := range distinct(relNames) {
		vr, validName := VerificationRelationships[vrn]
		if !validName {
			err = errorsmod.Wrapf(ErrInvalidInput, "unsupported verification relationship %s", vrn)
			return
		}
		vrs[i] = vr
	}
	return
}

/**
Regexp generated using this ABNF specs and using https://abnf.msweet.org/index.php

did-url            = (did | {id}) path-abempty [ "?" query ] [ "#" fragment ]
did                = "did:" method-name ":" method-specific-id
method-name        = 1*method-char
method-char        = %x61-7A / DIGIT
method-specific-id = *( *idchar ":" ) 1*idchar
idchar             = ALPHA / DIGIT / "." / "-" / "_" / pct-encoded
pct-encoded        = "%" HEXDIG HEXDIG
query              = *( pchar / "/" / "?" )
fragment           = *( pchar / "/" / "?" )
path-abempty       = *( "/" segment )
segment            = *pchar
unreserved         = ALPHA / DIGIT / "-" / "." / "_" / "~"
pchar              = unreserved / pct-encoded / sub-delims / ":" / "@"
sub-delims         = "!" / "$" / "&" / "'" / "(" / ")"
                 / "*" / "+" / "," / ";" / "="
*/

const (
	contextDIDBase            = "https://www.w3.org/ns/did/v1"
	didValidationRegexpStr    = `^(did\:[a-z0-9]+\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+)|\{id\}$`
	didURLValidationRegexpStr = `^((did\:[a-z0-9]+\:(([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])*\:)*([A-Z.a-z0-9]|\-|_|%[0-9A-Fa-f][0-9A-Fa-f])+)|\{id\})(/(([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)*)*(\?(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?(#(((([-A-Z._a-z0-9]|~)|%[0-9A-Fa-f][0-9A-Fa-f]|(\!|\$|&|'|\(|\)|\*|\+|,|;|\=)|\:|@)|/|\?)*))?$`
	rfc3986RegexpStr          = `^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?$`
)

var (
	didValidationRegexp    = regexp.MustCompile(didValidationRegexpStr)
	didURLValidationRegexp = regexp.MustCompile(didURLValidationRegexpStr)
	rfc3986Regexp          = regexp.MustCompile(rfc3986RegexpStr)
)

// DID as typed string
type DID string

// NewChainDID format a DID from a method specific did
// cfr.https://www.w3.org/TR/did-core/#did
func NewChainDID(chainName, didID string) DID {
	return DID(fmt.Sprint(DidChainPrefix, chainName, ":", didID))
}

// func UnmarshalMetadata(meta []byte) IidMetadata {
// 	var metadata IidMetadata
// 	json.Unmarshal(meta, &metadata)
// 	return metadata
// }

// String return the string representation of the did
func (did DID) String() string {
	return string(did)
}

// NewVerificationMethodID compose a verification method id from an account address
func (did DID) NewVerificationMethodID(vmID string) string {
	return fmt.Sprint(did, "#", vmID)
}

// IsValidDID validate the input string according to the
// did specification (cfr. https://www.w3.org/TR/did-core/#did-syntax ).
func IsValidDID(input string) bool {
	return didValidationRegexp.MatchString(input)
}

// IsValidDIDURL validate the input string according to the
// did url specification (cfr. https://www.w3.org/TR/did-core/#did-url-syntax  ).
func IsValidDIDURL(input string) bool {
	return didURLValidationRegexp.MatchString(input)
}

// IsValidRFC3986Uri checks if the input string is a valid RFC3986 URI
// (cfr https://datatracker.ietf.org/doc/html/rfc3986#page-50)
func IsValidRFC3986Uri(input string) bool {
	return rfc3986Regexp.MatchString(input)
}

// IsValidDIDDocument tells if a DID document is valid,
// that is if it has the default context and a valid subject
func IsValidDIDDocument(didDoc *IidDocument) bool {
	if didDoc == nil {
		return false
	}

	if !IsValidDID(didDoc.Id) {
		return false
	}

	for _, element := range didDoc.Controller {
		if !IsValidIIDKeyFormat(element) {
			return false
		}
	}

	//for _, c := range didDoc.Context {
	//	if c.Key != "" && c.Value != "" {
	//		return true
	//	}
	//}

	return true
}

func IsValidIIDKeyFormat(did string) bool {
	//if _, err := sdk.AccAddressFromBech32(strings.TrimPrefix(did, DidKeyPrefix)); err != nil {
	//	bech32 := strings.Split(did, ":")
	//	if _, err := sdk.AccAddressFromBech32(bech32[len(bech32)-1]); err != nil {
	//		return false
	//	}
	//}
	return true
}

// IsValidDIDMetadata tells if a DID metadata is valid,
// that is if it has a non empty versionId and a non-zero create date
func IsValidDIDMetadata(didMeta *IidMetadata) bool {
	if didMeta == nil {
		return false
	}
	if IsEmpty(didMeta.VersionId) {
		return false
	}
	if didMeta.Created == nil || didMeta.Created.IsZero() {
		return false
	}
	return true
}

// ValidateVerification perform basic validation on a verification struct
// optionally validating the validation method controller against a list
// of allowed controllers.
// in case of error returns an cosmos-sdk wrapped error
// XXX: this pattern creates a ambiguous semantic (but maybe is not too severe (use WithCredentials and array of credentials))
func ValidateVerification(v *Verification, allowedControllers ...string) (err error) {
	if v == nil {
		err = errorsmod.Wrap(ErrInvalidInput, "verification is not defined")
		return err
	}
	// verify that the method id is correct
	if !IsValidDIDURL(v.Method.Id) {
		err = errorsmod.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", v.Method.Id)
		return err
	}

	// if the controller is not set return error
	if !IsValidDID(v.Method.Controller) {
		err = errorsmod.Wrapf(ErrInvalidDIDFormat, "verification method controller %v", v.Method.Controller)
		return err
	}

	// check for empty method type
	if IsEmpty(v.Method.Type) {
		err = errorsmod.Wrapf(ErrInvalidInput, "verification method type not set for verification method %s", v.Method.Id)
		return err
	}

	// check the verification material
	switch x := v.Method.VerificationMaterial.(type) {
	case *VerificationMethod_BlockchainAccountID:
		if IsEmpty(x.BlockchainAccountID) {
			err = errorsmod.Wrapf(ErrInvalidInput, "verification material blockchain account id invalid for verification method %s", v.Method.Id)
			return err
		}
	case *VerificationMethod_PublicKeyMultibase:
		if IsEmpty(x.PublicKeyMultibase) {
			err = errorsmod.Wrapf(ErrInvalidInput, "verification material multibase pubkey invalid for verification method %s", v.Method.Id)
			return err
		}
	case *VerificationMethod_PublicKeyHex:
		if IsEmpty(x.PublicKeyHex) {
			err = errorsmod.Wrapf(ErrInvalidInput, "verification material pubkey invalid for verification method %s", v.Method.Id)
			return err
		}
	case *VerificationMethod_PublicKeyBase58:
		if IsEmpty(x.PublicKeyBase58) {
			err = errorsmod.Wrapf(ErrInvalidInput, "verification material pubkey invalid for verification method %s", v.Method.Id)
			return err
		}
	default:
		err = errorsmod.Wrapf(ErrInvalidInput, "verification material %s not set for verification method %s", x, v.Method.Id)
		return err
	}

	// check for empty publickey
	if v.Method.VerificationMaterial.Size() == 0 {
		err = errorsmod.Wrapf(ErrInvalidInput, "verification material not set for verification method %s", v.Method.Id)
		return err
	}

	// check that there is at least a relationship
	if len(v.Relationships) == 0 {
		err = errorsmod.Wrap(ErrEmptyRelationships, "at least a verification relationship is required")
		return err
	}
	return nil
}

// ValidateService performs basic on a service struct
func ValidateService(s *Service) (err error) {
	if s == nil {
		err = errorsmod.Wrap(ErrInvalidInput, "service is not defined")
		return err
	}
	// verify that the id is not empty and is a valid url (according to RFC3986)
	if IsEmpty(s.Id) {
		err = errorsmod.Wrap(ErrInvalidInput, "service id cannot be empty")
		return err
	}

	if !IsValidRFC3986Uri(s.Id) {
		err = errorsmod.Wrapf(ErrInvalidRFC3986UriFormat, "service id %s is not a valid RFC3986 uri", s.Id)
		return err
	}

	// verify that the endpoint is not empty and is a valid url (according to RFC3986)
	if IsEmpty(s.ServiceEndpoint) {
		err = errorsmod.Wrap(ErrInvalidInput, "service endpoint cannot be empty;")
		return err
	}

	if !IsValidRFC3986Uri(s.ServiceEndpoint) {
		err = errorsmod.Wrapf(ErrInvalidRFC3986UriFormat, "service endpoint %s is not a valid RFC3986 uri", s.ServiceEndpoint)
		return err
	}

	// check that the service type is not empty
	if IsEmpty(s.Type) {
		err = errorsmod.Wrap(ErrInvalidInput, "service type cannot be empty")
		return err
	}

	return nil
}

// IsEmpty tells if the trimmed input is empty
func IsEmpty(input string) bool {
	return strings.TrimSpace(input) == ""
}

// DidDocumentOption implements variadic pattern for optional did document fields
type DidDocumentOption func(*IidDocument) error

// WithVerifications add optional verifications
func WithVerifications(verifications ...*Verification) DidDocumentOption {
	return func(did *IidDocument) error {
		return did.AddVerifications(verifications...)
	}
}

// WithServices add optional services
func WithServices(services ...*Service) DidDocumentOption {
	return func(did *IidDocument) error {
		return did.AddServices(services...)
	}
}

// WithRights add optional Rights
func WithRights(rights ...*AccordedRight) DidDocumentOption {
	return func(did *IidDocument) error {
		return did.AddAccordedRight(rights...)
	}
}

// WithResources add optional resources
func WithResources(resources ...*LinkedResource) DidDocumentOption {
	return func(did *IidDocument) error {
		return did.AddLinkedResource(resources...)
	}
}

// WithClaims add optional claims
func WithClaims(claims ...*LinkedClaim) DidDocumentOption {
	return func(did *IidDocument) error {
		return did.AddLinkedClaim(claims...)
	}
}

func WithEntities(entities ...*LinkedEntity) DidDocumentOption {
	return func(did *IidDocument) error {
		return did.AddLinkedEntity(entities...)
	}
}

// WithControllers add optional did controller
func WithControllers(controllers ...string) DidDocumentOption {
	return func(did *IidDocument) (err error) {
		return did.AddControllers(controllers...)
	}
}

// WithControllers add optional did controller
func WithAlsoKnownAs(alsoKnownAs string) DidDocumentOption {
	return func(did *IidDocument) (err error) {
		return did.AddAlsoKnownAs(alsoKnownAs)
	}
}

// WithControllers add optional did controller
func WithContexts(contexts ...*Context) DidDocumentOption {
	return func(did *IidDocument) (err error) {
		return did.AddDidContext(contexts...)
	}
}

// NewDidDocument constructs a new DidDocument
func NewDidDocument(ctx sdk.Context, id string, options ...DidDocumentOption) (did IidDocument, err error) {
	if !IsValidDID(id) {
		err = errorsmod.Wrapf(ErrInvalidDIDFormat, "did %s", id)
		return
	}

	metadata := NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
	did = IidDocument{
		Id:       id,
		Metadata: &metadata,
	}
	// apply all the options
	for _, fn := range options {
		if err = fn(&did); err != nil {
			return
		}
	}
	return
}

// AddControllers add a controller to a did document if not exists
func (didDoc *IidDocument) AddAlsoKnownAs(alsoKnownAs string) error {
	didDoc.AlsoKnownAs = alsoKnownAs
	return nil
}

// AddControllers add a controller to a did document if not exists
func (didDoc *IidDocument) AddControllers(controllers ...string) error {
	if len(controllers) == 0 {
		return nil
	}
	// join the exiting controllers with the new ones
	dc := distinct(append(didDoc.Controller, controllers...))
	for _, c := range dc {
		if !IsValidDID(c) {
			return errorsmod.Wrapf(ErrInvalidDIDFormat, "did document controller validation error '%s'", c)
		}
		if !IsValidIIDKeyFormat(c) {
			return errorsmod.Wrapf(ErrInvalidInput, "did document controller '%s' must be of type key", c)
		}
	}

	// remove duplicates
	didDoc.Controller = dc
	return nil
}

// DeleteControllers delete controllers from a did document
func (didDoc *IidDocument) DeleteControllers(controllers ...string) error {
	if len(controllers) == 0 {
		return nil
	}
	dc := distinct(controllers)
	for _, c := range dc {
		if !IsValidDID(c) {
			return errorsmod.Wrapf(ErrInvalidDIDFormat, "did document controller validation error '%s'", c)
		}
	}
	// remove existing
	didDoc.Controller = subtraction(didDoc.Controller, controllers)
	return nil
}

// Deactivate sets the deactivated field in IidMetadata to true
func (didDoc *IidDocument) Deactivate() error {
	didDoc.Metadata.Deactivated = true
	return nil
}

// AddVerifications add one or more verification method and relations to a did document
func (didDoc *IidDocument) AddVerifications(verifications ...*Verification) (err error) {
	// verify that there are no duplicates in method ids
	index := make(map[string]struct{}, len(didDoc.VerificationMethod))
	// load existing verifications if any
	for _, v := range didDoc.VerificationMethod {
		index[v.Id] = struct{}{}
	}

	// loop through the verifications and look for problems
	for _, v := range verifications {
		// perform base validation checks
		if err = ValidateVerification(v); err != nil {
			return err
		}

		// verify that there are no duplicates in method ids
		if _, found := index[v.Method.Id]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated verification method id %s", v.Method.Id)
			return err
		}
		index[v.Method.Id] = struct{}{}

		// first add the method to the list of methods
		didDoc.VerificationMethod = append(didDoc.VerificationMethod, v.GetMethod())

		// now add the relationships
		vrs, err := parseRelationshipLabels(v.Relationships...)
		if err != nil {
			return err
		}
		didDoc.setRelationships(v.Method.Id, vrs...)
	}
	return nil
}

// RevokeVerification revoke a verification method
// and all relationships associated with it
func (didDoc *IidDocument) RevokeVerification(methodID string) error {
	del := func(x int) {
		lastIdx := len(didDoc.VerificationMethod) - 1
		switch lastIdx {
		case 0:
			didDoc.VerificationMethod = nil
		case x:
			didDoc.VerificationMethod = didDoc.VerificationMethod[:lastIdx]
		default:
			didDoc.VerificationMethod[x] = didDoc.VerificationMethod[lastIdx]
			didDoc.VerificationMethod = didDoc.VerificationMethod[:lastIdx]
		}
	}

	// remove relationships
	didDoc.setRelationships(methodID)

	// now remove the method
	for i, vm := range didDoc.VerificationMethod {
		if vm.Id == methodID {
			del(i)
			return nil
		}
	}
	return errorsmod.Wrapf(ErrVerificationMethodNotFound, "verification method id: %v", methodID)
}

// SetVerificationRelationships for a did document
func (didDoc *IidDocument) SetVerificationRelationships(methodID string, relationships ...string) error {
	// verify that the method id is correct
	if !IsValidDIDURL(methodID) {
		return errorsmod.Wrapf(ErrInvalidDIDURLFormat, "verification method id: %v", methodID)
	}
	// check that the methodID exists
	hasVM := false
	for _, vm := range didDoc.VerificationMethod {
		if vm.Id == methodID {
			hasVM = true
			break
		}
	}
	if !hasVM {
		return errorsmod.Wrapf(ErrVerificationMethodNotFound, "verification method %v not found", methodID)
	}
	// check that there is at least a relationship
	if len(relationships) == 0 {
		return errorsmod.Wrap(ErrEmptyRelationships, "at least a verification relationship is required")
	}
	// check that the provided relationships are valid
	vrs, err := parseRelationshipLabels(relationships...)
	if err != nil {
		return err
	}
	// update the relationships
	didDoc.setRelationships(methodID, vrs...)
	return nil
}

// setRelationships overwrite relationships for a did document
func (didDoc *IidDocument) setRelationships(methodID string, relationships ...VerificationRelationship) {
	// first remove existing relationships
	for _, vr := range VerificationRelationships {
		vrs := didDoc.getRelationships(vr)
		for i, vmID := range *vrs {
			if vmID == methodID {
				lastIdx := len(*vrs) - 1 // get the last index of the current relationship list
				switch lastIdx {
				case 0: // remove the item since there is no elements left
					*vrs = nil
				case i: // if it's at the last position, just drop the last position
					*vrs = (*vrs)[:lastIdx]
				default: // swap and drop last position
					(*vrs)[i] = (*vrs)[lastIdx]
					(*vrs) = (*vrs)[:lastIdx]
				}
			}
		}
	}

	// then assign the new ones
	for _, vr := range relationships {
		vrs := didDoc.getRelationships(vr)
		*vrs = append(*vrs, methodID)
	}
}

// GetVerificationMethodBlockchainAddress returns the verification method cosmos blockchain address of a verification method.
// it fails if the verification method is not supported or if the verification method is not found
func (didDoc IidDocument) GetVerificationMethodBlockchainAddress(methodID string) (sdk.AccAddress, error) {
	// switch m := verificationMethod.GetVerificationMaterial().(type) {
	// case *iidtypes.VerificationMethod_PublicKeyHex:

	// 	pubKey, err := hex.DecodeString(m.PublicKeyHex)
	// 	if err != nil {
	// 		return nil, err

	// 	}

	// 	payerPubKey := ed25519.PubKey{Key: pubKey}
	// 	return accountKeeper.GetAccount(ctx, payerPubKey.Address().Bytes()), nil

	// case *iidtypes.VerificationMethod_BlockchainAccountID:
	// 	addr, err := sdk.AccAddressFromBech32(iidtypes.BlockchainAccountID(m.BlockchainAccountID).GetAddress())
	// 	if err != nil {
	// 		return nil, err
	// 	}

	for _, vm := range didDoc.VerificationMethod {
		if vm.Id == methodID {
			var (
				address string
				err     error
			)
			switch k := vm.VerificationMaterial.(type) {
			case *VerificationMethod_BlockchainAccountID:
				address = BlockchainAccountID(k.BlockchainAccountID).GetAddress()
			case *VerificationMethod_PublicKeyMultibase:
				if VerificationMaterialType(vm.Type) == DIDVMethodTypeEd25519VerificationKey2018 {
					address, err = toEd25519Address(k.PublicKeyMultibase[1:])
				} else {
					address, err = toSecp256k1Address(k.PublicKeyMultibase[1:])
				}
			case *VerificationMethod_PublicKeyHex:
				if VerificationMaterialType(vm.Type) == DIDVMethodTypeEd25519VerificationKey2018 {
					address, err = toEd25519Address(k.PublicKeyHex)
				} else {
					address, err = toSecp256k1Address(k.PublicKeyHex)
				}
			case *VerificationMethod_PublicKeyBase58:
				if VerificationMaterialType(vm.Type) == DIDVMethodTypeEd25519VerificationKey2018 {
					address, err = toEd25519AddressFromBase58(k.PublicKeyBase58)
				} else {
					address, err = toSecp256k1AddressFromBase58(k.PublicKeyBase58)
				}
			}

			if err != nil {
				return nil, err
			}

			return sdk.AccAddressFromBech32(address)
		}
	}
	return nil, ErrVerificationMethodNotFound
}

// GetVerificationRelationships returns the relationships associated with the
// verification method id.
func (didDoc IidDocument) GetVerificationRelationships(methodID string) []string {
	relationships := []string{}
	for vrn, vr := range VerificationRelationships {
		for _, vmID := range *didDoc.getRelationships(vr) {
			if vmID == methodID {
				relationships = append(relationships, vrn)
			}
		}
	}
	return relationships
}

// HasRelationship verifies if a controller did
// exist for at least one of the relationships in the did document
func (didDoc IidDocument) HasRelationship(
	signer BlockchainAccountID,
	relationships ...string,
) bool {
	// first check if the controller exists
	for _, vm := range didDoc.VerificationMethod {
		switch k := vm.VerificationMaterial.(type) {
		case *VerificationMethod_BlockchainAccountID:
			if k.BlockchainAccountID != signer.EncodeToString() {
				continue
			}
		case *VerificationMethod_PublicKeyMultibase:
			addr, err := toSecp256k1Address(k.PublicKeyMultibase[1:])
			if err != nil || !signer.MatchAddress(addr) {
				continue
			}
		case *VerificationMethod_PublicKeyHex:
			addr, err := toSecp256k1Address(k.PublicKeyHex)
			if err != nil || !signer.MatchAddress(addr) {
				continue
			}
		case *VerificationMethod_PublicKeyBase58:
			addr, err := toSecp256k1AddressFromBase58(k.PublicKeyBase58)
			if err != nil || !signer.MatchAddress(addr) {
				continue
			}
		}
		vrs := didDoc.GetVerificationRelationships(vm.Id)
		if len(intersection(vrs, relationships)) > 0 {
			return true
		}
	}
	return false
}

// HasPublicKey validates if a public key is contained in a DidDocument
func (didDoc IidDocument) HasPublicKey(pubkey cryptotypes.PubKey) bool {
	for _, vm := range didDoc.VerificationMethod {
		switch key := vm.VerificationMaterial.(type) {
		case *VerificationMethod_BlockchainAccountID:
			address := sdk.MustBech32ifyAddressBytes(
				sdk.GetConfig().GetBech32AccountAddrPrefix(),
				pubkey.Address().Bytes(),
			)
			if BlockchainAccountID(key.BlockchainAccountID).MatchAddress(address) {
				return true
			}
		case *VerificationMethod_PublicKeyMultibase:
			if key.PublicKeyMultibase == fmt.Sprint("F", hex.EncodeToString(pubkey.Bytes())) {
				return true
			}

		case *VerificationMethod_PublicKeyHex:
			if key.PublicKeyHex == hex.EncodeToString(pubkey.Bytes()) {
				return true
			}
		case *VerificationMethod_PublicKeyBase58:
			if key.PublicKeyBase58 == base58.Encode(pubkey.Bytes()) {
				return true
			}
		}
	}
	return false
}

// HasController returns true if the DID document has the input DID as a controller, false otherwise
func (didDoc *IidDocument) HasController(controller DID) bool {
	ctrl := controller.String()
	for _, c := range didDoc.Controller {
		if c == ctrl {
			return true
		}
	}
	return false
}

// AddServices add services to a did document
func (didDoc *IidDocument) AddServices(services ...*Service) (err error) {
	if didDoc.Service == nil {
		didDoc.Service = []*Service{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.Service))

	// load existing services
	for _, s := range didDoc.Service {
		index[s.Id] = struct{}{}
	}

	// services must be unique
	for _, s := range services {
		if err = ValidateService(s); err != nil {
			return
		}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated service id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.Service = append(didDoc.Service, s)
	}
	return
}

func (didDoc *IidDocument) AddLinkedResource(linkedResources ...*LinkedResource) (err error) {
	if didDoc.LinkedResource == nil {
		didDoc.LinkedResource = []*LinkedResource{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.LinkedResource))

	// load existing resources
	for _, s := range didDoc.LinkedResource {
		index[s.Id] = struct{}{}
	}

	// resources must be unique
	for _, s := range linkedResources {
		//if err = ValidateService(s); err != nil {
		//	return
		//}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated resource id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.LinkedResource = append(didDoc.LinkedResource, s)
	}
	return
}

func (didDoc *IidDocument) AddLinkedClaim(linkedClaims ...*LinkedClaim) (err error) {
	if didDoc.LinkedClaim == nil {
		didDoc.LinkedClaim = []*LinkedClaim{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.LinkedClaim))

	// load existing Claims
	for _, s := range didDoc.LinkedClaim {
		index[s.Id] = struct{}{}
	}

	// Claims must be unique
	for _, s := range linkedClaims {
		//if err = ValidateService(s); err != nil {
		//	return
		//}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated Claim id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.LinkedClaim = append(didDoc.LinkedClaim, s)
	}
	return
}

func (didDoc *IidDocument) DeleteLinkedResource(resourceID string) {
	del := func(x int) {
		lastIdx := len(didDoc.LinkedResource) - 1
		switch lastIdx {
		case 0: // remove the item since there is no elements left
			didDoc.LinkedResource = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.LinkedResource = didDoc.LinkedResource[:lastIdx]
		default: // swap and drop last position
			didDoc.LinkedResource[x] = didDoc.LinkedResource[lastIdx]
			didDoc.LinkedResource = didDoc.LinkedResource[:lastIdx]
		}
	}

	for i, s := range didDoc.LinkedResource {
		if s.Id == resourceID {
			del(i)
			break
		}
	}
}

func (didDoc *IidDocument) DeleteLinkedClaim(claimID string) {
	del := func(x int) {
		lastIdx := len(didDoc.LinkedClaim) - 1
		switch lastIdx {
		case 0: // remove the item since there is no elements left
			didDoc.LinkedClaim = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.LinkedClaim = didDoc.LinkedClaim[:lastIdx]
		default: // swap and drop last position
			didDoc.LinkedClaim[x] = didDoc.LinkedClaim[lastIdx]
			didDoc.LinkedClaim = didDoc.LinkedClaim[:lastIdx]
		}
	}

	for i, s := range didDoc.LinkedClaim {
		if s.Id == claimID {
			del(i)
			break
		}
	}
}

func (didDoc *IidDocument) AddLinkedEntity(linkedEntities ...*LinkedEntity) (err error) {
	if didDoc.LinkedEntity == nil {
		didDoc.LinkedEntity = []*LinkedEntity{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.LinkedEntity))

	// load existing resources
	for _, s := range didDoc.LinkedEntity {
		index[s.Id] = struct{}{}
	}

	// resources must be unique
	for _, s := range linkedEntities {
		//if err = ValidateService(s); err != nil {
		//	return
		//}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated linked entity id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.LinkedEntity = append(didDoc.LinkedEntity, s)
	}
	return
}

func (didDoc *IidDocument) DeleteLinkedEntity(entityID string) {
	del := func(x int) {
		lastIdx := len(didDoc.LinkedEntity) - 1
		switch lastIdx {
		case 0: // remove the item since there is no elements left
			didDoc.LinkedEntity = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.LinkedEntity = didDoc.LinkedEntity[:lastIdx]
		default: // swap and drop last position
			didDoc.LinkedEntity[x] = didDoc.LinkedEntity[lastIdx]
			didDoc.LinkedEntity = didDoc.LinkedEntity[:lastIdx]
		}
	}

	for i, s := range didDoc.LinkedEntity {
		if s.Id == entityID {
			del(i)
			break
		}
	}
}

func (didDoc *IidDocument) AddAccordedRight(accordedRights ...*AccordedRight) (err error) {
	if didDoc.AccordedRight == nil {
		didDoc.AccordedRight = []*AccordedRight{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.AccordedRight))

	// load existing resources
	for _, s := range didDoc.AccordedRight {
		index[s.Id] = struct{}{}
	}

	// resources must be unique
	for _, s := range accordedRights {
		//if err = ValidateService(s); err != nil {
		//	return
		//}

		// verify that there are no duplicates in method ids
		if _, found := index[s.Id]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated accorded right id %s", s.Id)
			return
		}
		index[s.Id] = struct{}{}

		didDoc.AccordedRight = append(didDoc.AccordedRight, s)
	}
	return
}

func (didDoc *IidDocument) DeleteAccordedRight(rightID string) {
	del := func(x int) {
		lastIdx := len(didDoc.AccordedRight) - 1
		switch lastIdx {
		case 0: // remove the item since there is no elements left
			didDoc.AccordedRight = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.AccordedRight = didDoc.AccordedRight[:lastIdx]
		default: // swap and drop last position
			didDoc.AccordedRight[x] = didDoc.AccordedRight[lastIdx]
			didDoc.AccordedRight = didDoc.AccordedRight[:lastIdx]
		}
	}

	for i, s := range didDoc.AccordedRight {
		if s.Id == rightID {
			del(i)
			break
		}
	}
}

func (didDoc *IidDocument) AddDidContext(contexts ...*Context) (err error) {
	if didDoc.Context == nil {
		didDoc.Context = []*Context{}
	}

	// used to check duplicates
	index := make(map[string]struct{}, len(didDoc.Context))

	//load existing resources
	for _, c := range didDoc.Context {
		index[c.Key] = struct{}{}
	}

	// resources must be unique
	for _, c := range contexts {
		if _, found := index[c.Key]; found {
			err = errorsmod.Wrapf(ErrInvalidInput, "duplicated context key found %s", c.Key)
			return
		}
		index[c.Key] = struct{}{}
		didDoc.Context = append(didDoc.Context, c)
	}
	return
}

func (didDoc *IidDocument) DeleteDidContext(contextKey string) {
	del := func(x int) {
		lastIdx := len(didDoc.Context) - 1
		switch lastIdx {
		case 0: // remove the item since there is no elements left
			didDoc.Context = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.Context = didDoc.Context[:lastIdx]
		default: // swap and drop last position
			didDoc.Context[x] = didDoc.Context[lastIdx]
			didDoc.Context = didDoc.Context[:lastIdx]
		}
	}

	for i, s := range didDoc.Context {
		if s.Key == contextKey {
			del(i)
			break
		}
	}
}

// DeleteService delete an existing service from a did document
func (didDoc *IidDocument) DeleteService(serviceID string) {
	del := func(x int) {
		lastIdx := len(didDoc.Service) - 1
		switch lastIdx {
		case 0: // remove the item since there is no elements left
			didDoc.Service = nil
		case x: // if it's at the last position, just drop the last position
			didDoc.Service = didDoc.Service[:lastIdx]
		default: // swap and drop last position
			didDoc.Service[x] = didDoc.Service[lastIdx]
			didDoc.Service = didDoc.Service[:lastIdx]
		}
	}

	for i, s := range didDoc.Service {
		if s.Id == serviceID {
			del(i)
			break
		}
	}
}

// GetBytes is a helper for serializing
func (didDoc IidDocument) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&didDoc))
}

func (didDoc IidDocument) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&didDoc))
}

// Verifications is a list of verification
type Verifications []*Verification

// NewVerification build a new verification to be
// attached to a did document
func NewVerification(
	method VerificationMethod,
	relationships []string,
	contexts []string,
) *Verification {
	return &Verification{
		Context:       contexts,
		Method:        &method,
		Relationships: relationships,
	}
}

// NewAccountVerification is a shortcut to create a verification based on comsos address
// func NewAccountVerification(did DID, chainID, accountAddress string, verificationMethods ...string) *Verification {
// 	return NewVerification(
// 		NewVerificationMethod(
// 			did.NewVerificationMethodID(accountAddress),
// 			did,
// 			NewBlockchainAccountID(chainID, accountAddress),
// 		),
// 		verificationMethods,
// 		nil,
// 	)
// }

// NewVerificationMethod build a new verification method
func NewVerificationMethod(id string, controller DID, vmr VerificationMaterial) VerificationMethod {
	vm := VerificationMethod{
		Id:         id,
		Controller: controller.String(),
		Type:       string(vmr.Type()),
	}
	switch vmr.(type) {
	case BlockchainAccountID:
		vm.VerificationMaterial = &VerificationMethod_BlockchainAccountID{vmr.EncodeToString()}
	case PublicKeyMultibase:
		vm.VerificationMaterial = &VerificationMethod_PublicKeyMultibase{vmr.EncodeToString()}
	case PublicKeyHex:
		vm.VerificationMaterial = &VerificationMethod_PublicKeyHex{vmr.EncodeToString()}
	case PublicKeyBase58:
		vm.VerificationMaterial = &VerificationMethod_PublicKeyBase58{vmr.EncodeToString()}
	}
	return vm
}

// GetBytes is a helper for serializing
func (did Verification) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&did))
}

func (did Verification) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleAminoCdc.MustMarshalJSON(&did))
}

// Services are a list of services
type Services []*Service
type AccordedRights []*AccordedRight
type LinkedResources []*LinkedResource
type LinkedEntities []*LinkedEntity
type Contexts []*Context

// NewService creates a new service
func NewService(id string, serviceType string, serviceEndpoint string) *Service {
	return &Service{
		Id:              id,
		Type:            serviceType,
		ServiceEndpoint: serviceEndpoint,
	}
}

func NewLinkedResource(id string, resourceType string, description string, mediaType string, serviceEndpoint string, proof string, encrypted string, privacy string) *LinkedResource {
	return &LinkedResource{
		Type:            resourceType,
		Id:              id,
		Description:     description,
		MediaType:       mediaType,
		ServiceEndpoint: serviceEndpoint,
		Proof:           proof,
		Encrypted:       encrypted,
		Right:           privacy,
	}
}

func NewLinkedClaim(id string, claimType string, description string, serviceEndpoint string, proof string, encrypted string, privacy string) *LinkedClaim {
	return &LinkedClaim{
		Type:            claimType,
		Id:              id,
		Description:     description,
		ServiceEndpoint: serviceEndpoint,
		Proof:           proof,
		Encrypted:       encrypted,
		Right:           privacy,
	}
}

func NewLinkedEntity(id, entityType, relationship, service string) *LinkedEntity {
	return &LinkedEntity{
		Id:           id,
		Type:         entityType,
		Relationship: relationship,
		Service:      service,
	}
}

func NewAccordedRight(id string, rightType string, mechanism string, message string, endpoint string) *AccordedRight {
	return &AccordedRight{
		Type:      rightType,
		Id:        id,
		Mechanism: mechanism,
		Message:   message,
		Service:   endpoint,
	}
}

func NewDidContext(key string, val string) *Context {
	return &Context{
		Key: key,
		Val: val,
	}
}

// UpdateDidMetadata updates a DID metadata time and version id
func UpdateDidMetadata(meta *IidMetadata, versionData []byte, updated time.Time) {
	txH := sha256.Sum256(versionData)
	meta.VersionId = hex.EncodeToString(txH[:])
	meta.Updated = &updated
}

// // helper function to update the did metadata on IidDocument
// func (didDoc *IidDocument) updateDidMetadata(ctx sdk.Context) {
// 	UpdateDidMetadata(didDoc.Metadata, ctx.TxBytes(), ctx.BlockTime())
// }

func NewDidMetadata(versionData []byte, created time.Time) IidMetadata {
	m := IidMetadata{
		Created:     &created,
		Deactivated: false,
	}
	UpdateDidMetadata(&m, versionData, created)
	return m
}

// toAddress encode a kexKey string to cosmos based address
func toSecp256k1Address(hexKey string) (addr string, err error) {
	// decode the hex string
	pkb, err := hex.DecodeString(hexKey)
	if err != nil {
		return
	}
	// check the size of the decoded byte slice, otherwise the pk.Address will panic
	if len(pkb) != secp256k1.PubKeySize {
		err = fmt.Errorf("invalid public key size")
		return
	}
	// load the public key
	pk := &secp256k1.PubKey{Key: pkb}
	// generate the address
	addr, err = sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pk.Address())
	return
}

func toEd25519Address(hexKey string) (addr string, err error) {
	// decode the hex string
	pkb, err := hex.DecodeString(hexKey)
	if err != nil {
		return
	}
	// check the size of the decoded byte slice, otherwise the pk.Address will panic
	if len(pkb) != ed25519.PubKeySize {
		err = fmt.Errorf("invalid public key size")
		return
	}
	// load the public key
	pk := &ed25519.PubKey{Key: pkb}
	// generate the address
	addr, err = sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pk.Address())
	return
}

func toSecp256k1AddressFromBase58(base58Key string) (addr string, err error) {
	// decode the base58 string
	pkb := base58.Decode(base58Key)

	// check the size of the decoded byte slice, otherwise the pk.Address will panic
	if len(pkb) != secp256k1.PubKeySize {
		err = fmt.Errorf("invalid public key size")
		return
	}
	// load the public key
	pk := &secp256k1.PubKey{Key: pkb}
	// generate the address
	addr, err = sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pk.Address())
	return
}

func toEd25519AddressFromBase58(base58Key string) (addr string, err error) {
	// decode the base58 string
	pkb := base58.Decode(base58Key)

	// check the size of the decoded byte slice, otherwise the pk.Address will panic
	if len(pkb) != ed25519.PubKeySize {
		err = fmt.Errorf("invalid public key size")
		return
	}
	// load the public key
	pk := &ed25519.PubKey{Key: pkb}
	// generate the address
	addr, err = sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pk.Address())
	return
}

// nolint:unused
// union perform union, distinct amd sort operation between two slices
// duplicated element in list a are
func union(a, b []string) []string {
	if len(b) == 0 {
		return a
	}
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	for _, item := range b {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
		}
	}
	u := make([]string, 0, len(m))
	for k := range m {
		u = append(u, k)
	}
	sort.Strings(u)
	return u
}

func intersection(a, b []string) []string {
	m := make(map[string]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	var i []string
	for _, item := range distinct(b) {
		if _, ok := m[item]; ok {
			i = append(i, item)
		}
	}
	sort.Strings(i)
	return i
}

// distinct remove duplicates and sorts from a list of strings
func distinct(a []string) []string {
	m := make(map[string]bool) // Use bool for readability; struct{} is traditionally used to signal intent of set.
	var d []string

	// Remove duplicates and collect unique items.
	for _, item := range a {
		if _, exists := m[item]; !exists {
			d = append(d, item)
			m[item] = true
		}
	}

	// Sort the unique items.
	sort.Strings(d)
	return d
}

// subtraction remove set b from a
func subtraction(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range b {
		m[item] = true // Create a map to identify items in b for quick lookup.
	}
	var s []string
	for _, item := range a {
		if !m[item] {
			s = append(s, item) // Append items not found in b to the slice s.
		}
	}
	sort.Strings(s)
	return s
}

func GenerateVerificationsFromJson(json VerificationsJSON) ([]*Verification, error) {
	verification := make([]*Verification, 0, len(json.Verifications))
	for _, v := range json.Verifications {
		var verificationMethodMaterial VerificationMaterial

		if v.Method.BlockchainAccountID != "" {
			verificationMethodMaterial = BlockchainAccountID(v.Method.BlockchainAccountID)
		}
		if v.Method.PublicKeyMultibase != "" {
			verificationMethodMaterial, _ = NewPublicKeyMultibaseFromHex(v.Method.PublicKeyMultibase[1:], VerificationMaterialType(v.Method.Type))
		}
		if v.Method.PublicKeyHex != "" {
			verificationMethodMaterial, _ = NewPublicKeyHexFromString(v.Method.PublicKeyHex, VerificationMaterialType(v.Method.Type))
		}
		if v.Method.PublicKeyBase58 != "" {
			verificationMethodMaterial, _ = NewPublicKeyBase58FromString(v.Method.PublicKeyBase58, VerificationMaterialType(v.Method.Type))
		}

		verification = append(verification, NewVerification(
			NewVerificationMethod(
				v.Method.Id,
				DID(v.Method.Controller),
				verificationMethodMaterial,
			),
			v.Relationships,
			v.Context,
		))
	}

	return verification, nil
}

type VerificationsJSON struct {
	Verifications []struct {
		Relationships []string
		Context       []string
		Method        struct {
			Id                  string
			Type                string
			Controller          string
			PublicKeyBase58     string
			BlockchainAccountID string
			PublicKeyMultibase  string
			PublicKeyHex        string
		}
	}
}
