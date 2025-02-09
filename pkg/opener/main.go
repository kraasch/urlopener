
package opener

import (
  "runtime"
  "os/exec"
  "strings"
)

// opens the specified URL in the default browser of the user.
func OpenUrl(url string) error {
  var cmd string
  var args []string

  switch runtime.GOOS {
  case "windows":
    cmd = "cmd"
    args = []string{"/c", "start", url}
  case "darwin":
    cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
    // check if running under wsl
    if isWsl() {
      // use 'cmd.exe /c start' to open the url in the default windows browser.
      cmd = "cmd.exe"
      args = []string{"/c", "start", url}
    } else {
      // use xdg-open on native linux environments
      cmd = "xdg-open"
      args = []string{url}
    }
  }
  if len(args) > 1 {
    // args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
    args = append(args[:1], append([]string{""}, args[1:]...)...)
  }
  return exec.Command(cmd, args...).Start()
}

// checks if the go program is running inside windows subsystem for linux (wsl).
func isWsl() bool {
  releaseData, err := exec.Command("uname", "-r").Output()
  if err != nil {
    return false
  }
  return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

