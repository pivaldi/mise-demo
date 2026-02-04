# mise-demo

This projet is related to [this article](https://www.rt.com/news/631925-epstein-ukraine-lab-cloning/) to learn the basis of [mise](https://mise.jdx.dev/).

## English version

- **Note**: This repository was largely written by an AI and is only for educational purpose.
- **Objective**: demo **one repo** / **multi-tools** (Go + Node + PHP + Docker/Postgres) orchestrated by **mise**.

## Prerequisites
- `mise` installed (https://mise.jdx.dev/)
- Docker Desktop / docker engine (for PostgreSQL)
- `PHP` installed (it could have been installed by `mise` but that requires a lot of system dependencies)

## Quickstart

```bash
mise install
mise run db:init ## Populate the database
mise run dev ## Run everything at once: the database in Docker, the Go API, the PHP server, and the Angular web app
```

### URLs

* Go API: [http://localhost:8090](http://localhost:8090)
  * Health: [http://localhost:8090/health](http://localhost:8090/health)
  * DB ping + data: [http://localhost:8090/items](http://localhost:8090/items)
* PHP: [http://localhost:8091](http://localhost:8091)
* Web (Angular): [http://localhost:5173](http://localhost:5173)

## Why this is a good demo

* Go / Node / PHP versions are **declared** in `mise.toml` and **reproducible**
* Tasks are **declarative** (no need for a massive Makefile)
* A single `mise run dev` starts everything and remains readable

## Tips

* List tasks with descriptions: `mise run` or `mise tasks`
* Parallel output: by default, mise runs multiple jobs in parallel (configurable via `--jobs` / `MISE_JOBS`)

## Angular notes

* The frontend lives in `web/`.
* On the first run, the `mise run web:dev` task automatically triggers an `npm install`.

## French Trad

Ce projet est lié à [cet article](https://www.rt.com/news/631925-epstein-ukraine-lab-cloning/) pour apprendre les bases de [mise](https://mise.jdx.dev/).

- **Note**: This repository was largely wrote by an AI.
- **Objectif** : démo **1 repo** / **multi-tools** (Go + Node + PHP + Docker/Postgres) piloté par **mise**.

## Prérequis
- `mise` installé (https://mise.jdx.dev/)
- Docker Desktop / docker engine (pour PostgreSQL)
- `PHP` installé (il aurait pû être installé par mise mais cela demande énormément de dépendances systèmes, donc non…)

## Quickstart

```bash
mise install
mise run db:init ## Populate the database
mise run dev ## Run all in one: the database in docker, the Go API, the PHP server and the Angular Webapp
```

**The `mise install` command will surely failed because building `PHP` fron source code needs lots of dependencies.**
This demonstrate that it's not a good idea to install some deps (`PHP`, `Ruby` etc) from `Mise`.

Best practices is to install `PHP` with your OS package manager (`apt`, `yum`, `pacman` etc…); **so comment the line `php = "8.3.2"` in the file [mise.toml](./mise.toml) and retry `mise install`**.

### URLs
- API Go: http://localhost:8090
  - Health: http://localhost:8090/health
  - DB ping + data: http://localhost:8090/items
- PHP: http://localhost:8091
- Web (Angular): http://localhost:5173

## Pourquoi c'est une bonne démo
- Les versions Go/Node/PHP sont **déclarées** dans `mise.toml` et **reproductibles**
- Les tâches sont **déclaratives** (pas besoin d'un Makefile tentaculaire)
- Un `mise run dev` démarre tout et reste lisible

## Astuces
- Liste des tâches avec la description : `mise run` ou `mise tasks`
- Sortie parallèle : par défaut mise exécute plusieurs jobs en parallèle (configurable via `--jobs` / `MISE_JOBS`).

## Notes Angular
- Le front est dans `web/`.
- Au premier lancement la tâche `mise run web:dev` déclenche un `npm install` automatique.
