Components
=====

Organize your components within this folder. You will reference your components from your main `openapi.yaml` entrypoint file.

Good conventions to adopt are:

* Name of the components file in `PascalCase` (ex: ItemRef.yaml)
* Name of the fields in the body and the query and uri (path) parameters in `camelCase` or `snake_case` (the second one being recommended as more human-readable)
* Name of the uri (path) parameters in `camelCase` or `kebab-case` (the first one being preferred)
