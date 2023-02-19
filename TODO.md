### Todo
- Look into setting JWT as a cookie
- Add postgres DB (and probably GORM) to store what reports users have already requested
- turn cors handling into middleware: `http.Handle('/endpoint', corsHandler(endPointHandler))


#### Important
- Make it so requests from the front end go via the server, and not the browser!!! This is why the native fetch api is failing

### Login In
- 86 log in page, have inbound traffic to /submit check for a cookie set by squarespace, if no cookie exists, redirect to some url

- add /api/auth endpoint that checks jwt in cookie and if its ok it redirects over /login to /submit
