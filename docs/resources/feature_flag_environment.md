---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "launchdarkly_feature_flag_environment Resource - launchdarkly"
subcategory: ""
description: |-
  Provides a LaunchDarkly environment-specific feature flag resource.
  This resource allows you to create and manage environment-specific feature flags attributes within your LaunchDarkly organization.
  -> Note: If you intend to attach a feature flag to any experiments, we do not recommend configuring environment-specific flag settings using Terraform. Subsequent applies may overwrite the changes made by experiments and break your experiment. An alternate workaround is to use the lifecycle.ignore_changes https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes Terraform meta-argument on the fallthrough field to prevent potential overwrites.
---

# launchdarkly_feature_flag_environment (Resource)

Provides a LaunchDarkly environment-specific feature flag resource.

This resource allows you to create and manage environment-specific feature flags attributes within your LaunchDarkly organization.

-> **Note:** If you intend to attach a feature flag to any experiments, we do _not_ recommend configuring environment-specific flag settings using Terraform. Subsequent applies may overwrite the changes made by experiments and break your experiment. An alternate workaround is to use the [lifecycle.ignore_changes](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#ignore_changes) Terraform meta-argument on the `fallthrough` field to prevent potential overwrites.

## Example Usage

```terraform
resource "launchdarkly_feature_flag_environment" "number_env" {
  flag_id = launchdarkly_feature_flag.number.id
  env_key = launchdarkly_environment.staging.key

  on = true

  prerequisites {
    flag_key  = launchdarkly_feature_flag.basic.key
    variation = 0
  }

  targets {
    values    = ["user0"]
    variation = 0
  }
  targets {
    values    = ["user1", "user2"]
    variation = 1
  }
  context_targets {
    values       = ["accountX"]
    variation    = 1
    context_kind = "account"
  }

  rules {
    description = "example targeting rule with two clauses"
    clauses {
      attribute = "country"
      op        = "startsWith"
      values    = ["aus", "de", "united"]
      negate    = false
    }
    clauses {
      attribute = "segmentMatch"
      op        = "segmentMatch"
      values    = [launchdarkly_segment.example.key]
      negate    = false
    }
    variation = 0
  }

  fallthrough {
    rollout_weights = [60000, 40000, 0]
    context_kind    = "account"
    bucket_by       = "accountId"
  }
  off_variation = 2
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `env_key` (String) The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
- `fallthrough` (Block List, Min: 1, Max: 1) Nested block describing the default variation to serve if no `prerequisites`, `target`, or `rules` apply. (see [below for nested schema](#nestedblock--fallthrough))
- `flag_id` (String) The feature flag's unique `id` in the format `project_key/flag_key`. A change in this field will force the destruction of the existing resource and the creation of a new one.
- `off_variation` (Number) The index of the variation to serve if targeting is disabled.

### Optional

- `context_targets` (Block Set) The set of nested blocks describing the individual targets for non-user context kinds for each variation. (see [below for nested schema](#nestedblock--context_targets))
- `on` (Boolean) Whether targeting is enabled. Defaults to `false` if not set.
- `prerequisites` (Block List) List of nested blocks describing prerequisite feature flags rules. (see [below for nested schema](#nestedblock--prerequisites))
- `rules` (Block List) List of logical targeting rules. (see [below for nested schema](#nestedblock--rules))
- `targets` (Block Set) Set of nested blocks describing the individual user targets for each variation. (see [below for nested schema](#nestedblock--targets))
- `track_events` (Boolean) Whether to send event data back to LaunchDarkly. Defaults to `false` if not set.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--fallthrough"></a>
### Nested Schema for `fallthrough`

Optional:

- `bucket_by` (String) Group percentage rollout by a custom attribute. This argument is only valid if rollout_weights is also specified.
- `context_kind` (String) The context kind associated with the specified rollout. This argument is only valid if rollout_weights is also specified. If omitted, defaults to `user`.
- `rollout_weights` (List of Number) List of integer percentage rollout weights (in thousandths of a percent) to apply to each variation if the rule clauses evaluates to `true`. The sum of the `rollout_weights` must equal 100000 and the number of rollout weights specified in the array must match the number of flag variations. You must specify either `variation` or `rollout_weights`.
- `variation` (Number) The default integer variation index to serve if no `prerequisites`, `target`, or `rules` apply. You must specify either `variation` or `rollout_weights`.


<a id="nestedblock--context_targets"></a>
### Nested Schema for `context_targets`

Required:

- `context_kind` (String) The context kind on which the flag should target in this environment. User (`user`) targets should be specified as `targets` attribute blocks.
- `values` (List of String) List of `user` strings to target.
- `variation` (Number) The index of the variation to serve if a user target value is matched.


<a id="nestedblock--prerequisites"></a>
### Nested Schema for `prerequisites`

Required:

- `flag_key` (String) The prerequisite feature flag's `key`.
- `variation` (Number) The index of the prerequisite feature flag's variation to target.


<a id="nestedblock--rules"></a>
### Nested Schema for `rules`

Optional:

- `bucket_by` (String) Group percentage rollout by a custom attribute. This argument is only valid if `rollout_weights` is also specified.
- `clauses` (Block List) List of nested blocks specifying the logical clauses to evaluate (see [below for nested schema](#nestedblock--rules--clauses))
- `description` (String) A human-readable description of the targeting rule.
- `rollout_weights` (List of Number) List of integer percentage rollout weights (in thousandths of a percent) to apply to each variation if the rule clauses evaluates to `true`. The sum of the `rollout_weights` must equal 100000 and the number of rollout weights specified in the array must match the number of flag variations. You must specify either `variation` or `rollout_weights`.
- `variation` (Number) The integer variation index to serve if the rule clauses evaluate to `true`. You must specify either `variation` or `rollout_weights`

<a id="nestedblock--rules--clauses"></a>
### Nested Schema for `rules.clauses`

Required:

- `attribute` (String) The user attribute to operate on
- `op` (String) The operator associated with the rule clause. Available options are `in`, `endsWith`, `startsWith`, `matches`, `contains`, `lessThan`, `lessThanOrEqual`, `greaterThanOrEqual`, `before`, `after`, `segmentMatch`, `semVerEqual`, `semVerLessThan`, and `semVerGreaterThan`.
- `values` (List of String) The list of values associated with the rule clause.

Optional:

- `context_kind` (String) The context kind associated with this rule clause. This argument is only valid if `rollout_weights` is also specified. If omitted, defaults to `user`.
- `negate` (Boolean) Whether to negate the rule clause.
- `value_type` (String) The type for each of the clause's values. Available types are `boolean`, `string`, and `number`. If omitted, `value_type` defaults to `string`.



<a id="nestedblock--targets"></a>
### Nested Schema for `targets`

Required:

- `values` (List of String) List of `user` strings to target.
- `variation` (Number) The index of the variation to serve if a user target value is matched.

## Import

Import is supported using the following syntax:

```shell
# LaunchDarkly feature flag environments can be imported using the resource's ID in the form `project_key/env_key/flag_key`
terraform import launchdarkly_feature_flag_environment.example example-project/example-env/example-flag-key
```
