package middleware

// middleware that checks every hour or so if the session id is still valid.
// uses in memory database, size can be set at initialisation and it will remove the oldest entry if it's full
// to increase performance, get requests can be checked asynchronous, so the request and the check will run in paralell


