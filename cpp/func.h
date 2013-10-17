#ifndef __FUNC_H__
#define __FUNC_H__

#include <pair>

namespace snd {

typedef interval pair<double, double>;

class function {
 public:
  interval domain, co_domain;
  function(interval domain, interval co_domain) : domain(domain),
                                                  co_domain(co_domain) {}
  virtual ~function() = 0;
  virtual double compute(double x) = 0;
};

class sin : public function {
 public:
  sin(
};

}

#endif
