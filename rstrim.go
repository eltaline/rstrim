package rstrim

import (
	"regexp"
	"strings"
)

var (
	wOnly  = regexp.MustCompile("(?m)^[ \t]+$")
	lWhite = regexp.MustCompile("(?m)(^[ \t]*)(?:[^ \t\n])")
)

// Rstrim: function for remove whitespaces
func Rstrim(inline string) string {

  var m string

  inline = wOnly.ReplaceAllString(inline, "")
  inds := lWhite.FindAllStringSubmatch(inline, -1)
	
  for i, ind := range inds {

    if i == 0 {
		  m = ind[1]
		} else if strings.HasPrefix(ind[1], m) {
      continue
		} else if strings.HasPrefix(m, ind[1]) {
			m = ind[1]
		} else {
			m = ""
			break
		}

}

	if m != "" {
		  inline = regexp.MustCompile("(?m)^"+m).ReplaceAllString(inline, "")
	}
	
  return inline

}
