/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package common

type MessageType int

const (
	InvitationType MessageType = iota
	ConnectionRequestType
	ConnectionResponseType
	ConnectionAckType
	SendDisconnectType
	DisconnectType

	SendProposalCredentialType
	ProposalCredentialType
	OfferCredentialType
	SendRequestCredentialType
	RequestCredentialType
	IssueCredentialType
	CredentialAckType
	DeleteCredentialType
	QueryCredentialType

	SendRequestPresentationType
	RequestPresentationType
	PresentationType
	PresentationAckType
	DeletePresentationType
	QueryPresentationType

	SendBasicMsgType
	ReceiveBasicMsgType
	QueryBasicMessageType
	QueryConnectionsType
)

type Message struct {
	MessageType `json:"type"`
	Content     interface{} `json:"content"`
}

const (
	InviteApi                    = "/api/v1/invitation"
	ConnectRequestApi            = "/api/v1/connectionrequest"
	ConnectResponseApi           = "/api/v1/connectionresponse"
	ConnectAckApi                = "/api/v1/connectionack"
	SendDisconnectApi            = "/api/v1/senddisconnect"
	DisconnectApi                = "/api/v1/disconnect"
	SendProposalCredentialReqApi = "/api/v1/sendproposalcredential"
	ProposalCredentialReqApi     = "/api/v1/proposalcredential"
	OfferCredentialApi           = "/api/v1/offercredential"
	SendRequestCredentialApi     = "/api/v1/sendrequestcredential"
	RequestCredentialApi         = "/api/v1/requestcredential"
	IssueCredentialApi           = "/api/v1/issuecredentail"
	CredentialAckApi             = "/api/v1/credentialack"
	DeleteCredentialApi          = "/api/v1/deletecredential"
	QueryCredentialApi           = "/api/v1/querycredential"
	SendRequestPresentationApi   = "/api/v1/sendrequestpresentation"
	RequestPresentationApi       = "/api/v1/requestpresentation"
	PresentationProofApi         = "/api/v1/presentproof"
	PresentationAckApi           = "/api/v1/presentationack"
	QueryPresentationApi         = "/api/v1/querypresentation"
	DeletePresentationApi        = "/api/v1/deletepresentation"
	SendBasicMsgApi              = "/api/v1/sendbasicmsg"
	ReceiveBasicMsgApi           = "/api/v1/receivebasicmsg"
	QueryBasicMsgApi             = "/api/v1/querybasicmsg"
	QueryConnectionsApi          = "/api/v1/queryconnections"
)

func GetApiName(msgType MessageType) string {
	switch msgType {
	case InvitationType:
		return InviteApi
	case ConnectionRequestType:
		return ConnectRequestApi
	case ConnectionResponseType:
		return ConnectResponseApi
	case ConnectionAckType:
		return ConnectAckApi
	case DisconnectType:
		return DisconnectApi
	case ProposalCredentialType:
		return ProposalCredentialReqApi
	case OfferCredentialType:
		return OfferCredentialApi
	case RequestCredentialType:
		return RequestCredentialApi
	case IssueCredentialType:
		return IssueCredentialApi
	case CredentialAckType:
		return CredentialAckApi
	case RequestPresentationType:
		return RequestPresentationApi
	case PresentationType:
		return PresentationProofApi
	case PresentationAckType:
		return PresentationAckApi
	case SendBasicMsgType:
		return SendBasicMsgApi
	case ReceiveBasicMsgType:
		return ReceiveBasicMsgApi
	case QueryBasicMessageType:
		return QueryBasicMsgApi
	case QueryCredentialType:
		return QueryCredentialApi
	case QueryPresentationType:
		return QueryPresentationApi
	case QueryConnectionsType:
		return QueryConnectionsApi
	default:
		return ""
	}
}

func TransferForwardMsgType(msgType MessageType) MessageType {
	switch msgType {
	case SendBasicMsgType:
		return ReceiveBasicMsgType
	case SendRequestCredentialType:
		return RequestCredentialType
	case SendProposalCredentialType:
		return ProposalCredentialType
	case SendDisconnectType:
		return DisconnectType
	case SendRequestPresentationType:
		return RequestPresentationType
	default:
		return msgType
	}
}
