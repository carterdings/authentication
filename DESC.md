# Instructions to the candidate for the take home task

Write a simple (and possibly not very secure) authentication and authorization service. The service allows users to be authenticated, and authorizes different behavior. No real knowledge of cryptography is assumed. You should complete this by yourself and you are free to use whatever resources/reference materials you see fit. The completed work should be returned via email or public Github no more than 48 hours after receiving this assignment. The exercise itself is not meant to take more than 1-2 hours to complete. Once the submission has been reviewed, a follow-up discussion will be arranged.

## Directions:

1. Write In Java.
2. Don't use any existing security classes from Java (other than hashing/encryption if you need to).
3. The main entities are:
    (1) Users.
    (2) Roles - special entities, that can be associated with individual users. Each user can have multiple roles assigned to it.
4. Keep all data in memory, no persistence storage is required.
5. No need to sign tokens in any special way, it is assumed that the communication channel between the API and the consumer is secure.
6. The main points to address are:
    (1) Clean API.
    (2) Performance of all the main operations.
    (3) Thorough testing (yes, that includes token expiry).
7. The deliverable should be a self-contained project we can easily open and run the tests in IntelliJ.
8. Use Maven or Gradle.
9. Remember, the auth token is something that can be passed around outside your service. While you are welcome to implement it internally the way you like, the value you pass around between calls should be some primitive type, long or string.
10. If you use external libraries outside of the standard JDK, please mention in the README and explain their purpose.

## API to implement:

Feel free to name your functions as you see fit, as long as the action is clearly stated.

1. Create user:
    (1) User name.
    (2) Password - to be stored in some encrypted form.
    (3) Should fail if the user already exists.
2. Delete user:
    (1) User.
    (2) Should fail if the user doesn't exist.
3. Create role:
    (1) Role name.
    (2) Should fail if the role already exists.
4. Delete role:
    (1) Role.
    (2) Should fail if the role doesn't exist.
5. Add role to user:
    (1) User.
    (2) Role.
    (3) If the role is already associated with the user, nothing should happen.
6.Authenticate:
    (1) User name.
    (2) Password.
    (3) Return a special "secret" auth token or error, if not found. The token is only valid for pre-configured time (2h).
7. Invalidate:
    (1) Auth token.
    (2) Returns nothing, the token is no longer valid after the call. Handles correctly the case of invalid token given as input.
8. Check role:
    (1) Auth token.
    (2) Role.
    (3) Returns true if the user, identified by the token, belongs to the role, false otherwise; error if token is invalid, expired, etc.
9. All roles:
    (1) Auth token.
    (2) Returns all roles for the user, error if token is invalid.
