## Test


I. Define a full project around the provided contrat: `swagger.yml`
- target of this service is micro-service
- use a postgres database
- rules :
    - register the new user is the invitation uuid exist
    - body and parameters validation, login / provider tuple is unique
    - jwt access and refresh tokens are sent back

II. Review code provided here: `MR`

III.Resolving technical issue about gin buffer: context: `ginRewriteResponse.go`
- on a single json object
- on an array of objects
