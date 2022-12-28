# simple-CRUD-example
Simple CRUD Example

[NOT FOR REDISTRIBUTION, NOT FOR COMMERCIAL]
Why ?
1. This is just an example, it is meant to show what I can do to Recruiter / Client.
2. I make this Project from scratch, so if you want to sell this, please appreciate my effort (please contact me if you want code like this, I can create an even better Code than this Project)
3. It's FINE if you copy paste ONLY THE LOGIC,NOT FINE if you just COPY PASTE THE WHOLE PROJECT to your project.
4. don't want my code to be redistributed.

Commands :
- First Time Installation : 
  - make new
- Start / Restart Docker :
  - make docker
- Migrate :
  - make database

Hello, this is an example app created by me, in this scenario this app can :
- Create Game (Insert Game to Database)
- Read Games (Read Games from Database)
- Update Games (Update Game to Database)
- Delete Games (Delete Game to Database)

This app have an Advanced and Simple Version : 
- Advanced : 
  - Characteristics :
    - More structured, 
    - Complex but Recommended for scaling up, 
    - good for Company / Startup, 
    - Hard to make,
    - Less code reuse (because there's constructor)
  - Structure : 
    - main.go
    -  | 
    - route.go
    -  | -> Initialize New Controller
    - controller.go
        -  |
      - [optional] validator.go
      -  |
      - [optional] mapper.go : DTO.go -> model.go
      -  | -> Initialize New DB
      -  | -> Initialize New Service
    - service.go
    -  | -> Initialize New Repository
    - repository.go
    -  | -> Initialize Query
    - query.go
      -  | -> Initialize New QueryTools
  
- Simple :
  - Characteristics :
    - Less Structured, 
    - Simple but Less recommended for scaling up, 
    - good for Small business / Small project,
    - Easier to make,
    - Many code reuse (because no constructor and need to declare at the start of function)
  - Structure :
    - main.go
    -  |
    - route.go
    -  |
    - controller.go
    -  | -> Initialize DB 
    - [optional] mapper.go : DTO.go -> model.go
    -  |
    - query.go

Tech Stack :
- Golang :
  - Pgx v5
  - Go-Playground Validator v10
  - Labstack Echo v4
  - Goqu v9
  - Godotenv
- Database : 
  - Postgresql
- DevOps :
  - Makefile
  - Docker Compose
  - Go Migrate (Database)
  - Air (Hot Reload)
- Tools Used for Development :
  - Linux Ubuntu 22.04
  - Terminator (Linux-Terminal)
  - Dbeaver
  - Oh-my-zsh
  - Postman

Made by Steffen Edlin.