# Grouper

Grouping tool to randomly group people into evenly sized groups

## Parameters

| Flag           | Meaning                                                                                                         | Default Value |
| -------------- | --------------------------------------------------------------------------------------------------------------- | ------------- |
| `-people-file` | The text file where your names are held. Should be separated by new lines, see `people_test.txt` for an example | `people.txt`  |
| `-num-groups`  | The number of groups to split into                                                                              | 5             |

## Very vague instructions for development

Sorry, being a little lazy here.

```bash
go get -u github.com/emman27/grouper
go run cmd/main.go
```
