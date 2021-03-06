package wzlib_subprocess

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type EnvBufferedExec struct {
	env map[string]string
}

func NewEnvBufferedExec() *EnvBufferedExec {
	envx := new(EnvBufferedExec)
	envx.env = make(map[string]string)
	return envx
}

// SetEnv sets an environment variable
func (envx *EnvBufferedExec) SetEnv(key string, value string) *EnvBufferedExec {
	envx.env[key] = value
	return envx
}

// Exec a command with the specific environment
func (envx *EnvBufferedExec) Exec(name string, args ...string) (*BufferedCmd, error) {
	cmd := ExecCommand(name, args...)
	cmd.Env = os.Environ()
	for ek, ev := range envx.env {
		if strings.Contains(ev, " ") {
			ev = fmt.Sprintf("\"%s\"", ev)
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", strings.TrimSpace(ek), ev))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return &BufferedCmd{cmd, stdin, bufio.NewReaderSize(stdout, stdoutBufSize),
		bufio.NewReaderSize(stderr, stdoutBufSize)}, nil
}

// BufferedExec starts up a command and creates a stdin pipe and a buffered
// stdout & stderr pipes, wrapped in a BufferedCmd. The stdout buffer will be
// of stdoutBufSize bytes.
func BufferedExec(name string, args ...string) (*BufferedCmd, error) {
	cmd := ExecCommand(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return &BufferedCmd{
		cmd,
		stdin,
		bufio.NewReaderSize(stdout, stdoutBufSize),
		bufio.NewReaderSize(stderr, stdoutBufSize),
	}, nil
}

// SimpleExec is a small wrapper around os/exec.Command.
func SimpleExec(name string, args ...string) (string, error) {
	Trace(name, args...)
	return Output(ExecCommand(name, args...))
}

func Output(cmd *Cmd) (string, error) {
	out, err := cmd.Output()

	if exitError, ok := err.(*exec.ExitError); ok {
		errorOutput := strings.TrimSpace(string(exitError.Stderr))
		if errorOutput == "" {
			// some commands might write nothing to stderr but something to stdout in error-conditions, in which case, we'll use that
			// in the error string
			errorOutput = strings.TrimSpace(string(out))
		}

		ran := cmd.Path
		if len(cmd.Args) > 1 {
			ran = fmt.Sprintf("%s %s", cmd.Path, quotedArgs(cmd.Args[1:]))
		}
		formattedErr := fmt.Errorf("error running %s: '%s' '%s'", ran, errorOutput, strings.TrimSpace(exitError.Error()))

		// return "" as output in error case, for callers that don't care about errors but rely on "" returned, in-case stdout != ""
		return "", formattedErr
	}

	return strings.Trim(string(out), " \n"), err
}

var shellWordRe = regexp.MustCompile(`\A[A-Za-z0-9_@/.-]+\z`)

// ShellQuoteSingle returns a string which is quoted suitably for sh.
func ShellQuoteSingle(str string) string {
	// Quote anything that looks slightly complicated.
	if shellWordRe.FindStringIndex(str) == nil {
		return "'" + strings.Replace(str, "'", "'\\''", -1) + "'"
	}
	return str
}

// ShellQuote returns a copied string slice where each element is quoted
// suitably for sh.
func ShellQuote(strs []string) []string {
	dup := make([]string, 0, len(strs))

	for _, str := range strs {
		dup = append(dup, ShellQuoteSingle(str))
	}
	return dup
}

// FormatForShell takes a command name and an argument string and returns a
// command and arguments that pass this command to the shell.  Note that neither
// the command nor the arguments are quoted.  Consider FormatForShellQuoted
// instead.
func FormatForShell(name string, args string) (string, []string) {
	return "sh", []string{"-c", name + " " + args}
}

// FormatForShellQuotedArgs takes a command name and an argument string and
// returns a command and arguments that pass this command to the shell.  The
// arguments are escaped, but the name of the command is not.
func FormatForShellQuotedArgs(name string, args []string) (string, []string) {
	return FormatForShell(name, strings.Join(ShellQuote(args), " "))
}

func Trace(name string, args ...string) {
	fmt.Printf("exec: %s %s", name, quotedArgs(args))
}

func quotedArgs(args []string) string {
	if len(args) == 0 {
		return ""
	}

	quoted := make([]string, len(args))
	for i, arg := range args {
		quoted[i] = fmt.Sprintf("'%s'", arg)
	}
	return strings.Join(quoted, " ")
}

// CallStreaming returns also STDOUT to a custom output writer.
// Usually done that this module accepts a filename to which the output is piped,
// so tailing that file would give an STDOUT streaming integration,
// while keeping JSON output safe.
func StreamedExec(pipe ProcessStream, name string, args ...string) (string, string) {
	cmd := ExecCommand(name, args...)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(pipe, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&stderrBuf)
	cmd.Run()
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())

	return outStr, errStr
}
