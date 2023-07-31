# Table: aws_acm_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| failure_reason | string | X | √ |  | 
| imported_at | timestamp | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| subject_alternative_names | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| domain_name | string | X | √ |  | 
| domain_validation_options | json | X | √ |  | 
| in_use_by | string_array | X | √ |  | 
| renewal_summary | json | X | √ |  | 
| signature_algorithm | string | X | √ |  | 
| certificate_authority_arn | string | X | √ |  | 
| serial | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| key_usages | json | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| renewal_eligibility | string | X | √ |  | 
| type | string | X | √ |  | 
| certificate_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| options | json | X | √ |  | 
| subject | string | X | √ |  | 
| extended_key_usages | json | X | √ |  | 
| issuer | string | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| created_at | timestamp | X | √ |  | 


