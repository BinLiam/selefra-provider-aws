# Table: aws_lightsail_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| domain_validation_records | json | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| serial_number | string | X | √ |  | 
| support_code | string | X | √ |  | 
| account_id | string | X | √ |  | 
| issuer_ca | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| revocation_reason | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| in_use_resource_count | big_int | X | √ |  | 
| name | string | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| renewal_summary | json | X | √ |  | 
| request_failure_reason | string | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| subject_alternative_names | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | X | √ |  | 
| eligible_to_renew | string | X | √ |  | 
| key_algorithm | string | X | √ |  | 


