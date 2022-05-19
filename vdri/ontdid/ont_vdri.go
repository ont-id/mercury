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

package ontdid

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ont-id/mercury/common/message"
	"github.com/ont-id/mercury/service/common"
	"github.com/ont-id/mercury/service/controller"
	"github.com/ont-id/mercury/store"
	"github.com/ont-id/mercury/utils"
	"github.com/ont-id/mercury/vdri"
	sdk "github.com/ontio/ontology-go-sdk"
)

var (
	contexts = []string{"context1", "context2"}
	types    = []string{"mercury"}
)

type SampleSubject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//todo use sdk query from smart contract
type OntVDRI struct {
	ontSdk *sdk.OntologySdk
	acct   *sdk.Account
	did    string
}

func NewOntVDRI(ontsdk *sdk.OntologySdk, acct *sdk.Account, did string) *OntVDRI {
	return &OntVDRI{
		ontSdk: ontsdk,
		acct:   acct,
		did:    did,
	}
}

func (ontVdri *OntVDRI) OfferCredential(req *message.ProposalCredential) (*message.OfferCredential, error) {
	offer := new(message.OfferCredential)
	offer.Type = vdri.OfferCredentialSpec
	offer.Id = utils.GenUUID()
	offer.Connection = common.ReverseConnection(req.Connection)
	offer.CredentialPreview = message.CredentialPreview{Type: "sample", Attributes: []message.Attribute{message.Attribute{
		Name:     "name1",
		MimeType: "json",
		Value:    "{abc}",
	}}}
	offer.Thread = message.Thread{
		ID: req.Id,
	}
	return offer, nil
}
func (ontVdri *OntVDRI) IssueCredential(req *message.RequestCredential) (*message.IssueCredential, error) {

	//for test
	subs := make([]*SampleSubject, 0)
	for _, attach := range req.RequestsAttach {
		s := attach.Data.JSON.(map[string]interface{})
		sample := new(SampleSubject)
		sample.Name = s["name"].(string)
		sample.Value = s["value"].(string)
		subs = append(subs, sample)
	}

	expirationDate := time.Now().UTC().Unix() + 86400

	vc, err := ontVdri.ontSdk.Credential.CreateJWTCredential(contexts, types, subs, req.Connection.TheirDid, expirationDate, "", nil, ontVdri.acct)
	if err != nil {
		return nil, err
	}

	//fixme
	credential := &message.IssueCredential{
		Type:    vdri.IssueCredentialSpec,
		Id:      uuid.New().String(),
		Comment: "ontdid issueCredential",
		Formats: []message.Format{message.Format{
			AttachID: "1",
			Format:   "base64",
		}},
		CredentialsAttach: []message.Attachment{message.Attachment{
			Id:          "1",
			LastModTime: time.Now(),
			Data: message.Data{
				Base64: vc,
			},
		}},
		Connection: common.ReverseConnection(req.Connection),
		Thread: message.Thread{
			ID: req.Id,
		},
	}
	//todo do we need to commit credential to blockchain?

	return credential, nil
}

func (ontVdri *OntVDRI) PresentProof(req *message.RequestPresentation, db store.Store) (*message.Presentation, error) {

	//holderdid := req.Connection.MyDid
	holderdid := req.Connection.TheirDid
	creds := make([]string, 0)
	for _, attachment := range req.RequestPresentationAttach {
		b64 := attachment.Data.Base64
		//should be cred id
		bts, err := utils.Base64Decode(b64)
		if err != nil {
			return nil, err
		}
		credid := string(bts)

		key := []byte(fmt.Sprintf("%s_%s_%s", controller.CredentialKey, holderdid, credid))
		data, err := db.Get(key)
		if err != nil {
			return nil, err
		}

		credrec := new(message.CredentialRec)
		err = json.Unmarshal(data, credrec)
		if err != nil {
			return nil, err
		}

		//todo check with format and related id
		s := credrec.Credential.CredentialsAttach[0].Data.Base64
		creds = append(creds, s)
	}

	presentation := new(message.Presentation)
	ps, err := ontVdri.ontSdk.Credential.CreateJWTPresentation(creds, contexts, types, holderdid, "", "", ontVdri.acct)
	if err != nil {
		return nil, err
	}
	presentation.Type = vdri.PresentationProofSpec
	presentation.Id = utils.GenUUID()
	presentation.Connection = common.ReverseConnection(req.Connection)
	presentation.Formats = []message.Format{message.Format{
		AttachID: "1", //fixed index
		Format:   "base64",
	}}
	presentation.PresentationAttach = []message.Attachment{
		{
			Id:          "1", //fixed index
			LastModTime: time.Now(),
			Data: message.Data{
				Base64: ps,
			},
		},
	}
	presentation.Thread = message.Thread{
		ID: req.Id,
	}

	return presentation, nil
}
func (o OntVDRI) GetDIDDoc(did string) (vdri.CommonDIDDoc, error) {
	bts, err := o.ontSdk.Native.OntId.GetDocumentJson(utils.CutDId(did))
	if err != nil {
		return nil, err
	}
	doc := new(message.DIDDoc)
	err = json.Unmarshal(bts, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (o OntVDRI) VerifyDID(did string) bool {
	return strings.HasPrefix(utils.CutDId(did), "did:ont:")
}
