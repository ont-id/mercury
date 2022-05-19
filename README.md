# Mercury

Mercury is a trustless, peer to peer decentralized communication protocol that makes entities be able to securely transmit messages, verifiable credentials, and verifiable presentations with each other.

Using Mercury, each entity has a decentralized identity and is identified using decentralized identifiers. Mercury is decentralized identifier agnostic, that is, supports various decentralized identifier methods defined in the W3C DID method registries. Now it supports ONT ID which is developed by the Ontology team, in the future, it will support more decentralized identifiers, it will even support ENS and other decentralized domain naming systems. Mercury employs cryptographic encryption and signature schemes to protect communication in a secure and privacy preserving manner.

Mercury defines several decentralized identifiers based sub-protocols for building communication connection, transmitting messages, verifiable credentials, and verifiable presentations between entities. The sub-protocol family defined by Mercury includes connection protocol, general (encrypted) message exchange protocol as well as verifiable credential and presentation transmission protocol.

## Basic Architecture

In Mercury, entities talk to each other through agents. There are three kinds of agents in this system:

### User Agent

A user agent is under the control of some end entity, and it can be built as or embedded into mobile apps or other rich clients. It is worth noting that user agents have a most important feature that they cannot be online for 24/7.

Entities can keep the corresponding secret keys and store their verifiable credentials in their local storage and then use user agents to initiate communication or transmit messages securely.

### Cloud Agent

In some business cases where user agents cannot be online continuously, messages and credentials need to be relayed, forwarded and stored temporarily. The cloud agents play an important role in routing the messages. The cloud agent itself has a public decentralized identity and an attribute of service endpoint, so its corresponding secret key needs to be properly kept.


### Service Agent

The service agent itself is also a cloud agent, and it also provides services such as the issuance of some verifiable credentials (such as diplomas from third-party institutions). The service agent should also have a public and certified DID , so it also needs to keep the secret key properly.


## Sub-Protocols

### Connection Protocol

The entities who want to talk with others should establish connections with the communication partners. In Mercury, entities could establish connections using connection protocol.

### General Message Exchange Protocol

After establishing a connection between two entities, they can send messages to each other. Messages which are sent to communication partners could be encrypted and signed using some cryptographic schemes for the privacy preservation purpose.

### Verifiable Credential and Presentation Transmission Protocol

In a verifiable credentials system, there are three roles: the holder, the issuer, and the verifiers. The issuer can issue a verifiable credential to the holder at the holder's request, the holder could generate verifiable presentations from their credentials for some proof purposes. The verifiers obtaining the verifiable presentation can verify the presentations cryptographically.

The verifiable credential and presentation transmission protocol defines the methods how the three above roles interact with each other.


## Details and Guide

More Details can be found [here](doc/Design.md).

Developers can follow this [instruction](doc/Usage.md) to try Mercury.
