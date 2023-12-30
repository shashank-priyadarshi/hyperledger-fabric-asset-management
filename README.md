# Hyperledger Fabric Asset Management

Fabric Asset Transfer is a Hyperledger Fabric-based application that enables asset creation and management between multiple organizations. This repository contains the chaincodes, backend, and frontend interfaces, and the deployment scripts, showcasing the power of blockchain technology in facilitating secure and transparent asset transactions. Get started with Fabric Asset Transfer and explore the potential of Hyperledger Fabric in a multi-organization asset management.

## Overview

The project aims to demonstrate the asset creation, transfer and management process within a two-organization environment using Hyperledger Fabric. It showcases the seamless transfer of assets from one organization to another, leveraging the power of blockchain technology.

## Features

- Asset Creation
- Multi-organization asset transfer
- Chaincode implementation for asset creation, transfer and deletion
- Backend for managing transactions
- Frontend for user interaction
- Components

### Chaincodes

The chaincodes contain the smart contract logic for validating and executing asset creation, transfer and deletion transactions on the Hyperledger Fabric network.

### Backends

The backend component provides the necessary APIs and services to facilitate asset transfer, manage transactions, and interact with the blockchain network.

### Frontend

The frontend component offers a user-friendly interface for initiating and monitoring asset transfer transactions between organizations.

## Setup

To set up and run the Fabric Asset Transfer project, follow the instructions in the [project setup documentation](./SETUP.md).

## Usage

- Start the Hyperledger Fabric network.
- Deploy the chaincode on the network.
- Run the backend services.
- Launch the frontend application.
- Contribution.
- Contributions to the Fabric Asset Transfer project are welcome. If you'd like to contribute, please follow our [contribution guidelines](./CONTRIBUTE.md).

## License

This project is licensed under the [MIT License](./LICENSE).

## Wallet


### Generate an unencrypted private key :

```
openssl genpkey -algorithm RSA -out private_key.pem

```

### Generate a certificate signing request (CSR) without a passphrase :

```
openssl req -new -key private_key.pem -out certificate.csr

```
### Generate a self-signed certificate using the CSR:

```

openssl x509 -req -days 365 -in certificate.csr -signkey private_key.pem -out certificate.pem

```