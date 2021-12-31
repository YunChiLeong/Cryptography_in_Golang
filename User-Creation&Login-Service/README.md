[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-f059dc9a6f8d3a56e377f745f24479a46679e63a5d9fe6f495e02850cd0d8118.svg)](https://classroom.github.com/online_ide?assignment_repo_id=5643195&assignment_repo_type=AssignmentRepo)
# User Login Service
Homework 2, implementing a simple user creation and login service.

## Requirements
The goal of this assignment is for you to create a service that allows users to create accounts and then login to those accounts.

- The service should be managed through Docker Compose and consist of two main components:
    - A Go webserver that performs the application logic
    - A relational database to store the password hashes and salts (You can use what you want, if you don't have strong feelings, just use Postgres)
- When a user registers for the service, their password should be salted and hashed, then the hash, salt, and username should be stored in the database
- When users login, you should retrieve that user's salt from the database, salt and hash password, then check it against the hash stored in the database
- On a successful login, you can just send back the text "Login Successful" through the webserver
- On a failed login attempt, just return "Login Unsuccessful" and maybe a reason why it wasn't successful (no user with that name, password was incorrect, etc)

In a "real" application you would probably generate a token for a user's session once they've logged in and send that token to the user's browser as a cookie, then manage and store that token in the database or some other store (like Redis). Then that token would be sent in each request the user makes and you can check that the token is valid by looking in Redis for it. For this homework we won't do anything useful when a user successfully logs in, but if there's interest we can build it out further.