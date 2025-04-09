# Laelapa/GoHome
A personal webpage+blog+portfolio using go, templ and htmx. 

## Development Environment
To set up a live reloading environment so that you can edit your templates and tailwindcss and have the changes propagate in realtime to your browser run:

```
$make run-watch
```
This uses the functionality of the templ library for [live reloading](https://templ.guide/developer-tools/live-reload), and also runs `tailwindcss --watch` on the background, which you will need to terminate with `pkill tailwindcss` after you are done.
