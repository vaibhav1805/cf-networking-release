# Dynamic Egress Policies and Destinations APIs - EXPERIMENTAL

NOTE: These APIs are EXPERIMENTAL.

## Turn on the Dynamic Egress Feature

The dynamic egress feature set is not turned on by default. In order for these
policies to be enforced, the cloud foundry operator must set the
`enforce_experimental_dynamic_egress_policies` property on the
`policy-server-internal` job to true. See spec file
[here](https://github.com/cloudfoundry/cf-networking-release/blob/develop/jobs/policy-server-internal/spec).
Currently this feature only works with silk.

## Purpose

These APIs are for creating, deleting, listing, and updating dynamic egress
destinations and policies.

## API Authorization

In order to communicate with the policy server API, a UAA oauth token with valid
`network.admin`.  The CF admin by default has `network.admin` scope, other users
will need to have the proper scope granted by an admin.

### Option 1: cf curl

Use the `cf curl` command as admin

Example

```bash
$ cf curl /networking/v1/external/egress_policies
{"total_egress_policies":2,"egress_policies":[{"source":{...}]}
```

### Option 2: curl

When using curl the token must be explicitly provided in the `Authorization` header.

Example

```bash
$ export TOKEN=`cf oauth-token` # as CF admin
$ curl http://api.bosh-lite.com/networking/v1/external/egress_policies -H "Authorization: $TOKEN"
{"total_egress_policies":2,"egress_policies":[{"source":{...}]}
```

## Typical Workflows

Adding an Egress Policy
1. Create an egress destination.
1. Create an egress policy linking the destination and an app or space.
1. See policy apply. No app restarts needed.

Updating an Egress Policy when an IP changes
1. List all egress destinations to find the one you want to change.
1. Update the egress destination.
1. See updated policy apply. No app restarts needed.

<hr>

## Egress Destination API

| Method | Path |  Description |
| :----- | :--- |  :----------- |
| GET | /networking/v1/external/destinations |  List Destinations |
| POST | /networking/v1/external/destinations | Create Destinations |
| PUT | /networking/v1/external/destinations | Update Destinations |
| DELETE | /networking/v1/external/destinations/GUID | Delete Destination |


### List Egress Destinations
#### GET /networking/v1/external/destinations
#### Arguments:

[optionally] `id`: comma-separated id values. This cannot be used with `name`.\
[optionally] `name`: comma-separated name values. This cannot be used with `id`.\

Will return all egress destinations.

#### Response Body:

```json
{
  "total_destinations": 2,
  "destinations": [
   {
      "name": "oracle database",
      "id": "90be9c1f-b694-4463-9f1f-6ce71904440d",
      "description": "db for user accounts",
      "rules": [
        {
          "ips": "1.9.9.9-1.9.9.20",
          "ports": "8000-9000",
          "protocol": "tcp",
          "description": "our user db"
        }
      ]
   },
   {
      "name": "AWS",
      "id": "72813418-bd38-49e0-ace0-7bf5b7c54687",
      "rules": [
        {
          "ips": "1.8.8.8-1.8.8.8",
          "ports": "8000-9000",
          "protocol": "udp",
          "description": "our user db on aws"
        },
        {
          "ips": "1.8.8.9-1.8.8.9",
          "protocol": "tcp",
          "description": "our frontend app app "
        },
        {
          "ips": "1.8.9.0-1.8.9.1",
          "ports": "8000-9000",
          "icmp_type": 14,
          "icmp_code": 0,
          "protocol": "all",
          "description": ""
        }
      ]
   }
  ]
}
```

### Create Egress Destinations
#### POST /networking/v1/external/destinations

#### Request Body:

```json
{
  "destinations": [
   {
      "name": "oracle database",
      "description": "db for user accounts",
      "rules": [
        {
          "ips": "1.9.9.9-1.9.9.20",
          "ports": "8000-9000",
          "protocol": "tcp",
          "description": "our user db"
        }
      ]
   },
   {
      "name": "AWS",
      "rules": [
        {
          "ips": "1.8.8.8-1.8.8.8",
          "ports": "8000-9000",
          "protocol": "udp"
        },
        {
          "ips": "9.9.9.9-10.10.10.10",
          "ports": "8001-9001",
          "protocol": "tcp",
          "description": "frontend app"
        },
        {
         "ips": "1.8.9.0-1.8.9.1",
          "ports": "8000-9000",
          "icmp_type": 14,
          "icmp_code": 0,
          "protocol": "all"
        }
      ]
   }
  ]
}
```

| Field | Required? | Description |
| :---- | :-------: | :------ |
| destinations.name | Y | The name of the destination. Must be globally unique.
| destinations.description | N | A description of the destination.
| destinations.rules.ips* | Y | The destination IP range. Must take the form X.X.X.X-X.X.X.X. Must be IPV4. Start and end IPs may be equal.
| destinations.rules.ports* | Y | The destination port range. Must take the form <start>-<end>. Start and end ports may be equal. Ports are not applicable for ICMP protocol.
| destinations.rules.protocol | Y | The protocol (`tcp`, `udp`, `icmp`, or `all`). The value `all` will allow tcp, udp, udplite, icmp, esp, ah, sctp protocols.
| destinations.rules.icmp_type | N | The icmp type to allow when using the icmp protocol. Default is all icmp types, represented by -1.
| destinations.rules.icmp_code | N | The icmp code to allow when using the icmp protocol. Default is all icmp codes, represented by -1.
| destinations.rules.description | N | A text field to describe the rule. Not unique. Not required.

> Note: Currently only one ip range and one port range is supported.
> In the future, a destination will be able to support multiple ip ranges and port ranges.

### Update Egress Destinations
#### PUT /networking/v1/external/destinations

#### Request Body:

```json
{
  "destinations": [
    {
      "id": "90be9c1f-b694-4463-9f1f-6ce71904440d",
      "name": "oracle database",
      "description": "db for user accounts",
      "rules": [
        {
          "ips": "1.9.9.9-1.9.9.20",
          "ports": "8000-9000",
          "protocol": "tcp"
        }
      ]
    },
    {
      "id": "72813418-bd38-49e0-ace0-7bf5b7c54687",
      "name": "AWS",
      "rules": [
        {
          "ips": "1.8.8.8-1.8.8.8",
          "ports": "8000-9000",
          "protocol": "udp",
          "description": "frontend app"
        }
      ]
    }
  ]
}
```

| Field | Required? | Description |
| :---- | :-------: | :------ |
| destinations.id | Y | The id of the destination. This id is returned in the destinations create response, as well as in the destinations index response.
| destinations.name | Y | The name of the destination. Must be globally unique.
| destinations.description | N | A description of the destination.
| destinations.rules.ips* | Y | The destination IP range. Must take the form X.X.X.X-X.X.X.X. Must be IPV4. Start and end IPs may be equal.
| destinations.rules.ports* | Y | The destination port range. Must take the form <start>-<end>. Start and end ports may be equal. Ports are not applicable for ICMP protocol.
| destinations.rules.protocol | Y |The protocol (`tcp`, `udp`, `icmp`, or `all`). The value `all` will allow tcp, udp, udplite, icmp, esp, ah, sctp protocols.
| destinations.rules.icmp_type | N | The icmp type to allow when using the icmp protocol. Default is all icmp types, represented by -1.
| destinations.rules.icmp_code | N | The icmp code to allow when using the icmp protocol. Default is all icmp codes, represented by -1.
| destinations.rules.description | N | A text field to describe the rule. Not unique. Not required.

> Note: Currently only one ip range and one port range is supported.
> In the future, a destination will be able to support multiple ip ranges and port ranges.

### Delete an Egress Destination

### DELETE /networking/v1/external/destinations/GUID

#### Response Body:

This endpoint returns the json of the deleted destination object.

```json
{
  "total_destinations": 1,
  "destinations": [
   {
      "name": "oracle database",
      "id": "90be9c1f-b694-4463-9f1f-6ce71904440d",
      "description": "db for user accounts",
      "rules": [
        {
          "ips": "1.9.9.9-1.9.9.20",
          "ports": "8000-9000",
          "protocol": "tcp"
        },
        {
          "ips": "1.1.1.1-1.1.1.1",
          "ports": "80-80",
          "protocol": "udp"
        }
      ]
   }
  ]
}
```

<hr>

## Egress Policy API

| Method | Path |  Description|
| :----- | :--- |  :----------- |
| GET | /networking/v1/external/egress_policies |   List Egress Policies |
| POST | /networking/v1/external/egress_policies |  Create EgressPolicies |
| DELETE | /networking/v1/external/egress_policies/GUID | Delete Egress Policy |

### List Egress Policies
#### GET /networking/v1/external/egress_policies
#### Arguments: None

Will return all egress policies.

#### Response Body:

```json
{
  "total_egress_policies": 1,
  "egress_policies": [{
    "id": "dynamic-egress-guid",
    "source": {
      "type": "app",
      "id": "SOURCE-APP-GUID"
     },
     "destination": {
        "id": "guid-abc-123",
        "name": "AWS",
        "description": "AWS",
        "id": "72813418-bd38-49e0-ace0-7bf5b7c54687",
        "rules": [
          {
            "ips": "1.8.8.8-1.8.8.8",
            "ports": "8000-9000",
            "protocol": "udp",
            "description": "frontend app"
          }
        ]
     },
     "app_lifecycle": "all"
   }]
}
```

### Create Egress Policies
#### POST /networking/v1/external/egress_policies

#### Request Body:

```json
{
  "egress_policies": [{
    "source": {
      "type": "space",
      "id": "SOURCE-SPACE-GUID"
    },
    "destination": {
      "id": "EGRESS-DESTINATION-GUID"
    },
    "app_lifecycle": "running"
  }]
}
```

| Field | Required? | Description |
| :---- | :-------: | :------ |
| source.type | N | The type of source. Must be `app`, `space`, or `default`. If `default` is provided, then the policy applies to every app in the foundation. If nothing is provided, defaults to `app`.
| source.id | Y/N | The guid of the source app or space. This is not allowed when the source type is `default`.
| destination.id | Y | The guid of the egress destination.
| app_lifecycle | N | The part of the app lifecycle you want the policy to apply to. Must be `running`, `staging`, or `all`. The `running` value applies to an app once it has started and to tasks. The `staging` value applies to apps when they are staging e.g when an app is being built by a buildpack during `cf push`. And the `all` value applies to both `running` and `staging`. Defaults to `all`.

### Delete an Egress Policy

### DELETE /networking/v1/external/egress_policies/GUID

#### Response Body:

This endpoint returns the json of the deleted egress policy object.


```json
{
  "total_egress_policies": 1,
  "egress_policies": [{
    "id": "dynamic-egress-guid",
    "source": {
      "type": "app",
      "id": "SOURCE-APP-GUID"
    },
    "destination": {
      "name": "AWS",
      "description": "AWS",
      "id": "72813418-bd38-49e0-ace0-7bf5b7c54687",
      "rules": [
        {
          "ips": "1.8.8.8-1.8.8.8",
          "ports": "8000-9000",
          "protocol": "udp",
          "description": "frontend app"
        }
      ]
     },
     "app_lifecycle": "staging"
   }]
}
```
