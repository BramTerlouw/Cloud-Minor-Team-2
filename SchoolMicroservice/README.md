## Auth Headers (REQUIRED)
```
{
    "Authorization": "Bearer <token>"
}
```

## Create School
```graphql
mutation CreateSchool{
    createSchool(input: {
        name: "Sample School",
        location: "Sample location",
        has_openai_access: true,
        openai_key: "<OpenAI Key>"
    }) {
        id
        name
        location
        has_openai_access
        openai_key
        join_code
        made_by
        created_at
        updated_at
        soft_deleted
    }
}
```

## Get one School
```graphql
query GetSchool {
    getSchool(id: "your-school-id") {
        id
        name
        location
        has_openai_access
        openai_key
        join_code
        made_by
        created_at
        updated_at
        soft_deleted
    }
}
```

## Get all schools
```graphql
query ListSchools {
    listSchools(
        filter:{
            softDelete:false,
            name: {
                input:"Sample",
                type: starts
            },
            location: {
                input: "Sample location",
                type: eq
            }
            made_by: "6c1ce448-670f-47b2-83f7-4d771b01775b"
            has_openai_access: true
        },
        paginate:{
            Step: 0,
            amount: 10
        }
    ) {
        id
        name
        location
        made_by
        has_openai_access
    }
}
```

## Update a school
```graphql
mutation UpdateSchool {
    updateSchool(
        id: "your-school-id-to-update",
        input: {
            name: "Updated School Name",
            location: "Updated School Location",
            has_openai_access: true,
            openai_key: "<New OpenAI Key>"
        }
    ) {
        id
        name,
        location,
        created_at,
        updated_at,
        soft_deleted
    }
}
```

## Delete a school
```graphql
mutation DeleteSchool {
    deleteSchool(id: "your-school-id")
}
```