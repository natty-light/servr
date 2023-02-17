### Todo
- Look into setting JWT as a cookie
- Add Redis cache for seeing what reports have been generated for that day? already
- Add postgres DB (and probably GORM) to store what reports users have already requested


#### Important
- Make it so requests from the front end go via the server, and not the browser!!! This is why the native fetch api is failing