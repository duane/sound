#include <math.h>
#include <stdint.h>
#include <stdio.h>
#include <errno.h>

#include <ao/ao.h>

const double freq = 440.0;
const int bits = 16;
const int channels = 2;
const int rate = 44100;

const char* ao_error(void) {
  switch(errno) {
    case AO_ENODRIVER:
      return "No driver corresponds to driver_id.";
    case AO_ENOTLIVE:
      return "This driver is not a live output device.";
    case AO_EBADOPTION:
      return "A valid option key has an invalid value.";
    case AO_EOPENDEVICE:
      return "Cannot open the device (for example, if /dev/dsp cannot be opened for writing).";
    case AO_EFAIL:
      return "Any other cause of failure in libao.";
  }
  return "Unknown failure.";
}

int main (int argc, char **argv) {
  ao_initialize();

  int driver = ao_default_driver_id();

  ao_sample_format fmt = {
    bits,
    rate,
    channels,
    AO_FMT_LITTLE,
    (char*)"L,R",
  };

  ao_device *dev = ao_open_live(driver, &fmt, NULL /* options */);
  if (!dev) {
    fprintf(stderr, "Unable to open device %d: %s\n", driver, ao_error());
    return 1;
  }

  uint8_t buf[bits/sizeof(char) * channels*rate];
  
  for (int i = 0; i < rate; ++i) {
    unsigned sample = (unsigned)(0.75 * 32768.0 * sin(2 * M_PI * freq * double(i)/rate));

    buf[i * 4] = buf[i * 4 + 2] = (uint8_t)(sample & 0xff);
    buf[i * 4 + 1] = buf[i * 4 + 3] = (uint8_t)((sample >> 8) & 0xff);
  }

  ao_play(dev, (char*)buf, sizeof(buf));
  ao_close(dev);
  ao_shutdown();
  return 0;
}
