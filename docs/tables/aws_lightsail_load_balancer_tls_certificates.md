# Table: aws_lightsail_load_balancer_tls_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| subject | string | X | √ |  | 
| name | string | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| status | string | X | √ |  | 
| domain_validation_records | json | X | √ |  | 
| is_attached | bool | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| serial | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| aws_lightsail_load_balancers_selefra_id | string | X | X | fk to aws_lightsail_load_balancers.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| issued_at | timestamp | X | √ |  | 
| location | json | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| resource_type | string | X | √ |  | 
| signature_algorithm | string | X | √ |  | 
| subject_alternative_names | string_array | X | √ |  | 
| issuer | string | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| renewal_summary | json | X | √ |  | 


