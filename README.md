# Lens Locked
Website created based on Jon Calhoun's [Web Development with Go](https://www.usegolang.com/) course.

Code has been implemented through Chapter 10.

## Start PostgreSQL in Docker Container
To run PostgreSQL in a Docker container in the background: `podman run --rm --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres`

## Dynamic Reloading
To use [fresh](https://github.com/gravityblast/fresh) for dynamic reloading when files change: `fresh`
