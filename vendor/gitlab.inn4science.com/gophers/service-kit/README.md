# Service Kit

Common libraries for go services:

- [Auth](./auth/README.md) - methods for the services authorization.
- [Crypto](./crypto/README.md) - wrappers for hashing, signing, random values generation etc.
- [Currency](types/currency/README.md) - types and methods for all kind of values, that represents "money" in our services.
- [TX](types/txst/README.md) - types and methods for all kind of values, that came from TX in our services.
- [DB](./db/README.md) - connector for the ORMless interaction with the PostgreSQL databases. 
- [Log](./log/README.md) - simple wrapper for logrus with some useful perks.
- [Routines](./routines/README.md) - implementation, running and controlling workers.
- **Api**
    - [Render](./api/render/README.md) - response helper, base responses
    - [Natswrap](./natswrap/README.md) - simple wrapper for Nats