# bcryptr

bcryptr is an extremely simple tool to generate bcrypt hashes on the command line.

## Installation

    go install github.com/bigwhoop/bcryptr

## Instructions

    bcryptr [-c COST] PASSWORD

    $ bcryptr somepassword
    $2a$10$cW0P/yg1CKkky0QXBbmEnubzxVT8YS39JmBZ8CeEfQ4Sk5GXCe6im
    
    $ bcryptr -c 14 somepassword
    $2a$14$oS5E7WxSQF9xNo2CKAiA6e52SQPKo83WKJUt/ZgygJ2gBHxuV5r.6