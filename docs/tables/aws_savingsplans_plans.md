# Table: aws_savingsplans_plans

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| currency | string | X | √ |  | 
| description | string | X | √ |  | 
| product_types | string_array | X | √ |  | 
| savings_plan_arn | string | X | √ |  | 
| savings_plan_type | string | X | √ |  | 
| state | string | X | √ |  | 
| upfront_payment_amount | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| commitment | string | X | √ |  | 
| recurring_payment_amount | string | X | √ |  | 
| term_duration_in_seconds | big_int | X | √ |  | 
| ec2_instance_family | string | X | √ |  | 
| end | string | X | √ |  | 
| payment_option | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) of the Savings Plan.` | 
| offering_id | string | X | √ |  | 
| savings_plan_id | string | X | √ |  | 
| start | string | X | √ |  | 


