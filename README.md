# AWS S3 Glacier manager
## Support functions

- [ListVaults](https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#Glacier.ListVaults)
- [ListJobs](https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#Glacier.ListJobs)
- [InitiateJob](https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#Glacier.InitiateJob)
  - type `inventory-retrieval`
  - type `archive-retrieval`
- [GetJobOutput](https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#Glacier.GetJobOutput)
  - list archives (`inventory-retrieval`)
- [DeleteArchive](https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#Glacier.DeleteArchive)

## Config

```yml
region: region-name
access_key_id: key_id
secret_access_key: key
vault: vault-name
```
