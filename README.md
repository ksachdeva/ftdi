## FTDI MPSSE

<TODO>


## Example - Arduino as a SPI Slave

There is a very simple SPI master program [hello_arduino.go](spi/example/hello_arduino.go) in the example directory that illustrates how to use the API.

Also included is the Sketch [hello_sketch.ino](spi/example/hello_sketch.ino) for Arduino. The sketch is quite simple; all it does is that it waits for a message with '\n' as the 
last character to arrive and if it sees it then it sends it over the serial line.

### Hardware used to test the implementation

adaFruit FT232H breakout board
https://learn.adafruit.com/adafruit-ft232h-breakout

#### Wiring

| Signal| Arduino Pin | FT232H Pin |
| ------| ------------| ---------- |
| Clock | 13          | D0         | 
| MISO  | 12          | D2         |
| MOSI  | 11          | D1         |
| CS    | 10          | D3         |
