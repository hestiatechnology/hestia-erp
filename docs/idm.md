# Identity Management Service
ref: methods/IdentityManagement.go, utils/idm/idm.go

The Identity Management Service (IDM) is a service that provides a way to manage user identities and their access to various resources. It is a key component of the platform, and is used to manage user accounts, roles, and permissions.

# Features
## User Management
TBD

## Password Management
TBD

### Password Encryption
A password is encrypted using the SHA-256 algorithm, then SHA-512 algorithm and again SHA-256 algorithm. This is done to ensure that the password is secure and cannot be easily decrypted. A random salt is also generated and stored with the password to ensure that the same password will not have the same hash.

#### Avoiding MITM Attacks
To avoid MITM attacks, the password is hashed using SHA-256 on the client side before being sent to the server. This ensures that the password is not sent in plain text over the network.