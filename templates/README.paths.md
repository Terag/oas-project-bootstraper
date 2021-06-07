Paths
=====

Organize your path definitions within this folder.  You will reference your paths from your main `openapi.yaml` entrypoint file.

Good conventions to adopt are:

* path separator token (replace `/` with `@`)
* uri parameters (e.g. `{example}`)
* file-per-path

## Each path in a separate file

Use a predefined "path separator" (`@` is highly recommended) and keep all of your path files in the top level of the `paths` folder.

example
```
/items
    GET
    POST
/items/{item_ref}
    GET
    PUT
    PATCH
    DELETE
/items/{item_ref}/actions/do-something
    POST
```
would be defined in files
```
/items.yaml
/items@{item_ref}.yaml
/items@{item_ref}@actions@do-something.yaml
```

## Content

An example of content of a `path` file

```yaml
parameters:
  - $ref: "../components/parameters/path/ItemRef.yaml"
get:
  operationId: GetItemProducts
  tags:
    - ItemV1
  summary: Return the list of all products for an item.
  description: |
    Return the list of all the product for an item. The list can be filtered and the use pagination is recommended

    Error | Code | Description
    ----- | ----- | -----
    400   | BAD_REQUEST        | The input doesn't respect the contract expected (required fields, type, etc.)
    401   | UNAUTHORIZED       | Missing, invalid or expired token. To fix, you should re-authenticate the user.
    403   | FORBIDDEN          | Access is forbidden
    404   | NOT_FOUND          | The requested resource doesn't exist.
    405   | METHOD_NOT_ALLOWED | A request was made of a resource using a request method not supported by that resource; for example, using GET on a form which requires data to be presented via POST, or using PUT on a read-only resource.
    406   | NOT_ACCEPTABLE     | The requested resource is only capable of generating content not acceptable according to the Accept headers sent in the request.
    500   | INTERNAL_ERROR     | We had a problem with our server. Please to try again later.
    501   | NOT_IMPLEMENTED    | For the context of the current business unit, this feature is not supported.
    502   | BAD_GATEWAY        | We had a problem with one of our backends that returns a http 500 status. Please to try again later.
  parameters:
    - $ref: "../../modules/core/resources/parameters/query/pageIndex.yaml"
    - $ref: "../../modules/core/resources/parameters/query/pageSize.yaml"
  responses:
    '200':
      $ref: "../components/responses/GetServiceCardProductsResponse.yaml"
    '400':
      $ref: "../../modules/core/resources/schemas/errors/400.yaml"
    '401':
      $ref: "../../modules/core/resources/schemas/errors/401.yaml"
    '403':
      $ref: "../../modules/core/resources/schemas/errors/403.yaml"
    '404':
      $ref: "../../modules/core/resources/schemas/errors/404.yaml"
    '405':
      $ref: "../../modules/core/resources/schemas/errors/405.yaml"
    '406':
      $ref: "../../modules/core/resources/schemas/errors/406.yaml"
    '500':
      $ref: "../../modules/core/resources/schemas/errors/500.yaml"
    '501':
      $ref: "../../modules/core/resources/schemas/errors/501.yaml"
    '502':
      $ref: "../../modules/core/resources/schemas/errors/502.yaml"
post:
  operationId: AddItemProduct
  tags:
    - ItemV1
  summary: Add an Item Product.
  description: |
    Add a Service Card Product.

    Reference the possible error codes here
    Error | Code | Description
    ----- | ----- | -----
    400   | BAD_REQUEST        | The input doesn't respect the contract expected (required fields, type, etc.)
    401   | UNAUTHORIZED       | Missing, invalid or expired token. To fix, you should re-authenticate the user.
    403   | FORBIDDEN          | Access is forbidden
    404   | NOT_FOUND          | The requested resource doesn't exist.
    405   | METHOD_NOT_ALLOWED | A request was made of a resource using a request method not supported by that resource; for example, using GET on a form which requires data to be presented via POST, or using PUT on a read-only resource.
    406   | NOT_ACCEPTABLE     | The requested resource is only capable of generating content not acceptable according to the Accept headers sent in the request.
    500   | INTERNAL_ERROR     | We had a problem with our server. Please to try again later.
    501   | NOT_IMPLEMENTED    | For the context of the current business unit, this feature is not supported.
    502   | BAD_GATEWAY        | We had a problem with one of our backends that returns a http 500 status. Please to try again later.
  requestBody:
    content:
      application/json:
        schema:
          $ref: "../components/entities/Product.yaml"
        examples:
          AddItemProductExample:
            $ref: "../components/examples/requests/AddItemProductExample.yaml"
    required: true
  responses:
    '200':
      $ref: "../components/responses/GetItemProductsResponse.yaml"
    '400':
      $ref: "../../modules/core/resources/schemas/errors/400.yaml"
    '401':
      $ref: "../../modules/core/resources/schemas/errors/401.yaml"
    '403':
      $ref: "../../modules/core/resources/schemas/errors/403.yaml"
    '404':
      $ref: "../../modules/core/resources/schemas/errors/404.yaml"
    '405':
      $ref: "../../modules/core/resources/schemas/errors/405.yaml"
    '406':
      $ref: "../../modules/core/resources/schemas/errors/406.yaml"
    '500':
      $ref: "../../modules/core/resources/schemas/errors/500.yaml"
    '501':
      $ref: "../../modules/core/resources/schemas/errors/501.yaml"
    '502':
      $ref: "../../modules/core/resources/schemas/errors/502.yaml"
```
