package ao

// #include <stdint.h>
// #include <errno.h>
// #include <ao/ao.h>
// static int get_errno(void) { return errno; }
// 
// #cgo LDFLAGS: -lao
import "C"
import "unsafe"

const (
  FMT_LITTLE int = C.AO_FMT_LITTLE
  FMT_BIG    int = C.AO_FMT_BIG
  FMT_NATIVE int = C.AO_FMT_NATIVE
)

type Device struct {
  dev *C.ao_device
}

type Info struct {
  info *C.ao_info
}

type Option struct {
  option *C.ao_option
}

type SampleFormat struct {
  format *C.ao_sample_format
}

func Initialize() {
  C.ao_initialize()
}

func Shutdown() {
  C.ao_shutdown()
}

func Error() string {
  switch C.get_errno() {
    case C.AO_ENODRIVER:
      return "No driver corresponds to driver_id."
    case C.AO_ENOTLIVE:
      return "This driver is not a live output device."
    case C.AO_EBADOPTION:
      return "A valid option key has an invalid value."
    case C.AO_EOPENDEVICE:
      return "Cannot open the device (for example, if /dev/dsp cannot be opened for writing)."
    case C.AO_EFAIL:
      return "Any other cause of failure in libao."
  }
  return "Unknown failure."
}

func MakeFormat(bits, rate, channels, byte_format int, matrix string) SampleFormat {
  cmatrix := C.CString(matrix)
  defer C.free(unsafe.Pointer(cmatrix))
  cfmt := C.ao_sample_format{C.int(bits),
                             C.int(rate),
                             C.int(channels),
                             C.int(byte_format),
                             cmatrix}
  return SampleFormat{&cfmt}
}

func AppendOption(options *Option, key string, value string) int {
  ckey := C.CString(key)
  cval := C.CString(value)
  defer C.free(unsafe.Pointer(ckey))
  defer C.free(unsafe.Pointer(cval))
  status := C.ao_append_option(&options.option, ckey, cval)
  return int(status)
}

func AppendGlobalOption(key string, value string) int {
  ckey := C.CString(key)
  cval := C.CString(value)
  defer C.free(unsafe.Pointer(ckey))
  defer C.free(unsafe.Pointer(cval))
  status := C.ao_append_global_option(ckey, cval)
  return int(status)
}

func FreeOptions(options *Option) {
  C.ao_free_options(options.option)
}

func OpenLive(driver_id int, format *SampleFormat, options *Option) *Device {
  var option *C.ao_option
  if options == nil {
    option = nil
  } else {
    option = options.option
  }
  dev := Device{C.ao_open_live(C.int(driver_id), format.format, option)}
  if dev.dev == nil {

    return nil
  }
  return &dev
}

func OpenFile(driver_id int,
              filename string,
              overwrite int,
              format *SampleFormat,
              options *Option) *Device {
  cfile := C.CString(filename)
  defer C.free(unsafe.Pointer(cfile))
  dev := Device{C.ao_open_file(C.int(driver_id), cfile, C.int(overwrite), format.format, options.option)}
  if dev.dev == nil {
    return nil
  }
  return &dev
}

func Play(device *Device,
          output_samples []byte) int {
  samples_ptr := (*C.char)(unsafe.Pointer(&output_samples[0]))
  return int(C.ao_play(device.dev, samples_ptr, C.uint_32(len(output_samples))))
}

func Close(device *Device) int {
  return int(C.ao_close(device.dev))
}

func DefaultDriverId() int {
  return int(C.ao_default_driver_id())
}
