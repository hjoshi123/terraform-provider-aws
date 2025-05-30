---
subcategory: "EMR"
layout: "aws"
page_title: "AWS: aws_emr_release_labels"
description: |-
  Retrieve information about EMR Release Labels
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_emr_release_labels

Retrieve information about EMR Release Labels.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEmrReleaseLabels } from "./.gen/providers/aws/data-aws-emr-release-labels";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEmrReleaseLabels(this, "example", {
      filters: {
        application: "spark@2.1.0",
        prefix: "emr-5",
      },
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `filters` – (Optional) Filters the results of the request. Prefix specifies the prefix of release labels to return. Application specifies the application (with/without version) of release labels to return. See [Filters](#filters).

### Filters

* `application` - (Optional) Optional release label application filter. For example, `Spark@2.1.0` or `Spark`.
* `prefix` - (Optional) Optional release label version prefix filter. For example, `emr-5`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `releaseLabels` - Returned release labels.

<!-- cache-key: cdktf-0.20.8 input-fb78ebfad85778813e9c267d6d243c73637d6545ffda61bd740d4ef131d668e1 -->