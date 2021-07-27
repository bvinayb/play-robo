# Play Robo Simulator

## Description
- The application is a simulation of a toy robot moving on a square tabletop,
  of dimensions 5 units x 5 units.
- There are no other obstructions on the table surface.
- The robot is free to roam around the surface of the table, but must be
  prevented from falling to destruction. Any movement that would result in the
  robot falling from the table must be prevented, however further valid
  movement commands must still be allowed.
  
## Development
* The source is already compiled and the binary is sitting under bin folder, so probably `make build` not needed unless source is changed.
* Run `make play` - Compiles and builds binary and runs the binary
* Run  `make tests` - Run all the unit test

## Example Input and Output
```
### Example a
 
    PLACE 0,0,NORTH
    MOVE
    REPORT

Expected output:

    0,1,NORTH

### Example b

    PLACE 0,0,NORTH
    LEFT
    REPORT

Expected output:

    0,0,WEST

### Example c

    PLACE 1,2,EAST
    MOVE
    MOVE
    LEFT
    MOVE
    REPORT

Expected output

    3,3,NORTH
```
## Commands

* Build: This will build the compile the source to create binary.
  ```sh
  make build
  ```
* Run play command
  ```sh
  make play
  ```
* Run Tests:
  ```sh
  make tests
  ```
* Check Linting Issues.
  ```sh
  make lint
  ```
* Check Code Issues.
  ```sh
  make vet
  ```

