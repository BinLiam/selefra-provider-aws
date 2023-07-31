# Table: aws_elbv1_load_balancer_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_type_name | string | X | √ |  | 
| region | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_elbv1_load_balancers_selefra_id | string | X | X | fk to aws_elbv1_load_balancers.selefra_id | 
| policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| policy_attribute_descriptions | json | X | √ |  | 


