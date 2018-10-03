package main

// #cgo CFLAGS:  -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lwiringPi
// #include <wiringPi.h>
import "C"

const PIN_R1 = 9
const PIN_G1 = 8
const PIN_B1 = 7
const PIN_R2 = 0
const PIN_G2 = 15
const PIN_B2 = 2
const PIN_A = 3
const PIN_B = 5
const PIN_C = 12
const PIN_D = 6
const PIN_CLK = 13
const PIN_LAT = 10
const PIN_OE = 14

func main() {
        C.wiringPiSetup();
        C.pinMode(PIN_R1, C.OUTPUT);
        C.pinMode(PIN_G1, C.OUTPUT);
        C.pinMode(PIN_B1, C.OUTPUT);
        C.pinMode(PIN_R2, C.OUTPUT);
        C.pinMode(PIN_G2, C.OUTPUT);
        C.pinMode(PIN_B2, C.OUTPUT);
        C.pinMode(PIN_A,  C.OUTPUT);
        C.pinMode(PIN_B,  C.OUTPUT);
        C.pinMode(PIN_C,  C.OUTPUT);
        C.pinMode(PIN_D,  C.OUTPUT);
        C.pinMode(PIN_CLK,C.OUTPUT);
        C.pinMode(PIN_LAT,C.OUTPUT);
        C.pinMode(PIN_OE, C.OUTPUT);

        for pwm := 0;; pwm = (pwm + 1) % 16 {
            for i := 0; i<16; i++ {
                for j := 0; j<16; j++ {
                    C.digitalWrite(PIN_R1, C.HIGH);
                    C.digitalWrite(PIN_G1, C.LOW);
                    C.digitalWrite(PIN_B1, C.LOW);
                    C.digitalWrite(PIN_R2, C.LOW);
                    C.digitalWrite(PIN_G2, C.LOW);
                    C.digitalWrite(PIN_B2, C.HIGH);
                }
                C.digitalWrite(PIN_CLK, C.HIGH);
                C.digitalWrite(PIN_CLK, C.LOW);

                C.digitalWrite(PIN_A, C.int((i & 0x1) >> 0));
                C.digitalWrite(PIN_B, C.int((i & 0x2) >> 1));
                C.digitalWrite(PIN_C, C.int((i & 0x4) >> 2));
                C.digitalWrite(PIN_D, C.int((i & 0x8) >> 3));
            }

            C.digitalWrite(PIN_OE,  C.HIGH);
            C.digitalWrite(PIN_LAT, C.HIGH);
            C.digitalWrite(PIN_LAT, C.LOW);
            C.digitalWrite(PIN_OE,  C.LOW);

        }
}
