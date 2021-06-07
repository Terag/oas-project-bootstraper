Responses
=====

Organize your responses within this folder.

## Content

An example of content of a `response` file

```
description: "Item"
content:
  application/json:
    schema:
      allOf:
        - $ref: "../../../modules/core/schemas/response.yaml"
        - $ref: "../entities/PagedResponse.yaml"
        - type: object
          properties:
            data: 
              $ref: ../entities/Order.yaml
      discriminator:
        propertyName: meta
    examples:
      success:
        $ref: ../examples/responses/GetItemsExample.yaml
```
