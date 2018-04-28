# blazr server

(Not graphql anymore just normal rest endpoints for now)

Layers:

api - The endpoints and request handlers <br />
services - The business logic that handles matchmaking and geolocation logic <br />
daos - The database abstractions used for profiles and matches <br />

Api endpoints: 

GET /profile/:userID    - Gets the requested profile <br />
DELETE /profile/:userID - Deletes the requested profile <br />
POST /profile/          - Creates a profile, must have new profile in request body <br />
PUT /profile/           - Updates a profile, must have profile id and profile in request body <br />
GET /profiles/          - Gets profiles from around the user, must have current location (rad) and search radius (mi) as request params <br />
GET /matches/:userID    - Gets a list of a users' matches <br />
PUT /location/          - Updates a users' location, must have user id and new location (rad) as request params <br />

RabbitMQ: 

Listens to messages in the following format: userID A <space> userID B whenever a tile on the client is swiped, then attempts to make a match
 
Sends two messages in the following format: userID A <space> userID B / userID B <space> userID A when a match event occurs
  
Changes / fixes

- RabbitMQ and MongoDB currently set up to run on localhost <br />
- No authentication set up for using api endpoints <br />
- No data validation done on create / update endpoints <br />
- Not coded in most efficient/ scalable manner <br />
