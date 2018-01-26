#include <SPI.h>

char buf[100];
volatile byte pos;
volatile boolean process_it;

void setup(void)
{
    Serial.begin(115200); // debugging

    // have to send on master in, *slave out*
    pinMode(MISO, OUTPUT);

    // turn on SPI in slave mode
    SPCR |= _BV(SPE);

    // get ready for an interrupt
    pos = 0; // buffer empty
    process_it = false;

    // now turn on interrupts
    SPI.attachInterrupt();

} // end of setup

// SPI interrupt routine
ISR(SPI_STC_vect)
{
    byte c = SPDR; // grab byte from SPI Data Register

    // add to buffer if room
    if (pos < sizeof buf)
    {
        buf[pos++] = c;

        // newline means time to process buffer
        if (c == '\n')
            process_it = true;

    } // end of room available
} // end of interrupt routine SPI_STC_vect

// main loop - wait for flag set in interrupt routine
void loop(void)
{
    if (process_it)
    {
        buf[pos] = 0;
        Serial.println(buf);
        pos = 0;
        process_it = false;
    } // end of flag set

} // end of loop
