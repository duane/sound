#ifndef __SOUND_H__
#define __SOUND_H__

namespace snd {

struct frame {
  uint16_t channel a, channel b;
};

struct sample {
  uint32_t samples;
  frame *frames;
}

class Sound {
 public:
  Sound() {}
  ~Sound() {}
}
}

#endif
