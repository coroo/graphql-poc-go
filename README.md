Installation

```
go run main.go
```

```
# Write your query or mutation here
{
  product(slug: "super-care-protection") {
    name,
    summary,
    available_claim_methods,
    bundling_with_rider,
    slug,
    riders {
      name,
      slug,
      is_active
    },
    category {
      name
    },
    insurance_type {
      name
    },
    benefit_groups {
      order,
      name,
      benefits {
        name,
        order
      }
    }
  }
}
```

