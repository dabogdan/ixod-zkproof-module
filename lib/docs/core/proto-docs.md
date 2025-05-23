# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [ixo/bonds/v1beta1/bonds.proto](#ixo/bonds/v1beta1/bonds.proto)
    - [BaseOrder](#ixo.bonds.v1beta1.BaseOrder)
    - [Batch](#ixo.bonds.v1beta1.Batch)
    - [Bond](#ixo.bonds.v1beta1.Bond)
    - [BondDetails](#ixo.bonds.v1beta1.BondDetails)
    - [BuyOrder](#ixo.bonds.v1beta1.BuyOrder)
    - [FunctionParam](#ixo.bonds.v1beta1.FunctionParam)
    - [Params](#ixo.bonds.v1beta1.Params)
    - [SellOrder](#ixo.bonds.v1beta1.SellOrder)
    - [SwapOrder](#ixo.bonds.v1beta1.SwapOrder)
  
- [ixo/bonds/v1beta1/event.proto](#ixo/bonds/v1beta1/event.proto)
    - [BondBuyOrderCancelledEvent](#ixo.bonds.v1beta1.BondBuyOrderCancelledEvent)
    - [BondBuyOrderEvent](#ixo.bonds.v1beta1.BondBuyOrderEvent)
    - [BondBuyOrderFulfilledEvent](#ixo.bonds.v1beta1.BondBuyOrderFulfilledEvent)
    - [BondCreatedEvent](#ixo.bonds.v1beta1.BondCreatedEvent)
    - [BondEditAlphaFailedEvent](#ixo.bonds.v1beta1.BondEditAlphaFailedEvent)
    - [BondEditAlphaSuccessEvent](#ixo.bonds.v1beta1.BondEditAlphaSuccessEvent)
    - [BondMakeOutcomePaymentEvent](#ixo.bonds.v1beta1.BondMakeOutcomePaymentEvent)
    - [BondSellOrderEvent](#ixo.bonds.v1beta1.BondSellOrderEvent)
    - [BondSellOrderFulfilledEvent](#ixo.bonds.v1beta1.BondSellOrderFulfilledEvent)
    - [BondSetNextAlphaEvent](#ixo.bonds.v1beta1.BondSetNextAlphaEvent)
    - [BondSwapOrderEvent](#ixo.bonds.v1beta1.BondSwapOrderEvent)
    - [BondSwapOrderFulfilledEvent](#ixo.bonds.v1beta1.BondSwapOrderFulfilledEvent)
    - [BondUpdatedEvent](#ixo.bonds.v1beta1.BondUpdatedEvent)
    - [BondWithdrawReserveEvent](#ixo.bonds.v1beta1.BondWithdrawReserveEvent)
    - [BondWithdrawShareEvent](#ixo.bonds.v1beta1.BondWithdrawShareEvent)
  
- [ixo/bonds/v1beta1/genesis.proto](#ixo/bonds/v1beta1/genesis.proto)
    - [GenesisState](#ixo.bonds.v1beta1.GenesisState)
  
- [ixo/bonds/v1beta1/query.proto](#ixo/bonds/v1beta1/query.proto)
    - [QueryAlphaMaximumsRequest](#ixo.bonds.v1beta1.QueryAlphaMaximumsRequest)
    - [QueryAlphaMaximumsResponse](#ixo.bonds.v1beta1.QueryAlphaMaximumsResponse)
    - [QueryAvailableReserveRequest](#ixo.bonds.v1beta1.QueryAvailableReserveRequest)
    - [QueryAvailableReserveResponse](#ixo.bonds.v1beta1.QueryAvailableReserveResponse)
    - [QueryBatchRequest](#ixo.bonds.v1beta1.QueryBatchRequest)
    - [QueryBatchResponse](#ixo.bonds.v1beta1.QueryBatchResponse)
    - [QueryBondRequest](#ixo.bonds.v1beta1.QueryBondRequest)
    - [QueryBondResponse](#ixo.bonds.v1beta1.QueryBondResponse)
    - [QueryBondsDetailedRequest](#ixo.bonds.v1beta1.QueryBondsDetailedRequest)
    - [QueryBondsDetailedResponse](#ixo.bonds.v1beta1.QueryBondsDetailedResponse)
    - [QueryBondsRequest](#ixo.bonds.v1beta1.QueryBondsRequest)
    - [QueryBondsResponse](#ixo.bonds.v1beta1.QueryBondsResponse)
    - [QueryBuyPriceRequest](#ixo.bonds.v1beta1.QueryBuyPriceRequest)
    - [QueryBuyPriceResponse](#ixo.bonds.v1beta1.QueryBuyPriceResponse)
    - [QueryCurrentPriceRequest](#ixo.bonds.v1beta1.QueryCurrentPriceRequest)
    - [QueryCurrentPriceResponse](#ixo.bonds.v1beta1.QueryCurrentPriceResponse)
    - [QueryCurrentReserveRequest](#ixo.bonds.v1beta1.QueryCurrentReserveRequest)
    - [QueryCurrentReserveResponse](#ixo.bonds.v1beta1.QueryCurrentReserveResponse)
    - [QueryCustomPriceRequest](#ixo.bonds.v1beta1.QueryCustomPriceRequest)
    - [QueryCustomPriceResponse](#ixo.bonds.v1beta1.QueryCustomPriceResponse)
    - [QueryLastBatchRequest](#ixo.bonds.v1beta1.QueryLastBatchRequest)
    - [QueryLastBatchResponse](#ixo.bonds.v1beta1.QueryLastBatchResponse)
    - [QueryParamsRequest](#ixo.bonds.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.bonds.v1beta1.QueryParamsResponse)
    - [QuerySellReturnRequest](#ixo.bonds.v1beta1.QuerySellReturnRequest)
    - [QuerySellReturnResponse](#ixo.bonds.v1beta1.QuerySellReturnResponse)
    - [QuerySwapReturnRequest](#ixo.bonds.v1beta1.QuerySwapReturnRequest)
    - [QuerySwapReturnResponse](#ixo.bonds.v1beta1.QuerySwapReturnResponse)
  
    - [Query](#ixo.bonds.v1beta1.Query)
  
- [ixo/bonds/v1beta1/tx.proto](#ixo/bonds/v1beta1/tx.proto)
    - [MsgBuy](#ixo.bonds.v1beta1.MsgBuy)
    - [MsgBuyResponse](#ixo.bonds.v1beta1.MsgBuyResponse)
    - [MsgCreateBond](#ixo.bonds.v1beta1.MsgCreateBond)
    - [MsgCreateBondResponse](#ixo.bonds.v1beta1.MsgCreateBondResponse)
    - [MsgEditBond](#ixo.bonds.v1beta1.MsgEditBond)
    - [MsgEditBondResponse](#ixo.bonds.v1beta1.MsgEditBondResponse)
    - [MsgMakeOutcomePayment](#ixo.bonds.v1beta1.MsgMakeOutcomePayment)
    - [MsgMakeOutcomePaymentResponse](#ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse)
    - [MsgSell](#ixo.bonds.v1beta1.MsgSell)
    - [MsgSellResponse](#ixo.bonds.v1beta1.MsgSellResponse)
    - [MsgSetNextAlpha](#ixo.bonds.v1beta1.MsgSetNextAlpha)
    - [MsgSetNextAlphaResponse](#ixo.bonds.v1beta1.MsgSetNextAlphaResponse)
    - [MsgSwap](#ixo.bonds.v1beta1.MsgSwap)
    - [MsgSwapResponse](#ixo.bonds.v1beta1.MsgSwapResponse)
    - [MsgUpdateBondState](#ixo.bonds.v1beta1.MsgUpdateBondState)
    - [MsgUpdateBondStateResponse](#ixo.bonds.v1beta1.MsgUpdateBondStateResponse)
    - [MsgWithdrawReserve](#ixo.bonds.v1beta1.MsgWithdrawReserve)
    - [MsgWithdrawReserveResponse](#ixo.bonds.v1beta1.MsgWithdrawReserveResponse)
    - [MsgWithdrawShare](#ixo.bonds.v1beta1.MsgWithdrawShare)
    - [MsgWithdrawShareResponse](#ixo.bonds.v1beta1.MsgWithdrawShareResponse)
  
    - [Msg](#ixo.bonds.v1beta1.Msg)
  
- [ixo/claims/v1beta1/claims.proto](#ixo/claims/v1beta1/claims.proto)
    - [CW20Output](#ixo.claims.v1beta1.CW20Output)
    - [CW20Payment](#ixo.claims.v1beta1.CW20Payment)
    - [Claim](#ixo.claims.v1beta1.Claim)
    - [ClaimPayments](#ixo.claims.v1beta1.ClaimPayments)
    - [Collection](#ixo.claims.v1beta1.Collection)
    - [Contract1155Payment](#ixo.claims.v1beta1.Contract1155Payment)
    - [Dispute](#ixo.claims.v1beta1.Dispute)
    - [DisputeData](#ixo.claims.v1beta1.DisputeData)
    - [Evaluation](#ixo.claims.v1beta1.Evaluation)
    - [Intent](#ixo.claims.v1beta1.Intent)
    - [Params](#ixo.claims.v1beta1.Params)
    - [Payment](#ixo.claims.v1beta1.Payment)
    - [Payments](#ixo.claims.v1beta1.Payments)
  
    - [CollectionIntentOptions](#ixo.claims.v1beta1.CollectionIntentOptions)
    - [CollectionState](#ixo.claims.v1beta1.CollectionState)
    - [EvaluationStatus](#ixo.claims.v1beta1.EvaluationStatus)
    - [IntentStatus](#ixo.claims.v1beta1.IntentStatus)
    - [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus)
    - [PaymentType](#ixo.claims.v1beta1.PaymentType)
  
- [ixo/claims/v1beta1/authz.proto](#ixo/claims/v1beta1/authz.proto)
    - [CreateClaimAuthorizationAuthorization](#ixo.claims.v1beta1.CreateClaimAuthorizationAuthorization)
    - [CreateClaimAuthorizationConstraints](#ixo.claims.v1beta1.CreateClaimAuthorizationConstraints)
    - [EvaluateClaimAuthorization](#ixo.claims.v1beta1.EvaluateClaimAuthorization)
    - [EvaluateClaimConstraints](#ixo.claims.v1beta1.EvaluateClaimConstraints)
    - [SubmitClaimAuthorization](#ixo.claims.v1beta1.SubmitClaimAuthorization)
    - [SubmitClaimConstraints](#ixo.claims.v1beta1.SubmitClaimConstraints)
    - [WithdrawPaymentAuthorization](#ixo.claims.v1beta1.WithdrawPaymentAuthorization)
    - [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints)
  
    - [CreateClaimAuthorizationType](#ixo.claims.v1beta1.CreateClaimAuthorizationType)
  
- [ixo/claims/v1beta1/event.proto](#ixo/claims/v1beta1/event.proto)
    - [ClaimAuthorizationCreatedEvent](#ixo.claims.v1beta1.ClaimAuthorizationCreatedEvent)
    - [ClaimDisputedEvent](#ixo.claims.v1beta1.ClaimDisputedEvent)
    - [ClaimEvaluatedEvent](#ixo.claims.v1beta1.ClaimEvaluatedEvent)
    - [ClaimSubmittedEvent](#ixo.claims.v1beta1.ClaimSubmittedEvent)
    - [ClaimUpdatedEvent](#ixo.claims.v1beta1.ClaimUpdatedEvent)
    - [CollectionCreatedEvent](#ixo.claims.v1beta1.CollectionCreatedEvent)
    - [CollectionUpdatedEvent](#ixo.claims.v1beta1.CollectionUpdatedEvent)
    - [IntentSubmittedEvent](#ixo.claims.v1beta1.IntentSubmittedEvent)
    - [IntentUpdatedEvent](#ixo.claims.v1beta1.IntentUpdatedEvent)
    - [PaymentWithdrawCreatedEvent](#ixo.claims.v1beta1.PaymentWithdrawCreatedEvent)
    - [PaymentWithdrawnEvent](#ixo.claims.v1beta1.PaymentWithdrawnEvent)
  
- [ixo/claims/v1beta1/genesis.proto](#ixo/claims/v1beta1/genesis.proto)
    - [GenesisState](#ixo.claims.v1beta1.GenesisState)
  
- [ixo/claims/v1beta1/query.proto](#ixo/claims/v1beta1/query.proto)
    - [QueryClaimListRequest](#ixo.claims.v1beta1.QueryClaimListRequest)
    - [QueryClaimListResponse](#ixo.claims.v1beta1.QueryClaimListResponse)
    - [QueryClaimRequest](#ixo.claims.v1beta1.QueryClaimRequest)
    - [QueryClaimResponse](#ixo.claims.v1beta1.QueryClaimResponse)
    - [QueryCollectionListRequest](#ixo.claims.v1beta1.QueryCollectionListRequest)
    - [QueryCollectionListResponse](#ixo.claims.v1beta1.QueryCollectionListResponse)
    - [QueryCollectionRequest](#ixo.claims.v1beta1.QueryCollectionRequest)
    - [QueryCollectionResponse](#ixo.claims.v1beta1.QueryCollectionResponse)
    - [QueryDisputeListRequest](#ixo.claims.v1beta1.QueryDisputeListRequest)
    - [QueryDisputeListResponse](#ixo.claims.v1beta1.QueryDisputeListResponse)
    - [QueryDisputeRequest](#ixo.claims.v1beta1.QueryDisputeRequest)
    - [QueryDisputeResponse](#ixo.claims.v1beta1.QueryDisputeResponse)
    - [QueryIntentListRequest](#ixo.claims.v1beta1.QueryIntentListRequest)
    - [QueryIntentListResponse](#ixo.claims.v1beta1.QueryIntentListResponse)
    - [QueryIntentRequest](#ixo.claims.v1beta1.QueryIntentRequest)
    - [QueryIntentResponse](#ixo.claims.v1beta1.QueryIntentResponse)
    - [QueryParamsRequest](#ixo.claims.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.claims.v1beta1.QueryParamsResponse)
  
    - [Query](#ixo.claims.v1beta1.Query)
  
- [ixo/claims/v1beta1/tx.proto](#ixo/claims/v1beta1/tx.proto)
    - [MsgClaimIntent](#ixo.claims.v1beta1.MsgClaimIntent)
    - [MsgClaimIntentResponse](#ixo.claims.v1beta1.MsgClaimIntentResponse)
    - [MsgCreateClaimAuthorization](#ixo.claims.v1beta1.MsgCreateClaimAuthorization)
    - [MsgCreateClaimAuthorizationResponse](#ixo.claims.v1beta1.MsgCreateClaimAuthorizationResponse)
    - [MsgCreateCollection](#ixo.claims.v1beta1.MsgCreateCollection)
    - [MsgCreateCollectionResponse](#ixo.claims.v1beta1.MsgCreateCollectionResponse)
    - [MsgDisputeClaim](#ixo.claims.v1beta1.MsgDisputeClaim)
    - [MsgDisputeClaimResponse](#ixo.claims.v1beta1.MsgDisputeClaimResponse)
    - [MsgEvaluateClaim](#ixo.claims.v1beta1.MsgEvaluateClaim)
    - [MsgEvaluateClaimResponse](#ixo.claims.v1beta1.MsgEvaluateClaimResponse)
    - [MsgSubmitClaim](#ixo.claims.v1beta1.MsgSubmitClaim)
    - [MsgSubmitClaimResponse](#ixo.claims.v1beta1.MsgSubmitClaimResponse)
    - [MsgUpdateCollectionDates](#ixo.claims.v1beta1.MsgUpdateCollectionDates)
    - [MsgUpdateCollectionDatesResponse](#ixo.claims.v1beta1.MsgUpdateCollectionDatesResponse)
    - [MsgUpdateCollectionIntents](#ixo.claims.v1beta1.MsgUpdateCollectionIntents)
    - [MsgUpdateCollectionIntentsResponse](#ixo.claims.v1beta1.MsgUpdateCollectionIntentsResponse)
    - [MsgUpdateCollectionPayments](#ixo.claims.v1beta1.MsgUpdateCollectionPayments)
    - [MsgUpdateCollectionPaymentsResponse](#ixo.claims.v1beta1.MsgUpdateCollectionPaymentsResponse)
    - [MsgUpdateCollectionState](#ixo.claims.v1beta1.MsgUpdateCollectionState)
    - [MsgUpdateCollectionStateResponse](#ixo.claims.v1beta1.MsgUpdateCollectionStateResponse)
    - [MsgWithdrawPayment](#ixo.claims.v1beta1.MsgWithdrawPayment)
    - [MsgWithdrawPaymentResponse](#ixo.claims.v1beta1.MsgWithdrawPaymentResponse)
  
    - [Msg](#ixo.claims.v1beta1.Msg)
  
- [ixo/iid/v1beta1/types.proto](#ixo/iid/v1beta1/types.proto)
    - [AccordedRight](#ixo.iid.v1beta1.AccordedRight)
    - [Context](#ixo.iid.v1beta1.Context)
    - [IidMetadata](#ixo.iid.v1beta1.IidMetadata)
    - [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim)
    - [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity)
    - [LinkedResource](#ixo.iid.v1beta1.LinkedResource)
    - [Service](#ixo.iid.v1beta1.Service)
    - [VerificationMethod](#ixo.iid.v1beta1.VerificationMethod)
  
- [ixo/iid/v1beta1/iid.proto](#ixo/iid/v1beta1/iid.proto)
    - [IidDocument](#ixo.iid.v1beta1.IidDocument)
  
- [ixo/entity/v1beta1/entity.proto](#ixo/entity/v1beta1/entity.proto)
    - [Entity](#ixo.entity.v1beta1.Entity)
    - [EntityAccount](#ixo.entity.v1beta1.EntityAccount)
    - [EntityMetadata](#ixo.entity.v1beta1.EntityMetadata)
    - [Params](#ixo.entity.v1beta1.Params)
  
- [ixo/entity/v1beta1/event.proto](#ixo/entity/v1beta1/event.proto)
    - [EntityAccountAuthzCreatedEvent](#ixo.entity.v1beta1.EntityAccountAuthzCreatedEvent)
    - [EntityAccountAuthzRevokedEvent](#ixo.entity.v1beta1.EntityAccountAuthzRevokedEvent)
    - [EntityAccountCreatedEvent](#ixo.entity.v1beta1.EntityAccountCreatedEvent)
    - [EntityCreatedEvent](#ixo.entity.v1beta1.EntityCreatedEvent)
    - [EntityTransferredEvent](#ixo.entity.v1beta1.EntityTransferredEvent)
    - [EntityUpdatedEvent](#ixo.entity.v1beta1.EntityUpdatedEvent)
    - [EntityVerifiedUpdatedEvent](#ixo.entity.v1beta1.EntityVerifiedUpdatedEvent)
  
- [ixo/entity/v1beta1/genesis.proto](#ixo/entity/v1beta1/genesis.proto)
    - [GenesisState](#ixo.entity.v1beta1.GenesisState)
  
- [ixo/entity/v1beta1/proposal.proto](#ixo/entity/v1beta1/proposal.proto)
    - [InitializeNftContract](#ixo.entity.v1beta1.InitializeNftContract)
  
- [ixo/entity/v1beta1/query.proto](#ixo/entity/v1beta1/query.proto)
    - [QueryEntityIidDocumentRequest](#ixo.entity.v1beta1.QueryEntityIidDocumentRequest)
    - [QueryEntityIidDocumentResponse](#ixo.entity.v1beta1.QueryEntityIidDocumentResponse)
    - [QueryEntityListRequest](#ixo.entity.v1beta1.QueryEntityListRequest)
    - [QueryEntityListResponse](#ixo.entity.v1beta1.QueryEntityListResponse)
    - [QueryEntityMetadataRequest](#ixo.entity.v1beta1.QueryEntityMetadataRequest)
    - [QueryEntityMetadataResponse](#ixo.entity.v1beta1.QueryEntityMetadataResponse)
    - [QueryEntityRequest](#ixo.entity.v1beta1.QueryEntityRequest)
    - [QueryEntityResponse](#ixo.entity.v1beta1.QueryEntityResponse)
    - [QueryEntityVerifiedRequest](#ixo.entity.v1beta1.QueryEntityVerifiedRequest)
    - [QueryEntityVerifiedResponse](#ixo.entity.v1beta1.QueryEntityVerifiedResponse)
    - [QueryParamsRequest](#ixo.entity.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.entity.v1beta1.QueryParamsResponse)
  
    - [Query](#ixo.entity.v1beta1.Query)
  
- [ixo/iid/v1beta1/tx.proto](#ixo/iid/v1beta1/tx.proto)
    - [MsgAddAccordedRight](#ixo.iid.v1beta1.MsgAddAccordedRight)
    - [MsgAddAccordedRightResponse](#ixo.iid.v1beta1.MsgAddAccordedRightResponse)
    - [MsgAddController](#ixo.iid.v1beta1.MsgAddController)
    - [MsgAddControllerResponse](#ixo.iid.v1beta1.MsgAddControllerResponse)
    - [MsgAddIidContext](#ixo.iid.v1beta1.MsgAddIidContext)
    - [MsgAddIidContextResponse](#ixo.iid.v1beta1.MsgAddIidContextResponse)
    - [MsgAddLinkedClaim](#ixo.iid.v1beta1.MsgAddLinkedClaim)
    - [MsgAddLinkedClaimResponse](#ixo.iid.v1beta1.MsgAddLinkedClaimResponse)
    - [MsgAddLinkedEntity](#ixo.iid.v1beta1.MsgAddLinkedEntity)
    - [MsgAddLinkedEntityResponse](#ixo.iid.v1beta1.MsgAddLinkedEntityResponse)
    - [MsgAddLinkedResource](#ixo.iid.v1beta1.MsgAddLinkedResource)
    - [MsgAddLinkedResourceResponse](#ixo.iid.v1beta1.MsgAddLinkedResourceResponse)
    - [MsgAddService](#ixo.iid.v1beta1.MsgAddService)
    - [MsgAddServiceResponse](#ixo.iid.v1beta1.MsgAddServiceResponse)
    - [MsgAddVerification](#ixo.iid.v1beta1.MsgAddVerification)
    - [MsgAddVerificationResponse](#ixo.iid.v1beta1.MsgAddVerificationResponse)
    - [MsgCreateIidDocument](#ixo.iid.v1beta1.MsgCreateIidDocument)
    - [MsgCreateIidDocumentResponse](#ixo.iid.v1beta1.MsgCreateIidDocumentResponse)
    - [MsgDeactivateIID](#ixo.iid.v1beta1.MsgDeactivateIID)
    - [MsgDeactivateIIDResponse](#ixo.iid.v1beta1.MsgDeactivateIIDResponse)
    - [MsgDeleteAccordedRight](#ixo.iid.v1beta1.MsgDeleteAccordedRight)
    - [MsgDeleteAccordedRightResponse](#ixo.iid.v1beta1.MsgDeleteAccordedRightResponse)
    - [MsgDeleteController](#ixo.iid.v1beta1.MsgDeleteController)
    - [MsgDeleteControllerResponse](#ixo.iid.v1beta1.MsgDeleteControllerResponse)
    - [MsgDeleteIidContext](#ixo.iid.v1beta1.MsgDeleteIidContext)
    - [MsgDeleteIidContextResponse](#ixo.iid.v1beta1.MsgDeleteIidContextResponse)
    - [MsgDeleteLinkedClaim](#ixo.iid.v1beta1.MsgDeleteLinkedClaim)
    - [MsgDeleteLinkedClaimResponse](#ixo.iid.v1beta1.MsgDeleteLinkedClaimResponse)
    - [MsgDeleteLinkedEntity](#ixo.iid.v1beta1.MsgDeleteLinkedEntity)
    - [MsgDeleteLinkedEntityResponse](#ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse)
    - [MsgDeleteLinkedResource](#ixo.iid.v1beta1.MsgDeleteLinkedResource)
    - [MsgDeleteLinkedResourceResponse](#ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse)
    - [MsgDeleteService](#ixo.iid.v1beta1.MsgDeleteService)
    - [MsgDeleteServiceResponse](#ixo.iid.v1beta1.MsgDeleteServiceResponse)
    - [MsgRevokeVerification](#ixo.iid.v1beta1.MsgRevokeVerification)
    - [MsgRevokeVerificationResponse](#ixo.iid.v1beta1.MsgRevokeVerificationResponse)
    - [MsgSetVerificationRelationships](#ixo.iid.v1beta1.MsgSetVerificationRelationships)
    - [MsgSetVerificationRelationshipsResponse](#ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse)
    - [MsgUpdateIidDocument](#ixo.iid.v1beta1.MsgUpdateIidDocument)
    - [MsgUpdateIidDocumentResponse](#ixo.iid.v1beta1.MsgUpdateIidDocumentResponse)
    - [Verification](#ixo.iid.v1beta1.Verification)
  
    - [Msg](#ixo.iid.v1beta1.Msg)
  
- [ixo/entity/v1beta1/tx.proto](#ixo/entity/v1beta1/tx.proto)
    - [MsgCreateEntity](#ixo.entity.v1beta1.MsgCreateEntity)
    - [MsgCreateEntityAccount](#ixo.entity.v1beta1.MsgCreateEntityAccount)
    - [MsgCreateEntityAccountResponse](#ixo.entity.v1beta1.MsgCreateEntityAccountResponse)
    - [MsgCreateEntityResponse](#ixo.entity.v1beta1.MsgCreateEntityResponse)
    - [MsgGrantEntityAccountAuthz](#ixo.entity.v1beta1.MsgGrantEntityAccountAuthz)
    - [MsgGrantEntityAccountAuthzResponse](#ixo.entity.v1beta1.MsgGrantEntityAccountAuthzResponse)
    - [MsgRevokeEntityAccountAuthz](#ixo.entity.v1beta1.MsgRevokeEntityAccountAuthz)
    - [MsgRevokeEntityAccountAuthzResponse](#ixo.entity.v1beta1.MsgRevokeEntityAccountAuthzResponse)
    - [MsgTransferEntity](#ixo.entity.v1beta1.MsgTransferEntity)
    - [MsgTransferEntityResponse](#ixo.entity.v1beta1.MsgTransferEntityResponse)
    - [MsgUpdateEntity](#ixo.entity.v1beta1.MsgUpdateEntity)
    - [MsgUpdateEntityResponse](#ixo.entity.v1beta1.MsgUpdateEntityResponse)
    - [MsgUpdateEntityVerified](#ixo.entity.v1beta1.MsgUpdateEntityVerified)
    - [MsgUpdateEntityVerifiedResponse](#ixo.entity.v1beta1.MsgUpdateEntityVerifiedResponse)
  
    - [Msg](#ixo.entity.v1beta1.Msg)
  
- [ixo/epochs/v1beta1/epoch.proto](#ixo/epochs/v1beta1/epoch.proto)
    - [EpochInfo](#ixo.epochs.v1beta1.EpochInfo)
  
- [ixo/epochs/v1beta1/event.proto](#ixo/epochs/v1beta1/event.proto)
    - [EpochEndEvent](#ixo.epochs.v1beta1.EpochEndEvent)
    - [EpochStartEvent](#ixo.epochs.v1beta1.EpochStartEvent)
  
- [ixo/epochs/v1beta1/genesis.proto](#ixo/epochs/v1beta1/genesis.proto)
    - [GenesisState](#ixo.epochs.v1beta1.GenesisState)
  
- [ixo/epochs/v1beta1/query.proto](#ixo/epochs/v1beta1/query.proto)
    - [QueryCurrentEpochRequest](#ixo.epochs.v1beta1.QueryCurrentEpochRequest)
    - [QueryCurrentEpochResponse](#ixo.epochs.v1beta1.QueryCurrentEpochResponse)
    - [QueryEpochsInfoRequest](#ixo.epochs.v1beta1.QueryEpochsInfoRequest)
    - [QueryEpochsInfoResponse](#ixo.epochs.v1beta1.QueryEpochsInfoResponse)
  
    - [Query](#ixo.epochs.v1beta1.Query)
  
- [ixo/iid/v1beta1/event.proto](#ixo/iid/v1beta1/event.proto)
    - [IidDocumentCreatedEvent](#ixo.iid.v1beta1.IidDocumentCreatedEvent)
    - [IidDocumentUpdatedEvent](#ixo.iid.v1beta1.IidDocumentUpdatedEvent)
  
- [ixo/iid/v1beta1/genesis.proto](#ixo/iid/v1beta1/genesis.proto)
    - [GenesisState](#ixo.iid.v1beta1.GenesisState)
  
- [ixo/iid/v1beta1/query.proto](#ixo/iid/v1beta1/query.proto)
    - [QueryIidDocumentRequest](#ixo.iid.v1beta1.QueryIidDocumentRequest)
    - [QueryIidDocumentResponse](#ixo.iid.v1beta1.QueryIidDocumentResponse)
    - [QueryIidDocumentsRequest](#ixo.iid.v1beta1.QueryIidDocumentsRequest)
    - [QueryIidDocumentsResponse](#ixo.iid.v1beta1.QueryIidDocumentsResponse)
  
    - [Query](#ixo.iid.v1beta1.Query)
  
- [ixo/liquidstake/v1beta1/liquidstake.proto](#ixo/liquidstake/v1beta1/liquidstake.proto)
    - [LiquidValidator](#ixo.liquidstake.v1beta1.LiquidValidator)
    - [LiquidValidatorState](#ixo.liquidstake.v1beta1.LiquidValidatorState)
    - [NetAmountState](#ixo.liquidstake.v1beta1.NetAmountState)
    - [Params](#ixo.liquidstake.v1beta1.Params)
    - [WeightedAddress](#ixo.liquidstake.v1beta1.WeightedAddress)
    - [WhitelistedValidator](#ixo.liquidstake.v1beta1.WhitelistedValidator)
  
    - [ValidatorStatus](#ixo.liquidstake.v1beta1.ValidatorStatus)
  
- [ixo/liquidstake/v1beta1/event.proto](#ixo/liquidstake/v1beta1/event.proto)
    - [AddLiquidValidatorEvent](#ixo.liquidstake.v1beta1.AddLiquidValidatorEvent)
    - [AutocompoundStakingRewardsEvent](#ixo.liquidstake.v1beta1.AutocompoundStakingRewardsEvent)
    - [LiquidStakeEvent](#ixo.liquidstake.v1beta1.LiquidStakeEvent)
    - [LiquidStakeParamsUpdatedEvent](#ixo.liquidstake.v1beta1.LiquidStakeParamsUpdatedEvent)
    - [LiquidUnstakeEvent](#ixo.liquidstake.v1beta1.LiquidUnstakeEvent)
    - [RebalancedLiquidStakeEvent](#ixo.liquidstake.v1beta1.RebalancedLiquidStakeEvent)
  
- [ixo/liquidstake/v1beta1/genesis.proto](#ixo/liquidstake/v1beta1/genesis.proto)
    - [GenesisState](#ixo.liquidstake.v1beta1.GenesisState)
  
- [ixo/liquidstake/v1beta1/query.proto](#ixo/liquidstake/v1beta1/query.proto)
    - [QueryLiquidValidatorsRequest](#ixo.liquidstake.v1beta1.QueryLiquidValidatorsRequest)
    - [QueryLiquidValidatorsResponse](#ixo.liquidstake.v1beta1.QueryLiquidValidatorsResponse)
    - [QueryParamsRequest](#ixo.liquidstake.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.liquidstake.v1beta1.QueryParamsResponse)
    - [QueryStatesRequest](#ixo.liquidstake.v1beta1.QueryStatesRequest)
    - [QueryStatesResponse](#ixo.liquidstake.v1beta1.QueryStatesResponse)
  
    - [Query](#ixo.liquidstake.v1beta1.Query)
  
- [ixo/liquidstake/v1beta1/tx.proto](#ixo/liquidstake/v1beta1/tx.proto)
    - [MsgLiquidStake](#ixo.liquidstake.v1beta1.MsgLiquidStake)
    - [MsgLiquidStakeResponse](#ixo.liquidstake.v1beta1.MsgLiquidStakeResponse)
    - [MsgLiquidUnstake](#ixo.liquidstake.v1beta1.MsgLiquidUnstake)
    - [MsgLiquidUnstakeResponse](#ixo.liquidstake.v1beta1.MsgLiquidUnstakeResponse)
    - [MsgSetModulePaused](#ixo.liquidstake.v1beta1.MsgSetModulePaused)
    - [MsgSetModulePausedResponse](#ixo.liquidstake.v1beta1.MsgSetModulePausedResponse)
    - [MsgUpdateParams](#ixo.liquidstake.v1beta1.MsgUpdateParams)
    - [MsgUpdateParamsResponse](#ixo.liquidstake.v1beta1.MsgUpdateParamsResponse)
    - [MsgUpdateWeightedRewardsReceivers](#ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceivers)
    - [MsgUpdateWeightedRewardsReceiversResponse](#ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceiversResponse)
    - [MsgUpdateWhitelistedValidators](#ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidators)
    - [MsgUpdateWhitelistedValidatorsResponse](#ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidatorsResponse)
  
    - [Msg](#ixo.liquidstake.v1beta1.Msg)
  
- [ixo/mint/v1beta1/event.proto](#ixo/mint/v1beta1/event.proto)
    - [MintEpochProvisionsMintedEvent](#ixo.mint.v1beta1.MintEpochProvisionsMintedEvent)
  
- [ixo/mint/v1beta1/mint.proto](#ixo/mint/v1beta1/mint.proto)
    - [DistributionProportions](#ixo.mint.v1beta1.DistributionProportions)
    - [Minter](#ixo.mint.v1beta1.Minter)
    - [Params](#ixo.mint.v1beta1.Params)
    - [WeightedAddress](#ixo.mint.v1beta1.WeightedAddress)
  
- [ixo/mint/v1beta1/genesis.proto](#ixo/mint/v1beta1/genesis.proto)
    - [GenesisState](#ixo.mint.v1beta1.GenesisState)
  
- [ixo/mint/v1beta1/query.proto](#ixo/mint/v1beta1/query.proto)
    - [QueryEpochProvisionsRequest](#ixo.mint.v1beta1.QueryEpochProvisionsRequest)
    - [QueryEpochProvisionsResponse](#ixo.mint.v1beta1.QueryEpochProvisionsResponse)
    - [QueryParamsRequest](#ixo.mint.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.mint.v1beta1.QueryParamsResponse)
  
    - [Query](#ixo.mint.v1beta1.Query)
  
- [ixo/smartaccount/crypto/crypto.proto](#ixo/smartaccount/crypto/crypto.proto)
    - [AuthnPubKey](#ixo.smartaccount.crypto.AuthnPubKey)
  
- [ixo/token/v1beta1/token.proto](#ixo/token/v1beta1/token.proto)
    - [CreditsTransferred](#ixo.token.v1beta1.CreditsTransferred)
    - [Params](#ixo.token.v1beta1.Params)
    - [Token](#ixo.token.v1beta1.Token)
    - [TokenData](#ixo.token.v1beta1.TokenData)
    - [TokenProperties](#ixo.token.v1beta1.TokenProperties)
    - [TokensCancelled](#ixo.token.v1beta1.TokensCancelled)
    - [TokensRetired](#ixo.token.v1beta1.TokensRetired)
  
- [ixo/token/v1beta1/tx.proto](#ixo/token/v1beta1/tx.proto)
    - [MintBatch](#ixo.token.v1beta1.MintBatch)
    - [MsgCancelToken](#ixo.token.v1beta1.MsgCancelToken)
    - [MsgCancelTokenResponse](#ixo.token.v1beta1.MsgCancelTokenResponse)
    - [MsgCreateToken](#ixo.token.v1beta1.MsgCreateToken)
    - [MsgCreateTokenResponse](#ixo.token.v1beta1.MsgCreateTokenResponse)
    - [MsgMintToken](#ixo.token.v1beta1.MsgMintToken)
    - [MsgMintTokenResponse](#ixo.token.v1beta1.MsgMintTokenResponse)
    - [MsgPauseToken](#ixo.token.v1beta1.MsgPauseToken)
    - [MsgPauseTokenResponse](#ixo.token.v1beta1.MsgPauseTokenResponse)
    - [MsgRetireToken](#ixo.token.v1beta1.MsgRetireToken)
    - [MsgRetireTokenResponse](#ixo.token.v1beta1.MsgRetireTokenResponse)
    - [MsgStopToken](#ixo.token.v1beta1.MsgStopToken)
    - [MsgStopTokenResponse](#ixo.token.v1beta1.MsgStopTokenResponse)
    - [MsgTransferCredit](#ixo.token.v1beta1.MsgTransferCredit)
    - [MsgTransferCreditResponse](#ixo.token.v1beta1.MsgTransferCreditResponse)
    - [MsgTransferToken](#ixo.token.v1beta1.MsgTransferToken)
    - [MsgTransferTokenResponse](#ixo.token.v1beta1.MsgTransferTokenResponse)
    - [TokenBatch](#ixo.token.v1beta1.TokenBatch)
  
    - [Msg](#ixo.token.v1beta1.Msg)
  
- [ixo/smartaccount/v1beta1/event.proto](#ixo/smartaccount/v1beta1/event.proto)
    - [AuthenticatorAddedEvent](#ixo.smartaccount.v1beta1.AuthenticatorAddedEvent)
    - [AuthenticatorRemovedEvent](#ixo.smartaccount.v1beta1.AuthenticatorRemovedEvent)
    - [AuthenticatorSetActiveStateEvent](#ixo.smartaccount.v1beta1.AuthenticatorSetActiveStateEvent)
  
- [ixo/smartaccount/v1beta1/params.proto](#ixo/smartaccount/v1beta1/params.proto)
    - [Params](#ixo.smartaccount.v1beta1.Params)
  
- [ixo/smartaccount/v1beta1/models.proto](#ixo/smartaccount/v1beta1/models.proto)
    - [AccountAuthenticator](#ixo.smartaccount.v1beta1.AccountAuthenticator)
  
- [ixo/smartaccount/v1beta1/genesis.proto](#ixo/smartaccount/v1beta1/genesis.proto)
    - [AuthenticatorData](#ixo.smartaccount.v1beta1.AuthenticatorData)
    - [GenesisState](#ixo.smartaccount.v1beta1.GenesisState)
  
- [ixo/smartaccount/v1beta1/query.proto](#ixo/smartaccount/v1beta1/query.proto)
    - [GetAuthenticatorRequest](#ixo.smartaccount.v1beta1.GetAuthenticatorRequest)
    - [GetAuthenticatorResponse](#ixo.smartaccount.v1beta1.GetAuthenticatorResponse)
    - [GetAuthenticatorsRequest](#ixo.smartaccount.v1beta1.GetAuthenticatorsRequest)
    - [GetAuthenticatorsResponse](#ixo.smartaccount.v1beta1.GetAuthenticatorsResponse)
    - [QueryParamsRequest](#ixo.smartaccount.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.smartaccount.v1beta1.QueryParamsResponse)
  
    - [Query](#ixo.smartaccount.v1beta1.Query)
  
- [ixo/smartaccount/v1beta1/tx.proto](#ixo/smartaccount/v1beta1/tx.proto)
    - [MsgAddAuthenticator](#ixo.smartaccount.v1beta1.MsgAddAuthenticator)
    - [MsgAddAuthenticatorResponse](#ixo.smartaccount.v1beta1.MsgAddAuthenticatorResponse)
    - [MsgRemoveAuthenticator](#ixo.smartaccount.v1beta1.MsgRemoveAuthenticator)
    - [MsgRemoveAuthenticatorResponse](#ixo.smartaccount.v1beta1.MsgRemoveAuthenticatorResponse)
    - [MsgSetActiveState](#ixo.smartaccount.v1beta1.MsgSetActiveState)
    - [MsgSetActiveStateResponse](#ixo.smartaccount.v1beta1.MsgSetActiveStateResponse)
    - [TxExtension](#ixo.smartaccount.v1beta1.TxExtension)
  
    - [Msg](#ixo.smartaccount.v1beta1.Msg)
  
- [ixo/token/v1beta1/authz.proto](#ixo/token/v1beta1/authz.proto)
    - [MintAuthorization](#ixo.token.v1beta1.MintAuthorization)
    - [MintConstraints](#ixo.token.v1beta1.MintConstraints)
  
- [ixo/token/v1beta1/event.proto](#ixo/token/v1beta1/event.proto)
    - [CreditsTransferredEvent](#ixo.token.v1beta1.CreditsTransferredEvent)
    - [TokenCancelledEvent](#ixo.token.v1beta1.TokenCancelledEvent)
    - [TokenCreatedEvent](#ixo.token.v1beta1.TokenCreatedEvent)
    - [TokenMintedEvent](#ixo.token.v1beta1.TokenMintedEvent)
    - [TokenPausedEvent](#ixo.token.v1beta1.TokenPausedEvent)
    - [TokenRetiredEvent](#ixo.token.v1beta1.TokenRetiredEvent)
    - [TokenStoppedEvent](#ixo.token.v1beta1.TokenStoppedEvent)
    - [TokenTransferredEvent](#ixo.token.v1beta1.TokenTransferredEvent)
    - [TokenUpdatedEvent](#ixo.token.v1beta1.TokenUpdatedEvent)
  
- [ixo/token/v1beta1/genesis.proto](#ixo/token/v1beta1/genesis.proto)
    - [GenesisState](#ixo.token.v1beta1.GenesisState)
  
- [ixo/token/v1beta1/proposal.proto](#ixo/token/v1beta1/proposal.proto)
    - [SetTokenContractCodes](#ixo.token.v1beta1.SetTokenContractCodes)
  
- [ixo/token/v1beta1/query.proto](#ixo/token/v1beta1/query.proto)
    - [QueryParamsRequest](#ixo.token.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#ixo.token.v1beta1.QueryParamsResponse)
    - [QueryTokenDocRequest](#ixo.token.v1beta1.QueryTokenDocRequest)
    - [QueryTokenDocResponse](#ixo.token.v1beta1.QueryTokenDocResponse)
    - [QueryTokenListRequest](#ixo.token.v1beta1.QueryTokenListRequest)
    - [QueryTokenListResponse](#ixo.token.v1beta1.QueryTokenListResponse)
    - [QueryTokenMetadataRequest](#ixo.token.v1beta1.QueryTokenMetadataRequest)
    - [QueryTokenMetadataResponse](#ixo.token.v1beta1.QueryTokenMetadataResponse)
    - [TokenMetadataProperties](#ixo.token.v1beta1.TokenMetadataProperties)
  
    - [Query](#ixo.token.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="ixo/bonds/v1beta1/bonds.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/bonds.proto



<a name="ixo.bonds.v1beta1.BaseOrder"></a>

### BaseOrder
BaseOrder defines a base order type. It contains all the necessary fields for
specifying the general details about a buy, sell, or swap order.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| cancelled | [bool](#bool) |  |  |
| cancel_reason | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.Batch"></a>

### Batch
Batch holds a collection of outstanding buy, sell, and swap orders on a
particular bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| blocks_remaining | [string](#string) |  |  |
| next_public_alpha | [string](#string) |  |  |
| total_buy_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| total_sell_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| buy_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| sell_prices | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| buys | [BuyOrder](#ixo.bonds.v1beta1.BuyOrder) | repeated |  |
| sells | [SellOrder](#ixo.bonds.v1beta1.SellOrder) | repeated |  |
| swaps | [SwapOrder](#ixo.bonds.v1beta1.SwapOrder) | repeated |  |






<a name="ixo.bonds.v1beta1.Bond"></a>

### Bond
Bond defines a token bonding curve type with all of its parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#ixo.bonds.v1beta1.FunctionParam) | repeated |  |
| reserve_tokens | [string](#string) | repeated |  |
| tx_fee_percentage | [string](#string) |  |  |
| exit_fee_percentage | [string](#string) |  |  |
| fee_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| current_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| current_outcome_payment_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| allow_sells | [bool](#bool) |  |  |
| allow_reserve_withdrawals | [bool](#bool) |  |  |
| alpha_bond | [bool](#bool) |  |  |
| batch_blocks | [string](#string) |  |  |
| outcome_payment | [string](#string) |  |  |
| state | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondDetails"></a>

### BondDetails
BondDetails contains details about the current state of a given bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| spot_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.BuyOrder"></a>

### BuyOrder
BuyOrder defines a type for submitting a buy order on a bond, together with
the maximum amount of reserve tokens the buyer is willing to pay.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#ixo.bonds.v1beta1.BaseOrder) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.FunctionParam"></a>

### FunctionParam
FunctionParam is a key-value pair used for specifying a specific bond
parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| param | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.Params"></a>

### Params
Params defines the parameters for the bonds module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reserved_bond_tokens | [string](#string) | repeated |  |






<a name="ixo.bonds.v1beta1.SellOrder"></a>

### SellOrder
SellOrder defines a type for submitting a sell order on a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#ixo.bonds.v1beta1.BaseOrder) |  |  |






<a name="ixo.bonds.v1beta1.SwapOrder"></a>

### SwapOrder
SwapOrder defines a type for submitting a swap order between two tokens on a
bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_order | [BaseOrder](#ixo.bonds.v1beta1.BaseOrder) |  |  |
| to_token | [string](#string) |  |  |





 

 

 

 



<a name="ixo/bonds/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/event.proto



<a name="ixo.bonds.v1beta1.BondBuyOrderCancelledEvent"></a>

### BondBuyOrderCancelledEvent
BondBuyOrderCancelledEvent is an event triggered on a Bond buy order
cancellation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| order | [BuyOrder](#ixo.bonds.v1beta1.BuyOrder) |  |  |






<a name="ixo.bonds.v1beta1.BondBuyOrderEvent"></a>

### BondBuyOrderEvent
BondBuyOrderEvent is an event triggered on a Bond buy order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [BuyOrder](#ixo.bonds.v1beta1.BuyOrder) |  |  |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondBuyOrderFulfilledEvent"></a>

### BondBuyOrderFulfilledEvent
BondBuyOrderFulfilledEvent is an event triggered on a Bond buy order
fullfillment


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| order | [BuyOrder](#ixo.bonds.v1beta1.BuyOrder) |  |  |
| charged_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| charged_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| returned_to_address | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| new_bond_token_balance | [string](#string) |  |  |
| charged_prices_of_which_reserve | [string](#string) |  |  |
| charged_prices_of_which_funding | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.BondCreatedEvent"></a>

### BondCreatedEvent
BondCreatedEvent is an event triggered on a Bond creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#ixo.bonds.v1beta1.Bond) |  |  |






<a name="ixo.bonds.v1beta1.BondEditAlphaFailedEvent"></a>

### BondEditAlphaFailedEvent
BondEditAlphaFailedEvent is an event triggered on a failed attempt to edit of
Bond alpha value


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| token | [string](#string) |  |  |
| cancel_reason | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondEditAlphaSuccessEvent"></a>

### BondEditAlphaSuccessEvent
BondEditAlphaSuccessEvent is an event triggered on a successful edit of Bond
alpha value


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| token | [string](#string) |  |  |
| public_alpha | [string](#string) |  |  |
| system_alpha | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondMakeOutcomePaymentEvent"></a>

### BondMakeOutcomePaymentEvent
BondMakeOutcomePaymentEvent is an event triggered on a Bond make outcome
payment


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| outcome_payment | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sender_did | [string](#string) |  |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondSellOrderEvent"></a>

### BondSellOrderEvent
BondSellOrderEvent is an event triggered on a Bond sell order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [SellOrder](#ixo.bonds.v1beta1.SellOrder) |  |  |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondSellOrderFulfilledEvent"></a>

### BondSellOrderFulfilledEvent
BondSellOrderFulfilledEvent is an event triggered on a Bond sell order
fullfillment


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| order | [SellOrder](#ixo.bonds.v1beta1.SellOrder) |  |  |
| charged_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| returned_to_address | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| new_bond_token_balance | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondSetNextAlphaEvent"></a>

### BondSetNextAlphaEvent
BondSetNextAlphaEvent is an event triggered when next batch alpha is set


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| next_alpha | [string](#string) |  |  |
| signer | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondSwapOrderEvent"></a>

### BondSwapOrderEvent
BondSwapOrderEvent is an event triggered on a Bond swap order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [SwapOrder](#ixo.bonds.v1beta1.SwapOrder) |  |  |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondSwapOrderFulfilledEvent"></a>

### BondSwapOrderFulfilledEvent
BondSwapOrderFulfilledEvent is an event triggered on a Bond swap order
fullfillment


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| order | [SwapOrder](#ixo.bonds.v1beta1.SwapOrder) |  |  |
| charged_fee | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| returned_to_address | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tokens_swapped | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ixo.bonds.v1beta1.BondUpdatedEvent"></a>

### BondUpdatedEvent
BondUpdatedEvent is an event triggered on a Bond update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#ixo.bonds.v1beta1.Bond) |  |  |






<a name="ixo.bonds.v1beta1.BondWithdrawReserveEvent"></a>

### BondWithdrawReserveEvent
BondWithdrawReserveEvent is an event triggered on a Bond reserve withdrawal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| withdraw_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| withdrawer_did | [string](#string) |  |  |
| withdrawer_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.BondWithdrawShareEvent"></a>

### BondWithdrawShareEvent
BondWithdrawShareEvent is an event triggered on a Bond share withdrawal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| withdraw_payment | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| recipient_did | [string](#string) |  |  |
| recipient_address | [string](#string) |  |  |





 

 

 

 



<a name="ixo/bonds/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/genesis.proto



<a name="ixo.bonds.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the bonds module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [Bond](#ixo.bonds.v1beta1.Bond) | repeated |  |
| batches | [Batch](#ixo.bonds.v1beta1.Batch) | repeated |  |
| params | [Params](#ixo.bonds.v1beta1.Params) |  |  |





 

 

 

 



<a name="ixo/bonds/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/query.proto



<a name="ixo.bonds.v1beta1.QueryAlphaMaximumsRequest"></a>

### QueryAlphaMaximumsRequest
QueryAlphaMaximumsRequest is the request type for the Query/AlphaMaximums RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryAlphaMaximumsResponse"></a>

### QueryAlphaMaximumsResponse
QueryAlphaMaximumsResponse is the response type for the Query/AlphaMaximums
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_system_alpha_increase | [string](#string) |  |  |
| max_system_alpha | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryAvailableReserveRequest"></a>

### QueryAvailableReserveRequest
QueryAvailableReserveRequest is the request type for the
Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryAvailableReserveResponse"></a>

### QueryAvailableReserveResponse
QueryAvailableReserveResponse is the response type for the
Query/AvailableReserve RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| available_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryBatchRequest"></a>

### QueryBatchRequest
QueryBatchRequest is the request type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryBatchResponse"></a>

### QueryBatchResponse
QueryBatchResponse is the response type for the Query/Batch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| batch | [Batch](#ixo.bonds.v1beta1.Batch) |  |  |






<a name="ixo.bonds.v1beta1.QueryBondRequest"></a>

### QueryBondRequest
QueryBondRequest is the request type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryBondResponse"></a>

### QueryBondResponse
QueryBondResponse is the response type for the Query/Bond RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond | [Bond](#ixo.bonds.v1beta1.Bond) |  |  |






<a name="ixo.bonds.v1beta1.QueryBondsDetailedRequest"></a>

### QueryBondsDetailedRequest
QueryBondsDetailedRequest is the request type for the Query/BondsDetailed RPC
method.






<a name="ixo.bonds.v1beta1.QueryBondsDetailedResponse"></a>

### QueryBondsDetailedResponse
QueryBondsDetailedResponse is the response type for the Query/BondsDetailed
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds_detailed | [BondDetails](#ixo.bonds.v1beta1.BondDetails) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryBondsRequest"></a>

### QueryBondsRequest
QueryBondsRequest is the request type for the Query/Bonds RPC method.






<a name="ixo.bonds.v1beta1.QueryBondsResponse"></a>

### QueryBondsResponse
QueryBondsResponse is the response type for the Query/Bonds RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bonds | [string](#string) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryBuyPriceRequest"></a>

### QueryBuyPriceRequest
QueryCustomPriceRequest is the request type for the Query/BuyPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryBuyPriceResponse"></a>

### QueryBuyPriceResponse
QueryCustomPriceResponse is the response type for the Query/BuyPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryCurrentPriceRequest"></a>

### QueryCurrentPriceRequest
QueryCurrentPriceRequest is the request type for the Query/CurrentPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryCurrentPriceResponse"></a>

### QueryCurrentPriceResponse
QueryCurrentPriceResponse is the response type for the Query/CurrentPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryCurrentReserveRequest"></a>

### QueryCurrentReserveRequest
QueryCurrentReserveRequest is the request type for the Query/CurrentReserve
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryCurrentReserveResponse"></a>

### QueryCurrentReserveResponse
QueryCurrentReserveResponse is the response type for the Query/CurrentReserve
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_reserve | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryCustomPriceRequest"></a>

### QueryCustomPriceRequest
QueryCustomPriceRequest is the request type for the Query/CustomPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryCustomPriceResponse"></a>

### QueryCustomPriceResponse
QueryCustomPriceResponse is the response type for the Query/CustomPrice RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="ixo.bonds.v1beta1.QueryLastBatchRequest"></a>

### QueryLastBatchRequest
QueryLastBatchRequest is the request type for the Query/LastBatch RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QueryLastBatchResponse"></a>

### QueryLastBatchResponse
QueryLastBatchResponse is the response type for the Query/LastBatch RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_batch | [Batch](#ixo.bonds.v1beta1.Batch) |  |  |






<a name="ixo.bonds.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="ixo.bonds.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.bonds.v1beta1.Params) |  |  |






<a name="ixo.bonds.v1beta1.QuerySellReturnRequest"></a>

### QuerySellReturnRequest
QuerySellReturnRequest is the request type for the Query/SellReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| bond_amount | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QuerySellReturnResponse"></a>

### QuerySellReturnResponse
QuerySellReturnResponse is the response type for the Query/SellReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adjusted_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| tx_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| exit_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="ixo.bonds.v1beta1.QuerySwapReturnRequest"></a>

### QuerySwapReturnRequest
QuerySwapReturnRequest is the request type for the Query/SwapReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| from_token_with_amount | [string](#string) |  |  |
| to_token | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.QuerySwapReturnResponse"></a>

### QuerySwapReturnResponse
QuerySwapReturnResponse is the response type for the Query/SwapReturn RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_returns | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| total_fees | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 

 

 


<a name="ixo.bonds.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Bonds | [QueryBondsRequest](#ixo.bonds.v1beta1.QueryBondsRequest) | [QueryBondsResponse](#ixo.bonds.v1beta1.QueryBondsResponse) | Bonds returns all existing bonds. |
| BondsDetailed | [QueryBondsDetailedRequest](#ixo.bonds.v1beta1.QueryBondsDetailedRequest) | [QueryBondsDetailedResponse](#ixo.bonds.v1beta1.QueryBondsDetailedResponse) | BondsDetailed returns a list of all existing bonds with some details about their current state. |
| Params | [QueryParamsRequest](#ixo.bonds.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.bonds.v1beta1.QueryParamsResponse) | Params queries the parameters of x/bonds module. |
| Bond | [QueryBondRequest](#ixo.bonds.v1beta1.QueryBondRequest) | [QueryBondResponse](#ixo.bonds.v1beta1.QueryBondResponse) | Bond queries info of a specific bond. |
| Batch | [QueryBatchRequest](#ixo.bonds.v1beta1.QueryBatchRequest) | [QueryBatchResponse](#ixo.bonds.v1beta1.QueryBatchResponse) | Batch queries info of a specific bond&#39;s current batch. |
| LastBatch | [QueryLastBatchRequest](#ixo.bonds.v1beta1.QueryLastBatchRequest) | [QueryLastBatchResponse](#ixo.bonds.v1beta1.QueryLastBatchResponse) | LastBatch queries info of a specific bond&#39;s last batch. |
| CurrentPrice | [QueryCurrentPriceRequest](#ixo.bonds.v1beta1.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#ixo.bonds.v1beta1.QueryCurrentPriceResponse) | CurrentPrice queries the current price/s of a specific bond. |
| CurrentReserve | [QueryCurrentReserveRequest](#ixo.bonds.v1beta1.QueryCurrentReserveRequest) | [QueryCurrentReserveResponse](#ixo.bonds.v1beta1.QueryCurrentReserveResponse) | CurrentReserve queries the current balance/s of the reserve pool for a specific bond. |
| AvailableReserve | [QueryAvailableReserveRequest](#ixo.bonds.v1beta1.QueryAvailableReserveRequest) | [QueryAvailableReserveResponse](#ixo.bonds.v1beta1.QueryAvailableReserveResponse) | AvailableReserve queries current available balance/s of the reserve pool for a specific bond. |
| CustomPrice | [QueryCustomPriceRequest](#ixo.bonds.v1beta1.QueryCustomPriceRequest) | [QueryCustomPriceResponse](#ixo.bonds.v1beta1.QueryCustomPriceResponse) | CustomPrice queries price/s of a specific bond at a specific supply. |
| BuyPrice | [QueryBuyPriceRequest](#ixo.bonds.v1beta1.QueryBuyPriceRequest) | [QueryBuyPriceResponse](#ixo.bonds.v1beta1.QueryBuyPriceResponse) | BuyPrice queries price/s of buying an amount of tokens from a specific bond. |
| SellReturn | [QuerySellReturnRequest](#ixo.bonds.v1beta1.QuerySellReturnRequest) | [QuerySellReturnResponse](#ixo.bonds.v1beta1.QuerySellReturnResponse) | SellReturn queries return/s on selling an amount of tokens of a specific bond. |
| SwapReturn | [QuerySwapReturnRequest](#ixo.bonds.v1beta1.QuerySwapReturnRequest) | [QuerySwapReturnResponse](#ixo.bonds.v1beta1.QuerySwapReturnResponse) | SwapReturn queries return/s on swapping an amount of tokens to another token of a specific bond. |
| AlphaMaximums | [QueryAlphaMaximumsRequest](#ixo.bonds.v1beta1.QueryAlphaMaximumsRequest) | [QueryAlphaMaximumsResponse](#ixo.bonds.v1beta1.QueryAlphaMaximumsResponse) | AlphaMaximums queries alpha maximums for a specific augmented bonding curve. |

 



<a name="ixo/bonds/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/bonds/v1beta1/tx.proto



<a name="ixo.bonds.v1beta1.MsgBuy"></a>

### MsgBuy
MsgBuy defines a message for buying from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buyer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| max_prices | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |
| buyer_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgBuyResponse"></a>

### MsgBuyResponse
MsgBuyResponse defines the Msg/Buy response type.






<a name="ixo.bonds.v1beta1.MsgCreateBond"></a>

### MsgCreateBond
MsgCreateBond defines a message for creating a new bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| function_type | [string](#string) |  |  |
| function_parameters | [FunctionParam](#ixo.bonds.v1beta1.FunctionParam) | repeated |  |
| creator_did | [string](#string) |  |  |
| controller_did | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |
| reserve_tokens | [string](#string) | repeated |  |
| tx_fee_percentage | [string](#string) |  |  |
| exit_fee_percentage | [string](#string) |  |  |
| fee_address | [string](#string) |  |  |
| reserve_withdrawal_address | [string](#string) |  |  |
| max_supply | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| order_quantity_limits | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| allow_sells | [bool](#bool) |  |  |
| allow_reserve_withdrawals | [bool](#bool) |  |  |
| alpha_bond | [bool](#bool) |  |  |
| batch_blocks | [string](#string) |  |  |
| outcome_payment | [string](#string) |  |  |
| creator_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgCreateBondResponse"></a>

### MsgCreateBondResponse
MsgCreateBondResponse defines the Msg/CreateBond response type.






<a name="ixo.bonds.v1beta1.MsgEditBond"></a>

### MsgEditBond
MsgEditBond defines a message for editing an existing bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| order_quantity_limits | [string](#string) |  |  |
| sanity_rate | [string](#string) |  |  |
| sanity_margin_percentage | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |
| editor_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgEditBondResponse"></a>

### MsgEditBondResponse
MsgEditBondResponse defines the Msg/EditBond response type.






<a name="ixo.bonds.v1beta1.MsgMakeOutcomePayment"></a>

### MsgMakeOutcomePayment
MsgMakeOutcomePayment defines a message for making an outcome payment to a
bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_did | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| sender_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse"></a>

### MsgMakeOutcomePaymentResponse
MsgMakeOutcomePaymentResponse defines the Msg/MakeOutcomePayment response
type.






<a name="ixo.bonds.v1beta1.MsgSell"></a>

### MsgSell
MsgSell defines a message for selling from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seller_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| bond_did | [string](#string) |  |  |
| seller_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgSellResponse"></a>

### MsgSellResponse
MsgSellResponse defines the Msg/Sell response type.






<a name="ixo.bonds.v1beta1.MsgSetNextAlpha"></a>

### MsgSetNextAlpha
MsgSetNextAlpha defines a message for editing a bond&#39;s alpha parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| alpha | [string](#string) |  |  |
| delta | [string](#string) |  |  |
| oracle_did | [string](#string) |  |  |
| oracle_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgSetNextAlphaResponse"></a>

### MsgSetNextAlphaResponse







<a name="ixo.bonds.v1beta1.MsgSwap"></a>

### MsgSwap
MsgSwap defines a message for swapping from one reserve bond token to
another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| swapper_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| from | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| to_token | [string](#string) |  |  |
| swapper_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgSwapResponse"></a>

### MsgSwapResponse
MsgSwapResponse defines the Msg/Swap response type.






<a name="ixo.bonds.v1beta1.MsgUpdateBondState"></a>

### MsgUpdateBondState
MsgUpdateBondState defines a message for updating a bond&#39;s current state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bond_did | [string](#string) |  |  |
| state | [string](#string) |  |  |
| editor_did | [string](#string) |  |  |
| editor_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgUpdateBondStateResponse"></a>

### MsgUpdateBondStateResponse
MsgUpdateBondStateResponse defines the Msg/UpdateBondState response type.






<a name="ixo.bonds.v1beta1.MsgWithdrawReserve"></a>

### MsgWithdrawReserve
MsgWithdrawReserve defines a message for withdrawing reserve from a bond.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdrawer_did | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| bond_did | [string](#string) |  |  |
| withdrawer_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgWithdrawReserveResponse"></a>

### MsgWithdrawReserveResponse
MsgWithdrawReserveResponse defines the Msg/WithdrawReserve response type.






<a name="ixo.bonds.v1beta1.MsgWithdrawShare"></a>

### MsgWithdrawShare
MsgWithdrawShare defines a message for withdrawing a share from a bond that
is in the SETTLE stage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient_did | [string](#string) |  |  |
| bond_did | [string](#string) |  |  |
| recipient_address | [string](#string) |  |  |






<a name="ixo.bonds.v1beta1.MsgWithdrawShareResponse"></a>

### MsgWithdrawShareResponse
MsgWithdrawShareResponse defines the Msg/WithdrawShare response type.





 

 

 


<a name="ixo.bonds.v1beta1.Msg"></a>

### Msg
Msg defines the bonds Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBond | [MsgCreateBond](#ixo.bonds.v1beta1.MsgCreateBond) | [MsgCreateBondResponse](#ixo.bonds.v1beta1.MsgCreateBondResponse) | CreateBond defines a method for creating a bond. |
| EditBond | [MsgEditBond](#ixo.bonds.v1beta1.MsgEditBond) | [MsgEditBondResponse](#ixo.bonds.v1beta1.MsgEditBondResponse) | EditBond defines a method for editing a bond. |
| SetNextAlpha | [MsgSetNextAlpha](#ixo.bonds.v1beta1.MsgSetNextAlpha) | [MsgSetNextAlphaResponse](#ixo.bonds.v1beta1.MsgSetNextAlphaResponse) | SetNextAlpha defines a method for editing a bond&#39;s alpha parameter. |
| UpdateBondState | [MsgUpdateBondState](#ixo.bonds.v1beta1.MsgUpdateBondState) | [MsgUpdateBondStateResponse](#ixo.bonds.v1beta1.MsgUpdateBondStateResponse) | UpdateBondState defines a method for updating a bond&#39;s current state. |
| Buy | [MsgBuy](#ixo.bonds.v1beta1.MsgBuy) | [MsgBuyResponse](#ixo.bonds.v1beta1.MsgBuyResponse) | Buy defines a method for buying from a bond. |
| Sell | [MsgSell](#ixo.bonds.v1beta1.MsgSell) | [MsgSellResponse](#ixo.bonds.v1beta1.MsgSellResponse) | Sell defines a method for selling from a bond. |
| Swap | [MsgSwap](#ixo.bonds.v1beta1.MsgSwap) | [MsgSwapResponse](#ixo.bonds.v1beta1.MsgSwapResponse) | Swap defines a method for swapping from one reserve bond token to another. |
| MakeOutcomePayment | [MsgMakeOutcomePayment](#ixo.bonds.v1beta1.MsgMakeOutcomePayment) | [MsgMakeOutcomePaymentResponse](#ixo.bonds.v1beta1.MsgMakeOutcomePaymentResponse) | MakeOutcomePayment defines a method for making an outcome payment to a bond. |
| WithdrawShare | [MsgWithdrawShare](#ixo.bonds.v1beta1.MsgWithdrawShare) | [MsgWithdrawShareResponse](#ixo.bonds.v1beta1.MsgWithdrawShareResponse) | WithdrawShare defines a method for withdrawing a share from a bond that is in the SETTLE stage. |
| WithdrawReserve | [MsgWithdrawReserve](#ixo.bonds.v1beta1.MsgWithdrawReserve) | [MsgWithdrawReserveResponse](#ixo.bonds.v1beta1.MsgWithdrawReserveResponse) | WithdrawReserve defines a method for withdrawing reserve from a bond. |

 



<a name="ixo/claims/v1beta1/claims.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/claims.proto



<a name="ixo.claims.v1beta1.CW20Output"></a>

### CW20Output
CW20Output represents a CW20 token output for split payments


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is the address of the recipient |
| contract_address | [string](#string) |  | contract_address is the address of the contract |
| amount | [uint64](#uint64) |  | amount is the amount of the token to transfer chose uint64 for now as amounts should be small enough to fit in a uint64(max 18446744073709551615) |






<a name="ixo.claims.v1beta1.CW20Payment"></a>

### CW20Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| amount | [uint64](#uint64) |  | chose uint64 for now as amounts should be small enough to fit in a uint64(max 18446744073709551615) |






<a name="ixo.claims.v1beta1.Claim"></a>

### Claim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| agent_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| agent_address | [string](#string) |  |  |
| submission_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submissionDate is the date and time that the claim was submitted on-chain |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim in the cid hash format |
| evaluation | [Evaluation](#ixo.claims.v1beta1.Evaluation) |  | evaluation is the result of one or more claim evaluations |
| payments_status | [ClaimPayments](#ixo.claims.v1beta1.ClaimPayments) |  | payments_status is the status of the payments for the claim |
| use_intent | [bool](#bool) |  | intent_id is the id of the intent for this claim, if any |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by service agent for claim approval NOTE: if both amount and cw20 amount are empty then collection default is used |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | custom cw20 payments specified by service agent for claim approval NOTE: if both amount and cw20 amount are empty then collection default is used |






<a name="ixo.claims.v1beta1.ClaimPayments"></a>

### ClaimPayments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  |  |
| evaluation | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  |  |
| approval | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  |  |
| rejection | [PaymentStatus](#ixo.claims.v1beta1.PaymentStatus) |  | PaymentStatus penalty = 5; |






<a name="ixo.claims.v1beta1.Collection"></a>

### Collection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | collection id is the incremented internal id for the collection of claims |
| entity | [string](#string) |  | entity is the DID of the entity for which the claims are being created |
| admin | [string](#string) |  | admin is the account address that will authorize or revoke agents and payments (the grantor), and can update the collection |
| protocol | [string](#string) |  | protocol is the DID of the claim protocol |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| quota | [uint64](#uint64) |  | quota is the maximum number of claims that may be submitted, 0 is unlimited |
| count | [uint64](#uint64) |  | count is the number of claims already submitted (internally calculated) |
| evaluated | [uint64](#uint64) |  | evaluated is the number of claims that have been evaluated (internally calculated) |
| approved | [uint64](#uint64) |  | approved is the number of claims that have been evaluated and approved (internally calculated) |
| rejected | [uint64](#uint64) |  | rejected is the number of claims that have been evaluated and rejected (internally calculated) |
| disputed | [uint64](#uint64) |  | disputed is the number of claims that have disputed status (internally calculated) |
| state | [CollectionState](#ixo.claims.v1beta1.CollectionState) |  | state is the current state of this Collection (open, paused, closed) |
| payments | [Payments](#ixo.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |
| signer | [string](#string) |  | signer address |
| invalidated | [uint64](#uint64) |  | invalidated is the number of claims that have been evaluated as invalid (internally calculated) |
| escrow_account | [string](#string) |  | escrow_account is the escrow account address for this collection created at collection creation, current purpose is to transfer payments to escrow account for GUARANTEED payments through intents |
| intents | [CollectionIntentOptions](#ixo.claims.v1beta1.CollectionIntentOptions) |  | intents is the option for intents for this collection (allow, deny, required) |






<a name="ixo.claims.v1beta1.Contract1155Payment"></a>

### Contract1155Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| token_id | [string](#string) |  |  |
| amount | [uint32](#uint32) |  |  |






<a name="ixo.claims.v1beta1.Dispute"></a>

### Dispute



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subject_id | [string](#string) |  |  |
| type | [int32](#int32) |  | type is expressed as an integer, interpreted by the client |
| data | [DisputeData](#ixo.claims.v1beta1.DisputeData) |  |  |






<a name="ixo.claims.v1beta1.DisputeData"></a>

### DisputeData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | dispute link ***.ipfs |
| type | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [bool](#bool) |  |  |






<a name="ixo.claims.v1beta1.Evaluation"></a>

### Evaluation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id indicates which Claim this evaluation is for |
| collection_id | [string](#string) |  | collection_id indicates to which Collection the claim being evaluated belongs to |
| oracle | [string](#string) |  | oracle is the DID of the Oracle entity that evaluates the claim |
| agent_did | [string](#string) |  | agent is the DID of the agent that submits the evaluation |
| agent_address | [string](#string) |  |  |
| status | [EvaluationStatus](#ixo.claims.v1beta1.EvaluationStatus) |  | status is the evaluation status expressed as an integer (2=approved, 3=rejected, ...) |
| reason | [uint32](#uint32) |  | reason is the code expressed as an integer, for why the evaluation result was given (codes defined by evaluator) |
| verification_proof | [string](#string) |  | verificationProof is the cid of the evaluation Verfiable Credential |
| evaluation_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | evaluationDate is the date and time that the claim evaluation was submitted on-chain |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by evaluator for claim approval NOTE: if both amount and cw20 amount are empty then collection default is used |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | custom cw20 payments specified by evaluator for claim approval NOTE: if both amount and cw20 amount are empty then collection default is used |






<a name="ixo.claims.v1beta1.Intent"></a>

### Intent
Intent defines the structure for a service agent&#39;s claim intent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the incremented internal id for the intent |
| agent_did | [string](#string) |  | The service agent&#39;s DID (Decentralized Identifier). |
| agent_address | [string](#string) |  | The service agent&#39;s address. |
| collection_id | [string](#string) |  | The id of the collection this intent is linked to. |
| claim_id | [string](#string) |  | claim_id (optional, set when claim is submitted) |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The time the intent was created. |
| expire_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timeout period for the intent. If the claim is not submitted by this time, the intent expires. |
| status | [IntentStatus](#ixo.claims.v1beta1.IntentStatus) |  | Status of the intent (e.g., &#34;ACTIVE&#34; or &#34;FULFILLED&#34;). |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | The payment amount the agent intends to claim, if any. |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | The CW20Payment amount the agent intends to claim, if any. |
| from_address | [string](#string) |  | the address the escrow payment came from |
| escrow_address | [string](#string) |  | the escrow account address |






<a name="ixo.claims.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_sequence | [uint64](#uint64) |  |  |
| ixo_account | [string](#string) |  |  |
| network_fee_percentage | [string](#string) |  |  |
| node_fee_percentage | [string](#string) |  |  |
| intent_sequence | [uint64](#uint64) |  |  |






<a name="ixo.claims.v1beta1.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account is the entity account address from which the payment will be made |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| contract_1155_payment | [Contract1155Payment](#ixo.claims.v1beta1.Contract1155Payment) |  | if empty(nil) then no contract payment, not allowed for Evaluation Payment |
| timeout_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  | timeout after claim/evaluation to create authZ for payment, if 0 then immediate direct payment |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | cw20 payments, can be empty or multiple |
| is_oracle_payment | [bool](#bool) |  | boolean to indicate if the payment is for oracle payments, aka it will go through network fees split, only allowed for APPROVED payment types. NOTE: if true and the payment contains cw20 payments, the claim will only be successfully if an intent exists to ensure immediate cw20 payment split, since there is no WithdrawalAuthorization to manage the cw20 payment split for delayed payments |






<a name="ixo.claims.v1beta1.Payments"></a>

### Payments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| submission | [Payment](#ixo.claims.v1beta1.Payment) |  |  |
| evaluation | [Payment](#ixo.claims.v1beta1.Payment) |  |  |
| approval | [Payment](#ixo.claims.v1beta1.Payment) |  |  |
| rejection | [Payment](#ixo.claims.v1beta1.Payment) |  | Payment penalty = 5; |





 


<a name="ixo.claims.v1beta1.CollectionIntentOptions"></a>

### CollectionIntentOptions


| Name | Number | Description |
| ---- | ------ | ----------- |
| ALLOW | 0 | Allow: Intents can be made for claims, but claims can also be made without intents. |
| DENY | 1 | Deny: Intents cannot be made for claims for the collection. |
| REQUIRED | 2 | Required: Claims cannot be made without an associated intent. An intent is mandatory before a claim can be submitted. |



<a name="ixo.claims.v1beta1.CollectionState"></a>

### CollectionState


| Name | Number | Description |
| ---- | ------ | ----------- |
| OPEN | 0 |  |
| PAUSED | 1 |  |
| CLOSED | 2 |  |



<a name="ixo.claims.v1beta1.EvaluationStatus"></a>

### EvaluationStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 |  |
| APPROVED | 1 |  |
| REJECTED | 2 |  |
| DISPUTED | 3 |  |
| INVALIDATED | 4 |  |



<a name="ixo.claims.v1beta1.IntentStatus"></a>

### IntentStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTIVE | 0 | Active: Intent is created and active, payments have been transferred to escrow if there is any |
| FULFILLED | 1 | Fulfilled: Intent is fulfilled, was used to create a claim and funds will be released on claim APPROVAL, or funds will be reverted on claim REJECTION or DISPUTE |
| EXPIRED | 2 | Expired: Intent has expired, payments have been transferred back out of escrow |



<a name="ixo.claims.v1beta1.PaymentStatus"></a>

### PaymentStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| NO_PAYMENT | 0 |  |
| PROMISED | 1 | Promised: Agent is contracted to receive payment |
| AUTHORIZED | 2 | Authorized: Authz set up, no guarantee |
| GUARANTEED | 3 | Guaranteed: Escrow set up with funds blocked |
| PAID | 4 | Paid: Funds have been paid |
| FAILED | 5 | Failed: Payment failed, most probably due to insufficient funds |
| DISPUTED_PAYMENT | 6 | DisputedPayment: Payment disputed |



<a name="ixo.claims.v1beta1.PaymentType"></a>

### PaymentType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SUBMISSION | 0 |  |
| APPROVAL | 1 |  |
| EVALUATION | 2 |  |
| REJECTION | 3 |  |


 

 

 



<a name="ixo/claims/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/authz.proto



<a name="ixo.claims.v1beta1.CreateClaimAuthorizationAuthorization"></a>

### CreateClaimAuthorizationAuthorization
CreateClaimAuthorizationAuthorization allows a grantee to create
SubmitClaimAuthorization and EvaluateClaimAuthorization for specific
collections(constraints)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin (entity admin module account) |
| constraints | [CreateClaimAuthorizationConstraints](#ixo.claims.v1beta1.CreateClaimAuthorizationConstraints) | repeated | Constraints on the authorizations that can be created |






<a name="ixo.claims.v1beta1.CreateClaimAuthorizationConstraints"></a>

### CreateClaimAuthorizationConstraints
CreateClaimAuthorizationConstraints defines the constraints for creating
claim authorizations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_authorizations | [uint64](#uint64) |  | Maximum number of authorizations that can be created through this meta-authorization, 0 means no quota |
| max_agent_quota | [uint64](#uint64) |  | Maximum quota that can be set in created authorizations 0 means no quota maximum quota per authorization |
| max_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Maximum amount that can be set in created authorizations, if empty then any custom amount is allowed in the created authorizations explicitly set to 0 to disallow any custom amount in the created authorizations |
| max_cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | Maximum cw20 payment that can be set in created authorizations, if empty then any cw20 payment is allowed in the created authorizations explicitly set to 0 to disallow any cw20 payment in the created authorizations |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Expiration of this meta-authorization(specific constraint), if not set then no expiration |
| collection_ids | [string](#string) | repeated | Collection IDs the grantee can create authorizations for, if empty then all collections for the admin are allowed |
| allowed_auth_types | [CreateClaimAuthorizationType](#ixo.claims.v1beta1.CreateClaimAuthorizationType) |  | Types of authorizations the grantee can create (submit, evaluate, or all(both)) |
| max_intent_duration_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  | Maximum intent duration for the authorization allowed (for submit) |






<a name="ixo.claims.v1beta1.EvaluateClaimAuthorization"></a>

### EvaluateClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin (entity admin module account) |
| constraints | [EvaluateClaimConstraints](#ixo.claims.v1beta1.EvaluateClaimConstraints) | repeated |  |






<a name="ixo.claims.v1beta1.EvaluateClaimConstraints"></a>

### EvaluateClaimConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| claim_ids | [string](#string) | repeated | either collection_id or claim_ids is needed |
| agent_quota | [uint64](#uint64) |  |  |
| before_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | if null then no before_date validation done |
| max_custom_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | max custom amount evaluator can change, if empty then no custom amount is allowed, and default payments from Collection payments are used |
| max_custom_cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | max custom cw20 payment evaluator can change, if empty then no custom amount is allowed, and default payments from Collection payments are used |






<a name="ixo.claims.v1beta1.SubmitClaimAuthorization"></a>

### SubmitClaimAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin (entity admin module account) |
| constraints | [SubmitClaimConstraints](#ixo.claims.v1beta1.SubmitClaimConstraints) | repeated |  |






<a name="ixo.claims.v1beta1.SubmitClaimConstraints"></a>

### SubmitClaimConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| agent_quota | [uint64](#uint64) |  |  |
| max_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom max_amount allowed to be specified by service agent for claim approval, if empty then no custom amount is allowed, and default payments from Collection payments are used |
| max_cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | custom max_cw20_payment allowed to be specified by service agent for claim approval, if empty then no custom amount is allowed, and default payments from Collection payments are used |
| intent_duration_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  | intent_duration_ns is the duration for which the intent is active, after which it will expire (in nanoseconds) |






<a name="ixo.claims.v1beta1.WithdrawPaymentAuthorization"></a>

### WithdrawPaymentAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | address of admin |
| constraints | [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints) | repeated |  |






<a name="ixo.claims.v1beta1.WithdrawPaymentConstraints"></a>

### WithdrawPaymentConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id the withdrawal is for |
| inputs | [cosmos.bank.v1beta1.Input](#cosmos.bank.v1beta1.Input) | repeated | Inputs to the multisend tx to run to withdraw payment |
| outputs | [cosmos.bank.v1beta1.Output](#cosmos.bank.v1beta1.Output) | repeated | Outputs for the multisend tx to run to withdraw payment |
| payment_type | [PaymentType](#ixo.claims.v1beta1.PaymentType) |  | payment type to keep track what payment is for and mark claim payment accordingly |
| contract_1155_payment | [Contract1155Payment](#ixo.claims.v1beta1.Contract1155Payment) |  | if empty(nil) then no contract payment |
| toAddress | [string](#string) |  | for contract payment |
| fromAddress | [string](#string) |  | for contract payment |
| release_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date that grantee can execute authorization, calculated from created date plus the timeout on Collection payments, if null then none |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | cw20 payments, can be empty or multiple |





 


<a name="ixo.claims.v1beta1.CreateClaimAuthorizationType"></a>

### CreateClaimAuthorizationType
AuthorizationType defines the types of claim authorizations that can be
created

| Name | Number | Description |
| ---- | ------ | ----------- |
| ALL | 0 | both submit and evaluate |
| SUBMIT | 1 | submit only |
| EVALUATE | 2 | evaluate only |


 

 

 



<a name="ixo/claims/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/event.proto



<a name="ixo.claims.v1beta1.ClaimAuthorizationCreatedEvent"></a>

### ClaimAuthorizationCreatedEvent
ClaimAuthorizationCreatedEvent is an event triggered on a Claim authorization
creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator | [string](#string) |  |  |
| creator_did | [string](#string) |  |  |
| grantee | [string](#string) |  |  |
| admin | [string](#string) |  |  |
| collection_id | [string](#string) |  |  |
| auth_type | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.ClaimDisputedEvent"></a>

### ClaimDisputedEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dispute | [Dispute](#ixo.claims.v1beta1.Dispute) |  |  |






<a name="ixo.claims.v1beta1.ClaimEvaluatedEvent"></a>

### ClaimEvaluatedEvent
ClaimEvaluatedEvent is an event triggered on a Claim evaluation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| evaluation | [Evaluation](#ixo.claims.v1beta1.Evaluation) |  |  |






<a name="ixo.claims.v1beta1.ClaimSubmittedEvent"></a>

### ClaimSubmittedEvent
CollectionCreatedEvent is an event triggered on a Claim submission


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#ixo.claims.v1beta1.Claim) |  |  |






<a name="ixo.claims.v1beta1.ClaimUpdatedEvent"></a>

### ClaimUpdatedEvent
ClaimUpdatedEvent is an event triggered on a Claim update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#ixo.claims.v1beta1.Claim) |  |  |






<a name="ixo.claims.v1beta1.CollectionCreatedEvent"></a>

### CollectionCreatedEvent
CollectionCreatedEvent is an event triggered on a Collection creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#ixo.claims.v1beta1.Collection) |  |  |






<a name="ixo.claims.v1beta1.CollectionUpdatedEvent"></a>

### CollectionUpdatedEvent
CollectionUpdatedEvent is an event triggered on a Collection update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#ixo.claims.v1beta1.Collection) |  |  |






<a name="ixo.claims.v1beta1.IntentSubmittedEvent"></a>

### IntentSubmittedEvent
IntentSubmittedEvent is an event triggered on an Intent submission


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| intent | [Intent](#ixo.claims.v1beta1.Intent) |  |  |






<a name="ixo.claims.v1beta1.IntentUpdatedEvent"></a>

### IntentUpdatedEvent
IntentUpdatedEvent is an event triggered on an Intent update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| intent | [Intent](#ixo.claims.v1beta1.Intent) |  |  |






<a name="ixo.claims.v1beta1.PaymentWithdrawCreatedEvent"></a>

### PaymentWithdrawCreatedEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw | [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints) |  |  |






<a name="ixo.claims.v1beta1.PaymentWithdrawnEvent"></a>

### PaymentWithdrawnEvent
ClaimDisputedEvent is an event triggered on a Claim dispute


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw | [WithdrawPaymentConstraints](#ixo.claims.v1beta1.WithdrawPaymentConstraints) |  |  |
| cw20_outputs | [CW20Output](#ixo.claims.v1beta1.CW20Output) | repeated |  |





 

 

 

 



<a name="ixo/claims/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/genesis.proto



<a name="ixo.claims.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the claims module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.claims.v1beta1.Params) |  |  |
| collections | [Collection](#ixo.claims.v1beta1.Collection) | repeated |  |
| claims | [Claim](#ixo.claims.v1beta1.Claim) | repeated |  |
| disputes | [Dispute](#ixo.claims.v1beta1.Dispute) | repeated |  |
| intents | [Intent](#ixo.claims.v1beta1.Intent) | repeated |  |





 

 

 

 



<a name="ixo/claims/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/query.proto



<a name="ixo.claims.v1beta1.QueryClaimListRequest"></a>

### QueryClaimListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryClaimListResponse"></a>

### QueryClaimListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claims | [Claim](#ixo.claims.v1beta1.Claim) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryClaimRequest"></a>

### QueryClaimRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryClaimResponse"></a>

### QueryClaimResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim | [Claim](#ixo.claims.v1beta1.Claim) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionListRequest"></a>

### QueryCollectionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionListResponse"></a>

### QueryCollectionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [Collection](#ixo.claims.v1beta1.Collection) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionRequest"></a>

### QueryCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryCollectionResponse"></a>

### QueryCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#ixo.claims.v1beta1.Collection) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeListRequest"></a>

### QueryDisputeListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeListResponse"></a>

### QueryDisputeListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disputes | [Dispute](#ixo.claims.v1beta1.Dispute) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeRequest"></a>

### QueryDisputeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proof | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryDisputeResponse"></a>

### QueryDisputeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dispute | [Dispute](#ixo.claims.v1beta1.Dispute) |  |  |






<a name="ixo.claims.v1beta1.QueryIntentListRequest"></a>

### QueryIntentListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ixo.claims.v1beta1.QueryIntentListResponse"></a>

### QueryIntentListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| intents | [Intent](#ixo.claims.v1beta1.Intent) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.claims.v1beta1.QueryIntentRequest"></a>

### QueryIntentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| agentAddress | [string](#string) |  |  |
| collectionId | [string](#string) |  |  |






<a name="ixo.claims.v1beta1.QueryIntentResponse"></a>

### QueryIntentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| intent | [Intent](#ixo.claims.v1beta1.Intent) |  |  |






<a name="ixo.claims.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ixo.claims.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.claims.v1beta1.Params) |  | params holds all the parameters of this module. |





 

 

 


<a name="ixo.claims.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.claims.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.claims.v1beta1.QueryParamsResponse) | Parameters queries the parameters of the module. |
| Collection | [QueryCollectionRequest](#ixo.claims.v1beta1.QueryCollectionRequest) | [QueryCollectionResponse](#ixo.claims.v1beta1.QueryCollectionResponse) |  |
| CollectionList | [QueryCollectionListRequest](#ixo.claims.v1beta1.QueryCollectionListRequest) | [QueryCollectionListResponse](#ixo.claims.v1beta1.QueryCollectionListResponse) |  |
| Claim | [QueryClaimRequest](#ixo.claims.v1beta1.QueryClaimRequest) | [QueryClaimResponse](#ixo.claims.v1beta1.QueryClaimResponse) |  |
| ClaimList | [QueryClaimListRequest](#ixo.claims.v1beta1.QueryClaimListRequest) | [QueryClaimListResponse](#ixo.claims.v1beta1.QueryClaimListResponse) |  |
| Dispute | [QueryDisputeRequest](#ixo.claims.v1beta1.QueryDisputeRequest) | [QueryDisputeResponse](#ixo.claims.v1beta1.QueryDisputeResponse) |  |
| DisputeList | [QueryDisputeListRequest](#ixo.claims.v1beta1.QueryDisputeListRequest) | [QueryDisputeListResponse](#ixo.claims.v1beta1.QueryDisputeListResponse) |  |
| Intent | [QueryIntentRequest](#ixo.claims.v1beta1.QueryIntentRequest) | [QueryIntentResponse](#ixo.claims.v1beta1.QueryIntentResponse) |  |
| IntentList | [QueryIntentListRequest](#ixo.claims.v1beta1.QueryIntentListRequest) | [QueryIntentListResponse](#ixo.claims.v1beta1.QueryIntentListResponse) |  |

 



<a name="ixo/claims/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/claims/v1beta1/tx.proto



<a name="ixo.claims.v1beta1.MsgClaimIntent"></a>

### MsgClaimIntent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agent_did | [string](#string) |  | The service agent&#39;s DID (Decentralized Identifier). |
| agent_address | [string](#string) |  | The service agent&#39;s address (who submits this message). |
| collection_id | [string](#string) |  | The id of the collection this intent is linked to. |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | The desired claim amount, if any. NOTE: if both amount and cw20 amount are empty then default by Collection is used (APPROVAL payment). |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | The custom CW20 payment, if any. NOTE: if both amount and cw20 amount are empty then default by Collection is used (APPROVAL payment). |






<a name="ixo.claims.v1beta1.MsgClaimIntentResponse"></a>

### MsgClaimIntentResponse
MsgClaimIntentResponse defines the response after submitting an intent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| intent_id | [string](#string) |  | Resulting intent id. |
| expire_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timeout period for the intent. If the claim is not submitted by this time, the intent expires. |






<a name="ixo.claims.v1beta1.MsgCreateClaimAuthorization"></a>

### MsgCreateClaimAuthorization
MsgCreateClaimAuthorization defines a message for creating a claim
authorization on behalf of an entity admin account (SubmitClaimAuthorization
or EvaluateClaimAuthorization)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| creator_address | [string](#string) |  | Address of the creator (user with meta-authorization) |
| creator_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| grantee_address | [string](#string) |  | Address of the grantee (who will receive the authorization) |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |
| collection_id | [string](#string) |  | Collection ID the authorization applies to (for both submit and evaluate) |
| auth_type | [CreateClaimAuthorizationType](#ixo.claims.v1beta1.CreateClaimAuthorizationType) |  | Type of authorization to create (submit or evaluate, can&#39;t create both in a single request) |
| agent_quota | [uint64](#uint64) |  | Quota for the created authorization (for both submit and evaluate) |
| max_amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Maximum amount that can be specified in the authorization (for both submit and evaluate) |
| max_cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | Maximum CW20 payment that can be specified in the authorization (for both submit and evaluate) |
| expiration | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Expiration time for the authorization, be careful with this as it is the expiration of the authorization itself, not the constraints, meaning if the authorization expires all constraints will be removed with the authorization (standard authz behavior) |
| intent_duration_ns | [google.protobuf.Duration](#google.protobuf.Duration) |  | Maximum intent duration for the authorization allowed (for submit) |
| before_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | if null then no before_date validation done (for evaluate) |






<a name="ixo.claims.v1beta1.MsgCreateClaimAuthorizationResponse"></a>

### MsgCreateClaimAuthorizationResponse
MsgCreateClaimAuthorizationResponse defines the response for creating a claim
authorization






<a name="ixo.claims.v1beta1.MsgCreateCollection"></a>

### MsgCreateCollection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [string](#string) |  | entity is the DID of the entity for which the claims are being created |
| signer | [string](#string) |  | signer address |
| protocol | [string](#string) |  | protocol is the DID of the claim protocol |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| quota | [uint64](#uint64) |  | quota is the maximum number of claims that may be submitted, 0 is unlimited |
| state | [CollectionState](#ixo.claims.v1beta1.CollectionState) |  | state is the current state of this Collection (open, paused, closed) |
| payments | [Payments](#ixo.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |
| intents | [CollectionIntentOptions](#ixo.claims.v1beta1.CollectionIntentOptions) |  | intents is the option for intents for this collection (allow, deny, required) |






<a name="ixo.claims.v1beta1.MsgCreateCollectionResponse"></a>

### MsgCreateCollectionResponse







<a name="ixo.claims.v1beta1.MsgDisputeClaim"></a>

### MsgDisputeClaim
Agent laying dispute must be admin for Collection, or controller on
Collection entity, or have authz cap, aka is agent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subject_id | [string](#string) |  | subject_id for which this dispute is against, for now can only lay disputes against claims |
| agent_did | [string](#string) |  | agent is the DID of the agent disputing the claim, agent details won&#39;t be saved in kvStore |
| agent_address | [string](#string) |  |  |
| dispute_type | [int32](#int32) |  | type is expressed as an integer, interpreted by the client |
| data | [DisputeData](#ixo.claims.v1beta1.DisputeData) |  |  |






<a name="ixo.claims.v1beta1.MsgDisputeClaimResponse"></a>

### MsgDisputeClaimResponse







<a name="ixo.claims.v1beta1.MsgEvaluateClaim"></a>

### MsgEvaluateClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim to make evaluation against |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| oracle | [string](#string) |  | oracle is the DID of the Oracle entity that evaluates the claim |
| agent_did | [string](#string) |  | agent is the DID of the agent that submits the evaluation |
| agent_address | [string](#string) |  |  |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |
| status | [EvaluationStatus](#ixo.claims.v1beta1.EvaluationStatus) |  | status is the evaluation status expressed as an integer (2=approved, 3=rejected, ...) |
| reason | [uint32](#uint32) |  | reason is the code expressed as an integer, for why the evaluation result was given (codes defined by evaluator) |
| verification_proof | [string](#string) |  | verificationProof is the cid of the evaluation Verifiable Credential |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by evaluator for claim approval NOTE: if claim is using intent, then amount and cw20 amount are ignored and overridden with intent amounts NOTE: if both amount and cw20 amount are empty then collection default is used |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | custom cw20 payments specified by evaluator for claim approval NOTE: if claim is using intent, then amount and cw20 amount are ignored and overridden with intent amounts NOTE: if both amount and cw20 amount are empty then collection default is used |






<a name="ixo.claims.v1beta1.MsgEvaluateClaimResponse"></a>

### MsgEvaluateClaimResponse







<a name="ixo.claims.v1beta1.MsgSubmitClaim"></a>

### MsgSubmitClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates to which Collection this claim belongs |
| claim_id | [string](#string) |  | claimID is the unique identifier of the claim in the cid hash format |
| agent_did | [string](#string) |  | agent is the DID of the agent submitting the claim |
| agent_address | [string](#string) |  |  |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |
| use_intent | [bool](#bool) |  | use_intent is the option for using intent for this claim if it exists and is active. NOTE: if use_intent is true then amount and cw20 amount are ignored and overridden with intent amounts. NOTE: if use_intent is true and there is no active intent then will error |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | custom amount specified by service agent for claim approval NOTE: if both amount and cw20_payment are empty then collection default is used |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | custom cw20 payments specified by service agent for claim approval NOTE: if both amount and cw20 amount are empty then collection default is used |






<a name="ixo.claims.v1beta1.MsgSubmitClaimResponse"></a>

### MsgSubmitClaimResponse







<a name="ixo.claims.v1beta1.MsgUpdateCollectionDates"></a>

### MsgUpdateCollectionDates



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates which Collection to update |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | startDate is the date after which claims may be submitted |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | endDate is the date after which no more claims may be submitted (no endDate is allowed) |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="ixo.claims.v1beta1.MsgUpdateCollectionDatesResponse"></a>

### MsgUpdateCollectionDatesResponse







<a name="ixo.claims.v1beta1.MsgUpdateCollectionIntents"></a>

### MsgUpdateCollectionIntents



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates which Collection to update |
| intents | [CollectionIntentOptions](#ixo.claims.v1beta1.CollectionIntentOptions) |  | intents is the option for intents for this collection (allow, deny, required) |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="ixo.claims.v1beta1.MsgUpdateCollectionIntentsResponse"></a>

### MsgUpdateCollectionIntentsResponse







<a name="ixo.claims.v1beta1.MsgUpdateCollectionPayments"></a>

### MsgUpdateCollectionPayments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates which Collection to update |
| payments | [Payments](#ixo.claims.v1beta1.Payments) |  | payments is the amount paid for claim submission, evaluation, approval, or rejection |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="ixo.claims.v1beta1.MsgUpdateCollectionPaymentsResponse"></a>

### MsgUpdateCollectionPaymentsResponse







<a name="ixo.claims.v1beta1.MsgUpdateCollectionState"></a>

### MsgUpdateCollectionState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | collection_id indicates which Collection to update |
| state | [CollectionState](#ixo.claims.v1beta1.CollectionState) |  | state is the state of this Collection (open, paused, closed) you want to update to |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |






<a name="ixo.claims.v1beta1.MsgUpdateCollectionStateResponse"></a>

### MsgUpdateCollectionStateResponse







<a name="ixo.claims.v1beta1.MsgWithdrawPayment"></a>

### MsgWithdrawPayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| claim_id | [string](#string) |  | claim_id the withdrawal is for |
| inputs | [cosmos.bank.v1beta1.Input](#cosmos.bank.v1beta1.Input) | repeated | Inputs to the multi send tx to run to withdraw payment |
| outputs | [cosmos.bank.v1beta1.Output](#cosmos.bank.v1beta1.Output) | repeated | Outputs for the multi send tx to run to withdraw payment |
| payment_type | [PaymentType](#ixo.claims.v1beta1.PaymentType) |  | payment type to keep track what payment is for and mark claim payment accordingly |
| contract_1155_payment | [Contract1155Payment](#ixo.claims.v1beta1.Contract1155Payment) |  | if empty(nil) then no contract payment |
| toAddress | [string](#string) |  | for contract payment |
| fromAddress | [string](#string) |  | for contract payment |
| release_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | date that grantee can execute authorization, calculated from created date plus the timeout on Collection payments |
| admin_address | [string](#string) |  | admin address used to sign this message, validated against Collection Admin |
| cw20_payment | [CW20Payment](#ixo.claims.v1beta1.CW20Payment) | repeated | cw20 payments, can be empty or multiple |






<a name="ixo.claims.v1beta1.MsgWithdrawPaymentResponse"></a>

### MsgWithdrawPaymentResponse






 

 

 


<a name="ixo.claims.v1beta1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCollection | [MsgCreateCollection](#ixo.claims.v1beta1.MsgCreateCollection) | [MsgCreateCollectionResponse](#ixo.claims.v1beta1.MsgCreateCollectionResponse) |  |
| SubmitClaim | [MsgSubmitClaim](#ixo.claims.v1beta1.MsgSubmitClaim) | [MsgSubmitClaimResponse](#ixo.claims.v1beta1.MsgSubmitClaimResponse) |  |
| EvaluateClaim | [MsgEvaluateClaim](#ixo.claims.v1beta1.MsgEvaluateClaim) | [MsgEvaluateClaimResponse](#ixo.claims.v1beta1.MsgEvaluateClaimResponse) |  |
| DisputeClaim | [MsgDisputeClaim](#ixo.claims.v1beta1.MsgDisputeClaim) | [MsgDisputeClaimResponse](#ixo.claims.v1beta1.MsgDisputeClaimResponse) |  |
| WithdrawPayment | [MsgWithdrawPayment](#ixo.claims.v1beta1.MsgWithdrawPayment) | [MsgWithdrawPaymentResponse](#ixo.claims.v1beta1.MsgWithdrawPaymentResponse) |  |
| UpdateCollectionState | [MsgUpdateCollectionState](#ixo.claims.v1beta1.MsgUpdateCollectionState) | [MsgUpdateCollectionStateResponse](#ixo.claims.v1beta1.MsgUpdateCollectionStateResponse) |  |
| UpdateCollectionDates | [MsgUpdateCollectionDates](#ixo.claims.v1beta1.MsgUpdateCollectionDates) | [MsgUpdateCollectionDatesResponse](#ixo.claims.v1beta1.MsgUpdateCollectionDatesResponse) |  |
| UpdateCollectionPayments | [MsgUpdateCollectionPayments](#ixo.claims.v1beta1.MsgUpdateCollectionPayments) | [MsgUpdateCollectionPaymentsResponse](#ixo.claims.v1beta1.MsgUpdateCollectionPaymentsResponse) |  |
| UpdateCollectionIntents | [MsgUpdateCollectionIntents](#ixo.claims.v1beta1.MsgUpdateCollectionIntents) | [MsgUpdateCollectionIntentsResponse](#ixo.claims.v1beta1.MsgUpdateCollectionIntentsResponse) |  |
| ClaimIntent | [MsgClaimIntent](#ixo.claims.v1beta1.MsgClaimIntent) | [MsgClaimIntentResponse](#ixo.claims.v1beta1.MsgClaimIntentResponse) |  |
| CreateClaimAuthorization | [MsgCreateClaimAuthorization](#ixo.claims.v1beta1.MsgCreateClaimAuthorization) | [MsgCreateClaimAuthorizationResponse](#ixo.claims.v1beta1.MsgCreateClaimAuthorizationResponse) |  |

 



<a name="ixo/iid/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/types.proto



<a name="ixo.iid.v1beta1.AccordedRight"></a>

### AccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| mechanism | [string](#string) |  |  |
| message | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.Context"></a>

### Context



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| val | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.IidMetadata"></a>

### IidMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versionId | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| deactivated | [bool](#bool) |  |  |






<a name="ixo.iid.v1beta1.LinkedClaim"></a>

### LinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [string](#string) |  |  |
| right | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.LinkedEntity"></a>

### LinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| relationship | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.LinkedResource"></a>

### LinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| mediaType | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |
| proof | [string](#string) |  |  |
| encrypted | [string](#string) |  |  |
| right | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.Service"></a>

### Service



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| serviceEndpoint | [string](#string) |  |  |






<a name="ixo.iid.v1beta1.VerificationMethod"></a>

### VerificationMethod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| controller | [string](#string) |  |  |
| blockchainAccountID | [string](#string) |  |  |
| publicKeyHex | [string](#string) |  |  |
| publicKeyMultibase | [string](#string) |  |  |
| publicKeyBase58 | [string](#string) |  |  |





 

 

 

 



<a name="ixo/iid/v1beta1/iid.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/iid.proto



<a name="ixo.iid.v1beta1.IidDocument"></a>

### IidDocument
type entity account
relationship entity account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| context | [Context](#ixo.iid.v1beta1.Context) | repeated | @context is spec for did document. |
| id | [string](#string) |  | id represents the id for the did document. |
| controller | [string](#string) | repeated | A DID controller is an entity that is authorized to make changes to a DID document. cfr. https://www.w3.org/TR/did-core/#did-controller |
| verificationMethod | [VerificationMethod](#ixo.iid.v1beta1.VerificationMethod) | repeated | A DID document can express verification methods, such as cryptographic public keys, which can be used to authenticate or authorize interactions with the DID subject or associated parties. https://www.w3.org/TR/did-core/#verification-methods |
| service | [Service](#ixo.iid.v1beta1.Service) | repeated | Services are used in DID documents to express ways of communicating with the DID subject or associated entities. https://www.w3.org/TR/did-core/#services |
| authentication | [string](#string) | repeated | NOTE: below this line there are the relationships Authentication represents public key associated with the did document. cfr. https://www.w3.org/TR/did-core/#authentication |
| assertionMethod | [string](#string) | repeated | Used to specify how the DID subject is expected to express claims, such as for the purposes of issuing a Verifiable Credential. cfr. https://www.w3.org/TR/did-core/#assertion |
| keyAgreement | [string](#string) | repeated | used to specify how an entity can generate encryption material in order to transmit confidential information intended for the DID subject. https://www.w3.org/TR/did-core/#key-agreement |
| capabilityInvocation | [string](#string) | repeated | Used to specify a verification method that might be used by the DID subject to invoke a cryptographic capability, such as the authorization to update the DID Document. https://www.w3.org/TR/did-core/#capability-invocation |
| capabilityDelegation | [string](#string) | repeated | Used to specify a mechanism that might be used by the DID subject to delegate a cryptographic capability to another party. https://www.w3.org/TR/did-core/#capability-delegation |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated |  |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated |  |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated |  |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| metadata | [IidMetadata](#ixo.iid.v1beta1.IidMetadata) |  | Metadata concerning the IidDocument such as versionId, created, updated and deactivated |





 

 

 

 



<a name="ixo/entity/v1beta1/entity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/entity.proto



<a name="ixo.entity.v1beta1.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id represents the id for the entity document. |
| type | [string](#string) |  | Type of entity, eg protocol or asset |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| relayer_node | [string](#string) |  | Did of the operator through which the Entity was created |
| credentials | [string](#string) | repeated | Credentials of the entity to be verified |
| entity_verified | [bool](#bool) |  | Used as check whether the credentials of entity is verified |
| metadata | [EntityMetadata](#ixo.entity.v1beta1.EntityMetadata) |  | Metadata concerning the Entity such as versionId, created, updated and deactivated |
| accounts | [EntityAccount](#ixo.entity.v1beta1.EntityAccount) | repeated | module accounts created for entity |






<a name="ixo.entity.v1beta1.EntityAccount"></a>

### EntityAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| address | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityMetadata"></a>

### EntityMetadata
EntityMetadata defines metadata associated to a entity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version_id | [string](#string) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ixo.entity.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nftContractAddress | [string](#string) |  |  |
| nftContractMinter | [string](#string) |  |  |
| createSequence | [uint64](#uint64) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/event.proto



<a name="ixo.entity.v1beta1.EntityAccountAuthzCreatedEvent"></a>

### EntityAccountAuthzCreatedEvent
EntityAccountCreatedEvent is an event triggered on a entity account authz
creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| account_name | [string](#string) |  |  |
| granter | [string](#string) |  |  |
| grantee | [string](#string) |  |  |
| grant | [cosmos.authz.v1beta1.Grant](#cosmos.authz.v1beta1.Grant) |  |  |






<a name="ixo.entity.v1beta1.EntityAccountAuthzRevokedEvent"></a>

### EntityAccountAuthzRevokedEvent
EntityAccountAuthzRevokedEvent is an event triggered on a entity account
authz revocation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| account_name | [string](#string) |  |  |
| granter | [string](#string) |  |  |
| grantee | [string](#string) |  |  |
| msg_type_url | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityAccountCreatedEvent"></a>

### EntityAccountCreatedEvent
EntityAccountCreatedEvent is an event triggered on a entity account creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| account_name | [string](#string) |  |  |
| account_address | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityCreatedEvent"></a>

### EntityCreatedEvent
EntityCreatedEvent is an event triggered on a Entity creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |
| signer | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityTransferredEvent"></a>

### EntityTransferredEvent
EntityTransferredEvent is an event triggered on a entity transfer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityUpdatedEvent"></a>

### EntityUpdatedEvent
EntityUpdatedEvent is an event triggered on a entity document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |
| signer | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.EntityVerifiedUpdatedEvent"></a>

### EntityVerifiedUpdatedEvent
EntityVerifiedUpdatedEvent is an event triggered on a entity verified
document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| signer | [string](#string) |  |  |
| entity_verified | [bool](#bool) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/genesis.proto



<a name="ixo.entity.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the project module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entities | [Entity](#ixo.entity.v1beta1.Entity) | repeated |  |
| params | [Params](#ixo.entity.v1beta1.Params) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/proposal.proto



<a name="ixo.entity.v1beta1.InitializeNftContract"></a>

### InitializeNftContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| NftContractCodeId | [uint64](#uint64) |  |  |
| NftMinterAddress | [string](#string) |  |  |





 

 

 

 



<a name="ixo/entity/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/query.proto



<a name="ixo.entity.v1beta1.QueryEntityIidDocumentRequest"></a>

### QueryEntityIidDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityIidDocumentResponse"></a>

### QueryEntityIidDocumentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [ixo.iid.v1beta1.IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityListRequest"></a>

### QueryEntityListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | string type = 2; string status = 3; |






<a name="ixo.entity.v1beta1.QueryEntityListResponse"></a>

### QueryEntityListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entities | [Entity](#ixo.entity.v1beta1.Entity) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityMetadataRequest"></a>

### QueryEntityMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityMetadataResponse"></a>

### QueryEntityMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityRequest"></a>

### QueryEntityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityResponse"></a>

### QueryEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity | [Entity](#ixo.entity.v1beta1.Entity) |  |  |
| iidDocument | [ixo.iid.v1beta1.IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityVerifiedRequest"></a>

### QueryEntityVerifiedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.QueryEntityVerifiedResponse"></a>

### QueryEntityVerifiedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_verified | [bool](#bool) |  |  |






<a name="ixo.entity.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ixo.entity.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.entity.v1beta1.Params) |  | params holds all the parameters of this module. |





 

 

 


<a name="ixo.entity.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.entity.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.entity.v1beta1.QueryParamsResponse) |  |
| Entity | [QueryEntityRequest](#ixo.entity.v1beta1.QueryEntityRequest) | [QueryEntityResponse](#ixo.entity.v1beta1.QueryEntityResponse) |  |
| EntityMetaData | [QueryEntityMetadataRequest](#ixo.entity.v1beta1.QueryEntityMetadataRequest) | [QueryEntityMetadataResponse](#ixo.entity.v1beta1.QueryEntityMetadataResponse) |  |
| EntityIidDocument | [QueryEntityIidDocumentRequest](#ixo.entity.v1beta1.QueryEntityIidDocumentRequest) | [QueryEntityIidDocumentResponse](#ixo.entity.v1beta1.QueryEntityIidDocumentResponse) |  |
| EntityVerified | [QueryEntityVerifiedRequest](#ixo.entity.v1beta1.QueryEntityVerifiedRequest) | [QueryEntityVerifiedResponse](#ixo.entity.v1beta1.QueryEntityVerifiedResponse) |  |
| EntityList | [QueryEntityListRequest](#ixo.entity.v1beta1.QueryEntityListRequest) | [QueryEntityListResponse](#ixo.entity.v1beta1.QueryEntityListResponse) |  |

 



<a name="ixo/iid/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/tx.proto



<a name="ixo.iid.v1beta1.MsgAddAccordedRight"></a>

### MsgAddAccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) |  | the Accorded right to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddAccordedRightResponse"></a>

### MsgAddAccordedRightResponse







<a name="ixo.iid.v1beta1.MsgAddController"></a>

### MsgAddController



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did of the document |
| controller_did | [string](#string) |  | the did to add as a controller of the did document |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddControllerResponse"></a>

### MsgAddControllerResponse







<a name="ixo.iid.v1beta1.MsgAddIidContext"></a>

### MsgAddIidContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| context | [Context](#ixo.iid.v1beta1.Context) |  | the context to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddIidContextResponse"></a>

### MsgAddIidContextResponse







<a name="ixo.iid.v1beta1.MsgAddLinkedClaim"></a>

### MsgAddLinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) |  | the claim to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddLinkedClaimResponse"></a>

### MsgAddLinkedClaimResponse







<a name="ixo.iid.v1beta1.MsgAddLinkedEntity"></a>

### MsgAddLinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the iid |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) |  | the entity to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddLinkedEntityResponse"></a>

### MsgAddLinkedEntityResponse







<a name="ixo.iid.v1beta1.MsgAddLinkedResource"></a>

### MsgAddLinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) |  | the verification to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddLinkedResourceResponse"></a>

### MsgAddLinkedResourceResponse







<a name="ixo.iid.v1beta1.MsgAddService"></a>

### MsgAddService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| service_data | [Service](#ixo.iid.v1beta1.Service) |  | the service data to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddServiceResponse"></a>

### MsgAddServiceResponse







<a name="ixo.iid.v1beta1.MsgAddVerification"></a>

### MsgAddVerification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| verification | [Verification](#ixo.iid.v1beta1.Verification) |  | the verification to add |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgAddVerificationResponse"></a>

### MsgAddVerificationResponse







<a name="ixo.iid.v1beta1.MsgCreateIidDocument"></a>

### MsgCreateIidDocument
MsgCreateDidDocument defines a SDK message for creating a new did.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| controllers | [string](#string) | repeated | the list of controller DIDs |
| context | [Context](#ixo.iid.v1beta1.Context) | repeated |  |
| verifications | [Verification](#ixo.iid.v1beta1.Verification) | repeated | the list of verification methods and relationships |
| services | [Service](#ixo.iid.v1beta1.Service) | repeated |  |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated |  |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated |  |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| signer | [string](#string) |  | address of the account signing the message |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated |  |






<a name="ixo.iid.v1beta1.MsgCreateIidDocumentResponse"></a>

### MsgCreateIidDocumentResponse







<a name="ixo.iid.v1beta1.MsgDeactivateIID"></a>

### MsgDeactivateIID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| state | [bool](#bool) |  |  |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeactivateIIDResponse"></a>

### MsgDeactivateIIDResponse







<a name="ixo.iid.v1beta1.MsgDeleteAccordedRight"></a>

### MsgDeleteAccordedRight



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| right_id | [string](#string) |  | the Accorded right id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteAccordedRightResponse"></a>

### MsgDeleteAccordedRightResponse







<a name="ixo.iid.v1beta1.MsgDeleteController"></a>

### MsgDeleteController



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did of the document |
| controller_did | [string](#string) |  | the did to remove from the list of controllers of the did document |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteControllerResponse"></a>

### MsgDeleteControllerResponse







<a name="ixo.iid.v1beta1.MsgDeleteIidContext"></a>

### MsgDeleteIidContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| contextKey | [string](#string) |  | the context key |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteIidContextResponse"></a>

### MsgDeleteIidContextResponse







<a name="ixo.iid.v1beta1.MsgDeleteLinkedClaim"></a>

### MsgDeleteLinkedClaim



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| claim_id | [string](#string) |  | the claim id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteLinkedClaimResponse"></a>

### MsgDeleteLinkedClaimResponse







<a name="ixo.iid.v1beta1.MsgDeleteLinkedEntity"></a>

### MsgDeleteLinkedEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the iid |
| entity_id | [string](#string) |  | the entity id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse"></a>

### MsgDeleteLinkedEntityResponse







<a name="ixo.iid.v1beta1.MsgDeleteLinkedResource"></a>

### MsgDeleteLinkedResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| resource_id | [string](#string) |  | the service id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse"></a>

### MsgDeleteLinkedResourceResponse







<a name="ixo.iid.v1beta1.MsgDeleteService"></a>

### MsgDeleteService



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| service_id | [string](#string) |  | the service id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgDeleteServiceResponse"></a>

### MsgDeleteServiceResponse







<a name="ixo.iid.v1beta1.MsgRevokeVerification"></a>

### MsgRevokeVerification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| method_id | [string](#string) |  | the verification method id |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgRevokeVerificationResponse"></a>

### MsgRevokeVerificationResponse







<a name="ixo.iid.v1beta1.MsgSetVerificationRelationships"></a>

### MsgSetVerificationRelationships



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| method_id | [string](#string) |  | the verification method id |
| relationships | [string](#string) | repeated | the list of relationships to set |
| signer | [string](#string) |  | address of the account signing the message |






<a name="ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse"></a>

### MsgSetVerificationRelationshipsResponse







<a name="ixo.iid.v1beta1.MsgUpdateIidDocument"></a>

### MsgUpdateIidDocument
Updates the entity with all the fields, so if field empty will be updated
with default go type, aka never null


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | the did |
| controllers | [string](#string) | repeated | the list of controller DIDs |
| context | [Context](#ixo.iid.v1beta1.Context) | repeated |  |
| verifications | [Verification](#ixo.iid.v1beta1.Verification) | repeated | the list of verification methods and relationships |
| services | [Service](#ixo.iid.v1beta1.Service) | repeated |  |
| accordedRight | [AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated |  |
| linkedResource | [LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated |  |
| linkedEntity | [LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated |  |
| alsoKnownAs | [string](#string) |  |  |
| signer | [string](#string) |  | address of the account signing the message |
| linkedClaim | [LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated |  |






<a name="ixo.iid.v1beta1.MsgUpdateIidDocumentResponse"></a>

### MsgUpdateIidDocumentResponse







<a name="ixo.iid.v1beta1.Verification"></a>

### Verification
Verification is a message that allows to assign a verification method
to one or more verification relationships


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relationships | [string](#string) | repeated | verificationRelationships defines which relationships are allowed to use the verification method

relationships that the method is allowed into. |
| method | [VerificationMethod](#ixo.iid.v1beta1.VerificationMethod) |  | public key associated with the did document. |
| context | [string](#string) | repeated | additional contexts (json ld schemas) |





 

 

 


<a name="ixo.iid.v1beta1.Msg"></a>

### Msg
Msg defines the identity Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateIidDocument | [MsgCreateIidDocument](#ixo.iid.v1beta1.MsgCreateIidDocument) | [MsgCreateIidDocumentResponse](#ixo.iid.v1beta1.MsgCreateIidDocumentResponse) | CreateDidDocument defines a method for creating a new identity. |
| UpdateIidDocument | [MsgUpdateIidDocument](#ixo.iid.v1beta1.MsgUpdateIidDocument) | [MsgUpdateIidDocumentResponse](#ixo.iid.v1beta1.MsgUpdateIidDocumentResponse) | UpdateDidDocument defines a method for updating an identity. |
| AddVerification | [MsgAddVerification](#ixo.iid.v1beta1.MsgAddVerification) | [MsgAddVerificationResponse](#ixo.iid.v1beta1.MsgAddVerificationResponse) | AddVerificationMethod adds a new verification method |
| RevokeVerification | [MsgRevokeVerification](#ixo.iid.v1beta1.MsgRevokeVerification) | [MsgRevokeVerificationResponse](#ixo.iid.v1beta1.MsgRevokeVerificationResponse) | RevokeVerification remove the verification method and all associated verification Relations |
| SetVerificationRelationships | [MsgSetVerificationRelationships](#ixo.iid.v1beta1.MsgSetVerificationRelationships) | [MsgSetVerificationRelationshipsResponse](#ixo.iid.v1beta1.MsgSetVerificationRelationshipsResponse) | SetVerificationRelationships overwrite current verification relationships |
| AddService | [MsgAddService](#ixo.iid.v1beta1.MsgAddService) | [MsgAddServiceResponse](#ixo.iid.v1beta1.MsgAddServiceResponse) | AddService add a new service |
| DeleteService | [MsgDeleteService](#ixo.iid.v1beta1.MsgDeleteService) | [MsgDeleteServiceResponse](#ixo.iid.v1beta1.MsgDeleteServiceResponse) | DeleteService delete an existing service |
| AddController | [MsgAddController](#ixo.iid.v1beta1.MsgAddController) | [MsgAddControllerResponse](#ixo.iid.v1beta1.MsgAddControllerResponse) | AddService add a new service |
| DeleteController | [MsgDeleteController](#ixo.iid.v1beta1.MsgDeleteController) | [MsgDeleteControllerResponse](#ixo.iid.v1beta1.MsgDeleteControllerResponse) | DeleteService delete an existing service |
| AddLinkedResource | [MsgAddLinkedResource](#ixo.iid.v1beta1.MsgAddLinkedResource) | [MsgAddLinkedResourceResponse](#ixo.iid.v1beta1.MsgAddLinkedResourceResponse) | Add / Delete Linked Resource |
| DeleteLinkedResource | [MsgDeleteLinkedResource](#ixo.iid.v1beta1.MsgDeleteLinkedResource) | [MsgDeleteLinkedResourceResponse](#ixo.iid.v1beta1.MsgDeleteLinkedResourceResponse) |  |
| AddLinkedClaim | [MsgAddLinkedClaim](#ixo.iid.v1beta1.MsgAddLinkedClaim) | [MsgAddLinkedClaimResponse](#ixo.iid.v1beta1.MsgAddLinkedClaimResponse) | Add / Delete Linked Claims |
| DeleteLinkedClaim | [MsgDeleteLinkedClaim](#ixo.iid.v1beta1.MsgDeleteLinkedClaim) | [MsgDeleteLinkedClaimResponse](#ixo.iid.v1beta1.MsgDeleteLinkedClaimResponse) |  |
| AddLinkedEntity | [MsgAddLinkedEntity](#ixo.iid.v1beta1.MsgAddLinkedEntity) | [MsgAddLinkedEntityResponse](#ixo.iid.v1beta1.MsgAddLinkedEntityResponse) | Add / Delete Linked Entity |
| DeleteLinkedEntity | [MsgDeleteLinkedEntity](#ixo.iid.v1beta1.MsgDeleteLinkedEntity) | [MsgDeleteLinkedEntityResponse](#ixo.iid.v1beta1.MsgDeleteLinkedEntityResponse) |  |
| AddAccordedRight | [MsgAddAccordedRight](#ixo.iid.v1beta1.MsgAddAccordedRight) | [MsgAddAccordedRightResponse](#ixo.iid.v1beta1.MsgAddAccordedRightResponse) | Add / Delete Accorded Right |
| DeleteAccordedRight | [MsgDeleteAccordedRight](#ixo.iid.v1beta1.MsgDeleteAccordedRight) | [MsgDeleteAccordedRightResponse](#ixo.iid.v1beta1.MsgDeleteAccordedRightResponse) |  |
| AddIidContext | [MsgAddIidContext](#ixo.iid.v1beta1.MsgAddIidContext) | [MsgAddIidContextResponse](#ixo.iid.v1beta1.MsgAddIidContextResponse) | Add / Delete Context |
| DeactivateIID | [MsgDeactivateIID](#ixo.iid.v1beta1.MsgDeactivateIID) | [MsgDeactivateIIDResponse](#ixo.iid.v1beta1.MsgDeactivateIIDResponse) |  |
| DeleteIidContext | [MsgDeleteIidContext](#ixo.iid.v1beta1.MsgDeleteIidContext) | [MsgDeleteIidContextResponse](#ixo.iid.v1beta1.MsgDeleteIidContextResponse) |  |

 



<a name="ixo/entity/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/entity/v1beta1/tx.proto



<a name="ixo.entity.v1beta1.MsgCreateEntity"></a>

### MsgCreateEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_type | [string](#string) |  | An Entity Type as defined by the implementer |
| entity_status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| controller | [string](#string) | repeated | the list of controller DIDs |
| context | [ixo.iid.v1beta1.Context](#ixo.iid.v1beta1.Context) | repeated | JSON-LD contexts |
| verification | [ixo.iid.v1beta1.Verification](#ixo.iid.v1beta1.Verification) | repeated | Verification Methods and Verification Relationships |
| service | [ixo.iid.v1beta1.Service](#ixo.iid.v1beta1.Service) | repeated | Service endpoints |
| accorded_right | [ixo.iid.v1beta1.AccordedRight](#ixo.iid.v1beta1.AccordedRight) | repeated | Legal or Electronic Rights and associated Object Capabilities |
| linked_resource | [ixo.iid.v1beta1.LinkedResource](#ixo.iid.v1beta1.LinkedResource) | repeated | Digital resources associated with the Subject |
| linked_entity | [ixo.iid.v1beta1.LinkedEntity](#ixo.iid.v1beta1.LinkedEntity) | repeated | DID of a linked Entity and its relationship with the Subject |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| relayer_node | [string](#string) |  | Did of the operator through which the Entity was created |
| credentials | [string](#string) | repeated | Content ID or Hash of public Verifiable Credentials associated with the subject |
| owner_did | [string](#string) |  | Owner of the Entity NFT | The ownersdid used to sign this transaction. |
| owner_address | [string](#string) |  | The ownersdid address used to sign this transaction. |
| data | [bytes](#bytes) |  | Extension data |
| alsoKnownAs | [string](#string) |  |  |
| linked_claim | [ixo.iid.v1beta1.LinkedClaim](#ixo.iid.v1beta1.LinkedClaim) | repeated | Digital claims associated with the Subject |






<a name="ixo.entity.v1beta1.MsgCreateEntityAccount"></a>

### MsgCreateEntityAccount
create a module account for an entity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | entity id (did) to create account for |
| name | [string](#string) |  | name of account |
| owner_address | [string](#string) |  | The owner_address used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgCreateEntityAccountResponse"></a>

### MsgCreateEntityAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  | account address that was created for specific entity and account name |






<a name="ixo.entity.v1beta1.MsgCreateEntityResponse"></a>

### MsgCreateEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entity_id | [string](#string) |  |  |
| entity_type | [string](#string) |  |  |
| entity_status | [int32](#int32) |  |  |






<a name="ixo.entity.v1beta1.MsgGrantEntityAccountAuthz"></a>

### MsgGrantEntityAccountAuthz
Create a authz grant from entity account (as grantor) to recipient in msg as
grantee for the specific authorization


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | entity id (did) which has the account |
| name | [string](#string) |  | name of account |
| grantee_address | [string](#string) |  | the grantee address that will be able to execute the authz authorization |
| grant | [cosmos.authz.v1beta1.Grant](#cosmos.authz.v1beta1.Grant) |  | grant to be Authorized in authz grant |
| owner_address | [string](#string) |  | the owner_address used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgGrantEntityAccountAuthzResponse"></a>

### MsgGrantEntityAccountAuthzResponse







<a name="ixo.entity.v1beta1.MsgRevokeEntityAccountAuthz"></a>

### MsgRevokeEntityAccountAuthz
Revoke an existing authz grant from entity account (as grantor) to recipient


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | entity id (did) which has the account |
| name | [string](#string) |  | name of account |
| grantee_address | [string](#string) |  | the grantee address for which the authz grant will be revoked |
| msg_type_url | [string](#string) |  | the msg type url of the grant to be revoked for the specific grantee |
| owner_address | [string](#string) |  | the owner_address used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgRevokeEntityAccountAuthzResponse"></a>

### MsgRevokeEntityAccountAuthzResponse







<a name="ixo.entity.v1beta1.MsgTransferEntity"></a>

### MsgTransferEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| owner_did | [string](#string) |  | The owner_did used to sign this transaction. |
| owner_address | [string](#string) |  | The owner_address used to sign this transaction. |
| recipient_did | [string](#string) |  |  |






<a name="ixo.entity.v1beta1.MsgTransferEntityResponse"></a>

### MsgTransferEntityResponse







<a name="ixo.entity.v1beta1.MsgUpdateEntity"></a>

### MsgUpdateEntity
Updates the entity with all the fields, so if field empty will be updated
with default go type, aka never null


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id of entity to be updated |
| entity_status | [int32](#int32) |  | Status of the Entity as defined by the implementer and interpreted by Client applications |
| start_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start Date of the Entity as defined by the implementer and interpreted by Client applications |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End Date of the Entity as defined by the implementer and interpreted by Client applications |
| credentials | [string](#string) | repeated | Content ID or Hash of public Verifiable Credentials associated with the subject |
| controller_did | [string](#string) |  | The controllerDid used to sign this transaction. |
| controller_address | [string](#string) |  | The controllerAddress used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgUpdateEntityResponse"></a>

### MsgUpdateEntityResponse







<a name="ixo.entity.v1beta1.MsgUpdateEntityVerified"></a>

### MsgUpdateEntityVerified
Only relayer nodes can update entity field &#39;entityVerified&#39;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Id of entity to be updated |
| entity_verified | [bool](#bool) |  | Whether entity is verified or not based on credentials |
| relayer_node_did | [string](#string) |  | The relayer node&#39;s did used to sign this transaction. |
| relayer_node_address | [string](#string) |  | The relayer node&#39;s address used to sign this transaction. |






<a name="ixo.entity.v1beta1.MsgUpdateEntityVerifiedResponse"></a>

### MsgUpdateEntityVerifiedResponse






 

 

 


<a name="ixo.entity.v1beta1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateEntity | [MsgCreateEntity](#ixo.entity.v1beta1.MsgCreateEntity) | [MsgCreateEntityResponse](#ixo.entity.v1beta1.MsgCreateEntityResponse) | CreateEntity defines a method for creating a entity. |
| UpdateEntity | [MsgUpdateEntity](#ixo.entity.v1beta1.MsgUpdateEntity) | [MsgUpdateEntityResponse](#ixo.entity.v1beta1.MsgUpdateEntityResponse) | UpdateEntity defines a method for updating a entity |
| UpdateEntityVerified | [MsgUpdateEntityVerified](#ixo.entity.v1beta1.MsgUpdateEntityVerified) | [MsgUpdateEntityVerifiedResponse](#ixo.entity.v1beta1.MsgUpdateEntityVerifiedResponse) | UpdateEntityVerified defines a method for updating if an entity is verified |
| TransferEntity | [MsgTransferEntity](#ixo.entity.v1beta1.MsgTransferEntity) | [MsgTransferEntityResponse](#ixo.entity.v1beta1.MsgTransferEntityResponse) | Transfers an entity and its nft to the recipient |
| CreateEntityAccount | [MsgCreateEntityAccount](#ixo.entity.v1beta1.MsgCreateEntityAccount) | [MsgCreateEntityAccountResponse](#ixo.entity.v1beta1.MsgCreateEntityAccountResponse) | Create a module account for an entity, |
| GrantEntityAccountAuthz | [MsgGrantEntityAccountAuthz](#ixo.entity.v1beta1.MsgGrantEntityAccountAuthz) | [MsgGrantEntityAccountAuthzResponse](#ixo.entity.v1beta1.MsgGrantEntityAccountAuthzResponse) | Create an authz grant from entity account |
| RevokeEntityAccountAuthz | [MsgRevokeEntityAccountAuthz](#ixo.entity.v1beta1.MsgRevokeEntityAccountAuthz) | [MsgRevokeEntityAccountAuthzResponse](#ixo.entity.v1beta1.MsgRevokeEntityAccountAuthzResponse) | Revoke an authz grant from entity account |

 



<a name="ixo/epochs/v1beta1/epoch.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/epochs/v1beta1/epoch.proto



<a name="ixo.epochs.v1beta1.EpochInfo"></a>

### EpochInfo
EpochInfo is a struct that describes the data going into
a timer defined by the x/epochs module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  | identifier is a unique reference to this particular timer. |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time is the time at which the timer first ever ticks. If start_time is in the future, the epoch will not begin until the start time. |
| duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | duration is the time in between epoch ticks. In order for intended behavior to be met, duration should be greater than the chains expected block time. Duration must be non-zero. |
| current_epoch | [int64](#int64) |  | current_epoch is the current epoch number, or in other words, how many times has the timer &#39;ticked&#39;. The first tick (current_epoch=1) is defined as the first block whose blocktime is greater than the EpochInfo start_time. |
| current_epoch_start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | current_epoch_start_time describes the start time of the current timer interval. The interval is (current_epoch_start_time, current_epoch_start_time &#43; duration] When the timer ticks, this is set to current_epoch_start_time = last_epoch_start_time &#43; duration only one timer tick for a given identifier can occur per block.

NOTE! The current_epoch_start_time may diverge significantly from the wall-clock time the epoch began at. Wall-clock time of epoch start may be &gt;&gt; current_epoch_start_time. Suppose current_epoch_start_time = 10, duration = 5. Suppose the chain goes offline at t=14, and comes back online at t=30, and produces blocks at every successive time. (t=31, 32, etc.) * The t=30 block will start the epoch for (10, 15] * The t=31 block will start the epoch for (15, 20] * The t=32 block will start the epoch for (20, 25] * The t=33 block will start the epoch for (25, 30] * The t=34 block will start the epoch for (30, 35] * The **t=36** block will start the epoch for (35, 40] |
| epoch_counting_started | [bool](#bool) |  | epoch_counting_started is a boolean, that indicates whether this epoch timer has began yet. |
| current_epoch_start_height | [int64](#int64) |  | current_epoch_start_height is the block height at which the current epoch started. (The block height at which the timer last ticked) |





 

 

 

 



<a name="ixo/epochs/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/epochs/v1beta1/event.proto



<a name="ixo.epochs.v1beta1.EpochEndEvent"></a>

### EpochEndEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_number | [int64](#int64) |  | Epoch number, starting from 1. |






<a name="ixo.epochs.v1beta1.EpochStartEvent"></a>

### EpochStartEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_number | [int64](#int64) |  | Epoch number, starting from 1. |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The start timestamp of the epoch. |





 

 

 

 



<a name="ixo/epochs/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/epochs/v1beta1/genesis.proto



<a name="ixo.epochs.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the epochs module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochs | [EpochInfo](#ixo.epochs.v1beta1.EpochInfo) | repeated |  |





 

 

 

 



<a name="ixo/epochs/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/epochs/v1beta1/query.proto



<a name="ixo.epochs.v1beta1.QueryCurrentEpochRequest"></a>

### QueryCurrentEpochRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  |  |






<a name="ixo.epochs.v1beta1.QueryCurrentEpochResponse"></a>

### QueryCurrentEpochResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_epoch | [int64](#int64) |  |  |






<a name="ixo.epochs.v1beta1.QueryEpochsInfoRequest"></a>

### QueryEpochsInfoRequest







<a name="ixo.epochs.v1beta1.QueryEpochsInfoResponse"></a>

### QueryEpochsInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochs | [EpochInfo](#ixo.epochs.v1beta1.EpochInfo) | repeated |  |





 

 

 


<a name="ixo.epochs.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| EpochInfos | [QueryEpochsInfoRequest](#ixo.epochs.v1beta1.QueryEpochsInfoRequest) | [QueryEpochsInfoResponse](#ixo.epochs.v1beta1.QueryEpochsInfoResponse) | EpochInfos provide running epochInfos |
| CurrentEpoch | [QueryCurrentEpochRequest](#ixo.epochs.v1beta1.QueryCurrentEpochRequest) | [QueryCurrentEpochResponse](#ixo.epochs.v1beta1.QueryCurrentEpochResponse) | CurrentEpoch provide current epoch of specified identifier |

 



<a name="ixo/iid/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/event.proto



<a name="ixo.iid.v1beta1.IidDocumentCreatedEvent"></a>

### IidDocumentCreatedEvent
IidDocumentCreatedEvent is triggered when a new IidDocument is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.iid.v1beta1.IidDocumentUpdatedEvent"></a>

### IidDocumentUpdatedEvent
DidDocumentUpdatedEvent is an event triggered on a DID document update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |





 

 

 

 



<a name="ixo/iid/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/genesis.proto



<a name="ixo.iid.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the did module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iid_docs | [IidDocument](#ixo.iid.v1beta1.IidDocument) | repeated |  |





 

 

 

 



<a name="ixo/iid/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/iid/v1beta1/query.proto



<a name="ixo.iid.v1beta1.QueryIidDocumentRequest"></a>

### QueryIidDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | did id of iid document querying |






<a name="ixo.iid.v1beta1.QueryIidDocumentResponse"></a>

### QueryIidDocumentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocument | [IidDocument](#ixo.iid.v1beta1.IidDocument) |  |  |






<a name="ixo.iid.v1beta1.QueryIidDocumentsRequest"></a>

### QueryIidDocumentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="ixo.iid.v1beta1.QueryIidDocumentsResponse"></a>

### QueryIidDocumentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| iidDocuments | [IidDocument](#ixo.iid.v1beta1.IidDocument) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 

 

 


<a name="ixo.iid.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IidDocuments | [QueryIidDocumentsRequest](#ixo.iid.v1beta1.QueryIidDocumentsRequest) | [QueryIidDocumentsResponse](#ixo.iid.v1beta1.QueryIidDocumentsResponse) | IidDocuments queries all iid documents that match the given status. |
| IidDocument | [QueryIidDocumentRequest](#ixo.iid.v1beta1.QueryIidDocumentRequest) | [QueryIidDocumentResponse](#ixo.iid.v1beta1.QueryIidDocumentResponse) | IidDocument queries a iid documents with an id. |

 



<a name="ixo/liquidstake/v1beta1/liquidstake.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/liquidstake/v1beta1/liquidstake.proto



<a name="ixo.liquidstake.v1beta1.LiquidValidator"></a>

### LiquidValidator
LiquidValidator defines a Validator that can be the target of LiquidStaking
and LiquidUnstaking, Active, Weight, etc. fields are derived as functions to
deal with by maintaining consistency with the state of the staking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operator_address | [string](#string) |  | operator_address defines the address of the validator&#39;s operator; bech encoded in JSON. |






<a name="ixo.liquidstake.v1beta1.LiquidValidatorState"></a>

### LiquidValidatorState
LiquidValidatorState is type LiquidValidator with state added to return to
query results.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operator_address | [string](#string) |  | operator_address defines the address of the validator&#39;s operator; bech encoded in JSON. |
| weight | [string](#string) |  | weight specifies the weight for liquid staking, unstaking amount |
| status | [ValidatorStatus](#ixo.liquidstake.v1beta1.ValidatorStatus) |  | status is the liquid validator status |
| del_shares | [string](#string) |  | del_shares define the delegation shares of the validator |
| liquid_tokens | [string](#string) |  | liquid_tokens define the token amount worth of delegation shares of the validator (slashing applied amount) |






<a name="ixo.liquidstake.v1beta1.NetAmountState"></a>

### NetAmountState
NetAmountState is type for net amount raw data and mint rate, This is a value
that depends on the several module state every time, so it is used only for
calculation and query and is not stored in kv.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stake_rate | [string](#string) |  | stake_rate is the rate at which the liquid staking module mints stkIXO |
| unstake_rate | [string](#string) |  | unstake_rate is the rate at which the liquid staking module burns stkIXO |
| stkixo_total_supply | [string](#string) |  | btoken_total_supply returns the total supply of uzero (stkIXO denom) |
| net_amount | [string](#string) |  | net_amount is proxy account&#39;s total liquid tokens &#43; total unbonding balance |
| total_del_shares | [string](#string) |  | total_del_shares define the delegation shares of all liquid validators |
| total_liquid_tokens | [string](#string) |  | total_liquid_tokens define the token amount worth of delegation shares of all liquid validator (slashing applied amount) |
| total_remaining_rewards | [string](#string) |  | total_remaining_rewards define the sum of remaining rewards of proxy account by all liquid validators |
| total_unbonding_balance | [string](#string) |  | total_unbonding_balance define the unbonding balance of proxy account by all liquid validator (slashing applied amount) |
| proxy_acc_balance | [string](#string) |  | proxy_acc_balance define the balance of proxy account for the native token |






<a name="ixo.liquidstake.v1beta1.Params"></a>

### Params
Params defines the set of params for the liquidstake module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| liquid_bond_denom | [string](#string) |  | LiquidBondDenom specifies the denomination of the token receiving after liquid stake, The value is calculated through NetAmount. |
| whitelisted_validators | [WhitelistedValidator](#ixo.liquidstake.v1beta1.WhitelistedValidator) | repeated | WhitelistedValidators specifies the validators elected to become Active Liquid Validators. |
| unstake_fee_rate | [string](#string) |  | UnstakeFeeRate specifies the fee rate when liquid unstake is requested, unbonded by subtracting it from unbondingAmount |
| min_liquid_stake_amount | [string](#string) |  | MinLiquidStakingAmount specifies the minimum number of coins to be staked to the active liquid validators on liquid staking to minimize decimal loss and consider gas efficiency. |
| fee_account_address | [string](#string) |  | FeeAccountAddress defines the bech32-encoded address of an account responsible for accumulating protocol fees. |
| autocompound_fee_rate | [string](#string) |  | AutocompoundFeeRate specifies the fee rate for auto redelegating the stake rewards. The fee is taken in favour of the fee account (see FeeAccountAddress). |
| whitelist_admin_address | [string](#string) |  | WhitelistAdminAddress the bech32-encoded address of an admin authority that is allowed to update whitelisted validators or pause liquidstaking module entirely. It is also the only address that can update the weighted_rewards_receivers. The key is controlled by the ZERO dao. Pausing of the module can be required during important migrations or failures. |
| module_paused | [bool](#bool) |  | ModulePaused is a safety toggle that allows to stop main module functions such as stake/unstake/stake-to-lp and the BeginBlocker logic. |
| weighted_rewards_receivers | [WeightedAddress](#ixo.liquidstake.v1beta1.WeightedAddress) | repeated | weighted_rewards_receivers is the addresses to receive the staking rewards on autocompounding with weights assigned to each address. The total of weights in the list in not allowed to be greater than 1.

Eg. if the list has 1 address with weight 0.2, then on autocompounding the staking rewards will be split between 0.2 for the weighted receiver and 0.8 gets auto-compounded to the proxy account. |






<a name="ixo.liquidstake.v1beta1.WeightedAddress"></a>

### WeightedAddress
WeightedAddress represents an address with a weight assigned to it.
The weight is used to determine the proportion of autocompounding
rewards to be paid to the address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| weight | [string](#string) |  |  |






<a name="ixo.liquidstake.v1beta1.WhitelistedValidator"></a>

### WhitelistedValidator
WhitelistedValidator consists of the validator operator address and the
target weight, which is a value for calculating the real weight to be derived
according to the active status.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validator_address | [string](#string) |  | validator_address defines the bech32-encoded address of the whitelisted validator |
| target_weight | [string](#string) |  | target_weight specifies the target weight for liquid staking, unstaking amount, which is a value for calculating the real weight to be derived according to the active status |





 


<a name="ixo.liquidstake.v1beta1.ValidatorStatus"></a>

### ValidatorStatus
ValidatorStatus enumerates the status of a liquid validator.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VALIDATOR_STATUS_UNSPECIFIED | 0 | VALIDATOR_STATUS_UNSPECIFIED defines the unspecified invalid status. |
| VALIDATOR_STATUS_ACTIVE | 1 | VALIDATOR_STATUS_ACTIVE defines the active, valid status |
| VALIDATOR_STATUS_INACTIVE | 2 | VALIDATOR_STATUS_INACTIVE defines the inactive, invalid status |


 

 

 



<a name="ixo/liquidstake/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/liquidstake/v1beta1/event.proto



<a name="ixo.liquidstake.v1beta1.AddLiquidValidatorEvent"></a>

### AddLiquidValidatorEvent
LiquidRedelegateEvent is triggered when a liquid validator is added.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validator | [string](#string) |  |  |






<a name="ixo.liquidstake.v1beta1.AutocompoundStakingRewardsEvent"></a>

### AutocompoundStakingRewardsEvent
AutocompoundEvent is triggered after a epoch is triggered for autocompound.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator | [string](#string) |  |  |
| total_amount | [string](#string) |  |  |
| fee_amount | [string](#string) |  |  |
| redelegate_amount | [string](#string) |  |  |
| weighted_rewards_amount | [string](#string) |  |  |






<a name="ixo.liquidstake.v1beta1.LiquidStakeEvent"></a>

### LiquidStakeEvent
LiquidStakeEvent is triggered when a liquid stake is performed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator | [string](#string) |  |  |
| liquid_amount | [string](#string) |  |  |
| stk_ixo_minted_amount | [string](#string) |  |  |






<a name="ixo.liquidstake.v1beta1.LiquidStakeParamsUpdatedEvent"></a>

### LiquidStakeParamsUpdatedEvent
LiquidStakeParamsUpdatedEvent is triggered when a the Params is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.liquidstake.v1beta1.Params) |  |  |
| authority | [string](#string) |  |  |






<a name="ixo.liquidstake.v1beta1.LiquidUnstakeEvent"></a>

### LiquidUnstakeEvent
LiquidUnstakeEvent is triggered when a liquid unstake is performed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator | [string](#string) |  |  |
| unstake_amount | [string](#string) |  |  |
| unbonding_amount | [string](#string) |  |  |
| unbonded_amount | [string](#string) |  |  |
| completion_time | [string](#string) |  |  |






<a name="ixo.liquidstake.v1beta1.RebalancedLiquidStakeEvent"></a>

### RebalancedLiquidStakeEvent
RebalancedEvent is triggered after a rebalance is performed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator | [string](#string) |  |  |
| redelegation_count | [string](#string) |  |  |
| redelegation_fail_count | [string](#string) |  |  |





 

 

 

 



<a name="ixo/liquidstake/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/liquidstake/v1beta1/genesis.proto



<a name="ixo.liquidstake.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the liquidstake module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.liquidstake.v1beta1.Params) |  |  |
| liquid_validators | [LiquidValidator](#ixo.liquidstake.v1beta1.LiquidValidator) | repeated |  |





 

 

 

 



<a name="ixo/liquidstake/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/liquidstake/v1beta1/query.proto



<a name="ixo.liquidstake.v1beta1.QueryLiquidValidatorsRequest"></a>

### QueryLiquidValidatorsRequest







<a name="ixo.liquidstake.v1beta1.QueryLiquidValidatorsResponse"></a>

### QueryLiquidValidatorsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| liquid_validators | [LiquidValidatorState](#ixo.liquidstake.v1beta1.LiquidValidatorState) | repeated |  |






<a name="ixo.liquidstake.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ixo.liquidstake.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.liquidstake.v1beta1.Params) |  |  |






<a name="ixo.liquidstake.v1beta1.QueryStatesRequest"></a>

### QueryStatesRequest







<a name="ixo.liquidstake.v1beta1.QueryStatesResponse"></a>

### QueryStatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| net_amount_state | [NetAmountState](#ixo.liquidstake.v1beta1.NetAmountState) |  |  |





 

 

 


<a name="ixo.liquidstake.v1beta1.Query"></a>

### Query
Query defines the gRPC query service for the liquidstake module.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.liquidstake.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.liquidstake.v1beta1.QueryParamsResponse) | Params returns parameters of the liquidstake module. |
| LiquidValidators | [QueryLiquidValidatorsRequest](#ixo.liquidstake.v1beta1.QueryLiquidValidatorsRequest) | [QueryLiquidValidatorsResponse](#ixo.liquidstake.v1beta1.QueryLiquidValidatorsResponse) | LiquidValidators returns liquid validators with states of the liquidstake module. |
| States | [QueryStatesRequest](#ixo.liquidstake.v1beta1.QueryStatesRequest) | [QueryStatesResponse](#ixo.liquidstake.v1beta1.QueryStatesResponse) | States returns states of the liquidstake module. |

 



<a name="ixo/liquidstake/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/liquidstake/v1beta1/tx.proto



<a name="ixo.liquidstake.v1beta1.MsgLiquidStake"></a>

### MsgLiquidStake
MsgLiquidStake defines a SDK message for performing a liquid stake of coins
from a delegator to whitelisted validators.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator_address | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ixo.liquidstake.v1beta1.MsgLiquidStakeResponse"></a>

### MsgLiquidStakeResponse
MsgLiquidStakeResponse defines the MsgLiquidStake response type.






<a name="ixo.liquidstake.v1beta1.MsgLiquidUnstake"></a>

### MsgLiquidUnstake
MsgLiquidUnstake defines a SDK message for performing an undelegation of
liquid staking from a delegate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegator_address | [string](#string) |  |  |
| amount | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="ixo.liquidstake.v1beta1.MsgLiquidUnstakeResponse"></a>

### MsgLiquidUnstakeResponse
MsgLiquidUnstakeResponse defines the MsgLiquidUnstake response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| completion_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="ixo.liquidstake.v1beta1.MsgSetModulePaused"></a>

### MsgSetModulePaused



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authority | [string](#string) |  | Authority is the address that is allowed to update module&#39;s paused state, defined as admin address in params (WhitelistAdminAddress). |
| is_paused | [bool](#bool) |  | IsPaused represents the target state of the paused flag. |






<a name="ixo.liquidstake.v1beta1.MsgSetModulePausedResponse"></a>

### MsgSetModulePausedResponse
MsgSetModulePausedResponse defines the response structure for
executing a






<a name="ixo.liquidstake.v1beta1.MsgUpdateParams"></a>

### MsgUpdateParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authority | [string](#string) |  | authority is the address that controls the module (defaults to x/gov unless overwritten). |
| params | [Params](#ixo.liquidstake.v1beta1.Params) |  | params defines the parameters to update. NOTE: denom and whitelisted_validators and weighted_rewards_receivers are not updated. |






<a name="ixo.liquidstake.v1beta1.MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse defines the response structure for executing a






<a name="ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceivers"></a>

### MsgUpdateWeightedRewardsReceivers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authority | [string](#string) |  | Authority is the address that is allowed to update wieghted rewards receivers, defined as admin address in params (WhitelistAdminAddress). |
| weighted_rewards_receivers | [WeightedAddress](#ixo.liquidstake.v1beta1.WeightedAddress) | repeated | WhitelistedValidators specifies the validators elected to become Active Liquid Validators. |






<a name="ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceiversResponse"></a>

### MsgUpdateWeightedRewardsReceiversResponse
MsgUpdateWeightedRewardsReceiversResponse defines the response structure for
executing a






<a name="ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidators"></a>

### MsgUpdateWhitelistedValidators



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authority | [string](#string) |  | Authority is the address that is allowed to update whitelisted validators, defined as admin address in params (WhitelistAdminAddress). |
| whitelisted_validators | [WhitelistedValidator](#ixo.liquidstake.v1beta1.WhitelistedValidator) | repeated | WhitelistedValidators specifies the validators elected to become Active Liquid Validators. |






<a name="ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidatorsResponse"></a>

### MsgUpdateWhitelistedValidatorsResponse
MsgUpdateWhitelistedValidatorsResponse defines the response structure for
executing a





 

 

 


<a name="ixo.liquidstake.v1beta1.Msg"></a>

### Msg
Msg defines the liquid staking Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LiquidStake | [MsgLiquidStake](#ixo.liquidstake.v1beta1.MsgLiquidStake) | [MsgLiquidStakeResponse](#ixo.liquidstake.v1beta1.MsgLiquidStakeResponse) | LiquidStake defines a method for performing a delegation of coins from a delegator to whitelisted validators. |
| LiquidUnstake | [MsgLiquidUnstake](#ixo.liquidstake.v1beta1.MsgLiquidUnstake) | [MsgLiquidUnstakeResponse](#ixo.liquidstake.v1beta1.MsgLiquidUnstakeResponse) | LiquidUnstake defines a method for performing an undelegation of liquid staking from a delegate. |
| UpdateParams | [MsgUpdateParams](#ixo.liquidstake.v1beta1.MsgUpdateParams) | [MsgUpdateParamsResponse](#ixo.liquidstake.v1beta1.MsgUpdateParamsResponse) | UpdateParams defines a method to update the module params. |
| UpdateWhitelistedValidators | [MsgUpdateWhitelistedValidators](#ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidators) | [MsgUpdateWhitelistedValidatorsResponse](#ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidatorsResponse) | UpdateWhitelistedValidators defines a method to update the whitelisted validators list. |
| UpdateWeightedRewardsReceivers | [MsgUpdateWeightedRewardsReceivers](#ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceivers) | [MsgUpdateWeightedRewardsReceiversResponse](#ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceiversResponse) | UpdateWhitelistedValidators defines a method to update the whitelisted validators list. Only the whitelist admin address can update this list. |
| SetModulePaused | [MsgSetModulePaused](#ixo.liquidstake.v1beta1.MsgSetModulePaused) | [MsgSetModulePausedResponse](#ixo.liquidstake.v1beta1.MsgSetModulePausedResponse) | SetModulePaused defines a method to update the module&#39;s pause status, setting value of the safety flag in params. |

 



<a name="ixo/mint/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/mint/v1beta1/event.proto



<a name="ixo.mint.v1beta1.MintEpochProvisionsMintedEvent"></a>

### MintEpochProvisionsMintedEvent
MintEpochProvisionsMintedEvent is triggered after a epoch is triggered
minting module for inflation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_number | [string](#string) |  |  |
| epoch_provisions | [string](#string) |  |  |
| amount | [string](#string) |  |  |





 

 

 

 



<a name="ixo/mint/v1beta1/mint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/mint/v1beta1/mint.proto



<a name="ixo.mint.v1beta1.DistributionProportions"></a>

### DistributionProportions
DistributionProportions defines the distribution proportions of the minted
denom. In other words, defines which stakeholders will receive the minted
denoms and how much.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staking | [string](#string) |  | staking defines the proportion of the minted mint_denom that is to be allocated as staking rewards. |
| impact_rewards | [string](#string) |  | impact_rewards defines the proportion of the minted mint_denom that is to be allocated to impact rewards addresses. |
| community_pool | [string](#string) |  | community_pool defines the proportion of the minted mint_denom that is to be allocated to the community pool. |






<a name="ixo.mint.v1beta1.Minter"></a>

### Minter
Minter represents the minting state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_provisions | [string](#string) |  | epoch_provisions represent rewards for the current epoch. |






<a name="ixo.mint.v1beta1.Params"></a>

### Params
Params holds parameters for the x/mint module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mint_denom | [string](#string) |  | mint_denom is the denom of the coin to mint. |
| genesis_epoch_provisions | [string](#string) |  | genesis_epoch_provisions epoch provisions from the first epoch. |
| epoch_identifier | [string](#string) |  | epoch_identifier mint epoch identifier e.g. (day, week). |
| reduction_period_in_epochs | [int64](#int64) |  | reduction_period_in_epochs the number of epochs it takes to reduce the rewards. |
| reduction_factor | [string](#string) |  | reduction_factor is the reduction multiplier to execute at the end of each period set by reduction_period_in_epochs. |
| distribution_proportions | [DistributionProportions](#ixo.mint.v1beta1.DistributionProportions) |  | distribution_proportions defines the distribution proportions of the minted denom. In other words, defines which stakeholders will receive the minted denoms and how much. |
| weighted_impact_rewards_receivers | [WeightedAddress](#ixo.mint.v1beta1.WeightedAddress) | repeated | weighted_impact_rewards_receivers is the address to receive impact rewards with weights assigned to each address. The final amount that each address receives is: epoch_provisions * distribution_proportions.impact_rewards * Address&#39;s Weight. |
| minting_rewards_distribution_start_epoch | [int64](#int64) |  | minting_rewards_distribution_start_epoch start epoch to distribute minting rewards |






<a name="ixo.mint.v1beta1.WeightedAddress"></a>

### WeightedAddress
WeightedAddress represents an address with a weight assigned to it.
The weight is used to determine the proportion of the total minted
tokens to be minted to the address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| weight | [string](#string) |  |  |





 

 

 

 



<a name="ixo/mint/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/mint/v1beta1/genesis.proto



<a name="ixo.mint.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the mint module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [Minter](#ixo.mint.v1beta1.Minter) |  | minter is an abstraction for holding current rewards information. |
| params | [Params](#ixo.mint.v1beta1.Params) |  | params defines all the parameters of the mint module. |
| reduction_started_epoch | [int64](#int64) |  | reduction_started_epoch is the first epoch in which the reduction of mint begins. |





 

 

 

 



<a name="ixo/mint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/mint/v1beta1/query.proto



<a name="ixo.mint.v1beta1.QueryEpochProvisionsRequest"></a>

### QueryEpochProvisionsRequest
QueryEpochProvisionsRequest is the request type for the
Query/EpochProvisions RPC method.






<a name="ixo.mint.v1beta1.QueryEpochProvisionsResponse"></a>

### QueryEpochProvisionsResponse
QueryEpochProvisionsResponse is the response type for the
Query/EpochProvisions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_provisions | [string](#string) |  | epoch_provisions is the current minting per epoch provisions value. |






<a name="ixo.mint.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="ixo.mint.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.mint.v1beta1.Params) |  | params defines the parameters of the module. |





 

 

 


<a name="ixo.mint.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.mint.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.mint.v1beta1.QueryParamsResponse) | Params returns the total set of minting parameters. |
| EpochProvisions | [QueryEpochProvisionsRequest](#ixo.mint.v1beta1.QueryEpochProvisionsRequest) | [QueryEpochProvisionsResponse](#ixo.mint.v1beta1.QueryEpochProvisionsResponse) | EpochProvisions returns the current minting epoch provisions value. |

 



<a name="ixo/smartaccount/crypto/crypto.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/crypto/crypto.proto



<a name="ixo.smartaccount.crypto.AuthnPubKey"></a>

### AuthnPubKey
PubKey defines an authn public key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key_id | [string](#string) |  | The key_id (or credential ID) is a unique identifier for a passkey. This ID is provided by the authenticator when the passkey is created. As it is not possible to retrieve the public key from the authenticator after the passkey is created, if the user loses the public key - id association, the key_id can be used to retrieve the public key. |
| cose_algorithm | [int32](#int32) |  | Store the COSE algorithm identifier. Since authn allows multiple different public key credential algorithm, eg. -7(ES256) - ECDSA with SHA-256 on the P-256 curve -257(RS256) - RSASSA-PKCS1-v1_5 with SHA-256 we need to store the algorithm identifier to be able to verify the signature according to the algorithm the public key is using. |
| key | [bytes](#bytes) |  | The public key bytes |





 

 

 

 



<a name="ixo/token/v1beta1/token.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/token.proto



<a name="ixo.token.v1beta1.CreditsTransferred"></a>

### CreditsTransferred



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| jurisdiction | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| authorization_id | [string](#string) |  |  |






<a name="ixo.token.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo1155_contract_code | [uint64](#uint64) |  |  |






<a name="ixo.token.v1beta1.Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  | generated on token initiation |
| class | [string](#string) |  | class is the token protocol entity DID (validated) |
| name | [string](#string) |  | name is the token name, which must be unique (namespace) |
| description | [string](#string) |  | description is any arbitrary description |
| image | [string](#string) |  | image is the image url for the token |
| type | [string](#string) |  | type is the token type (eg ixo1155) |
| cap | [string](#string) |  | cap is the maximum number of tokens with this name that can be minted, 0 is unlimited |
| supply | [string](#string) |  | how much has already been minted for this Token type, aka the supply |
| paused | [bool](#bool) |  | stop allowance of token minter temporarily |
| stopped | [bool](#bool) |  | stop allowance of token minter permanently |
| retired | [TokensRetired](#ixo.token.v1beta1.TokensRetired) | repeated | tokens that has been retired for this Token with specific name and contract address |
| cancelled | [TokensCancelled](#ixo.token.v1beta1.TokensCancelled) | repeated | tokens that has been cancelled for this Token with specific name and contract address |
| transferred | [CreditsTransferred](#ixo.token.v1beta1.CreditsTransferred) | repeated | credits that has been transferred for this Token with specific name and contract address |






<a name="ixo.token.v1beta1.TokenData"></a>

### TokenData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | media type value should always be &#34;application/json&#34;

credential link ***.ipfs |
| encrypted | [bool](#bool) |  |  |
| proof | [string](#string) |  |  |
| type | [string](#string) |  |  |
| id | [string](#string) |  | did of entity to map token to |






<a name="ixo.token.v1beta1.TokenProperties"></a>

### TokenProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| name | [string](#string) |  | name is the token name, which is same as Token name |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| tokenData | [TokenData](#ixo.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |






<a name="ixo.token.v1beta1.TokensCancelled"></a>

### TokensCancelled



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="ixo.token.v1beta1.TokensRetired"></a>

### TokensRetired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| jurisdiction | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| owner | [string](#string) |  |  |





 

 

 

 



<a name="ixo/token/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/tx.proto



<a name="ixo.token.v1beta1.MintBatch"></a>

### MintBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name is the token name, which must be unique (namespace), will be verified against Token name provided on msgCreateToken |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| amount | [string](#string) |  | amount is the number of tokens to mint |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| token_data | [TokenData](#ixo.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |






<a name="ixo.token.v1beta1.MsgCancelToken"></a>

### MsgCancelToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |






<a name="ixo.token.v1beta1.MsgCancelTokenResponse"></a>

### MsgCancelTokenResponse







<a name="ixo.token.v1beta1.MsgCreateToken"></a>

### MsgCreateToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| class | [string](#string) |  | class is the token protocol entity DID (validated) |
| name | [string](#string) |  | name is the token name, which must be unique (namespace) |
| description | [string](#string) |  | description is any arbitrary description |
| image | [string](#string) |  | image is the image url for the token |
| token_type | [string](#string) |  | type is the token type (eg ixo1155) |
| cap | [string](#string) |  | cap is the maximum number of tokens with this name that can be minted, 0 is unlimited |






<a name="ixo.token.v1beta1.MsgCreateTokenResponse"></a>

### MsgCreateTokenResponse







<a name="ixo.token.v1beta1.MsgMintToken"></a>

### MsgMintToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |
| owner | [string](#string) |  | address of owner to mint for |
| mint_batch | [MintBatch](#ixo.token.v1beta1.MintBatch) | repeated |  |






<a name="ixo.token.v1beta1.MsgMintTokenResponse"></a>

### MsgMintTokenResponse







<a name="ixo.token.v1beta1.MsgPauseToken"></a>

### MsgPauseToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |
| paused | [bool](#bool) |  | pause or unpause Token Minting allowance |






<a name="ixo.token.v1beta1.MsgPauseTokenResponse"></a>

### MsgPauseTokenResponse







<a name="ixo.token.v1beta1.MsgRetireToken"></a>

### MsgRetireToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| jurisdiction | [string](#string) |  | jurisdiction is the jurisdiction of the token owner. A jurisdiction has the format: &lt;country-code&gt;[-&lt;sub-national-code&gt;[ &lt;postal-code&gt;]] The country-code must be 2 alphabetic characters, the sub-national-code can be 1-3 alphanumeric characters, and the postal-code can be up to 64 alphanumeric characters. Only the country-code is required, while the sub-national-code and postal-code are optional and can be added for increased precision. See the valid format for this below. |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |






<a name="ixo.token.v1beta1.MsgRetireTokenResponse"></a>

### MsgRetireTokenResponse







<a name="ixo.token.v1beta1.MsgStopToken"></a>

### MsgStopToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| contract_address | [string](#string) |  |  |






<a name="ixo.token.v1beta1.MsgStopTokenResponse"></a>

### MsgStopTokenResponse







<a name="ixo.token.v1beta1.MsgTransferCredit"></a>

### MsgTransferCredit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | tokens to retire, all tokens must be in same smart contract |
| jurisdiction | [string](#string) |  | jurisdiction is the jurisdiction of the token owner. A jurisdiction has the format: &lt;country-code&gt;[-&lt;sub-national-code&gt;[ &lt;postal-code&gt;]] The country-code must be 2 alphabetic characters, the sub-national-code can be 1-3 alphanumeric characters, and the postal-code can be up to 64 alphanumeric characters. Only the country-code is required, while the sub-national-code and postal-code are optional and can be added for increased precision. See the valid format for this below. |
| reason | [string](#string) |  | reason is any arbitrary string that specifies the reason for retiring tokens. |
| authorization_id | [string](#string) |  | authorization_id is the id of the authorization that was used for the credit transfer |






<a name="ixo.token.v1beta1.MsgTransferCreditResponse"></a>

### MsgTransferCreditResponse







<a name="ixo.token.v1beta1.MsgTransferToken"></a>

### MsgTransferToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | address of owner |
| recipient | [string](#string) |  | address of receiver |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated | all tokens must be in same smart contract |






<a name="ixo.token.v1beta1.MsgTransferTokenResponse"></a>

### MsgTransferTokenResponse







<a name="ixo.token.v1beta1.TokenBatch"></a>

### TokenBatch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id that identifies the token |
| amount | [string](#string) |  | amount is the number of tokens to transfer |





 

 

 


<a name="ixo.token.v1beta1.Msg"></a>

### Msg
Msg defines the project Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateToken | [MsgCreateToken](#ixo.token.v1beta1.MsgCreateToken) | [MsgCreateTokenResponse](#ixo.token.v1beta1.MsgCreateTokenResponse) |  |
| MintToken | [MsgMintToken](#ixo.token.v1beta1.MsgMintToken) | [MsgMintTokenResponse](#ixo.token.v1beta1.MsgMintTokenResponse) |  |
| TransferToken | [MsgTransferToken](#ixo.token.v1beta1.MsgTransferToken) | [MsgTransferTokenResponse](#ixo.token.v1beta1.MsgTransferTokenResponse) |  |
| RetireToken | [MsgRetireToken](#ixo.token.v1beta1.MsgRetireToken) | [MsgRetireTokenResponse](#ixo.token.v1beta1.MsgRetireTokenResponse) |  |
| TransferCredit | [MsgTransferCredit](#ixo.token.v1beta1.MsgTransferCredit) | [MsgTransferCreditResponse](#ixo.token.v1beta1.MsgTransferCreditResponse) |  |
| CancelToken | [MsgCancelToken](#ixo.token.v1beta1.MsgCancelToken) | [MsgCancelTokenResponse](#ixo.token.v1beta1.MsgCancelTokenResponse) |  |
| PauseToken | [MsgPauseToken](#ixo.token.v1beta1.MsgPauseToken) | [MsgPauseTokenResponse](#ixo.token.v1beta1.MsgPauseTokenResponse) |  |
| StopToken | [MsgStopToken](#ixo.token.v1beta1.MsgStopToken) | [MsgStopTokenResponse](#ixo.token.v1beta1.MsgStopTokenResponse) |  |

 



<a name="ixo/smartaccount/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/v1beta1/event.proto



<a name="ixo.smartaccount.v1beta1.AuthenticatorAddedEvent"></a>

### AuthenticatorAddedEvent
AuthenticatorAddedEvent is an event triggered on Authenticator addition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the account that added the authenticator |
| authenticator_type | [string](#string) |  | authenticator_type is the type of the authenticator that was added |
| authenticator_id | [string](#string) |  | authenticator_id is the id of the authenticator that was added |






<a name="ixo.smartaccount.v1beta1.AuthenticatorRemovedEvent"></a>

### AuthenticatorRemovedEvent
AuthenticatorRemovedEvent is an event triggered on Authenticator removal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the account that removed the authenticator |
| authenticator_id | [string](#string) |  | authenticator_id is the id of the authenticator that was removed |






<a name="ixo.smartaccount.v1beta1.AuthenticatorSetActiveStateEvent"></a>

### AuthenticatorSetActiveStateEvent
AuthenticatorSetActiveStateEvent is an event triggered on Authenticator
active state change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender is the address of the account that changed the active state |
| is_smart_account_active | [bool](#bool) |  | active is the new active state |





 

 

 

 



<a name="ixo/smartaccount/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/v1beta1/params.proto



<a name="ixo.smartaccount.v1beta1.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| maximum_unauthenticated_gas | [uint64](#uint64) |  | MaximumUnauthenticatedGas defines the maximum amount of gas that can be used to authenticate a transaction in ante handler without having fee payer authenticated. |
| is_smart_account_active | [bool](#bool) |  | IsSmartAccountActive defines the state of the authenticator. If set to false, the authenticator module will not be used and the classic cosmos sdk authentication will be used instead. |
| circuit_breaker_controllers | [string](#string) | repeated | CircuitBreakerControllers defines list of addresses that are allowed to set is_smart_account_active without going through governance. |





 

 

 

 



<a name="ixo/smartaccount/v1beta1/models.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/v1beta1/models.proto



<a name="ixo.smartaccount.v1beta1.AccountAuthenticator"></a>

### AccountAuthenticator
AccountAuthenticator represents a foundational model for all authenticators.
It provides extensibility by allowing concrete types to interpret and
validate transactions based on the encapsulated data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID uniquely identifies the authenticator instance. |
| type | [string](#string) |  | Type specifies the category of the AccountAuthenticator. This type information is essential for differentiating authenticators and ensuring precise data retrieval from the storage layer. |
| config | [bytes](#bytes) |  | Config is a versatile field used in conjunction with the specific type of account authenticator to facilitate complex authentication processes. The interpretation of this field is overloaded, enabling multiple authenticators to utilize it for their respective purposes. |





 

 

 

 



<a name="ixo/smartaccount/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/v1beta1/genesis.proto



<a name="ixo.smartaccount.v1beta1.AuthenticatorData"></a>

### AuthenticatorData
AuthenticatorData represents a genesis exported account with Authenticators.
The address is used as the key, and the account authenticators are stored in
the authenticators field.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address is an account address, one address can have many authenticators |
| authenticators | [AccountAuthenticator](#ixo.smartaccount.v1beta1.AccountAuthenticator) | repeated | authenticators are the account&#39;s authenticators, these can be multiple types including SignatureVerification, AllOfs, CosmWasmAuthenticators, etc |






<a name="ixo.smartaccount.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the authenticator module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.smartaccount.v1beta1.Params) |  | params define the parameters for the authenticator module. |
| next_authenticator_id | [uint64](#uint64) |  | next_authenticator_id is the next available authenticator ID. |
| authenticator_data | [AuthenticatorData](#ixo.smartaccount.v1beta1.AuthenticatorData) | repeated | authenticator_data contains the data for multiple accounts, each with their authenticators. |





 

 

 

 



<a name="ixo/smartaccount/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/v1beta1/query.proto



<a name="ixo.smartaccount.v1beta1.GetAuthenticatorRequest"></a>

### GetAuthenticatorRequest
MsgGetAuthenticatorRequest defines the Msg/GetAuthenticator request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| authenticator_id | [uint64](#uint64) |  |  |






<a name="ixo.smartaccount.v1beta1.GetAuthenticatorResponse"></a>

### GetAuthenticatorResponse
MsgGetAuthenticatorResponse defines the Msg/GetAuthenticator response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_authenticator | [AccountAuthenticator](#ixo.smartaccount.v1beta1.AccountAuthenticator) |  |  |






<a name="ixo.smartaccount.v1beta1.GetAuthenticatorsRequest"></a>

### GetAuthenticatorsRequest
MsgGetAuthenticatorsRequest defines the Msg/GetAuthenticators request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |






<a name="ixo.smartaccount.v1beta1.GetAuthenticatorsResponse"></a>

### GetAuthenticatorsResponse
MsgGetAuthenticatorsResponse defines the Msg/GetAuthenticators response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_authenticators | [AccountAuthenticator](#ixo.smartaccount.v1beta1.AccountAuthenticator) | repeated |  |






<a name="ixo.smartaccount.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="ixo.smartaccount.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.smartaccount.v1beta1.Params) |  | params holds all the parameters of this module. |





 

 

 


<a name="ixo.smartaccount.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.smartaccount.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.smartaccount.v1beta1.QueryParamsResponse) | Parameters queries the parameters of the module. |
| GetAuthenticator | [GetAuthenticatorRequest](#ixo.smartaccount.v1beta1.GetAuthenticatorRequest) | [GetAuthenticatorResponse](#ixo.smartaccount.v1beta1.GetAuthenticatorResponse) |  |
| GetAuthenticators | [GetAuthenticatorsRequest](#ixo.smartaccount.v1beta1.GetAuthenticatorsRequest) | [GetAuthenticatorsResponse](#ixo.smartaccount.v1beta1.GetAuthenticatorsResponse) |  |

 



<a name="ixo/smartaccount/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/smartaccount/v1beta1/tx.proto



<a name="ixo.smartaccount.v1beta1.MsgAddAuthenticator"></a>

### MsgAddAuthenticator
MsgAddAuthenticatorRequest defines the Msg/AddAuthenticator request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  |  |
| authenticator_type | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="ixo.smartaccount.v1beta1.MsgAddAuthenticatorResponse"></a>

### MsgAddAuthenticatorResponse
MsgAddAuthenticatorResponse defines the Msg/AddAuthenticator response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="ixo.smartaccount.v1beta1.MsgRemoveAuthenticator"></a>

### MsgRemoveAuthenticator
MsgRemoveAuthenticatorRequest defines the Msg/RemoveAuthenticator request
type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  |  |
| id | [uint64](#uint64) |  |  |






<a name="ixo.smartaccount.v1beta1.MsgRemoveAuthenticatorResponse"></a>

### MsgRemoveAuthenticatorResponse
MsgRemoveAuthenticatorResponse defines the Msg/RemoveAuthenticator response
type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="ixo.smartaccount.v1beta1.MsgSetActiveState"></a>

### MsgSetActiveState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  |  |
| active | [bool](#bool) |  |  |






<a name="ixo.smartaccount.v1beta1.MsgSetActiveStateResponse"></a>

### MsgSetActiveStateResponse







<a name="ixo.smartaccount.v1beta1.TxExtension"></a>

### TxExtension
TxExtension allows for additional authenticator-specific data in
transactions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selected_authenticators | [uint64](#uint64) | repeated | selected_authenticators holds the authenticator_id for the chosen authenticator per message. |





 

 

 


<a name="ixo.smartaccount.v1beta1.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddAuthenticator | [MsgAddAuthenticator](#ixo.smartaccount.v1beta1.MsgAddAuthenticator) | [MsgAddAuthenticatorResponse](#ixo.smartaccount.v1beta1.MsgAddAuthenticatorResponse) |  |
| RemoveAuthenticator | [MsgRemoveAuthenticator](#ixo.smartaccount.v1beta1.MsgRemoveAuthenticator) | [MsgRemoveAuthenticatorResponse](#ixo.smartaccount.v1beta1.MsgRemoveAuthenticatorResponse) |  |
| SetActiveState | [MsgSetActiveState](#ixo.smartaccount.v1beta1.MsgSetActiveState) | [MsgSetActiveStateResponse](#ixo.smartaccount.v1beta1.MsgSetActiveStateResponse) | SetActiveState sets the active state of the authenticator. Primarily used for circuit breaking. |

 



<a name="ixo/token/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/authz.proto



<a name="ixo.token.v1beta1.MintAuthorization"></a>

### MintAuthorization



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | address of minter |
| constraints | [MintConstraints](#ixo.token.v1beta1.MintConstraints) | repeated |  |






<a name="ixo.token.v1beta1.MintConstraints"></a>

### MintConstraints



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| name | [string](#string) |  | name is the token name, which must be unique (namespace), will be verified against Token name provided on msgCreateToken |
| index | [string](#string) |  | index is the unique identifier hexstring that identifies the token |
| collection | [string](#string) |  | did of collection (eg Supamoto Malawi) |
| tokenData | [TokenData](#ixo.token.v1beta1.TokenData) | repeated | tokenData is the linkedResources added to tokenMetadata when queried eg (credential link ***.ipfs) |





 

 

 

 



<a name="ixo/token/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/event.proto



<a name="ixo.token.v1beta1.CreditsTransferredEvent"></a>

### CreditsTransferredEvent
CreditsTransferredEvent is an event triggered on a Credit transfer
execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenCancelledEvent"></a>

### TokenCancelledEvent
TokenCancelledEvent is an event triggered on a Token cancel execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenCreatedEvent"></a>

### TokenCreatedEvent
TokenCreatedEvent is an event triggered on a Token creation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#ixo.token.v1beta1.Token) |  |  |






<a name="ixo.token.v1beta1.TokenMintedEvent"></a>

### TokenMintedEvent
TokenMintedEvent is an event triggered on a Token mint execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | the contract address of token contract being initialized |
| minter | [string](#string) |  | the token minter |
| owner | [string](#string) |  | the new tokens owner |
| amount | [string](#string) |  |  |
| tokenProperties | [TokenProperties](#ixo.token.v1beta1.TokenProperties) |  |  |






<a name="ixo.token.v1beta1.TokenPausedEvent"></a>

### TokenPausedEvent
TokenPausedEvent is an event triggered on a Token pause/unpause execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | the minter address |
| contract_address | [string](#string) |  |  |
| paused | [bool](#bool) |  |  |






<a name="ixo.token.v1beta1.TokenRetiredEvent"></a>

### TokenRetiredEvent
TokenRetiredEvent is an event triggered on a Token retire execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the token owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenStoppedEvent"></a>

### TokenStoppedEvent
TokenStoppedEvent is an event triggered on a Token stopped execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | the minter address |
| contract_address | [string](#string) |  |  |
| stopped | [bool](#bool) |  |  |






<a name="ixo.token.v1beta1.TokenTransferredEvent"></a>

### TokenTransferredEvent
TokenTransferedEvent is an event triggered on a Token transfer execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  | the old token owner |
| recipient | [string](#string) |  | the new tokens owner |
| tokens | [TokenBatch](#ixo.token.v1beta1.TokenBatch) | repeated |  |






<a name="ixo.token.v1beta1.TokenUpdatedEvent"></a>

### TokenUpdatedEvent
TokenUpdatedEvent is an event triggered on a Token update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#ixo.token.v1beta1.Token) |  |  |





 

 

 

 



<a name="ixo/token/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/genesis.proto



<a name="ixo.token.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.token.v1beta1.Params) |  |  |
| tokens | [Token](#ixo.token.v1beta1.Token) | repeated |  |
| token_properties | [TokenProperties](#ixo.token.v1beta1.TokenProperties) | repeated |  |





 

 

 

 



<a name="ixo/token/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/proposal.proto



<a name="ixo.token.v1beta1.SetTokenContractCodes"></a>

### SetTokenContractCodes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ixo1155_contract_code | [uint64](#uint64) |  |  |





 

 

 

 



<a name="ixo/token/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ixo/token/v1beta1/query.proto



<a name="ixo.token.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ixo.token.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#ixo.token.v1beta1.Params) |  | params holds all the parameters of this module. |






<a name="ixo.token.v1beta1.QueryTokenDocRequest"></a>

### QueryTokenDocRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minter | [string](#string) |  | minter address to get Token Doc for |
| contract_address | [string](#string) |  |  |






<a name="ixo.token.v1beta1.QueryTokenDocResponse"></a>

### QueryTokenDocResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokenDoc | [Token](#ixo.token.v1beta1.Token) |  |  |






<a name="ixo.token.v1beta1.QueryTokenListRequest"></a>

### QueryTokenListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |
| minter | [string](#string) |  | minter address to get list for |






<a name="ixo.token.v1beta1.QueryTokenListResponse"></a>

### QueryTokenListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |
| tokenDocs | [Token](#ixo.token.v1beta1.Token) | repeated |  |






<a name="ixo.token.v1beta1.QueryTokenMetadataRequest"></a>

### QueryTokenMetadataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="ixo.token.v1beta1.QueryTokenMetadataResponse"></a>

### QueryTokenMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| decimals | [string](#string) |  |  |
| image | [string](#string) |  |  |
| index | [string](#string) |  |  |
| properties | [TokenMetadataProperties](#ixo.token.v1beta1.TokenMetadataProperties) |  |  |






<a name="ixo.token.v1beta1.TokenMetadataProperties"></a>

### TokenMetadataProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [string](#string) |  |  |
| collection | [string](#string) |  |  |
| cap | [string](#string) |  |  |
| linkedResources | [TokenData](#ixo.token.v1beta1.TokenData) | repeated |  |





 

 

 


<a name="ixo.token.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#ixo.token.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#ixo.token.v1beta1.QueryParamsResponse) |  |
| TokenMetadata | [QueryTokenMetadataRequest](#ixo.token.v1beta1.QueryTokenMetadataRequest) | [QueryTokenMetadataResponse](#ixo.token.v1beta1.QueryTokenMetadataResponse) |  |
| TokenList | [QueryTokenListRequest](#ixo.token.v1beta1.QueryTokenListRequest) | [QueryTokenListResponse](#ixo.token.v1beta1.QueryTokenListResponse) |  |
| TokenDoc | [QueryTokenDocRequest](#ixo.token.v1beta1.QueryTokenDocRequest) | [QueryTokenDocResponse](#ixo.token.v1beta1.QueryTokenDocResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

