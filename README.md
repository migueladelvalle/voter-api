# voter-api


This is a simple API which allows the user to register voters and log voter history.

## Endpoints

**- ![##569B4F](https://placehold.co/15x15/569B4F/569B4F.png) GET** /voter/health

returns all available APIs and server uptime. 

**- ![##569B4F](https://placehold.co/15x15/569B4F/569B4F.png) GET** /voters

returns all registered voters

**- ![##DC9F31](https://placehold.co/15x15/DC9F31/DC9F31.png) POST**  /voters/:id

Registers a voter with the specified id. 

**- ![##313DDC](https://placehold.co/15x15/313DDC/313DDC.png) PUT**  /voters/:id

Updates a voter with the specified id. 

**- ![##F41D1D](https://placehold.co/15x15/F41D1D/F41D1D.png) DELETE**  /voters/:id

Removes a voter with the specified id.

**- ![##569B4F](https://placehold.co/15x15/569B4F/569B4F.png) GET** /voters/:id

Retrieves a voter with the specified id.

**- ![##DC9F31](https://placehold.co/15x15/DC9F31/DC9F31.png) POST**  /voters/:id/polls/:pollId

Records a Poll event for the specified voter. 

**- ![##313DDC](https://placehold.co/15x15/313DDC/313DDC.png) PUT**  /voters/:id/polls/:pollId

Upates a Poll event for the specified voter. 

**- ![##F41D1D](https://placehold.co/15x15/F41D1D/F41D1D.png) DELETE**  /voters/:id/polls/:pollId

Deletes a Poll event for the specified voter. 

**- ![##569B4F](https://placehold.co/15x15/569B4F/569B4F.png) GET** /voters/:id/polls/:pollId

Retrieves a specific Poll event with the :pollId from the specified id.

**- ![##569B4F](https://placehold.co/15x15/569B4F/569B4F.png) GET** /voters/:id/polls

Retrieves all Poll history for a specified voter.


## CLI Usage
<pre>
Usage:
  voter-api [flags]
  voter-api [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  restore     Restores the database to a backup file
  start       starts the server

Flags:
  -h, --help      help for voter-api
  -v, --version   print the version

</pre>

### Start

<pre>
Usage:
  voter-api start [flags]

Flags:
  -f, --filePath string   The file path to the Json DB (optional)
  -h, --help              help for start
  -p, --port int          The port on which to start the server (default 3000)

</pre>

### restore
<Pre>

Usage:
  voter-api restore [flags]

Flags:
  -f, --destination string   The file path to the Json DB (default "./Data")
  -h, --help                 help for restore
  -t, --target string        target a specific backup file (default "./Data.Bak")

</Pre>

## Supporting Screenshots

![Alt text](./screenshots/DELETE_voters_id.png?raw=true "Optional Title")
![Alt text](./screenshots/DELETE_voters_voterId_polls_pollId.png?raw=true "Optional Title")
![Alt text](./screenshots/GET_voters.png "Optional Title")
![Alt text](./screenshots/GET_voters_health.png "Optional Title")
![Alt text](./screenshots/GET_voters_id.png "Optional Title")
![Alt text](./screenshots/GET_voters_id_polls.png "Optional Title")
![Alt text](./screenshots/GET_voters_voterId_polls_pollId.png "Optional Title")
![Alt text](./screenshots/POST_voters_id.png "Optional Title") 
![Alt text](./screenshots/POST_voters_voterId_polls_pollId.png "Optional Title")
![Alt text](./screenshots/PUT_voters_id.png "Optional Title")
![Alt text](./screenshots/PUT_voters_voterId_polls_pollId.png "Optional Title")