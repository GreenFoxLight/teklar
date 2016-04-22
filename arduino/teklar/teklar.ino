#include <RH_ASK.h>
#include <SPI.h>

RH_ASK driver;

void setup() {
        pinMode(13, OUTPUT);
}

void loop() {<
        uint8_t buf[RH_ASK_MAX_MESSAGE_LEN];
        uint8_t buflen = sizeof(buf);
        if (driver.recv(buf, &buflen)) {
                int i;
                for (i = 0; i < buflen; i++) {
                        for (uint8_t j = 0; j < buf[i]; j++) {
                                digitalWrite(13, HIGH);
                                delay(100);
                                digitalWrite(13, LOW);
                                delay(100);
                        }
                }
        }
        
}
