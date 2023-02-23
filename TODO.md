### Todo
- Look into setting JWT as a cookie
- Add postgres DB (and probably GORM) to store what reports users have already requested
- Turn JWT into middleware instead of doing it in the handler
- Verify if new server side fetch request and buffer->blob handling allows for hiding the generate API behind the nginx container and that it doesnt ruin the /api/login endpoint


### Login In
- 86 log in page, have inbound traffic to /submit check for a cookie set by squarespace, if no cookie exists, redirect to some url

- add /api/auth endpoint that checks jwt in cookie and if its ok it redirects over /login to /submit
