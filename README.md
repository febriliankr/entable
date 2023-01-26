# Entable

Generate entity and repo pattern CRUD from PostgreSQL table.

Copy and paste your table schema to {table_name}.txt file and run `go run main.go table_name.txt`.

The program will generate in "generated" folder, with interfaces, entities, repos, and usecases.

## Example Usage

You can try this live.

```
go run main.go example/blog_post.txt
```

### Generated Files

#### Entities
![Entities](https://raw.githubusercontent.com/entable/entable/master/screenshoots/entities.png)

#### Interfaces
![Interfaces](https://raw.githubusercontent.com/entable/entable/master/screenshoots/interfaces.png)

#### Repo
![Repo](https://raw.githubusercontent.com/entable/entable/master/screenshoots/repo.png)

#### Usecase
![Usecase](https://raw.githubusercontent.com/entable/entable/master/screenshoots/usecase.png)

