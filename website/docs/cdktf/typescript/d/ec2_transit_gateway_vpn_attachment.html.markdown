---
subcategory: "Transit Gateway"
layout: "aws"
page_title: "AWS: aws_ec2_transit_gateway_vpn_attachment"
description: |-
  Get information on an EC2 Transit Gateway VPN Attachment
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ec2_transit_gateway_vpn_attachment

Get information on an EC2 Transit Gateway VPN Attachment.

-> EC2 Transit Gateway VPN Attachments are implicitly created by VPN Connections referencing an EC2 Transit Gateway so there is no managed resource. For ease, the [`awsVpnConnection` resource](/docs/providers/aws/r/vpn_connection.html) includes a `transitGatewayAttachmentId` attribute which can replace some usage of this data source. For tagging the attachment, see the [`awsEc2Tag` resource](/docs/providers/aws/r/ec2_tag.html).

## Example Usage

### By Transit Gateway and VPN Connection Identifiers

```typescript
// Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEc2TransitGatewayVpnAttachment } from "./.gen/providers/aws/data-aws-ec2-transit-gateway-vpn-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEc2TransitGatewayVpnAttachment(this, "example", {
      transitGatewayId: Token.asString(awsEc2TransitGatewayExample.id),
      vpnConnectionId: Token.asString(awsVpnConnectionExample.id),
    });
  }
}

```

### Filter

```typescript
// Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsEc2TransitGatewayVpnAttachment } from "./.gen/providers/aws/data-aws-ec2-transit-gateway-vpn-attachment";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsEc2TransitGatewayVpnAttachment(this, "test", {
      filter: [
        {
          name: "resource-id",
          values: ["some-resource"],
        },
      ],
    });
  }
}

```

## Argument Reference

The following arguments are supported:

* `transitGatewayId` - (Optional) Identifier of the EC2 Transit Gateway.
* `vpnConnectionId` - (Optional) Identifier of the EC2 VPN Connection.
* `filter` - (Optional) Configuration block(s) for filtering. Detailed below.
* `tags` - (Optional) Map of tags, each pair of which must exactly match a pair on the desired Transit Gateway VPN Attachment.

### filter Configuration Block

The following arguments are supported by the `filter` configuration block:

* `name` - (Required) Name of the filter field. Valid values can be found in the [EC2 DescribeTransitGatewayAttachments API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayAttachments.html).
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - EC2 Transit Gateway VPN Attachment identifier
* `tags` - Key-value tags for the EC2 Transit Gateway VPN Attachment

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20M`)

<!-- cache-key: cdktf-0.17.1 input-2149138dd4a425dc46526663f1027084972afb2d6643d8f1be249b8702bebeb1 -->